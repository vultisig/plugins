// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: Nano.proto

package nano

import (
	common "github.com/vultisig/vultisigner/walletcore/protos/common"
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

// Input data necessary to create a signed transaction.
type SigningInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The secret private key used for signing (32 bytes).
	PrivateKey []byte `protobuf:"bytes,1,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// Optional parent block hash
	ParentBlock []byte `protobuf:"bytes,2,opt,name=parent_block,json=parentBlock,proto3" json:"parent_block,omitempty"`
	// Receive/Send reference
	//
	// Types that are assignable to LinkOneof:
	//
	//	*SigningInput_LinkBlock
	//	*SigningInput_LinkRecipient
	LinkOneof isSigningInput_LinkOneof `protobuf_oneof:"link_oneof"`
	// Representative address
	Representative string `protobuf:"bytes,5,opt,name=representative,proto3" json:"representative,omitempty"`
	// Account balance (128-bit unsigned integer, as a string)
	Balance string `protobuf:"bytes,6,opt,name=balance,proto3" json:"balance,omitempty"`
	// Work
	Work string `protobuf:"bytes,7,opt,name=work,proto3" json:"work,omitempty"`
	// Pulic key used for building preImage (32 bytes).
	PublicKey []byte `protobuf:"bytes,8,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *SigningInput) Reset() {
	*x = SigningInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Nano_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningInput) ProtoMessage() {}

func (x *SigningInput) ProtoReflect() protoreflect.Message {
	mi := &file_Nano_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningInput.ProtoReflect.Descriptor instead.
func (*SigningInput) Descriptor() ([]byte, []int) {
	return file_Nano_proto_rawDescGZIP(), []int{0}
}

func (x *SigningInput) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

func (x *SigningInput) GetParentBlock() []byte {
	if x != nil {
		return x.ParentBlock
	}
	return nil
}

func (m *SigningInput) GetLinkOneof() isSigningInput_LinkOneof {
	if m != nil {
		return m.LinkOneof
	}
	return nil
}

func (x *SigningInput) GetLinkBlock() []byte {
	if x, ok := x.GetLinkOneof().(*SigningInput_LinkBlock); ok {
		return x.LinkBlock
	}
	return nil
}

func (x *SigningInput) GetLinkRecipient() string {
	if x, ok := x.GetLinkOneof().(*SigningInput_LinkRecipient); ok {
		return x.LinkRecipient
	}
	return ""
}

func (x *SigningInput) GetRepresentative() string {
	if x != nil {
		return x.Representative
	}
	return ""
}

func (x *SigningInput) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *SigningInput) GetWork() string {
	if x != nil {
		return x.Work
	}
	return ""
}

func (x *SigningInput) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type isSigningInput_LinkOneof interface {
	isSigningInput_LinkOneof()
}

type SigningInput_LinkBlock struct {
	// Hash of a block to receive from
	LinkBlock []byte `protobuf:"bytes,3,opt,name=link_block,json=linkBlock,proto3,oneof"`
}

type SigningInput_LinkRecipient struct {
	// Recipient address to send coins to
	LinkRecipient string `protobuf:"bytes,4,opt,name=link_recipient,json=linkRecipient,proto3,oneof"`
}

func (*SigningInput_LinkBlock) isSigningInput_LinkOneof() {}

func (*SigningInput_LinkRecipient) isSigningInput_LinkOneof() {}

// Result containing the signed and encoded transaction.
type SigningOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signature
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Block hash
	BlockHash []byte `protobuf:"bytes,2,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// JSON representation of the block
	Json string `protobuf:"bytes,3,opt,name=json,proto3" json:"json,omitempty"`
	// error code, 0 is ok, other codes will be treated as errors
	Error common.SigningError `protobuf:"varint,4,opt,name=error,proto3,enum=TW.Common.Proto.SigningError" json:"error,omitempty"`
	// error code description
	ErrorMessage string `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *SigningOutput) Reset() {
	*x = SigningOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Nano_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SigningOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SigningOutput) ProtoMessage() {}

func (x *SigningOutput) ProtoReflect() protoreflect.Message {
	mi := &file_Nano_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SigningOutput.ProtoReflect.Descriptor instead.
func (*SigningOutput) Descriptor() ([]byte, []int) {
	return file_Nano_proto_rawDescGZIP(), []int{1}
}

func (x *SigningOutput) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SigningOutput) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *SigningOutput) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

func (x *SigningOutput) GetError() common.SigningError {
	if x != nil {
		return x.Error
	}
	return common.SigningError(0)
}

func (x *SigningOutput) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_Nano_proto protoreflect.FileDescriptor

var file_Nano_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x4e, 0x61, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x54, 0x57,
	0x2e, 0x4e, 0x61, 0x6e, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x02, 0x0a, 0x0c, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f,
	0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x48, 0x00, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x27, 0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x6c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x0c, 0x0a,
	0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x22, 0xba, 0x01, 0x0a, 0x0d,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x12, 0x33,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e,
	0x54, 0x57, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x17, 0x0a, 0x15, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x6a, 0x6e, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Nano_proto_rawDescOnce sync.Once
	file_Nano_proto_rawDescData = file_Nano_proto_rawDesc
)

func file_Nano_proto_rawDescGZIP() []byte {
	file_Nano_proto_rawDescOnce.Do(func() {
		file_Nano_proto_rawDescData = protoimpl.X.CompressGZIP(file_Nano_proto_rawDescData)
	})
	return file_Nano_proto_rawDescData
}

var file_Nano_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_Nano_proto_goTypes = []interface{}{
	(*SigningInput)(nil),     // 0: TW.Nano.Proto.SigningInput
	(*SigningOutput)(nil),    // 1: TW.Nano.Proto.SigningOutput
	(common.SigningError)(0), // 2: TW.Common.Proto.SigningError
}
var file_Nano_proto_depIdxs = []int32{
	2, // 0: TW.Nano.Proto.SigningOutput.error:type_name -> TW.Common.Proto.SigningError
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Nano_proto_init() }
func file_Nano_proto_init() {
	if File_Nano_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Nano_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningInput); i {
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
		file_Nano_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SigningOutput); i {
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
	file_Nano_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SigningInput_LinkBlock)(nil),
		(*SigningInput_LinkRecipient)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Nano_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Nano_proto_goTypes,
		DependencyIndexes: file_Nano_proto_depIdxs,
		MessageInfos:      file_Nano_proto_msgTypes,
	}.Build()
	File_Nano_proto = out.File
	file_Nano_proto_rawDesc = nil
	file_Nano_proto_goTypes = nil
	file_Nano_proto_depIdxs = nil
}