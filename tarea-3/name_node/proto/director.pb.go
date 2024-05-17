// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: proto/director.proto

package proto

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg      string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	ClientIp string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_director_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_director_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_director_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Message) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res string `protobuf:"bytes,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_director_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_director_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_director_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetRes() string {
	if x != nil {
		return x.Res
	}
	return ""
}

type Decision struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option   string `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	ClientIp string `protobuf:"bytes,2,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	Floor    string `protobuf:"bytes,3,opt,name=floor,proto3" json:"floor,omitempty"`
}

func (x *Decision) Reset() {
	*x = Decision{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_director_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Decision) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Decision) ProtoMessage() {}

func (x *Decision) ProtoReflect() protoreflect.Message {
	mi := &file_proto_director_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Decision.ProtoReflect.Descriptor instead.
func (*Decision) Descriptor() ([]byte, []int) {
	return file_proto_director_proto_rawDescGZIP(), []int{2}
}

func (x *Decision) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

func (x *Decision) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *Decision) GetFloor() string {
	if x != nil {
		return x.Floor
	}
	return ""
}

var File_proto_director_proto protoreflect.FileDescriptor

var file_proto_director_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x22, 0x38, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x22, 0x1c, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f,
	0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x32,
	0xb4, 0x02, 0x0a, 0x15, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x11, 0x2e, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x50, 0x69, 0x63, 0x6b, 0x47, 0x75, 0x6e, 0x12, 0x11, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0b, 0x50, 0x69, 0x63, 0x6b, 0x48, 0x61,
	0x6c, 0x6c, 0x77, 0x61, 0x79, 0x12, 0x11, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x0a, 0x42, 0x6f, 0x73, 0x73, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x12, 0x11, 0x2e, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x14, 0x41, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x11, 0x2e,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x1a, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x43,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x1a, 0x12, 0x2e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_director_proto_rawDescOnce sync.Once
	file_proto_director_proto_rawDescData = file_proto_director_proto_rawDesc
)

func file_proto_director_proto_rawDescGZIP() []byte {
	file_proto_director_proto_rawDescOnce.Do(func() {
		file_proto_director_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_director_proto_rawDescData)
	})
	return file_proto_director_proto_rawDescData
}

var file_proto_director_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_director_proto_goTypes = []interface{}{
	(*Message)(nil),  // 0: director.Message
	(*Response)(nil), // 1: director.Response
	(*Decision)(nil), // 2: director.Decision
}
var file_proto_director_proto_depIdxs = []int32{
	0, // 0: director.DirectorCommunication.ConfirmReady:input_type -> director.Message
	0, // 1: director.DirectorCommunication.PickGun:input_type -> director.Message
	0, // 2: director.DirectorCommunication.PickHallway:input_type -> director.Message
	0, // 3: director.DirectorCommunication.BossBattle:input_type -> director.Message
	0, // 4: director.DirectorCommunication.AskForAccountBalance:input_type -> director.Message
	2, // 5: director.NodesCommunication.NotifyDecision:input_type -> director.Decision
	1, // 6: director.DirectorCommunication.ConfirmReady:output_type -> director.Response
	1, // 7: director.DirectorCommunication.PickGun:output_type -> director.Response
	1, // 8: director.DirectorCommunication.PickHallway:output_type -> director.Response
	1, // 9: director.DirectorCommunication.BossBattle:output_type -> director.Response
	1, // 10: director.DirectorCommunication.AskForAccountBalance:output_type -> director.Response
	1, // 11: director.NodesCommunication.NotifyDecision:output_type -> director.Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_director_proto_init() }
func file_proto_director_proto_init() {
	if File_proto_director_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_director_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_director_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_director_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Decision); i {
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
			RawDescriptor: file_proto_director_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_director_proto_goTypes,
		DependencyIndexes: file_proto_director_proto_depIdxs,
		MessageInfos:      file_proto_director_proto_msgTypes,
	}.Build()
	File_proto_director_proto = out.File
	file_proto_director_proto_rawDesc = nil
	file_proto_director_proto_goTypes = nil
	file_proto_director_proto_depIdxs = nil
}
