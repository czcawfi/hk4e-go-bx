package service

import (
	"context"
	"math"
	"sort"
	"strings"
	"sync"
	"time"

	"hk4e/common/mq"
	"hk4e/common/region"
	"hk4e/node/api"
	"hk4e/pkg/logger"
	"hk4e/pkg/random"

	"github.com/pkg/errors"
)

const (
	MaxGsId = 1000
)

var _ api.DiscoveryNATSRPCServer = (*DiscoveryService)(nil)

// ServerInstance 服务器实例
type ServerInstance struct {
	serverType        string   // 服务器类型
	appId             string   // appid
	gateServerKcpAddr string   // 网关kcp地址
	gateServerKcpPort uint32   // 网关kcp端口
	gateServerMqAddr  string   // 网关tcp直连消息队列地址
	gateServerMqPort  uint32   // 网关tcp直连消息队列端口
	version           []string // 网关支持的客户端协议版本
	lastAliveTime     int64    // 最后保活时间
	gsId              uint32   // 游戏服务器编号
	loadCount         uint32   // 负载数
}

// StopServerInfo 停服信息
type StopServerInfo struct {
	stopServer      bool                // 是否停服
	startTime       uint32              // 停服开始时间
	endTime         uint32              // 停服结束时间
	ipAddrWhiteList map[string]struct{} // ip地址白名单
}

type DiscoveryService struct {
	regionEc2b        *random.Ec2b         // 区服密钥信息
	serverInstanceMap map[string]*sync.Map // 全部服务器实例集合 key:服务器类型 value:服务器实例集合 -> key:appid value:服务器实例
	serverAppIdMap    *sync.Map            // 服务器appid集合 key:appid value:是否存在
	globalGsOnlineMap *sync.Map            // 全服玩家在线集合 key:uid value:gsAppid
	stopServerInfo    *StopServerInfo      // 停服信息
	messageQueue      *mq.MessageQueue     // 消息队列实例
}

func NewDiscoveryService(messageQueue *mq.MessageQueue) *DiscoveryService {
	r := new(DiscoveryService)
	r.regionEc2b = region.LoadRegionEc2b()
	logger.Info("region ec2b create ok, seed: %v", r.regionEc2b.Seed())
	r.serverInstanceMap = make(map[string]*sync.Map)
	r.serverInstanceMap[api.GATE] = new(sync.Map)
	r.serverInstanceMap[api.GS] = new(sync.Map)
	r.serverInstanceMap[api.ANTICHEAT] = new(sync.Map)
	r.serverInstanceMap[api.PATHFINDING] = new(sync.Map)
	r.serverInstanceMap[api.ROBOT] = new(sync.Map)
	r.serverInstanceMap[api.DISPATCH] = new(sync.Map)
	r.serverAppIdMap = new(sync.Map)
	r.globalGsOnlineMap = new(sync.Map)
	r.stopServerInfo = &StopServerInfo{
		stopServer:      false,
		startTime:       0,
		endTime:         0,
		ipAddrWhiteList: make(map[string]struct{}),
	}
	r.messageQueue = messageQueue
	go r.removeDeadServer()
	go r.broadcastReceiver()
	return r
}

func (s *DiscoveryService) broadcastReceiver() {
	for {
		netMsg := <-s.messageQueue.GetNetMsg()
		if netMsg.MsgType != mq.MsgTypeServer {
			continue
		}
		if netMsg.EventId != mq.ServerUserOnlineStateChangeNotify {
			continue
		}
		if netMsg.OriginServerType != api.GS {
			continue
		}
		serverMsg := netMsg.ServerMsg
		if serverMsg.IsOnline {
			s.globalGsOnlineMap.Store(serverMsg.UserId, netMsg.OriginServerAppId)
		} else {
			s.globalGsOnlineMap.Delete(serverMsg.UserId)
		}
	}
}

