// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: calculator/v1/calculator.proto

package calculatorv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CalculateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Operand1      float64                `protobuf:"fixed64,1,opt,name=operand1,proto3" json:"operand1,omitempty"`
	Operand2      float64                `protobuf:"fixed64,2,opt,name=operand2,proto3" json:"operand2,omitempty"`
	Operator      string                 `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculateRequest) Reset() {
	*x = CalculateRequest{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateRequest) ProtoMessage() {}

func (x *CalculateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateRequest.ProtoReflect.Descriptor instead.
func (*CalculateRequest) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *CalculateRequest) GetOperand1() float64 {
	if x != nil {
		return x.Operand1
	}
	return 0
}

func (x *CalculateRequest) GetOperand2() float64 {
	if x != nil {
		return x.Operand2
	}
	return 0
}

func (x *CalculateRequest) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type CalculateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        float64                `protobuf:"fixed64,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CalculateResponse) Reset() {
	*x = CalculateResponse{}
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CalculateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculateResponse) ProtoMessage() {}

func (x *CalculateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_v1_calculator_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculateResponse.ProtoReflect.Descriptor instead.
func (*CalculateResponse) Descriptor() ([]byte, []int) {
	return file_calculator_v1_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *CalculateResponse) GetResult() float64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_v1_calculator_proto protoreflect.FileDescriptor

const file_calculator_v1_calculator_proto_rawDesc = "" +
	"\n" +
	"\x1ecalculator/v1/calculator.proto\x12\rcalculator.v1\"f\n" +
	"\x10CalculateRequest\x12\x1a\n" +
	"\boperand1\x18\x01 \x01(\x01R\boperand1\x12\x1a\n" +
	"\boperand2\x18\x02 \x01(\x01R\boperand2\x12\x1a\n" +
	"\boperator\x18\x03 \x01(\tR\boperator\"+\n" +
	"\x11CalculateResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\x01R\x06result2c\n" +
	"\x11CalculatorService\x12N\n" +
	"\tCalculate\x12\x1f.calculator.v1.CalculateRequest\x1a .calculator.v1.CalculateResponseB`Z^github.com/WUZHUHE38/calculator-fullstack/calculator-backend/gen/go/calculator/v1;calculatorv1b\x06proto3"

var (
	file_calculator_v1_calculator_proto_rawDescOnce sync.Once
	file_calculator_v1_calculator_proto_rawDescData []byte
)

func file_calculator_v1_calculator_proto_rawDescGZIP() []byte {
	file_calculator_v1_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_v1_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)))
	})
	return file_calculator_v1_calculator_proto_rawDescData
}

var file_calculator_v1_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculator_v1_calculator_proto_goTypes = []any{
	(*CalculateRequest)(nil),  // 0: calculator.v1.CalculateRequest
	(*CalculateResponse)(nil), // 1: calculator.v1.CalculateResponse
}
var file_calculator_v1_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.v1.CalculatorService.Calculate:input_type -> calculator.v1.CalculateRequest
	1, // 1: calculator.v1.CalculatorService.Calculate:output_type -> calculator.v1.CalculateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_v1_calculator_proto_init() }
func file_calculator_v1_calculator_proto_init() {
	if File_calculator_v1_calculator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_calculator_v1_calculator_proto_rawDesc), len(file_calculator_v1_calculator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_v1_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_v1_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_v1_calculator_proto_msgTypes,
	}.Build()
	File_calculator_v1_calculator_proto = out.File
	file_calculator_v1_calculator_proto_goTypes = nil
	file_calculator_v1_calculator_proto_depIdxs = nil
}
