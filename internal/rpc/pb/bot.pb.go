// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: bot.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReadyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadyRequest) Reset() {
	*x = ReadyRequest{}
	mi := &file_bot_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyRequest) ProtoMessage() {}

func (x *ReadyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyRequest.ProtoReflect.Descriptor instead.
func (*ReadyRequest) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{0}
}

func (x *ReadyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Payload       string                 `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_bot_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bot_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_bot_proto_rawDescGZIP(), []int{1}
}

func (x *PingResponse) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

var File_bot_proto protoreflect.FileDescriptor

const file_bot_proto_rawDesc = "" +
	"\n" +
	"\tbot.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x1e\n" +
	"\fReadyRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"(\n" +
	"\fPingResponse\x12\x18\n" +
	"\apayload\x18\x01 \x01(\tR\apayload2:\n" +
	"\bAcceptor\x12.\n" +
	"\x05Ready\x12\r.ReadyRequest\x1a\x16.google.protobuf.Empty24\n" +
	"\x03Bot\x12-\n" +
	"\x04Ping\x12\x16.google.protobuf.Empty\x1a\r.PingResponseB7Z5github.com/mc-botnet/mc-botnet-server/internal/rpc/pbb\x06proto3"

var (
	file_bot_proto_rawDescOnce sync.Once
	file_bot_proto_rawDescData []byte
)

func file_bot_proto_rawDescGZIP() []byte {
	file_bot_proto_rawDescOnce.Do(func() {
		file_bot_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_bot_proto_rawDesc), len(file_bot_proto_rawDesc)))
	})
	return file_bot_proto_rawDescData
}

var file_bot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bot_proto_goTypes = []any{
	(*ReadyRequest)(nil),  // 0: ReadyRequest
	(*PingResponse)(nil),  // 1: PingResponse
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_bot_proto_depIdxs = []int32{
	0, // 0: Acceptor.Ready:input_type -> ReadyRequest
	2, // 1: Bot.Ping:input_type -> google.protobuf.Empty
	2, // 2: Acceptor.Ready:output_type -> google.protobuf.Empty
	1, // 3: Bot.Ping:output_type -> PingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bot_proto_init() }
func file_bot_proto_init() {
	if File_bot_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_bot_proto_rawDesc), len(file_bot_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bot_proto_goTypes,
		DependencyIndexes: file_bot_proto_depIdxs,
		MessageInfos:      file_bot_proto_msgTypes,
	}.Build()
	File_bot_proto = out.File
	file_bot_proto_goTypes = nil
	file_bot_proto_depIdxs = nil
}
