// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: rpc/bisky.proto

package bisky

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

type JudgeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncodedCode string      `protobuf:"bytes,1,opt,name=encoded_code,json=encodedCode,proto3" json:"encoded_code,omitempty"`
	Language    string      `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Problem     string      `protobuf:"bytes,3,opt,name=problem,proto3" json:"problem,omitempty"`
	OutputType  string      `protobuf:"bytes,4,opt,name=output_type,json=outputType,proto3" json:"output_type,omitempty"`
	Code        string      `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	TestCases   []*TestCase `protobuf:"bytes,6,rep,name=test_cases,json=testCases,proto3" json:"test_cases,omitempty"`
}

func (x *JudgeRequest) Reset() {
	*x = JudgeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bisky_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeRequest) ProtoMessage() {}

func (x *JudgeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bisky_proto_msgTypes[0]
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
	return file_rpc_bisky_proto_rawDescGZIP(), []int{0}
}

func (x *JudgeRequest) GetEncodedCode() string {
	if x != nil {
		return x.EncodedCode
	}
	return ""
}

func (x *JudgeRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *JudgeRequest) GetProblem() string {
	if x != nil {
		return x.Problem
	}
	return ""
}

func (x *JudgeRequest) GetOutputType() string {
	if x != nil {
		return x.OutputType
	}
	return ""
}

func (x *JudgeRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *JudgeRequest) GetTestCases() []*TestCase {
	if x != nil {
		return x.TestCases
	}
	return nil
}

type TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input          string   `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	ExpectedOutput []string `protobuf:"bytes,2,rep,name=expected_output,json=expectedOutput,proto3" json:"expected_output,omitempty"`
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bisky_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bisky_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_rpc_bisky_proto_rawDescGZIP(), []int{1}
}

func (x *TestCase) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestCase) GetExpectedOutput() []string {
	if x != nil {
		return x.ExpectedOutput
	}
	return nil
}

type TestCaseResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stdout         string `protobuf:"bytes,1,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Passed         bool   `protobuf:"varint,2,opt,name=passed,proto3" json:"passed,omitempty"`
	Elapsed        string `protobuf:"bytes,3,opt,name=elapsed,proto3" json:"elapsed,omitempty"`
	Input          string `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`                                         // test input
	ExpectedOutput string `protobuf:"bytes,5,opt,name=expected_output,json=expectedOutput,proto3" json:"expected_output,omitempty"` // test expected out
}

func (x *TestCaseResult) Reset() {
	*x = TestCaseResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bisky_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCaseResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCaseResult) ProtoMessage() {}

func (x *TestCaseResult) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bisky_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCaseResult.ProtoReflect.Descriptor instead.
func (*TestCaseResult) Descriptor() ([]byte, []int) {
	return file_rpc_bisky_proto_rawDescGZIP(), []int{2}
}

func (x *TestCaseResult) GetStdout() string {
	if x != nil {
		return x.Stdout
	}
	return ""
}

func (x *TestCaseResult) GetPassed() bool {
	if x != nil {
		return x.Passed
	}
	return false
}

func (x *TestCaseResult) GetElapsed() string {
	if x != nil {
		return x.Elapsed
	}
	return ""
}

func (x *TestCaseResult) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestCaseResult) GetExpectedOutput() string {
	if x != nil {
		return x.ExpectedOutput
	}
	return ""
}

type JudgeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*TestCaseResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *JudgeResponse) Reset() {
	*x = JudgeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_bisky_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JudgeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JudgeResponse) ProtoMessage() {}

func (x *JudgeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_bisky_proto_msgTypes[3]
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
	return file_rpc_bisky_proto_rawDescGZIP(), []int{3}
}

func (x *JudgeResponse) GetResults() []*TestCaseResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_rpc_bisky_proto protoreflect.FileDescriptor

var file_rpc_bisky_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x69, 0x73, 0x6b, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x72,
	0x70, 0x63, 0x22, 0xd6, 0x01, 0x0a, 0x0c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x38, 0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x09, 0x74, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x08, 0x54,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64,
	0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x64, 0x6f, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x70, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x61,
	0x70, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x61, 0x70,
	0x73, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x22, 0x4a, 0x0a, 0x0d, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x4f,
	0x0a, 0x05, 0x42, 0x69, 0x73, 0x6b, 0x79, 0x12, 0x46, 0x0a, 0x05, 0x4a, 0x75, 0x64, 0x67, 0x65,
	0x12, 0x1d, 0x2e, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x17, 0x5a, 0x15, 0x61, 0x6c, 0x67, 0x6f, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x62, 0x69, 0x73, 0x6b, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_bisky_proto_rawDescOnce sync.Once
	file_rpc_bisky_proto_rawDescData = file_rpc_bisky_proto_rawDesc
)

func file_rpc_bisky_proto_rawDescGZIP() []byte {
	file_rpc_bisky_proto_rawDescOnce.Do(func() {
		file_rpc_bisky_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_bisky_proto_rawDescData)
	})
	return file_rpc_bisky_proto_rawDescData
}

var file_rpc_bisky_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rpc_bisky_proto_goTypes = []interface{}{
	(*JudgeRequest)(nil),   // 0: algotrainer.rpc.JudgeRequest
	(*TestCase)(nil),       // 1: algotrainer.rpc.TestCase
	(*TestCaseResult)(nil), // 2: algotrainer.rpc.TestCaseResult
	(*JudgeResponse)(nil),  // 3: algotrainer.rpc.JudgeResponse
}
var file_rpc_bisky_proto_depIdxs = []int32{
	1, // 0: algotrainer.rpc.JudgeRequest.test_cases:type_name -> algotrainer.rpc.TestCase
	2, // 1: algotrainer.rpc.JudgeResponse.results:type_name -> algotrainer.rpc.TestCaseResult
	0, // 2: algotrainer.rpc.Bisky.Judge:input_type -> algotrainer.rpc.JudgeRequest
	3, // 3: algotrainer.rpc.Bisky.Judge:output_type -> algotrainer.rpc.JudgeResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_bisky_proto_init() }
func file_rpc_bisky_proto_init() {
	if File_rpc_bisky_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_bisky_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_bisky_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCase); i {
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
		file_rpc_bisky_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCaseResult); i {
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
		file_rpc_bisky_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_rpc_bisky_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_bisky_proto_goTypes,
		DependencyIndexes: file_rpc_bisky_proto_depIdxs,
		MessageInfos:      file_rpc_bisky_proto_msgTypes,
	}.Build()
	File_rpc_bisky_proto = out.File
	file_rpc_bisky_proto_rawDesc = nil
	file_rpc_bisky_proto_goTypes = nil
	file_rpc_bisky_proto_depIdxs = nil
}
