// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/meta_graph.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import tensorflow10 "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
import tensorflow7 "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
import tensorflow2 "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
import tensorflow3 "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NOTE: This protocol buffer is evolving, and will go through revisions in the
// coming months.
//
// Protocol buffer containing the following which are necessary to restart
// training, run inference. It can be used to serialize/de-serialize memory
// objects necessary for running computation in a graph when crossing the
// process boundary. It can be used for long term storage of graphs,
// cross-language execution of graphs, etc.
//   MetaInfoDef
//   GraphDef
//   SaverDef
//   CollectionDef
//   TensorInfo
//   SignatureDef
type MetaGraphDef struct {
	MetaInfoDef *MetaGraphDef_MetaInfoDef `protobuf:"bytes,1,opt,name=meta_info_def,json=metaInfoDef" json:"meta_info_def,omitempty"`
	// GraphDef.
	GraphDef *tensorflow10.GraphDef `protobuf:"bytes,2,opt,name=graph_def,json=graphDef" json:"graph_def,omitempty"`
	// SaverDef.
	SaverDef *SaverDef `protobuf:"bytes,3,opt,name=saver_def,json=saverDef" json:"saver_def,omitempty"`
	// collection_def: Map from collection name to collections.
	// See CollectionDef section for details.
	CollectionDef map[string]*CollectionDef `protobuf:"bytes,4,rep,name=collection_def,json=collectionDef" json:"collection_def,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// signature_def: Map from user supplied key for a signature to a single
	// SignatureDef.
	SignatureDef map[string]*SignatureDef `protobuf:"bytes,5,rep,name=signature_def,json=signatureDef" json:"signature_def,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Asset file def to be used with the defined graph.
	AssetFileDef []*AssetFileDef `protobuf:"bytes,6,rep,name=asset_file_def,json=assetFileDef" json:"asset_file_def,omitempty"`
}

func (m *MetaGraphDef) Reset()                    { *m = MetaGraphDef{} }
func (m *MetaGraphDef) String() string            { return proto.CompactTextString(m) }
func (*MetaGraphDef) ProtoMessage()               {}
func (*MetaGraphDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *MetaGraphDef) GetMetaInfoDef() *MetaGraphDef_MetaInfoDef {
	if m != nil {
		return m.MetaInfoDef
	}
	return nil
}

func (m *MetaGraphDef) GetGraphDef() *tensorflow10.GraphDef {
	if m != nil {
		return m.GraphDef
	}
	return nil
}

func (m *MetaGraphDef) GetSaverDef() *SaverDef {
	if m != nil {
		return m.SaverDef
	}
	return nil
}

func (m *MetaGraphDef) GetCollectionDef() map[string]*CollectionDef {
	if m != nil {
		return m.CollectionDef
	}
	return nil
}

func (m *MetaGraphDef) GetSignatureDef() map[string]*SignatureDef {
	if m != nil {
		return m.SignatureDef
	}
	return nil
}

func (m *MetaGraphDef) GetAssetFileDef() []*AssetFileDef {
	if m != nil {
		return m.AssetFileDef
	}
	return nil
}

// Meta information regarding the graph to be exported.  To be used by users
// of this protocol buffer to encode information regarding their meta graph.
type MetaGraphDef_MetaInfoDef struct {
	// User specified Version string. Can be the name of the model and revision,
	// steps this model has been trained to, etc.
	MetaGraphVersion string `protobuf:"bytes,1,opt,name=meta_graph_version,json=metaGraphVersion" json:"meta_graph_version,omitempty"`
	// A copy of the OpDefs used by the producer of this graph_def.
	// Descriptions and Ops not used in graph_def are stripped out.
	StrippedOpList *tensorflow7.OpList `protobuf:"bytes,2,opt,name=stripped_op_list,json=strippedOpList" json:"stripped_op_list,omitempty"`
	// A serialized protobuf. Can be the time this meta graph is created, or
	// modified, or name of the model.
	AnyInfo *google_protobuf.Any `protobuf:"bytes,3,opt,name=any_info,json=anyInfo" json:"any_info,omitempty"`
	// User supplied tag(s) on the meta_graph and included graph_def.
	//
	// MetaGraphDefs should be tagged with their capabilities or use-cases.
	// Examples: "train", "serve", "gpu", "tpu", etc.
	// These tags enable loaders to access the MetaGraph(s) appropriate for a
	// specific use-case or runtime environment.
	Tags []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	// The __version__ string of the tensorflow build used to write this graph.
	// This will be populated by the framework, which will overwrite any user
	// supplied value.
	TensorflowVersion string `protobuf:"bytes,5,opt,name=tensorflow_version,json=tensorflowVersion" json:"tensorflow_version,omitempty"`
	// The __git_version__ string of the tensorflow build used to write this
	// graph. This will be populated by the framework, which will overwrite any
	// user supplied value.
	TensorflowGitVersion string `protobuf:"bytes,6,opt,name=tensorflow_git_version,json=tensorflowGitVersion" json:"tensorflow_git_version,omitempty"`
	// A flag to denote whether default-valued attrs have been stripped from
	// the nodes in this graph_def.
	StrippedDefaultAttrs bool `protobuf:"varint,7,opt,name=stripped_default_attrs,json=strippedDefaultAttrs" json:"stripped_default_attrs,omitempty"`
}

