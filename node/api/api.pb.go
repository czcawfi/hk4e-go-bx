// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.0
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NullMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NullMsg) Reset() {
	*x = NullMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NullMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NullMsg) ProtoMessage() {}

func (x *NullMsg) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NullMsg.ProtoReflect.Descriptor instead.
func (*NullMsg) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type GetServerAppIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType string `protobuf:"bytes,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
}

func (x *GetServerAppIdReq) Reset() {
	*x = GetServerAppIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerAppIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerAppIdReq) ProtoMessage() {}

func (x *GetServerAppIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerAppIdReq.ProtoReflect.Descriptor instead.
func (*GetServerAppIdReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetServerAppIdReq) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

type GetServerAppIdRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetServerAppIdRsp) Reset() {
	*x = GetServerAppIdRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerAppIdRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerAppIdRsp) ProtoMessage() {}

func (x *GetServerAppIdRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerAppIdRsp.ProtoReflect.Descriptor instead.
func (*GetServerAppIdRsp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *GetServerAppIdRsp) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type RegisterServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType     string          `protobuf:"bytes,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	GateServerAddr *GateServerAddr `protobuf:"bytes,2,opt,name=gate_server_addr,json=gateServerAddr,proto3" json:"gate_server_addr,omitempty"`
	Version        string          `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *RegisterServerReq) Reset() {
	*x = RegisterServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServerReq) ProtoMessage() {}

func (x *RegisterServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServerReq.ProtoReflect.Descriptor instead.
func (*RegisterServerReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterServerReq) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *RegisterServerReq) GetGateServerAddr() *GateServerAddr {
	if x != nil {
		return x.GateServerAddr
	}
	return nil
}

func (x *RegisterServerReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RegisterServerRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	GsId  uint32 `protobuf:"varint,2,opt,name=gs_id,json=gsId,proto3" json:"gs_id,omitempty"`
}

func (x *RegisterServerRsp) Reset() {
	*x = RegisterServerRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServerRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServerRsp) ProtoMessage() {}

func (x *RegisterServerRsp) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServerRsp.ProtoReflect.Descriptor instead.
func (*RegisterServerRsp) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterServerRsp) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RegisterServerRsp) GetGsId() uint32 {
	if x != nil {
		return x.GsId
	}
	return 0
}

type CancelServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType string `protobuf:"bytes,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	AppId      string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *CancelServerReq) Reset() {
	*x = CancelServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelServerReq) ProtoMessage() {}

func (x *CancelServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelServerReq.ProtoReflect.Descriptor instead.
func (*CancelServerReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *CancelServerReq) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *CancelServerReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type KeepaliveServerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerType string `protobuf:"bytes,1,opt,name=server_type,json=serverType,proto3" json:"server_type,omitempty"`
	AppId      string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *KeepaliveServerReq) Reset() {
	*x = KeepaliveServerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeepaliveServerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeepaliveServerReq) ProtoMessage() {}

func (x *KeepaliveServerReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeepaliveServerReq.ProtoReflect.Descriptor instead.
func (*KeepaliveServerReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *KeepaliveServerReq) GetServerType() string {
	if x != nil {
		return x.ServerType
	}
	return ""
}

func (x *KeepaliveServerReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetGateServerAddrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *GetGateServerAddrReq) Reset() {
	*x = GetGateServerAddrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGateServerAddrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGateServerAddrReq) ProtoMessage() {}

func (x *GetGateServerAddrReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGateServerAddrReq.ProtoReflect.Descriptor instead.
func (*GetGateServerAddrReq) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *GetGateServerAddrReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type RegionEc2B struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RegionEc2B) Reset() {
	*x = RegionEc2B{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionEc2B) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionEc2B) ProtoMessage() {}

func (x *RegionEc2B) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionEc2B.ProtoReflect.Descriptor instead.
func (*RegionEc2B) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

func (x *RegionEc2B) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type GateServerAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddr string `protobuf:"bytes,1,opt,name=ip_addr,json=ipAddr,proto3" json:"ip_addr,omitempty"`
	Port   uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *GateServerAddr) Reset() {
	*x = GateServerAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GateServerAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GateServerAddr) ProtoMessage() {}

