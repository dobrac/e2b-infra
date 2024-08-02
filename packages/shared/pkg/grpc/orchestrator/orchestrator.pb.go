// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: orchestrator.proto

package orchestrator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SandboxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TemplateID         string `protobuf:"bytes,1,opt,name=templateID,proto3" json:"templateID,omitempty"`
	BuildID            string `protobuf:"bytes,2,opt,name=buildID,proto3" json:"buildID,omitempty"`
	KernelVersion      string `protobuf:"bytes,3,opt,name=kernelVersion,proto3" json:"kernelVersion,omitempty"`
	FirecrackerVersion string `protobuf:"bytes,4,opt,name=firecrackerVersion,proto3" json:"firecrackerVersion,omitempty"`
	HugePages          bool   `protobuf:"varint,5,opt,name=hugePages,proto3" json:"hugePages,omitempty"`
	TeamID             string `protobuf:"bytes,6,opt,name=teamID,proto3" json:"teamID,omitempty"`
	// Maximum length of the instance in Hours
	MaxInstanceLength int64             `protobuf:"varint,7,opt,name=maxInstanceLength,proto3" json:"maxInstanceLength,omitempty"`
	Alias             *string           `protobuf:"bytes,8,opt,name=alias,proto3,oneof" json:"alias,omitempty"`
	SandboxID         string            `protobuf:"bytes,9,opt,name=sandboxID,proto3" json:"sandboxID,omitempty"`
	Metadata          map[string]string `protobuf:"bytes,10,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SandboxConfig) Reset() {
	*x = SandboxConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxConfig) ProtoMessage() {}

func (x *SandboxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxConfig.ProtoReflect.Descriptor instead.
func (*SandboxConfig) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{0}
}

func (x *SandboxConfig) GetTemplateID() string {
	if x != nil {
		return x.TemplateID
	}
	return ""
}

func (x *SandboxConfig) GetBuildID() string {
	if x != nil {
		return x.BuildID
	}
	return ""
}

func (x *SandboxConfig) GetKernelVersion() string {
	if x != nil {
		return x.KernelVersion
	}
	return ""
}

func (x *SandboxConfig) GetFirecrackerVersion() string {
	if x != nil {
		return x.FirecrackerVersion
	}
	return ""
}

func (x *SandboxConfig) GetHugePages() bool {
	if x != nil {
		return x.HugePages
	}
	return false
}

func (x *SandboxConfig) GetTeamID() string {
	if x != nil {
		return x.TeamID
	}
	return ""
}

func (x *SandboxConfig) GetMaxInstanceLength() int64 {
	if x != nil {
		return x.MaxInstanceLength
	}
	return 0
}

func (x *SandboxConfig) GetAlias() string {
	if x != nil && x.Alias != nil {
		return *x.Alias
	}
	return ""
}

func (x *SandboxConfig) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

func (x *SandboxConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Data required for creating a new sandbox.
type SandboxCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sandbox   *SandboxConfig         `protobuf:"bytes,1,opt,name=sandbox,proto3" json:"sandbox,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *SandboxCreateRequest) Reset() {
	*x = SandboxCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreateRequest) ProtoMessage() {}

func (x *SandboxCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreateRequest.ProtoReflect.Descriptor instead.
func (*SandboxCreateRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{1}
}

func (x *SandboxCreateRequest) GetSandbox() *SandboxConfig {
	if x != nil {
		return x.Sandbox
	}
	return nil
}

func (x *SandboxCreateRequest) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *SandboxCreateRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Data about the sandbox.
type SandboxCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (x *SandboxCreateResponse) Reset() {
	*x = SandboxCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxCreateResponse) ProtoMessage() {}

func (x *SandboxCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxCreateResponse.ProtoReflect.Descriptor instead.
func (*SandboxCreateResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{2}
}

func (x *SandboxCreateResponse) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

// Data required for creating a new sandbox.
type SandboxUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxID string                 `protobuf:"bytes,1,opt,name=sandboxID,proto3" json:"sandboxID,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *SandboxUpdateRequest) Reset() {
	*x = SandboxUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxUpdateRequest) ProtoMessage() {}

func (x *SandboxUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxUpdateRequest.ProtoReflect.Descriptor instead.
func (*SandboxUpdateRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{3}
}

func (x *SandboxUpdateRequest) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

func (x *SandboxUpdateRequest) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Data required for action on a specified sandbox.
type SandboxRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SandboxID string `protobuf:"bytes,1,opt,name=sandboxID,proto3" json:"sandboxID,omitempty"`
}

func (x *SandboxRequest) Reset() {
	*x = SandboxRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxRequest) ProtoMessage() {}

func (x *SandboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxRequest.ProtoReflect.Descriptor instead.
func (*SandboxRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{4}
}

func (x *SandboxRequest) GetSandboxID() string {
	if x != nil {
		return x.SandboxID
	}
	return ""
}

type RunningSandbox struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config    *SandboxConfig         `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	ClientID  string                 `protobuf:"bytes,2,opt,name=clientID,proto3" json:"clientID,omitempty"`
	StartTime *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime   *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *RunningSandbox) Reset() {
	*x = RunningSandbox{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningSandbox) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningSandbox) ProtoMessage() {}

func (x *RunningSandbox) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningSandbox.ProtoReflect.Descriptor instead.
func (*RunningSandbox) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{5}
}