func (m *MetaGraphDef_MetaInfoDef) Reset()                    { *m = MetaGraphDef_MetaInfoDef{} }
func (m *MetaGraphDef_MetaInfoDef) String() string            { return proto.CompactTextString(m) }
func (*MetaGraphDef_MetaInfoDef) ProtoMessage()               {}
func (*MetaGraphDef_MetaInfoDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func (m *MetaGraphDef_MetaInfoDef) GetMetaGraphVersion() string {
	if m != nil {
		return m.MetaGraphVersion
	}
	return ""
}

func (m *MetaGraphDef_MetaInfoDef) GetStrippedOpList() *tensorflow7.OpList {
	if m != nil {
		return m.StrippedOpList
	}
	return nil
}

func (m *MetaGraphDef_MetaInfoDef) GetAnyInfo() *google_protobuf.Any {
	if m != nil {
		return m.AnyInfo
	}
	return nil
}

func (m *MetaGraphDef_MetaInfoDef) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MetaGraphDef_MetaInfoDef) GetTensorflowVersion() string {
	if m != nil {
		return m.TensorflowVersion
	}
	return ""
}

func (m *MetaGraphDef_MetaInfoDef) GetTensorflowGitVersion() string {
	if m != nil {
		return m.TensorflowGitVersion
	}
	return ""
}

func (m *MetaGraphDef_MetaInfoDef) GetStrippedDefaultAttrs() bool {
	if m != nil {
		return m.StrippedDefaultAttrs
	}
	return false
}

// CollectionDef should cover most collections.
// To add a user-defined collection, do one of the following:
// 1. For simple data types, such as string, int, float:
//      tf.add_to_collection("your_collection_name", your_simple_value)
//    strings will be stored as bytes_list.
//
// 2. For Protobuf types, there are three ways to add them:
//    1) tf.add_to_collection("your_collection_name",
//         your_proto.SerializeToString())
//
//       collection_def {
//         key: "user_defined_bytes_collection"
//         value {
//           bytes_list {
//             value: "queue_name: \"test_queue\"\n"
//           }
//         }
//       }
//
//  or
//
//    2) tf.add_to_collection("your_collection_name", str(your_proto))
//
//       collection_def {
//         key: "user_defined_string_collection"
//         value {
//          bytes_list {
//             value: "\n\ntest_queue"
//           }
//         }
//       }
//
//  or
//
//    3) any_buf = any_pb2.Any()
//       tf.add_to_collection("your_collection_name",
//         any_buf.Pack(your_proto))
//
//       collection_def {
//         key: "user_defined_any_collection"
//         value {
//           any_list {
//             value {
//               type_url: "type.googleapis.com/tensorflow.QueueRunnerDef"
//               value: "\n\ntest_queue"
//             }
//           }
//         }
//       }
//
// 3. For Python objects, implement to_proto() and from_proto(), and register
//    them in the following manner:
//    ops.register_proto_function("your_collection_name",
//                                proto_type,
//                                to_proto=YourPythonObject.to_proto,
//                                from_proto=YourPythonObject.from_proto)
//    These functions will be invoked to serialize and de-serialize the
//    collection. For example,
//    ops.register_proto_function(ops.GraphKeys.GLOBAL_VARIABLES,
//                                proto_type=variable_pb2.VariableDef,
//                                to_proto=Variable.to_proto,
//                                from_proto=Variable.from_proto)
type CollectionDef struct {
	// Types that are valid to be assigned to Kind:
	//	*CollectionDef_NodeList_
	//	*CollectionDef_BytesList_
	//	*CollectionDef_Int64List_
	//	*CollectionDef_FloatList_
	//	*CollectionDef_AnyList_
	Kind isCollectionDef_Kind `protobuf_oneof:"kind"`
}

func (m *CollectionDef) Reset()                    { *m = CollectionDef{} }
func (m *CollectionDef) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef) ProtoMessage()               {}
func (*CollectionDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isCollectionDef_Kind interface {
	isCollectionDef_Kind()
}

type CollectionDef_NodeList_ struct {
	NodeList *CollectionDef_NodeList `protobuf:"bytes,1,opt,name=node_list,json=nodeList,oneof"`
}
type CollectionDef_BytesList_ struct {
	BytesList *CollectionDef_BytesList `protobuf:"bytes,2,opt,name=bytes_list,json=bytesList,oneof"`
}
type CollectionDef_Int64List_ struct {
	Int64List *CollectionDef_Int64List `protobuf:"bytes,3,opt,name=int64_list,json=int64List,oneof"`
}
type CollectionDef_FloatList_ struct {
	FloatList *CollectionDef_FloatList `protobuf:"bytes,4,opt,name=float_list,json=floatList,oneof"`
}
type CollectionDef_AnyList_ struct {
	AnyList *CollectionDef_AnyList `protobuf:"bytes,5,opt,name=any_list,json=anyList,oneof"`
}

func (*CollectionDef_NodeList_) isCollectionDef_Kind()  {}
func (*CollectionDef_BytesList_) isCollectionDef_Kind() {}
func (*CollectionDef_Int64List_) isCollectionDef_Kind() {}
func (*CollectionDef_FloatList_) isCollectionDef_Kind() {}
func (*CollectionDef_AnyList_) isCollectionDef_Kind()   {}

func (m *CollectionDef) GetKind() isCollectionDef_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *CollectionDef) GetNodeList() *CollectionDef_NodeList {
	if x, ok := m.GetKind().(*CollectionDef_NodeList_); ok {
		return x.NodeList
	}
	return nil
}

