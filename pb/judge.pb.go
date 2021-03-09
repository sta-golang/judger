// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: judge.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//接口请求入参
type JudgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceCode []byte `protobuf:"bytes,1,opt,name=sourceCode,proto3" json:"sourceCode,omitempty"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ProblemID  int32  `protobuf:"varint,3,opt,name=problemID,proto3" json:"problemID,omitempty"`
	UserID     int32  `protobuf:"varint,4,opt,name=userID,proto3" json:"userID,omitempty"`
	IsUpdate   bool   `protobuf:"varint,5,opt,name=isUpdate,proto3" json:"isUpdate,omitempty"`
}

func (x *JudgeRequest) Reset() {
	*x = JudgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeRequest) ProtoMessage() {}

func (x *JudgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeRequest.ProtoReflect.Descriptor instead.
func (*JudgeRequest) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{0}
}

func (x *JudgeRequest) GetSourceCode() []byte {
	if x != nil {
		return x.SourceCode
	}
	return nil
}

func (x *JudgeRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *JudgeRequest) GetProblemID() int32 {
	if x != nil {
		return x.ProblemID
	}
	return 0
}

func (x *JudgeRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *JudgeRequest) GetIsUpdate() bool {
	if x != nil {
		return x.IsUpdate
	}
	return false
}

//接口返回出参
type JudgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *JudgeResponse) Reset() {
	*x = JudgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_judge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeResponse) ProtoMessage() {}

func (x *JudgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_judge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JudgeResponse.ProtoReflect.Descriptor instead.
func (*JudgeResponse) Descriptor() ([]byte, []int) {
	return file_judge_proto_rawDescGZIP(), []int{1}
}

func (x *JudgeResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_judge_proto protoreflect.FileDescriptor

var file_judge_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c,
	0x65, 0x6d, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x2b, 0x0a, 0x0d, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x12,
	0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x1a, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4a, 0x75, 0x64, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_judge_proto_rawDescOnce sync.Once
	file_judge_proto_rawDescData = file_judge_proto_rawDesc
)

func file_judge_proto_rawDescGZIP() []byte {
	file_judge_proto_rawDescOnce.Do(func() {
		file_judge_proto_rawDescData = protoimpl.X.CompressGZIP(file_judge_proto_rawDescData)
	})
	return file_judge_proto_rawDescData
}

var file_judge_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_judge_proto_goTypes = []interface{}{
	(*JudgeRequest)(nil),  // 0: pb.JudgeRequest
	(*JudgeResponse)(nil), // 1: pb.JudgeResponse
}
var file_judge_proto_depIdxs = []int32{
	0, // 0: pb.JudgeService.Judge:input_type -> pb.JudgeRequest
	0, // 1: pb.JudgeService.JudgeClientAndServerStream:input_type -> pb.JudgeRequest
	1, // 2: pb.JudgeService.Judge:output_type -> pb.JudgeResponse
	1, // 3: pb.JudgeService.JudgeClientAndServerStream:output_type -> pb.JudgeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_judge_proto_init() }
func file_judge_proto_init() {
	if File_judge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_judge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeRequest); i {
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
		file_judge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JudgeResponse); i {
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
			RawDescriptor: file_judge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_judge_proto_goTypes,
		DependencyIndexes: file_judge_proto_depIdxs,
		MessageInfos:      file_judge_proto_msgTypes,
	}.Build()
	File_judge_proto = out.File
	file_judge_proto_rawDesc = nil
	file_judge_proto_goTypes = nil
	file_judge_proto_depIdxs = nil
}