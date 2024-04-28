// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: spec/network.proto

package envd

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PortState int32

const (
	PortState_OPEN  PortState = 0
	PortState_CLOSE PortState = 1
)

// Enum value maps for PortState.
var (
	PortState_name = map[int32]string{
		0: "OPEN",
		1: "CLOSE",
	}
	PortState_value = map[string]int32{
		"OPEN":  0,
		"CLOSE": 1,
	}
)

func (x PortState) Enum() *PortState {
	p := new(PortState)
	*p = x
	return p
}

func (x PortState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PortState) Descriptor() protoreflect.EnumDescriptor {
	return file_spec_network_proto_enumTypes[0].Descriptor()
}

func (PortState) Type() protoreflect.EnumType {
	return &file_spec_network_proto_enumTypes[0]
}

func (x PortState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PortState.Descriptor instead.
func (PortState) EnumDescriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{0}
}

type ConfigurePortRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port   uint32         `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	State  PortState      `protobuf:"varint,2,opt,name=state,proto3,enum=network.PortState" json:"state,omitempty"`
	Access *AccessControl `protobuf:"bytes,3,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *ConfigurePortRequest) Reset() {
	*x = ConfigurePortRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigurePortRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigurePortRequest) ProtoMessage() {}

func (x *ConfigurePortRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigurePortRequest.ProtoReflect.Descriptor instead.
func (*ConfigurePortRequest) Descriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigurePortRequest) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ConfigurePortRequest) GetState() PortState {
	if x != nil {
		return x.State
	}
	return PortState_OPEN
}

func (x *ConfigurePortRequest) GetAccess() *AccessControl {
	if x != nil {
		return x.Access
	}
	return nil
}

type ListPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Access *AccessControl `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *ListPortsRequest) Reset() {
	*x = ListPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortsRequest) ProtoMessage() {}

func (x *ListPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortsRequest.ProtoReflect.Descriptor instead.
func (*ListPortsRequest) Descriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{1}
}

func (x *ListPortsRequest) GetAccess() *AccessControl {
	if x != nil {
		return x.Access
	}
	return nil
}

type ListPortsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports []*Port `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *ListPortsResponse) Reset() {
	*x = ListPortsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPortsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPortsResponse) ProtoMessage() {}

func (x *ListPortsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spec_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPortsResponse.ProtoReflect.Descriptor instead.
func (*ListPortsResponse) Descriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{2}
}

func (x *ListPortsResponse) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

type WatchPortsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ports  []uint32       `protobuf:"varint,1,rep,packed,name=ports,proto3" json:"ports,omitempty"`
	Access *AccessControl `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *WatchPortsRequest) Reset() {
	*x = WatchPortsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchPortsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchPortsRequest) ProtoMessage() {}

func (x *WatchPortsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchPortsRequest.ProtoReflect.Descriptor instead.
func (*WatchPortsRequest) Descriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{3}
}

func (x *WatchPortsRequest) GetPorts() []uint32 {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *WatchPortsRequest) GetAccess() *AccessControl {
	if x != nil {
		return x.Access
	}
	return nil
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port  uint32    `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	State PortState `protobuf:"varint,2,opt,name=state,proto3,enum=network.PortState" json:"state,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_spec_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_spec_network_proto_rawDescGZIP(), []int{4}
}

func (x *Port) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Port) GetState() PortState {
	if x != nil {
		return x.State
	}
	return PortState_OPEN
}

var File_spec_network_proto protoreflect.FileDescriptor

var file_spec_network_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x16, 0x73,
	0x70, 0x65, 0x63, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x46, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x06, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x22,
	0x5d, 0x0a, 0x11, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x44,
	0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x2a, 0x20, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x32, 0xd0, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x46, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x42, 0x0a, 0x09, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x50, 0x6f, 0x72, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x32, 0x62, 0x2d, 0x64, 0x65, 0x76, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x65,
	0x6e, 0x76, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_network_proto_rawDescOnce sync.Once
	file_spec_network_proto_rawDescData = file_spec_network_proto_rawDesc
)

func file_spec_network_proto_rawDescGZIP() []byte {
	file_spec_network_proto_rawDescOnce.Do(func() {
		file_spec_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_network_proto_rawDescData)
	})
	return file_spec_network_proto_rawDescData
}

var file_spec_network_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_spec_network_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spec_network_proto_goTypes = []interface{}{
	(PortState)(0),               // 0: network.PortState
	(*ConfigurePortRequest)(nil), // 1: network.ConfigurePortRequest
	(*ListPortsRequest)(nil),     // 2: network.ListPortsRequest
	(*ListPortsResponse)(nil),    // 3: network.ListPortsResponse
	(*WatchPortsRequest)(nil),    // 4: network.WatchPortsRequest
	(*Port)(nil),                 // 5: network.Port
	(*AccessControl)(nil),        // 6: permissions.AccessControl
	(*emptypb.Empty)(nil),        // 7: google.protobuf.Empty
}
var file_spec_network_proto_depIdxs = []int32{
	0, // 0: network.ConfigurePortRequest.state:type_name -> network.PortState
	6, // 1: network.ConfigurePortRequest.access:type_name -> permissions.AccessControl
	6, // 2: network.ListPortsRequest.access:type_name -> permissions.AccessControl
	5, // 3: network.ListPortsResponse.ports:type_name -> network.Port
	6, // 4: network.WatchPortsRequest.access:type_name -> permissions.AccessControl
	0, // 5: network.Port.state:type_name -> network.PortState
	1, // 6: network.Network.ConfigurePort:input_type -> network.ConfigurePortRequest
	2, // 7: network.Network.ListPorts:input_type -> network.ListPortsRequest
	4, // 8: network.Network.WatchPorts:input_type -> network.WatchPortsRequest
	7, // 9: network.Network.ConfigurePort:output_type -> google.protobuf.Empty
	3, // 10: network.Network.ListPorts:output_type -> network.ListPortsResponse
	5, // 11: network.Network.WatchPorts:output_type -> network.Port
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_spec_network_proto_init() }
func file_spec_network_proto_init() {
	if File_spec_network_proto != nil {
		return
	}
	file_spec_permissions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spec_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigurePortRequest); i {
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
		file_spec_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortsRequest); i {
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
		file_spec_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPortsResponse); i {
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
		file_spec_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchPortsRequest); i {
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
		file_spec_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
			RawDescriptor: file_spec_network_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_network_proto_goTypes,
		DependencyIndexes: file_spec_network_proto_depIdxs,
		EnumInfos:         file_spec_network_proto_enumTypes,
		MessageInfos:      file_spec_network_proto_msgTypes,
	}.Build()
	File_spec_network_proto = out.File
	file_spec_network_proto_rawDesc = nil
	file_spec_network_proto_goTypes = nil
	file_spec_network_proto_depIdxs = nil
}