func (m *CollectionDef) GetBytesList() *CollectionDef_BytesList {
	if x, ok := m.GetKind().(*CollectionDef_BytesList_); ok {
		return x.BytesList
	}
	return nil
}

func (m *CollectionDef) GetInt64List() *CollectionDef_Int64List {
	if x, ok := m.GetKind().(*CollectionDef_Int64List_); ok {
		return x.Int64List
	}
	return nil
}

func (m *CollectionDef) GetFloatList() *CollectionDef_FloatList {
	if x, ok := m.GetKind().(*CollectionDef_FloatList_); ok {
		return x.FloatList
	}
	return nil
}

func (m *CollectionDef) GetAnyList() *CollectionDef_AnyList {
	if x, ok := m.GetKind().(*CollectionDef_AnyList_); ok {
		return x.AnyList
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CollectionDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CollectionDef_OneofMarshaler, _CollectionDef_OneofUnmarshaler, _CollectionDef_OneofSizer, []interface{}{
		(*CollectionDef_NodeList_)(nil),
		(*CollectionDef_BytesList_)(nil),
		(*CollectionDef_Int64List_)(nil),
		(*CollectionDef_FloatList_)(nil),
		(*CollectionDef_AnyList_)(nil),
	}
}

func _CollectionDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CollectionDef)
	// kind
	switch x := m.Kind.(type) {
	case *CollectionDef_NodeList_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NodeList); err != nil {
			return err
		}
	case *CollectionDef_BytesList_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.BytesList); err != nil {
			return err
		}
	case *CollectionDef_Int64List_:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Int64List); err != nil {
			return err
		}
	case *CollectionDef_FloatList_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FloatList); err != nil {
			return err
		}
	case *CollectionDef_AnyList_:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AnyList); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CollectionDef.Kind has unexpected type %T", x)
	}
	return nil
}

func _CollectionDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CollectionDef)
	switch tag {
	case 1: // kind.node_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CollectionDef_NodeList)
		err := b.DecodeMessage(msg)
		m.Kind = &CollectionDef_NodeList_{msg}
		return true, err
	case 2: // kind.bytes_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CollectionDef_BytesList)
		err := b.DecodeMessage(msg)
		m.Kind = &CollectionDef_BytesList_{msg}
		return true, err
	case 3: // kind.int64_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CollectionDef_Int64List)
		err := b.DecodeMessage(msg)
		m.Kind = &CollectionDef_Int64List_{msg}
		return true, err
	case 4: // kind.float_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CollectionDef_FloatList)
		err := b.DecodeMessage(msg)
		m.Kind = &CollectionDef_FloatList_{msg}
		return true, err
	case 5: // kind.any_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CollectionDef_AnyList)
		err := b.DecodeMessage(msg)
		m.Kind = &CollectionDef_AnyList_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CollectionDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CollectionDef)
	// kind
	switch x := m.Kind.(type) {
	case *CollectionDef_NodeList_:
		s := proto.Size(x.NodeList)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CollectionDef_BytesList_:
		s := proto.Size(x.BytesList)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CollectionDef_Int64List_:
		s := proto.Size(x.Int64List)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CollectionDef_FloatList_:
		s := proto.Size(x.FloatList)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CollectionDef_AnyList_:
		s := proto.Size(x.AnyList)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// NodeList is used for collecting nodes in graph. For example
// collection_def {
//   key: "summaries"
//   value {
//     node_list {
//       value: "input_producer/ScalarSummary:0"
//       value: "shuffle_batch/ScalarSummary:0"
//       value: "ImageSummary:0"
//     }
//   }
type CollectionDef_NodeList struct {
	Value []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *CollectionDef_NodeList) Reset()                    { *m = CollectionDef_NodeList{} }
func (m *CollectionDef_NodeList) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef_NodeList) ProtoMessage()               {}
func (*CollectionDef_NodeList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 0} }

func (m *CollectionDef_NodeList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

// BytesList is used for collecting strings and serialized protobufs. For
// example:
// collection_def {
//   key: "trainable_variables"
//   value {
//     bytes_list {
//       value: "\n\017conv1/weights:0\022\024conv1/weights/Assign
//              \032\024conv1/weights/read:0"
//       value: "\n\016conv1/biases:0\022\023conv1/biases/Assign\032
//              \023conv1/biases/read:0"
//     }
//   }
// }
type CollectionDef_BytesList struct {
	Value [][]byte `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (m *CollectionDef_BytesList) Reset()                    { *m = CollectionDef_BytesList{} }
func (m *CollectionDef_BytesList) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef_BytesList) ProtoMessage()               {}
func (*CollectionDef_BytesList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 1} }

func (m *CollectionDef_BytesList) GetValue() [][]byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Int64List is used for collecting int, int64 and long values.
type CollectionDef_Int64List struct {
	Value []int64 `protobuf:"varint,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *CollectionDef_Int64List) Reset()                    { *m = CollectionDef_Int64List{} }
func (m *CollectionDef_Int64List) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef_Int64List) ProtoMessage()               {}
func (*CollectionDef_Int64List) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 2} }