func (x *GateServerAddr) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GateServerAddr.ProtoReflect.Descriptor instead.
func (*GateServerAddr) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *GateServerAddr) GetIpAddr() string {
	if x != nil {
		return x.IpAddr
	}
	return ""
}

func (x *GateServerAddr) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x22, 0x09, 0x0a, 0x07, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67,
	0x22, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x2a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x49, 0x64, 0x52, 0x73, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70,
	0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x10, 0x67, 0x61, 0x74,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x52, 0x0e, 0x67,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x67, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x67, 0x73, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x12, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x22, 0x30, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x20, 0x0a, 0x0a, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x63, 0x32,
	0x62, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x0e, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x32, 0xba, 0x03, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x0f, 0x4b, 0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4b,
	0x65, 0x65, 0x70, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x75, 0x6c,
	0x6c, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x70, 0x70, 0x49, 0x64, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x45, 0x63, 0x32, 0x62, 0x12, 0x11, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4e, 0x75, 0x6c, 0x6c, 0x4d, 0x73, 0x67, 0x1a, 0x14, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x45, 0x63, 0x32, 0x62, 0x22, 0x00,
	0x12, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x68, 0x6b, 0x34, 0x65, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_proto_goTypes = []interface{}{
	(*NullMsg)(nil),              // 0: node.api.NullMsg
	(*GetServerAppIdReq)(nil),    // 1: node.api.GetServerAppIdReq
	(*GetServerAppIdRsp)(nil),    // 2: node.api.GetServerAppIdRsp
	(*RegisterServerReq)(nil),    // 3: node.api.RegisterServerReq
	(*RegisterServerRsp)(nil),    // 4: node.api.RegisterServerRsp
	(*CancelServerReq)(nil),      // 5: node.api.CancelServerReq
	(*KeepaliveServerReq)(nil),   // 6: node.api.KeepaliveServerReq
	(*GetGateServerAddrReq)(nil), // 7: node.api.GetGateServerAddrReq
	(*RegionEc2B)(nil),           // 8: node.api.RegionEc2b
	(*GateServerAddr)(nil),       // 9: node.api.GateServerAddr
}
var file_api_proto_depIdxs = []int32{
	9, // 0: node.api.RegisterServerReq.gate_server_addr:type_name -> node.api.GateServerAddr
	3, // 1: node.api.Discovery.RegisterServer:input_type -> node.api.RegisterServerReq
	5, // 2: node.api.Discovery.CancelServer:input_type -> node.api.CancelServerReq
	6, // 3: node.api.Discovery.KeepaliveServer:input_type -> node.api.KeepaliveServerReq
	1, // 4: node.api.Discovery.GetServerAppId:input_type -> node.api.GetServerAppIdReq
	0, // 5: node.api.Discovery.GetRegionEc2b:input_type -> node.api.NullMsg
	7, // 6: node.api.Discovery.GetGateServerAddr:input_type -> node.api.GetGateServerAddrReq
	4, // 7: node.api.Discovery.RegisterServer:output_type -> node.api.RegisterServerRsp
	0, // 8: node.api.Discovery.CancelServer:output_type -> node.api.NullMsg
	0, // 9: node.api.Discovery.KeepaliveServer:output_type -> node.api.NullMsg
	2, // 10: node.api.Discovery.GetServerAppId:output_type -> node.api.GetServerAppIdRsp
	8, // 11: node.api.Discovery.GetRegionEc2b:output_type -> node.api.RegionEc2b
	9, // 12: node.api.Discovery.GetGateServerAddr:output_type -> node.api.GateServerAddr
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NullMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerAppIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerAppIdRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServerRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelServerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeepaliveServerReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGateServerAddrReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionEc2B); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GateServerAddr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
