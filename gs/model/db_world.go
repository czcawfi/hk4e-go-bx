package model

import (
	"hk4e/common/constant"
	"hk4e/gdconf"
)

type DbWorld struct {
	SceneMap    map[uint32]*DbScene
	MapMarkList []*MapMark
}

type DbScene struct {
	SceneId        uint32
	UnlockPointMap map[uint32]bool
	UnHidePointMap map[uint32]bool
	SceneGroupMap  map[uint32]*DbSceneGroup
	VehicleMap     map[uint32]*DbVehicle
}

type DbSceneGroup struct {
	VariableMap    map[string]int32
	KillConfigMap  map[uint32]bool
	GadgetStateMap map[uint32]uint8
}

type DbVehicle struct {
	VehicleId uint32  // 载具Id
	OwnerUid  uint32  // 所有者Id
	Pos       *Vector // 位置
	Rot       *Vector // 旋转
}

type MapMark struct {
	SceneId   uint32
	Pos       *Vector
	PointType uint32
	Name      string
}

func (p *Player) GetDbWorld() *DbWorld {
	if p.DbWorld == nil {
		p.DbWorld = new(DbWorld)
	}
	if p.DbWorld.SceneMap == nil {
		p.DbWorld.SceneMap = make(map[uint32]*DbScene)
	}
	if p.DbWorld.MapMarkList == nil {
		p.DbWorld.MapMarkList = make([]*MapMark, 0)
	}
	return p.DbWorld
}

func (w *DbWorld) GetSceneById(sceneId uint32) *DbScene {
	scene, exist := w.SceneMap[sceneId]
	// 不存在自动创建场景
	if !exist {
		// 拒绝创建配置表中不存在的非法场景
		sceneDataConfig := gdconf.GetSceneDataById(int32(sceneId))
		if sceneDataConfig == nil {
			return nil
		}
		scene = new(DbScene)
		w.SceneMap[sceneId] = scene
	}
	if scene.SceneId == 0 {
		scene.SceneId = sceneId
	}

	if scene.UnlockPointMap == nil {
		scene.UnlockPointMap = make(map[uint32]bool)
	}
	if scene.UnHidePointMap == nil {
		scene.UnHidePointMap = make(map[uint32]bool)
	}
	if scene.SceneGroupMap == nil {
		scene.SceneGroupMap = make(map[uint32]*DbSceneGroup)
	}
	if scene.VehicleMap == nil {
		scene.VehicleMap = make(map[uint32]*DbVehicle)
	}
	return scene
}

func (s *DbScene) GetUnHidePointList() []uint32 {
	unHidePointList := make([]uint32, 0, len(s.UnHidePointMap))
	for pointId := range s.UnHidePointMap {
		unHidePointList = append(unHidePointList, pointId)
	}
	return unHidePointList
}

func (s *DbScene) GetUnlockPointList() []uint32 {
	unlockPointList := make([]uint32, 0, len(s.UnlockPointMap))
	for pointId := range s.UnlockPointMap {
		unlockPointList = append(unlockPointList, pointId)
	}
	return unlockPointList
}

func (s *DbScene) UnlockPoint(pointId uint32) {
	pointDataConfig := gdconf.GetScenePointBySceneIdAndPointId(int32(s.SceneId), int32(pointId))
	if pointDataConfig == nil {
		return
	}
	s.UnlockPointMap[pointId] = true
	// 隐藏锚点取消隐藏
	if pointDataConfig.IsModelHidden {
		s.UnHidePointMap[pointId] = true
	}
}

func (s *DbScene) CheckPointUnlock(pointId uint32) bool {
	_, exist := s.UnlockPointMap[pointId]
	return exist
}

func (s *DbScene) GetSceneGroupById(groupId uint32) *DbSceneGroup {
	dbSceneGroup, exist := s.SceneGroupMap[groupId]
	if !exist {
		dbSceneGroup = &DbSceneGroup{
			VariableMap:    make(map[string]int32),
			KillConfigMap:  make(map[uint32]bool),
			GadgetStateMap: make(map[uint32]uint8),
		}
		s.SceneGroupMap[groupId] = dbSceneGroup
	}
	return dbSceneGroup
}

func (g *DbSceneGroup) GetVariableByName(name string) int32 {
	return g.VariableMap[name]
}

func (g *DbSceneGroup) SetVariable(name string, value int32) {
	g.VariableMap[name] = value
}

func (g *DbSceneGroup) CheckVariableExist(name string) bool {
	_, exist := g.VariableMap[name]
	return exist
}

func (g *DbSceneGroup) AddKill(configId uint32) {
	g.KillConfigMap[configId] = true
}

func (g *DbSceneGroup) CheckIsKill(configId uint32) bool {
	_, exist := g.KillConfigMap[configId]
	return exist
}

func (g *DbSceneGroup) RemoveAllKill() {
	g.KillConfigMap = make(map[uint32]bool)
}

func (g *DbSceneGroup) GetGadgetState(configId uint32) uint8 {
	state, exist := g.GadgetStateMap[configId]
	if !exist {
		return constant.GADGET_STATE_DEFAULT
	}
	return state
}

func (g *DbSceneGroup) ChangeGadgetState(configId uint32, state uint8) {
	g.GadgetStateMap[configId] = state
}

func (g *DbSceneGroup) CheckGadgetExist(configId uint32) bool {
	_, exist := g.GadgetStateMap[configId]
	return exist
}

func (s *DbScene) AddVehicle(vehicleId uint32, ownerId uint32, pos *Vector, rot *Vector) {
	s.VehicleMap[vehicleId] = &DbVehicle{
		VehicleId: vehicleId,
		OwnerUid:  ownerId,
		Pos:       pos,
		Rot:       rot,
	}
}
