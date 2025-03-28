// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: proto/solana.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Get the latest block hash
type GetLatestBlockHashRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLatestBlockHashRequest) Reset() {
	*x = GetLatestBlockHashRequest{}
	mi := &file_proto_solana_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestBlockHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockHashRequest) ProtoMessage() {}

func (x *GetLatestBlockHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockHashRequest.ProtoReflect.Descriptor instead.
func (*GetLatestBlockHashRequest) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{0}
}

type GetLatestBlockHashResponse struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	Blockhash            string                 `protobuf:"bytes,1,opt,name=blockhash,proto3" json:"blockhash,omitempty"`
	LastValidBlockHeight int64                  `protobuf:"varint,2,opt,name=last_valid_block_height,json=lastValidBlockHeight,proto3" json:"last_valid_block_height,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *GetLatestBlockHashResponse) Reset() {
	*x = GetLatestBlockHashResponse{}
	mi := &file_proto_solana_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLatestBlockHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockHashResponse) ProtoMessage() {}

func (x *GetLatestBlockHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockHashResponse.ProtoReflect.Descriptor instead.
func (*GetLatestBlockHashResponse) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatestBlockHashResponse) GetBlockhash() string {
	if x != nil {
		return x.Blockhash
	}
	return ""
}

func (x *GetLatestBlockHashResponse) GetLastValidBlockHeight() int64 {
	if x != nil {
		return x.LastValidBlockHeight
	}
	return 0
}

// Get the balance of an individual account
type GetAccountBalanceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       string                 `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountBalanceRequest) Reset() {
	*x = GetAccountBalanceRequest{}
	mi := &file_proto_solana_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountBalanceRequest) ProtoMessage() {}

func (x *GetAccountBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetAccountBalanceRequest) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountBalanceRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type GetAccountBalanceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Slot          int64                  `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	Value         int64                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAccountBalanceResponse) Reset() {
	*x = GetAccountBalanceResponse{}
	mi := &file_proto_solana_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountBalanceResponse) ProtoMessage() {}

func (x *GetAccountBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetAccountBalanceResponse) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{3}
}

func (x *GetAccountBalanceResponse) GetSlot() int64 {
	if x != nil {
		return x.Slot
	}
	return 0
}

func (x *GetAccountBalanceResponse) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

// Get the slot leader
type GetSlotLeaderRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Commitment     string                 `protobuf:"bytes,1,opt,name=commitment,proto3" json:"commitment,omitempty"`
	MinContextSlot int64                  `protobuf:"varint,2,opt,name=minContextSlot,proto3" json:"minContextSlot,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetSlotLeaderRequest) Reset() {
	*x = GetSlotLeaderRequest{}
	mi := &file_proto_solana_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSlotLeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlotLeaderRequest) ProtoMessage() {}

func (x *GetSlotLeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlotLeaderRequest.ProtoReflect.Descriptor instead.
func (*GetSlotLeaderRequest) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{4}
}

func (x *GetSlotLeaderRequest) GetCommitment() string {
	if x != nil {
		return x.Commitment
	}
	return ""
}

func (x *GetSlotLeaderRequest) GetMinContextSlot() int64 {
	if x != nil {
		return x.MinContextSlot
	}
	return 0
}

type GetSlotLeaderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pubkey        string                 `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSlotLeaderResponse) Reset() {
	*x = GetSlotLeaderResponse{}
	mi := &file_proto_solana_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSlotLeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSlotLeaderResponse) ProtoMessage() {}

func (x *GetSlotLeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_solana_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSlotLeaderResponse.ProtoReflect.Descriptor instead.
func (*GetSlotLeaderResponse) Descriptor() ([]byte, []int) {
	return file_proto_solana_proto_rawDescGZIP(), []int{5}
}

func (x *GetSlotLeaderResponse) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

var File_proto_solana_proto protoreflect.FileDescriptor

const file_proto_solana_proto_rawDesc = "" +
	"\n" +
	"\x12proto/solana.proto\x12\x05proto\x1a\x1cgoogle/api/annotations.proto\"\x1b\n" +
	"\x19GetLatestBlockHashRequest\"q\n" +
	"\x1aGetLatestBlockHashResponse\x12\x1c\n" +
	"\tblockhash\x18\x01 \x01(\tR\tblockhash\x125\n" +
	"\x17last_valid_block_height\x18\x02 \x01(\x03R\x14lastValidBlockHeight\"4\n" +
	"\x18GetAccountBalanceRequest\x12\x18\n" +
	"\aaccount\x18\x01 \x01(\tR\aaccount\"E\n" +
	"\x19GetAccountBalanceResponse\x12\x12\n" +
	"\x04slot\x18\x01 \x01(\x03R\x04slot\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x03R\x05value\"^\n" +
	"\x14GetSlotLeaderRequest\x12\x1e\n" +
	"\n" +
	"commitment\x18\x01 \x01(\tR\n" +
	"commitment\x12&\n" +
	"\x0eminContextSlot\x18\x02 \x01(\x03R\x0eminContextSlot\"/\n" +
	"\x15GetSlotLeaderResponse\x12\x16\n" +
	"\x06pubkey\x18\x01 \x01(\tR\x06pubkey2\xc4\x03\n" +
	"\rSolanaService\x12y\n" +
	"\x12GetLatestBlockHash\x12 .proto.GetLatestBlockHashRequest\x1a!.proto.GetLatestBlockHashResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/getLatestBlockHash\x12x\n" +
	"\x11GetAccountBalance\x12\x1f.proto.GetAccountBalanceRequest\x1a .proto.GetAccountBalanceResponse\" \x82\xd3\xe4\x93\x02\x1a:\x01*\"\x15/v1/getAccountBalance\x12h\n" +
	"\rGetSlotLeader\x12\x1b.proto.GetSlotLeaderRequest\x1a\x1c.proto.GetSlotLeaderResponse\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11/v1/getSlotLeader\x12T\n" +
	"\x13GetSlotLeaderStream\x12\x1b.proto.GetSlotLeaderRequest\x1a\x1c.proto.GetSlotLeaderResponse\"\x000\x01B\tZ\a./protob\x06proto3"

var (
	file_proto_solana_proto_rawDescOnce sync.Once
	file_proto_solana_proto_rawDescData []byte
)

func file_proto_solana_proto_rawDescGZIP() []byte {
	file_proto_solana_proto_rawDescOnce.Do(func() {
		file_proto_solana_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_solana_proto_rawDesc), len(file_proto_solana_proto_rawDesc)))
	})
	return file_proto_solana_proto_rawDescData
}

var file_proto_solana_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_solana_proto_goTypes = []any{
	(*GetLatestBlockHashRequest)(nil),  // 0: proto.GetLatestBlockHashRequest
	(*GetLatestBlockHashResponse)(nil), // 1: proto.GetLatestBlockHashResponse
	(*GetAccountBalanceRequest)(nil),   // 2: proto.GetAccountBalanceRequest
	(*GetAccountBalanceResponse)(nil),  // 3: proto.GetAccountBalanceResponse
	(*GetSlotLeaderRequest)(nil),       // 4: proto.GetSlotLeaderRequest
	(*GetSlotLeaderResponse)(nil),      // 5: proto.GetSlotLeaderResponse
}
var file_proto_solana_proto_depIdxs = []int32{
	0, // 0: proto.SolanaService.GetLatestBlockHash:input_type -> proto.GetLatestBlockHashRequest
	2, // 1: proto.SolanaService.GetAccountBalance:input_type -> proto.GetAccountBalanceRequest
	4, // 2: proto.SolanaService.GetSlotLeader:input_type -> proto.GetSlotLeaderRequest
	4, // 3: proto.SolanaService.GetSlotLeaderStream:input_type -> proto.GetSlotLeaderRequest
	1, // 4: proto.SolanaService.GetLatestBlockHash:output_type -> proto.GetLatestBlockHashResponse
	3, // 5: proto.SolanaService.GetAccountBalance:output_type -> proto.GetAccountBalanceResponse
	5, // 6: proto.SolanaService.GetSlotLeader:output_type -> proto.GetSlotLeaderResponse
	5, // 7: proto.SolanaService.GetSlotLeaderStream:output_type -> proto.GetSlotLeaderResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_solana_proto_init() }
func file_proto_solana_proto_init() {
	if File_proto_solana_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_solana_proto_rawDesc), len(file_proto_solana_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_solana_proto_goTypes,
		DependencyIndexes: file_proto_solana_proto_depIdxs,
		MessageInfos:      file_proto_solana_proto_msgTypes,
	}.Build()
	File_proto_solana_proto = out.File
	file_proto_solana_proto_goTypes = nil
	file_proto_solana_proto_depIdxs = nil
}