func (m *CollectionDef_Int64List) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

// FloatList is used for collecting float values.
type CollectionDef_FloatList struct {
	Value []float32 `protobuf:"fixed32,1,rep,packed,name=value" json:"value,omitempty"`
}

func (m *CollectionDef_FloatList) Reset()                    { *m = CollectionDef_FloatList{} }
func (m *CollectionDef_FloatList) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef_FloatList) ProtoMessage()               {}
func (*CollectionDef_FloatList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 3} }

func (m *CollectionDef_FloatList) GetValue() []float32 {
	if m != nil {
		return m.Value
	}
	return nil
}

// AnyList is used for collecting Any protos.
type CollectionDef_AnyList struct {
	Value []*google_protobuf.Any `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
}

func (m *CollectionDef_AnyList) Reset()                    { *m = CollectionDef_AnyList{} }
func (m *CollectionDef_AnyList) String() string            { return proto.CompactTextString(m) }
func (*CollectionDef_AnyList) ProtoMessage()               {}
func (*CollectionDef_AnyList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1, 4} }

func (m *CollectionDef_AnyList) GetValue() []*google_protobuf.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

// Information about a Tensor necessary for feeding or retrieval.
type TensorInfo struct {
	// Types that are valid to be assigned to Encoding:
	//	*TensorInfo_Name
	//	*TensorInfo_CooSparse_
	Encoding isTensorInfo_Encoding `protobuf_oneof:"encoding"`
	Dtype    tensorflow3.DataType  `protobuf:"varint,2,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// The static shape should be recorded here, to the extent that it can
	// be known in advance.  In the case of a SparseTensor, this field describes
	// the logical shape of the represented tensor (aka dense_shape).
	TensorShape *tensorflow2.TensorShapeProto `protobuf:"bytes,3,opt,name=tensor_shape,json=tensorShape" json:"tensor_shape,omitempty"`
}

func (m *TensorInfo) Reset()                    { *m = TensorInfo{} }
func (m *TensorInfo) String() string            { return proto.CompactTextString(m) }
func (*TensorInfo) ProtoMessage()               {}
func (*TensorInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type isTensorInfo_Encoding interface {
	isTensorInfo_Encoding()
}

type TensorInfo_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,oneof"`
}
type TensorInfo_CooSparse_ struct {
	CooSparse *TensorInfo_CooSparse `protobuf:"bytes,4,opt,name=coo_sparse,json=cooSparse,oneof"`
}

func (*TensorInfo_Name) isTensorInfo_Encoding()       {}
func (*TensorInfo_CooSparse_) isTensorInfo_Encoding() {}

func (m *TensorInfo) GetEncoding() isTensorInfo_Encoding {
	if m != nil {
		return m.Encoding
	}
	return nil
}

func (m *TensorInfo) GetName() string {
	if x, ok := m.GetEncoding().(*TensorInfo_Name); ok {
		return x.Name
	}
	return ""
}

func (m *TensorInfo) GetCooSparse() *TensorInfo_CooSparse {
	if x, ok := m.GetEncoding().(*TensorInfo_CooSparse_); ok {
		return x.CooSparse
	}
	return nil
}

func (m *TensorInfo) GetDtype() tensorflow3.DataType {
	if m != nil {
		return m.Dtype
	}
	return tensorflow3.DataType_DT_INVALID
}

func (m *TensorInfo) GetTensorShape() *tensorflow2.TensorShapeProto {
	if m != nil {
		return m.TensorShape
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TensorInfo) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TensorInfo_OneofMarshaler, _TensorInfo_OneofUnmarshaler, _TensorInfo_OneofSizer, []interface{}{
		(*TensorInfo_Name)(nil),
		(*TensorInfo_CooSparse_)(nil),
	}
}

func _TensorInfo_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TensorInfo)
	// encoding
	switch x := m.Encoding.(type) {
	case *TensorInfo_Name:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Name)
	case *TensorInfo_CooSparse_:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CooSparse); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TensorInfo.Encoding has unexpected type %T", x)
	}
	return nil
}