// RegisterServer 服务器启动注册获取appid
func (s *DiscoveryService) RegisterServer(ctx context.Context, req *api.RegisterServerReq) (*api.RegisterServerRsp, error) {
	logger.Info("register new server, server type: %v", req.ServerType)
	instMap, exist := s.serverInstanceMap[req.ServerType]
	if !exist {
		return nil, errors.New("server type not exist")
	}
	var appId string
	for {
		appId = strings.ToLower(random.GetRandomStr(8))
		_, exist := s.serverAppIdMap.Load(appId)
		if !exist {
			s.serverAppIdMap.Store(appId, true)
			break
		}
	}
	inst := &ServerInstance{
		serverType:    req.ServerType,
		appId:         appId,
		lastAliveTime: time.Now().Unix(),
		loadCount:     0,
	}
	if req.ServerType == api.GATE {
		logger.Info("register new gate server, ip: %v, port: %v", req.GateServerAddr.KcpAddr, req.GateServerAddr.KcpPort)
		inst.gateServerKcpAddr = req.GateServerAddr.KcpAddr
		inst.gateServerKcpPort = req.GateServerAddr.KcpPort
		inst.gateServerMqAddr = req.GateServerAddr.MqAddr
		inst.gateServerMqPort = req.GateServerAddr.MqPort
		inst.version = req.Version
	}
	instMap.Store(appId, inst)
	logger.Info("new server appid is: %v", appId)
	rsp := &api.RegisterServerRsp{
		AppId: appId,
	}
	if req.ServerType == api.GS {
		gsIdUseList := make([]bool, MaxGsId+1)
		gsIdUseList[0] = true
		instMap.Range(func(key, value any) bool {
			serverInstance := value.(*ServerInstance)
			if serverInstance.gsId > MaxGsId {
				logger.Error("invalid gs id inst: %v", serverInstance)
				return true
			}
			gsIdUseList[serverInstance.gsId] = true
			return true
		})
		newGsId := uint32(0)
		for gsId, use := range gsIdUseList {
			if !use {
				newGsId = uint32(gsId)
				break
			}
		}
		if newGsId == 0 {
			return nil, errors.New("no gs id can use")
		}
		inst.gsId = newGsId
		rsp.GsId = newGsId
	}
	return rsp, nil
}

// CancelServer 服务器关闭取消注册
func (s *DiscoveryService) CancelServer(ctx context.Context, req *api.CancelServerReq) (*api.NullMsg, error) {
	logger.Info("server cancel, server type: %v, appid: %v", req.ServerType, req.AppId)
	instMap, exist := s.serverInstanceMap[req.ServerType]
	if !exist {
		return nil, errors.New("server type not exist")
	}
	_, exist = instMap.Load(req.AppId)
	if !exist {
		logger.Error("recv not exist server cancel, server type: %v, appid: %v", req.ServerType, req.AppId)
		return nil, errors.New("server not exist")
	}
	instMap.Delete(req.AppId)
	return &api.NullMsg{}, nil
}

// KeepaliveServer 服务器在线心跳保持
func (s *DiscoveryService) KeepaliveServer(ctx context.Context, req *api.KeepaliveServerReq) (*api.NullMsg, error) {
	logger.Debug("server keepalive, server type: %v, appid: %v, load: %v", req.ServerType, req.AppId, req.LoadCount)
	instMap, exist := s.serverInstanceMap[req.ServerType]
	if !exist {
		return nil, errors.New("server type not exist")
	}
	inst, exist := instMap.Load(req.AppId)
	if !exist {
		logger.Error("recv not exist server keepalive, server type: %v, appid: %v", req.ServerType, req.AppId)
		return nil, errors.New("server not exist")
	}
	serverInstance := inst.(*ServerInstance)
	serverInstance.lastAliveTime = time.Now().Unix()
	serverInstance.loadCount = req.LoadCount
	return &api.NullMsg{}, nil
}

// GetServerAppId 获取负载最小的服务器的appid
func (s *DiscoveryService) GetServerAppId(ctx context.Context, req *api.GetServerAppIdReq) (*api.GetServerAppIdRsp, error) {
	logger.Debug("get server instance, server type: %v", req.ServerType)
	instMap, exist := s.serverInstanceMap[req.ServerType]
	if !exist {
		return nil, errors.New("server type not exist")
	}
	if s.getServerInstanceMapLen(instMap) == 0 {
		return nil, errors.New("no server found")
	}
	var inst *ServerInstance = nil
	if req.ServerType == api.GATE || req.ServerType == api.GS {
		inst = s.getMinLoadServerInstance(instMap)
	} else {
		inst = s.getRandomServerInstance(instMap)
	}
	logger.Debug("get server appid is: %v", inst.appId)
	return &api.GetServerAppIdRsp{
		AppId: inst.appId,
	}, nil
}