func (x *RunningSandbox) GetConfig() *SandboxConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *RunningSandbox) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *RunningSandbox) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *RunningSandbox) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

// Data returned after listing all the sandboxes.
type SandboxListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sandboxes []*RunningSandbox `protobuf:"bytes,1,rep,name=sandboxes,proto3" json:"sandboxes,omitempty"`
}

func (x *SandboxListResponse) Reset() {
	*x = SandboxListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SandboxListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SandboxListResponse) ProtoMessage() {}

func (x *SandboxListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SandboxListResponse.ProtoReflect.Descriptor instead.
func (*SandboxListResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_rawDescGZIP(), []int{6}
}

func (x *SandboxListResponse) GetSandboxes() []*RunningSandbox {
	if x != nil {
		return x.Sandboxes
	}
	return nil
}

var File_orchestrator_proto protoreflect.FileDescriptor

var file_orchestrator_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x0d, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x44, 0x12, 0x24,
	0x0a, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x65, 0x63, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x66, 0x69, 0x72, 0x65, 0x63, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x75, 0x67, 0x65, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x75, 0x67, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x61, 0x6d, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61,
	0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x44,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49,
	0x44, 0x12, 0x38, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x22, 0xb0, 0x01, 0x0a, 0x14, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x73,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x73, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x33, 0x0a, 0x15, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x22, 0x6a, 0x0a, 0x14, 0x53, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x49, 0x44,
	0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x61, 0x6e,
	0x64, 0x62, 0x6f, 0x78, 0x49, 0x44, 0x22, 0xc4, 0x01, 0x0a, 0x0e, 0x52, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x38, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x44, 0x0a,
	0x13, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x09, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x78, 0x65, 0x73, 0x32, 0xe4, 0x01, 0x0a, 0x07, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x12,
	0x37, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x53, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x34, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x14, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x0f, 0x2e, 0x53, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x2f, 0x5a, 0x2d, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x65, 0x32, 0x62, 0x2d, 0x64, 0x65, 0x76, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x2f, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_orchestrator_proto_rawDescOnce sync.Once
	file_orchestrator_proto_rawDescData = file_orchestrator_proto_rawDesc
)

func file_orchestrator_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_proto_rawDescData)
	})
	return file_orchestrator_proto_rawDescData
}

var file_orchestrator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_orchestrator_proto_goTypes = []interface{}{
	(*SandboxConfig)(nil),         // 0: SandboxConfig
	(*SandboxCreateRequest)(nil),  // 1: SandboxCreateRequest
	(*SandboxCreateResponse)(nil), // 2: SandboxCreateResponse
	(*SandboxUpdateRequest)(nil),  // 3: SandboxUpdateRequest
	(*SandboxRequest)(nil),        // 4: SandboxRequest
	(*RunningSandbox)(nil),        // 5: RunningSandbox
	(*SandboxListResponse)(nil),   // 6: SandboxListResponse
	nil,                           // 7: SandboxConfig.MetadataEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_orchestrator_proto_depIdxs = []int32{
	7,  // 0: SandboxConfig.metadata:type_name -> SandboxConfig.MetadataEntry
	0,  // 1: SandboxCreateRequest.sandbox:type_name -> SandboxConfig
	8,  // 2: SandboxCreateRequest.startTime:type_name -> google.protobuf.Timestamp
	8,  // 3: SandboxCreateRequest.endTime:type_name -> google.protobuf.Timestamp
	8,  // 4: SandboxUpdateRequest.endTime:type_name -> google.protobuf.Timestamp
	0,  // 5: RunningSandbox.config:type_name -> SandboxConfig
	8,  // 6: RunningSandbox.startTime:type_name -> google.protobuf.Timestamp
	8,  // 7: RunningSandbox.endTime:type_name -> google.protobuf.Timestamp
	5,  // 8: SandboxListResponse.sandboxes:type_name -> RunningSandbox
	1,  // 9: Sandbox.Create:input_type -> SandboxCreateRequest
	3,  // 10: Sandbox.Update:input_type -> SandboxUpdateRequest
	9,  // 11: Sandbox.List:input_type -> google.protobuf.Empty
	4,  // 12: Sandbox.Delete:input_type -> SandboxRequest
	2,  // 13: Sandbox.Create:output_type -> SandboxCreateResponse
	9,  // 14: Sandbox.Update:output_type -> google.protobuf.Empty
	6,  // 15: Sandbox.List:output_type -> SandboxListResponse
	9,  // 16: Sandbox.Delete:output_type -> google.protobuf.Empty
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_init() }
func file_orchestrator_proto_init() {
	if File_orchestrator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestrator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxConfig); i {
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
		file_orchestrator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxCreateRequest); i {
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
		file_orchestrator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxCreateResponse); i {
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
		file_orchestrator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxUpdateRequest); i {
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
		file_orchestrator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxRequest); i {
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
		file_orchestrator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunningSandbox); i {
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
		file_orchestrator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SandboxListResponse); i {
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
	file_orchestrator_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orchestrator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_depIdxs,
		MessageInfos:      file_orchestrator_proto_msgTypes,
	}.Build()
	File_orchestrator_proto = out.File
	file_orchestrator_proto_rawDesc = nil
	file_orchestrator_proto_goTypes = nil
	file_orchestrator_proto_depIdxs = nil
}