func _TensorInfo_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TensorInfo)
	switch tag {
	case 1: // encoding.name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Encoding = &TensorInfo_Name{x}
		return true, err
	case 4: // encoding.coo_sparse
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorInfo_CooSparse)
		err := b.DecodeMessage(msg)
		m.Encoding = &TensorInfo_CooSparse_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TensorInfo_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TensorInfo)
	// encoding
	switch x := m.Encoding.(type) {
	case *TensorInfo_Name:
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case *TensorInfo_CooSparse_:
		s := proto.Size(x.CooSparse)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// For sparse tensors, The COO encoding stores a triple of values, indices,
// and shape.
type TensorInfo_CooSparse struct {
	// The shape of the values Tensor is [?].  Its dtype must be the dtype of
	// the SparseTensor as a whole, given in the enclosing TensorInfo.
	ValuesTensorName string `protobuf:"bytes,1,opt,name=values_tensor_name,json=valuesTensorName" json:"values_tensor_name,omitempty"`
	// The indices Tensor must have dtype int64 and shape [?, ?].
	IndicesTensorName string `protobuf:"bytes,2,opt,name=indices_tensor_name,json=indicesTensorName" json:"indices_tensor_name,omitempty"`
	// The dynamic logical shape represented by the SparseTensor is recorded in
	// the Tensor referenced here.  It must have dtype int64 and shape [?].
	DenseShapeTensorName string `protobuf:"bytes,3,opt,name=dense_shape_tensor_name,json=denseShapeTensorName" json:"dense_shape_tensor_name,omitempty"`
}

func (m *TensorInfo_CooSparse) Reset()                    { *m = TensorInfo_CooSparse{} }
func (m *TensorInfo_CooSparse) String() string            { return proto.CompactTextString(m) }
func (*TensorInfo_CooSparse) ProtoMessage()               {}
func (*TensorInfo_CooSparse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

func (m *TensorInfo_CooSparse) GetValuesTensorName() string {
	if m != nil {
		return m.ValuesTensorName
	}
	return ""
}

func (m *TensorInfo_CooSparse) GetIndicesTensorName() string {
	if m != nil {
		return m.IndicesTensorName
	}
	return ""
}

func (m *TensorInfo_CooSparse) GetDenseShapeTensorName() string {
	if m != nil {
		return m.DenseShapeTensorName
	}
	return ""
}

// SignatureDef defines the signature of a computation supported by a TensorFlow
// graph.
//
// For example, a model with two loss computations, sharing a single input,
// might have the following signature_def map.
//
// Note that across the two SignatureDefs "loss_A" and "loss_B", the input key,
// output key, and method_name are identical, and will be used by system(s) that
// implement or rely upon this particular loss method. The output tensor names
// differ, demonstrating how different outputs can exist for the same method.
//
// signature_def {
//   key: "loss_A"
//   value {
//     inputs {
//       key: "input"
//       value {
//         name: "input:0"
//         dtype: DT_STRING
//         tensor_shape: ...
//       }
//     }
//     outputs {
//       key: "loss_output"
//       value {
//         name: "loss_output_A:0"
//         dtype: DT_FLOAT
//         tensor_shape: ...
//       }
//     }
//   }
//   ...
//   method_name: "some/package/compute_loss"
// }
// signature_def {
//   key: "loss_B"
//   value {
//     inputs {
//       key: "input"
//       value {
//         name: "input:0"
//         dtype: DT_STRING
//         tensor_shape: ...
//       }
//     }
//     outputs {
//       key: "loss_output"
//       value {
//         name: "loss_output_B:0"
//         dtype: DT_FLOAT
//         tensor_shape: ...
//       }
//     }
//   }
//   ...
//   method_name: "some/package/compute_loss"
// }
type SignatureDef struct {
	// Named input parameters.
	Inputs map[string]*TensorInfo `protobuf:"bytes,1,rep,name=inputs" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Named output parameters.
	Outputs map[string]*TensorInfo `protobuf:"bytes,2,rep,name=outputs" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Extensible method_name information enabling third-party users to mark a
	// SignatureDef as supporting a particular method. This enables producers and
	// consumers of SignatureDefs, e.g. a model definition library and a serving
	// library to have a clear hand-off regarding the semantics of a computation.
	//
	// Note that multiple SignatureDefs in a single MetaGraphDef may have the same
	// method_name. This is commonly used to support multi-headed computation,
	// where a single graph computation may return multiple results.
	MethodName string `protobuf:"bytes,3,opt,name=method_name,json=methodName" json:"method_name,omitempty"`
}

func (m *SignatureDef) Reset()                    { *m = SignatureDef{} }
func (m *SignatureDef) String() string            { return proto.CompactTextString(m) }
func (*SignatureDef) ProtoMessage()               {}
func (*SignatureDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SignatureDef) GetInputs() map[string]*TensorInfo {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *SignatureDef) GetOutputs() map[string]*TensorInfo {
	if m != nil {
		return m.Outputs
	}
	return nil
}

func (m *SignatureDef) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

// An asset file def for a single file or a set of sharded files with the same
// name.
type AssetFileDef struct {
	// The tensor to bind the asset filename to.
	TensorInfo *TensorInfo `protobuf:"bytes,1,opt,name=tensor_info,json=tensorInfo" json:"tensor_info,omitempty"`
	// The filename within an assets directory. Note: does not include the path
	// prefix, i.e. directories. For an asset at /tmp/path/vocab.txt, the filename
	// would be "vocab.txt".
	Filename string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
}

func (m *AssetFileDef) Reset()                    { *m = AssetFileDef{} }
func (m *AssetFileDef) String() string            { return proto.CompactTextString(m) }
func (*AssetFileDef) ProtoMessage()               {}
func (*AssetFileDef) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *AssetFileDef) GetTensorInfo() *TensorInfo {
	if m != nil {
		return m.TensorInfo
	}
	return nil
}

func (m *AssetFileDef) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func init() {
	proto.RegisterType((*MetaGraphDef)(nil), "tensorflow.MetaGraphDef")
	proto.RegisterType((*MetaGraphDef_MetaInfoDef)(nil), "tensorflow.MetaGraphDef.MetaInfoDef")
	proto.RegisterType((*CollectionDef)(nil), "tensorflow.CollectionDef")
	proto.RegisterType((*CollectionDef_NodeList)(nil), "tensorflow.CollectionDef.NodeList")
	proto.RegisterType((*CollectionDef_BytesList)(nil), "tensorflow.CollectionDef.BytesList")
	proto.RegisterType((*CollectionDef_Int64List)(nil), "tensorflow.CollectionDef.Int64List")
	proto.RegisterType((*CollectionDef_FloatList)(nil), "tensorflow.CollectionDef.FloatList")
	proto.RegisterType((*CollectionDef_AnyList)(nil), "tensorflow.CollectionDef.AnyList")
	proto.RegisterType((*TensorInfo)(nil), "tensorflow.TensorInfo")
	proto.RegisterType((*TensorInfo_CooSparse)(nil), "tensorflow.TensorInfo.CooSparse")
	proto.RegisterType((*SignatureDef)(nil), "tensorflow.SignatureDef")
	proto.RegisterType((*AssetFileDef)(nil), "tensorflow.AssetFileDef")
}

func init() { proto.RegisterFile("github.com/tensorflow/tensorflow/tensorflow/go/core/protobuf/meta_graph.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1016 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xc7, 0x2b, 0xdb, 0x89, 0xed, 0x63, 0x27, 0x73, 0x39, 0xa3, 0x73, 0x85, 0x01, 0x73, 0xbd,
	0x64, 0xc8, 0xba, 0x4c, 0x5e, 0xbb, 0x76, 0x1b, 0x86, 0xa2, 0x85, 0x5d, 0xaf, 0x4d, 0x80, 0xad,
	0xe9, 0x94, 0x62, 0xc0, 0xb0, 0x0b, 0x81, 0xb1, 0x28, 0x47, 0x88, 0x4c, 0x0a, 0x22, 0x9d, 0x42,
	0x97, 0x7b, 0x93, 0x3d, 0xc5, 0xf6, 0x24, 0xbb, 0xd8, 0x9b, 0xec, 0x72, 0x20, 0xa9, 0x0f, 0xaa,
	0x89, 0x96, 0x9b, 0xde, 0xe9, 0xf0, 0xfc, 0xcf, 0x8f, 0x87, 0x87, 0x87, 0xa4, 0xe0, 0x73, 0x41,
	0x28, 0x67, 0x49, 0x10, 0xb1, 0xb7, 0xd3, 0x25, 0x4b, 0xc8, 0x34, 0x4e, 0x98, 0x60, 0x67, 0x9b,
	0x60, 0xba, 0x26, 0x02, 0x7b, 0xab, 0x04, 0xc7, 0xe7, 0x8e, 0x1a, 0x43, 0x50, 0x4a, 0xed, 0xbb,
	0x2b, 0xc6, 0x56, 0x91, 0xa1, 0xc6, 0x34, 0xd5, 0x32, 0x7b, 0xff, 0x5d, 0x62, 0x90, 0xe0, 0x35,
	0x79, 0xcb, 0x92, 0x8b, 0xa9, 0x41, 0xb3, 0x3f, 0xab, 0x97, 0xb1, 0xd8, 0xf3, 0x49, 0x90, 0xe9,
	0x0e, 0xeb, 0x75, 0xda, 0xe3, 0xf1, 0x73, 0x1c, 0x93, 0x9b, 0x27, 0x17, 0x69, 0x4c, 0x78, 0x26,
	0xdb, 0xab, 0x5d, 0x35, 0xc7, 0x97, 0x24, 0xd1, 0xaa, 0xc9, 0x5f, 0x6d, 0xe8, 0xff, 0x44, 0x04,
	0x7e, 0x29, 0xd3, 0x5e, 0x90, 0x00, 0x1d, 0xc1, 0x8e, 0xaa, 0x4a, 0x48, 0x03, 0x26, 0x53, 0x1c,
	0x59, 0x63, 0xeb, 0xa0, 0xf7, 0x70, 0xcf, 0x29, 0x71, 0x8e, 0x19, 0xa0, 0x8c, 0x63, 0x1a, 0xb0,
	0x05, 0x09, 0xdc, 0xde, 0xba, 0x34, 0xd0, 0x03, 0xe8, 0xaa, 0x62, 0x28, 0x4a, 0x43, 0x51, 0x86,
	0x26, 0x25, 0x27, 0xb8, 0x9d, 0x55, 0x3e, 0xf9, 0x03, 0xe8, 0xaa, 0xe4, 0x54, 0x48, 0xf3, 0x6a,
	0xc8, 0xa9, 0x74, 0xaa, 0x10, 0x9e, 0x7d, 0x21, 0x17, 0x76, 0x97, 0x2c, 0x8a, 0xc8, 0x52, 0x84,
	0x8c, 0xaa, 0xb8, 0xd6, 0xb8, 0x79, 0xd0, 0x7b, 0xf8, 0x45, 0x6d, 0xc2, 0xcf, 0x0b, 0xf9, 0x82,
	0x04, 0x3f, 0x50, 0x91, 0xa4, 0xee, 0xce, 0xd2, 0x1c, 0x43, 0x27, 0xb0, 0xc3, 0xc3, 0x15, 0xc5,
	0x62, 0x93, 0x10, 0x85, 0xdc, 0x52, 0xc8, 0xfb, 0xb5, 0xc8, 0xd3, 0x5c, 0x5d, 0x10, 0xfb, 0xdc,
	0x18, 0x42, 0x4f, 0x61, 0x17, 0x73, 0x4e, 0x84, 0x17, 0x84, 0x91, 0x26, 0x6e, 0x2b, 0xe2, 0xc8,
	0x24, 0xce, 0xa4, 0xe2, 0x45, 0x18, 0xc9, 0x08, 0xb7, 0x8f, 0x0d, 0xcb, 0xfe, 0xa7, 0x01, 0x3d,
	0xa3, 0xce, 0xe8, 0x10, 0x50, 0xd9, 0xba, 0xde, 0x25, 0x49, 0x78, 0xc8, 0xa8, 0xda, 0xa9, 0xae,
	0x3b, 0x58, 0xe7, 0x99, 0xfd, 0xa2, 0xc7, 0xd1, 0x13, 0x18, 0x70, 0x91, 0x84, 0x71, 0x4c, 0x7c,
	0x8f, 0xc5, 0x5e, 0x14, 0x72, 0x91, 0xed, 0x07, 0x32, 0xe7, 0x3f, 0x89, 0x7f, 0x0c, 0xb9, 0x70,
	0x77, 0x73, 0xad, 0xb6, 0xd1, 0x14, 0x3a, 0x98, 0xa6, 0xaa, 0x1f, 0x8a, 0x2d, 0xd1, 0x27, 0xc3,
	0xc9, 0x3b, 0xca, 0x99, 0xd1, 0xd4, 0x6d, 0x63, 0x9a, 0xca, 0xfc, 0x10, 0x82, 0x96, 0xc0, 0x2b,
	0xae, 0xf6, 0xa1, 0xeb, 0xaa, 0x6f, 0xf4, 0x25, 0xa0, 0x72, 0xa6, 0x22, 0xe1, 0x2d, 0x95, 0xf0,
	0xed, 0xd2, 0x93, 0x67, 0xfc, 0x08, 0xee, 0x18, 0xf2, 0x55, 0x28, 0x8a, 0x90, 0x6d, 0x15, 0x32,
	0x2c, 0xbd, 0x2f, 0x43, 0x61, 0x44, 0x15, 0xeb, 0xf4, 0x49, 0x80, 0x37, 0x91, 0xf0, 0xb0, 0x10,
	0x09, 0x1f, 0xb5, 0xc7, 0xd6, 0x41, 0xc7, 0x1d, 0xe6, 0xde, 0x85, 0x76, 0xce, 0xa4, 0xcf, 0xfe,
	0x0d, 0xd0, 0xd5, 0x8e, 0x40, 0x03, 0x68, 0x5e, 0x90, 0x34, 0x2b, 0xa9, 0xfc, 0x44, 0x53, 0xd8,
	0xba, 0xc4, 0xd1, 0x86, 0x64, 0xa5, 0xbb, 0x6b, 0x96, 0xae, 0x02, 0x70, 0xb5, 0xee, 0xfb, 0xc6,
	0x77, 0x96, 0xfd, 0x2b, 0xdc, 0xbe, 0xd2, 0x1b, 0xd7, 0xb0, 0x9d, 0x2a, 0xbb, 0xd2, 0x16, 0x66,
	0xbc, 0x81, 0x9e, 0xfc, 0xd9, 0x82, 0x9d, 0xca, 0xbc, 0x68, 0x06, 0x5d, 0xca, 0x7c, 0xa2, 0x37,
	0x58, 0x1f, 0xdb, 0x49, 0x6d, 0x96, 0xce, 0x2b, 0xe6, 0x13, 0xb9, 0xc1, 0x47, 0xb7, 0xdc, 0x0e,
	0xcd, 0xbe, 0xd1, 0x02, 0xe0, 0x2c, 0x15, 0x84, 0x9b, 0x4d, 0xf2, 0x69, 0x3d, 0x63, 0x2e, 0xb5,
	0x19, 0xa4, 0x7b, 0x96, 0x1b, 0x92, 0x12, 0x52, 0xf1, 0xcd, 0x23, 0x4d, 0x69, 0xde, 0x44, 0x39,
	0x96, 0xda, 0x9c, 0x12, 0xe6, 0x86, 0xa4, 0x04, 0x11, 0xc3, 0x42, 0x53, 0x5a, 0x37, 0x51, 0x5e,
	0x48, 0x6d, 0x4e, 0x09, 0x72, 0x03, 0x3d, 0xd5, 0xed, 0xab, 0x18, 0x5b, 0x8a, 0x71, 0xaf, 0x9e,
	0x31, 0xa3, 0x69, 0x46, 0x90, 0xdd, 0x2c, 0x3f, 0xed, 0x31, 0x74, 0xf2, 0x4a, 0xa1, 0x61, 0xbe,
	0x4d, 0x96, 0x6a, 0x6d, 0x6d, 0xd8, 0xf7, 0xa0, 0x5b, 0xd4, 0xa1, 0x2a, 0xe9, 0xe7, 0x92, 0x7d,
	0xe8, 0x16, 0x8b, 0x44, 0x23, 0x53, 0xd2, 0x9c, 0x37, 0x06, 0x96, 0x21, 0x2b, 0x56, 0x51, 0x95,
	0x35, 0x4c, 0xd9, 0x63, 0x68, 0x67, 0x89, 0xa2, 0xfb, 0xa6, 0xa8, 0xee, 0x64, 0x6a, 0xc9, 0x7c,
	0x1b, 0x5a, 0x17, 0x21, 0xf5, 0x27, 0xbf, 0x37, 0x01, 0xde, 0xa8, 0x0a, 0xa8, 0xe3, 0x3a, 0x84,
	0x16, 0xc5, 0x6b, 0xa2, 0xdb, 0xf1, 0xe8, 0x96, 0xab, 0x2c, 0x34, 0x03, 0x58, 0x32, 0xe6, 0xf1,
	0x18, 0x27, 0x9c, 0x64, 0xc5, 0x1f, 0x9b, 0x85, 0x2b, 0x09, 0xce, 0x73, 0xc6, 0x4e, 0x95, 0x4e,
	0x56, 0x7e, 0x99, 0x1b, 0x32, 0x37, 0x5f, 0x3e, 0x48, 0xaa, 0x8d, 0x76, 0xab, 0x17, 0xf9, 0x02,
	0x0b, 0xfc, 0x26, 0x8d, 0x89, 0xab, 0x25, 0xe8, 0x19, 0xf4, 0xcd, 0x97, 0x2e, 0xeb, 0x99, 0x8f,
	0xaf, 0x4e, 0x78, 0x2a, 0xdd, 0xaf, 0xe5, 0xea, 0xdc, 0x9e, 0x28, 0x47, 0xec, 0x3f, 0x2c, 0xe8,
	0x16, 0x79, 0xc8, 0xfb, 0x51, 0xad, 0x99, 0x7b, 0x19, 0xb5, 0x5c, 0xa1, 0x3b, 0xd0, 0x1e, 0x8d,
	0x7b, 0x25, 0xd7, 0xea, 0xc0, 0x87, 0x21, 0xf5, 0xc3, 0xe5, 0x3b, 0xf2, 0x86, 0xbe, 0x9d, 0x32,
	0x97, 0xa1, 0x7f, 0x0c, 0x1f, 0xf9, 0x84, 0x72, 0xa2, 0x73, 0xad, 0xc4, 0x34, 0xf5, 0xf5, 0xa4,
	0xdc, 0x2a, 0xb1, 0x32, 0x6c, 0x0e, 0xd0, 0x21, 0x74, 0xc9, 0xfc, 0x90, 0xae, 0x26, 0x7f, 0x37,
	0xa0, 0x6f, 0x1e, 0x6c, 0xf4, 0x04, 0xb6, 0x43, 0x1a, 0x6f, 0x04, 0xcf, 0x76, 0x72, 0xaf, 0xee,
	0x0a, 0x70, 0x8e, 0x95, 0x4c, 0xbf, 0x32, 0x59, 0x0c, 0x7a, 0x06, 0x6d, 0xb6, 0x11, 0x2a, 0xbc,
	0xa1, 0xc2, 0xf7, 0x6b, 0xc3, 0x4f, 0xb4, 0x4e, 0xc7, 0xe7, 0x51, 0xe8, 0x13, 0x90, 0x4f, 0xf7,
	0x39, 0xf3, 0xcd, 0x65, 0x80, 0x1e, 0x92, 0xc9, 0xdb, 0x3f, 0x43, 0xcf, 0x98, 0xf8, 0x9a, 0x2b,
	0xec, 0xb0, 0x7a, 0x85, 0xdd, 0xb9, 0xbe, 0x57, 0xcc, 0xbb, 0xd1, 0x85, 0xbe, 0x99, 0xcc, 0xfb,
	0x60, 0x4e, 0x96, 0xd0, 0x37, 0x9f, 0x51, 0xf4, 0x2d, 0x64, 0x5d, 0xa2, 0xdf, 0x2f, 0xeb, 0x7f,
	0x39, 0xd9, 0xcf, 0x9f, 0x3a, 0x15, 0x36, 0x74, 0xe4, 0x5b, 0x6d, 0x34, 0x42, 0x61, 0xcf, 0xbf,
	0x82, 0x11, 0x4b, 0x56, 0x26, 0xa4, 0xf8, 0x03, 0x9b, 0x7f, 0x50, 0xfc, 0x17, 0xa8, 0x26, 0xe5,
	0xaf, 0xad, 0x7f, 0x2d, 0xeb, 0x6c, 0x5b, 0x9d, 0xc7, 0xaf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0xba, 0x6a, 0x3e, 0xe3, 0x8a, 0x0a, 0x00, 0x00,
}
