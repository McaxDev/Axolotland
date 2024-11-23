// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: gameapi.proto

package rpc

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

type BackupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *BackupRequest) Reset() {
	*x = BackupRequest{}
	mi := &file_gameapi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRequest) ProtoMessage() {}

func (x *BackupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRequest.ProtoReflect.Descriptor instead.
func (*BackupRequest) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{0}
}

func (x *BackupRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type BackupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BackupResponse) Reset() {
	*x = BackupResponse{}
	mi := &file_gameapi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupResponse) ProtoMessage() {}

func (x *BackupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupResponse.ProtoReflect.Descriptor instead.
func (*BackupResponse) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{1}
}

func (x *BackupResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CmdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server  string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Command string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Admin   bool   `protobuf:"varint,3,opt,name=admin,proto3" json:"admin,omitempty"`
}

func (x *CmdRequest) Reset() {
	*x = CmdRequest{}
	mi := &file_gameapi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRequest) ProtoMessage() {}

func (x *CmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRequest.ProtoReflect.Descriptor instead.
func (*CmdRequest) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{2}
}

func (x *CmdRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *CmdRequest) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *CmdRequest) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

type CmdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CmdResponse) Reset() {
	*x = CmdResponse{}
	mi := &file_gameapi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdResponse) ProtoMessage() {}

func (x *CmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdResponse.ProtoReflect.Descriptor instead.
func (*CmdResponse) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{3}
}

func (x *CmdResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   string `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Player   string `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"`
	Authcode string `protobuf:"bytes,3,opt,name=authcode,proto3" json:"authcode,omitempty"`
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	mi := &file_gameapi_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{4}
}

func (x *BindRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *BindRequest) GetPlayer() string {
	if x != nil {
		return x.Player
	}
	return ""
}

func (x *BindRequest) GetAuthcode() string {
	if x != nil {
		return x.Authcode
	}
	return ""
}

type BindResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BindResponse) Reset() {
	*x = BindResponse{}
	mi := &file_gameapi_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BindResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindResponse) ProtoMessage() {}

func (x *BindResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gameapi_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindResponse.ProtoReflect.Descriptor instead.
func (*BindResponse) Descriptor() ([]byte, []int) {
	return file_gameapi_proto_rawDescGZIP(), []int{5}
}

func (x *BindResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_gameapi_proto protoreflect.FileDescriptor

var file_gameapi_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49, 0x22, 0x27, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x22, 0x2a, 0x0a, 0x0e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x54, 0x0a,
	0x0a, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x22, 0x29, 0x0a, 0x0b, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59,
	0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x75, 0x74, 0x68, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x32, 0xb8, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49, 0x12,
	0x3e, 0x0a, 0x0b, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x16,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49,
	0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x34, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x6d, 0x64, 0x12, 0x13, 0x2e, 0x47, 0x61, 0x6d,
	0x65, 0x41, 0x50, 0x49, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x69, 0x6e,
	0x64, 0x12, 0x14, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50, 0x49, 0x2e, 0x42, 0x69, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x50,
	0x49, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gameapi_proto_rawDescOnce sync.Once
	file_gameapi_proto_rawDescData = file_gameapi_proto_rawDesc
)

func file_gameapi_proto_rawDescGZIP() []byte {
	file_gameapi_proto_rawDescOnce.Do(func() {
		file_gameapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_gameapi_proto_rawDescData)
	})
	return file_gameapi_proto_rawDescData
}

var file_gameapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gameapi_proto_goTypes = []any{
	(*BackupRequest)(nil),  // 0: GameAPI.BackupRequest
	(*BackupResponse)(nil), // 1: GameAPI.BackupResponse
	(*CmdRequest)(nil),     // 2: GameAPI.CmdRequest
	(*CmdResponse)(nil),    // 3: GameAPI.CmdResponse
	(*BindRequest)(nil),    // 4: GameAPI.BindRequest
	(*BindResponse)(nil),   // 5: GameAPI.BindResponse
}
var file_gameapi_proto_depIdxs = []int32{
	0, // 0: GameAPI.GameAPI.WorldBackup:input_type -> GameAPI.BackupRequest
	2, // 1: GameAPI.GameAPI.SendCmd:input_type -> GameAPI.CmdRequest
	4, // 2: GameAPI.GameAPI.GameBind:input_type -> GameAPI.BindRequest
	1, // 3: GameAPI.GameAPI.WorldBackup:output_type -> GameAPI.BackupResponse
	3, // 4: GameAPI.GameAPI.SendCmd:output_type -> GameAPI.CmdResponse
	5, // 5: GameAPI.GameAPI.GameBind:output_type -> GameAPI.BindResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gameapi_proto_init() }
func file_gameapi_proto_init() {
	if File_gameapi_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gameapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gameapi_proto_goTypes,
		DependencyIndexes: file_gameapi_proto_depIdxs,
		MessageInfos:      file_gameapi_proto_msgTypes,
	}.Build()
	File_gameapi_proto = out.File
	file_gameapi_proto_rawDesc = nil
	file_gameapi_proto_goTypes = nil
	file_gameapi_proto_depIdxs = nil
}
