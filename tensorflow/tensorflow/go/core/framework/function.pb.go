// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/function.proto

package framework

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A library is a set of named functions.
type FunctionDefLibrary struct {
	Function []*FunctionDef `protobuf:"bytes,1,rep,name=function" json:"function,omitempty"`
	Gradient []*GradientDef `protobuf:"bytes,2,rep,name=gradient" json:"gradient,omitempty"`
}

func (m *FunctionDefLibrary) Reset()                    { *m = FunctionDefLibrary{} }
func (m *FunctionDefLibrary) String() string            { return proto.CompactTextString(m) }
func (*FunctionDefLibrary) ProtoMessage()               {}
func (*FunctionDefLibrary) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *FunctionDefLibrary) GetFunction() []*FunctionDef {
	if m != nil {
		return m.Function
	}
	return nil
}

func (m *FunctionDefLibrary) GetGradient() []*GradientDef {
	if m != nil {
		return m.Gradient
	}
	return nil
}

// A function can be instantiated when the runtime can bind every attr
// with a value. When a GraphDef has a call to a function, it must
// have binding for every attr defined in the signature.
//
// TODO(zhifengc):
//   * device spec, etc.
type FunctionDef struct {
	// The definition of the function's name, arguments, return values,
	// attrs etc.
	Signature *OpDef `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Attributes specific to this function definition.
	Attr map[string]*AttrValue `protobuf:"bytes,5,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// By convention, "op" in node_def is resolved by consulting with a
	// user-defined library first. If not resolved, "func" is assumed to
	// be a builtin op.
	NodeDef []*NodeDef `protobuf:"bytes,3,rep,name=node_def,json=nodeDef" json:"node_def,omitempty"`
	// A mapping from the output arg names from `signature` to the
	// outputs from `node_def` that should be returned by the function.
	Ret map[string]string `protobuf:"bytes,4,rep,name=ret" json:"ret,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FunctionDef) Reset()                    { *m = FunctionDef{} }
func (m *FunctionDef) String() string            { return proto.CompactTextString(m) }
func (*FunctionDef) ProtoMessage()               {}
func (*FunctionDef) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *FunctionDef) GetSignature() *OpDef {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *FunctionDef) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *FunctionDef) GetNodeDef() []*NodeDef {
	if m != nil {
		return m.NodeDef
	}
	return nil
}

func (m *FunctionDef) GetRet() map[string]string {
	if m != nil {
		return m.Ret
	}
	return nil
}

// GradientDef defines the gradient function of a function defined in
// a function library.
//
// A gradient function g (specified by gradient_func) for a function f
// (specified by function_name) must follow the following:
//
// The function 'f' must be a numerical function which takes N inputs
// and produces M outputs. Its gradient function 'g', which is a
// function taking N + M inputs and produces N outputs.
//
// I.e. if we have
//    (y1, y2, ..., y_M) = f(x1, x2, ..., x_N),
// then, g is
//    (dL/dx1, dL/dx2, ..., dL/dx_N) = g(x1, x2, ..., x_N,
//                                      dL/dy1, dL/dy2, ..., dL/dy_M),
// where L is a scalar-value function of (x1, x2, ..., xN) (e.g., the
// loss function). dL/dx_i is the partial derivative of L with respect
// to x_i.
type GradientDef struct {
	FunctionName string `protobuf:"bytes,1,opt,name=function_name,json=functionName" json:"function_name,omitempty"`
	GradientFunc string `protobuf:"bytes,2,opt,name=gradient_func,json=gradientFunc" json:"gradient_func,omitempty"`
}

func (m *GradientDef) Reset()                    { *m = GradientDef{} }
func (m *GradientDef) String() string            { return proto.CompactTextString(m) }
func (*GradientDef) ProtoMessage()               {}
func (*GradientDef) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *GradientDef) GetFunctionName() string {
	if m != nil {
		return m.FunctionName
	}
	return ""
}

func (m *GradientDef) GetGradientFunc() string {
	if m != nil {
		return m.GradientFunc
	}
	return ""
}

func init() {
	proto.RegisterType((*FunctionDefLibrary)(nil), "tensorflow.FunctionDefLibrary")
	proto.RegisterType((*FunctionDef)(nil), "tensorflow.FunctionDef")
	proto.RegisterType((*GradientDef)(nil), "tensorflow.GradientDef")
}

func init() { proto.RegisterFile("github.com/tensorflow/tensorflow/tensorflow/go/core/framework/function.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xb1, 0x9d, 0x6c, 0xf1, 0x73, 0x36, 0x36, 0x6d, 0x63, 0xc2, 0xa7, 0x2c, 0x83, 0x11,
	0x36, 0xb0, 0x21, 0x61, 0xa5, 0xf4, 0xd6, 0xd0, 0x1f, 0x97, 0x92, 0x06, 0x1f, 0xda, 0xa3, 0x51,
	0x62, 0x39, 0x98, 0x24, 0x52, 0x50, 0x94, 0x86, 0x5c, 0xfa, 0xbf, 0xf6, 0xbf, 0xe8, 0xb1, 0x48,
	0xb6, 0x62, 0x41, 0xeb, 0xde, 0x84, 0xf4, 0xf9, 0x7e, 0xbf, 0x4f, 0xef, 0x3d, 0x18, 0x48, 0xca,
	0xb6, 0x5c, 0xe4, 0x2b, 0xbe, 0x8f, 0xe7, 0x5c, 0xd0, 0x38, 0x17, 0x64, 0x4d, 0xf7, 0x5c, 0x2c,
	0xe3, 0x7c, 0xc7, 0xe6, 0xb2, 0xe0, 0x2c, 0xda, 0x08, 0x2e, 0x39, 0x82, 0x9a, 0x0c, 0xff, 0x36,
	0xab, 0x88, 0x94, 0x22, 0x7d, 0x20, 0xab, 0x1d, 0x2d, 0x75, 0xe1, 0x3b, 0x09, 0x8c, 0x67, 0x34,
	0xcd, 0x68, 0x5e, 0x91, 0x7f, 0x9a, 0x49, 0xbe, 0xa9, 0xb9, 0xfe, 0x23, 0xa0, 0xab, 0xaa, 0xb6,
	0x0b, 0x9a, 0xdf, 0x14, 0x33, 0x41, 0xc4, 0x01, 0x8d, 0xa0, 0x63, 0x2a, 0xc6, 0x4e, 0xcf, 0x1b,
	0x04, 0xc3, 0x9f, 0x51, 0x6d, 0x18, 0x59, 0x8a, 0xe4, 0x08, 0x2a, 0xd1, 0x42, 0x90, 0xac, 0xa0,
	0x4c, 0x62, 0xf7, 0xb5, 0xe8, 0xba, 0x7a, 0xd3, 0x22, 0x03, 0xf6, 0x9f, 0x5c, 0x08, 0x2c, 0x3b,
	0x14, 0x83, 0xbf, 0x2d, 0x16, 0x8c, 0xc8, 0x9d, 0xa0, 0xd8, 0xe9, 0x39, 0x83, 0x60, 0xf8, 0xd5,
	0x76, 0xb9, 0xdd, 0x28, 0x7d, 0xcd, 0xa0, 0xff, 0xd0, 0x52, 0x6d, 0xc2, 0x6d, 0x9d, 0xf8, 0xab,
	0xa1, 0xcc, 0xe8, 0x5c, 0x4a, 0x71, 0xc9, 0xa4, 0x38, 0x24, 0x1a, 0x47, 0x11, 0x74, 0x4c, 0xc7,
	0xb0, 0xa7, 0xa5, 0xdf, 0x6c, 0xe9, 0x84, 0x67, 0x54, 0x05, 0x7d, 0x64, 0xe5, 0x01, 0x0d, 0xc1,
	0x13, 0x54, 0xe2, 0x96, 0x46, 0x7b, 0x4d, 0x29, 0x09, 0x95, 0x65, 0x88, 0x82, 0xc3, 0x09, 0xf8,
	0xc7, 0x58, 0xf4, 0x05, 0xbc, 0x25, 0x3d, 0xe8, 0x2f, 0xf9, 0x89, 0x3a, 0xa2, 0x7f, 0xd0, 0xd6,
	0xb3, 0xc5, 0xae, 0xfe, 0xe6, 0x0f, 0xdb, 0x54, 0xe9, 0xee, 0xd4, 0x63, 0x52, 0x32, 0x67, 0xee,
	0xa9, 0x13, 0x9e, 0x40, 0xc7, 0x04, 0xbc, 0x61, 0xf7, 0xdd, 0xb6, 0xf3, 0x2d, 0x5d, 0xff, 0x1e,
	0x02, 0xab, 0xf9, 0xe8, 0x37, 0x7c, 0x32, 0x33, 0x4b, 0x19, 0x59, 0xd3, 0xca, 0xa4, 0x6b, 0x2e,
	0x27, 0x64, 0x4d, 0x15, 0x64, 0x66, 0x94, 0xaa, 0x87, 0xca, 0xb5, 0x6b, 0x2e, 0xd5, 0xaf, 0xc7,
	0x31, 0x60, 0x2e, 0x16, 0x76, 0xdd, 0xc7, 0x2d, 0x1b, 0x7f, 0x36, 0x7d, 0x99, 0xaa, 0x3d, 0xdb,
	0x4e, 0x9d, 0x67, 0xc7, 0x99, 0x7d, 0xd0, 0x4b, 0x37, 0x7a, 0x09, 0x00, 0x00, 0xff, 0xff, 0xa1,
	0xb1, 0x76, 0x27, 0x2a, 0x03, 0x00, 0x00,
}