// GetRegionEc2B 获取区服密钥信息
func (s *DiscoveryService) GetRegionEc2B(ctx context.Context, req *api.NullMsg) (*api.RegionEc2B, error) {
	logger.Info("get region ec2b ok")
	return &api.RegionEc2B{
		Data: s.regionEc2b.Bytes(),
	}, nil
}

// GetGateServerAddr 获取负载最小的网关服务器的地址和端口
func (s *DiscoveryService) GetGateServerAddr(ctx context.Context, req *api.GetGateServerAddrReq) (*api.GateServerAddr, error) {
	logger.Debug("get gate server addr")
	instMap, exist := s.serverInstanceMap[api.GATE]
	if !exist {
		return nil, errors.New("gate server not exist")
	}
	if s.getServerInstanceMapLen(instMap) == 0 {
		return nil, errors.New("no gate server found")
	}
	versionInstMap := sync.Map{}
	instMap.Range(func(key, value any) bool {
		serverInstance := value.(*ServerInstance)
		for _, version := range serverInstance.version {
			if version == req.Version {
				versionInstMap.Store(key, serverInstance)
				return true
			}
		}
		return true
	})
	if s.getServerInstanceMapLen(&versionInstMap) == 0 {
		return nil, errors.New("no gate server found")
	}
	inst := s.getMinLoadServerInstance(&versionInstMap)
	logger.Debug("get gate server addr is, ip: %v, port: %v", inst.gateServerKcpAddr, inst.gateServerKcpPort)
	return &api.GateServerAddr{
		KcpAddr: inst.gateServerKcpAddr,
		KcpPort: inst.gateServerKcpPort,
	}, nil
}

// GetAllGateServerInfoList 获取全部网关服务器信息列表
func (s *DiscoveryService) GetAllGateServerInfoList(ctx context.Context, req *api.NullMsg) (*api.GateServerInfoList, error) {
	logger.Debug("get all gate server info list")
	instMap, exist := s.serverInstanceMap[api.GATE]
	if !exist {
		return nil, errors.New("gate server not exist")
	}
	if s.getServerInstanceMapLen(instMap) == 0 {
		return nil, errors.New("no gate server found")
	}
	gateServerInfoList := make([]*api.GateServerInfo, 0)
	instMap.Range(func(key, value any) bool {
		serverInstance := value.(*ServerInstance)
		gateServerInfoList = append(gateServerInfoList, &api.GateServerInfo{
			AppId:  serverInstance.appId,
			MqAddr: serverInstance.gateServerMqAddr,
			MqPort: serverInstance.gateServerMqPort,
		})
		return true
	})
	return &api.GateServerInfoList{
		GateServerInfoList: gateServerInfoList,
	}, nil
}

// GetMainGameServerAppId 获取主游戏服务器的appid
func (s *DiscoveryService) GetMainGameServerAppId(ctx context.Context, req *api.NullMsg) (*api.GetMainGameServerAppIdRsp, error) {
	logger.Debug("get main game server appid")
	instMap, exist := s.serverInstanceMap[api.GS]
	if !exist {
		return nil, errors.New("game server not exist")
	}
	if s.getServerInstanceMapLen(instMap) == 0 {
		return nil, errors.New("no game server found")
	}
	appid := ""
	mainGsId := uint32(1)
	instMap.Range(func(key, value any) bool {
		serverInstance := value.(*ServerInstance)
		if serverInstance.gsId == mainGsId {
			appid = serverInstance.appId
			return false
		}
		return true
	})
	if appid == "" {
		return nil, errors.New("main game server not found")
	}
	return &api.GetMainGameServerAppIdRsp{
		AppId: appid,
	}, nil
}

// GetGlobalGsOnlineMap 获取全服玩家GS在线列表
func (s *DiscoveryService) GetGlobalGsOnlineMap(ctx context.Context, req *api.NullMsg) (*api.GlobalGsOnlineMap, error) {
	copyMap := make(map[uint32]string)
	s.globalGsOnlineMap.Range(func(key, value any) bool {
		copyMap[key.(uint32)] = value.(string)
		return true
	})
	return &api.GlobalGsOnlineMap{
		OnlineMap: copyMap,
	}, nil
}

// GetStopServerInfo 获取停服维护信息
func (s *DiscoveryService) GetStopServerInfo(ctx context.Context, req *api.GetStopServerInfoReq) (*api.StopServerInfo, error) {
	stopServer := s.stopServerInfo.stopServer
	_, exist := s.stopServerInfo.ipAddrWhiteList[req.ClientIpAddr]
	if exist {
		stopServer = false
	}
	return &api.StopServerInfo{
		StopServer: stopServer,
		StartTime:  s.stopServerInfo.startTime,
		EndTime:    s.stopServerInfo.endTime,
	}, nil
}

// SetStopServerInfo 修改停服维护信息
func (s *DiscoveryService) SetStopServerInfo(ctx context.Context, req *api.StopServerInfo) (*api.NullMsg, error) {
	if s.stopServerInfo.stopServer == false && req.StopServer == true {
		s.messageQueue.SendToAll(&mq.NetMsg{
			MsgType: mq.MsgTypeServer,
			EventId: mq.ServerStopNotify,
		})
	}
	s.stopServerInfo.stopServer = req.StopServer
	s.stopServerInfo.startTime = req.StartTime
	s.stopServerInfo.endTime = req.EndTime
	return &api.NullMsg{}, nil
}

// GetWhiteList 获取停服维护白名单
func (s *DiscoveryService) GetWhiteList(ctx context.Context, req *api.NullMsg) (*api.GetWhiteListRsp, error) {
	ipAddrList := make([]string, 0)
	for ipAddr := range s.stopServerInfo.ipAddrWhiteList {
		ipAddrList = append(ipAddrList, ipAddr)
	}
	return &api.GetWhiteListRsp{IpAddrList: ipAddrList}, nil
}

// SetWhiteList 修改停服维护白名单
func (s *DiscoveryService) SetWhiteList(ctx context.Context, req *api.SetWhiteListReq) (*api.NullMsg, error) {
	if req.IsAdd {
		s.stopServerInfo.ipAddrWhiteList[req.IpAddr] = struct{}{}
	} else {
		delete(s.stopServerInfo.ipAddrWhiteList, req.IpAddr)
	}
	return &api.NullMsg{}, nil
}

func (s *DiscoveryService) getRandomServerInstance(instMap *sync.Map) *ServerInstance {
	instList := make([]*ServerInstance, 0)
	instMap.Range(func(key, value any) bool {
		instList = append(instList, value.(*ServerInstance))
		return true
	})
	sort.SliceStable(instList, func(i, j int) bool {
		return instList[i].appId < instList[j].appId
	})
	index := random.GetRandomInt32(0, int32(len(instList)-1))
	inst := instList[index]
	return inst
}

func (s *DiscoveryService) getMinLoadServerInstance(instMap *sync.Map) *ServerInstance {
	instList := make([]*ServerInstance, 0)
	instMap.Range(func(key, value any) bool {
		instList = append(instList, value.(*ServerInstance))
		return true
	})
	sort.SliceStable(instList, func(i, j int) bool {
		return instList[i].appId < instList[j].appId
	})
	minLoadInstIndex := 0
	minLoadInstCount := math.MaxUint32
	for index, inst := range instList {
		if inst.loadCount < uint32(minLoadInstCount) {
			minLoadInstCount = int(inst.loadCount)
			minLoadInstIndex = index
		}
	}
	inst := instList[minLoadInstIndex]
	return inst
}

func (s *DiscoveryService) getServerInstanceMapLen(instMap *sync.Map) int {
	count := 0
	instMap.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

// 定时移除掉线服务器
func (s *DiscoveryService) removeDeadServer() {
	ticker := time.NewTicker(time.Second * 10)
	for {
		<-ticker.C
		nowTime := time.Now().Unix()
		for _, instMap := range s.serverInstanceMap {
			instMap.Range(func(appid, inst any) bool {
				serverInstance := inst.(*ServerInstance)
				if nowTime-serverInstance.lastAliveTime > 30 {
					instMap.Delete(appid)
					s.handleServerDead(serverInstance)
				}
				return true
			})
		}
	}
}

func (s *DiscoveryService) handleServerDead(serverInstance *ServerInstance) {
	logger.Warn("remove dead server, server type: %v, appid: %v, last alive time: %v",
		serverInstance.serverType, serverInstance.appId, serverInstance.lastAliveTime)
	if serverInstance.serverType == api.GS {
		s.globalGsOnlineMap.Range(func(uid, gsAppid any) bool {
			if serverInstance.appId == gsAppid.(string) {
				s.globalGsOnlineMap.Delete(uid)
			}
			return true
		})
	}
}
