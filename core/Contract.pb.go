// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Contract.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core/Contract.proto
	core/Discover.proto
	core/Tron.proto
	core/TronInventoryItems.proto

It has these top-level messages:
	AccountCreateContract
	AccountUpdateContract
	TransferContract
	TransferAssetContract
	VoteAssetContract
	VoteWitnessContract
	WitnessCreateContract
	WitnessUpdateContract
	AssetIssueContract
	ParticipateAssetIssueContract
	DeployContract
	Endpoint
	PingMessage
	PongMessage
	FindNeighbours
	Neighbours
	Account
	Witness
	TXOutput
	TXInput
	TXOutputs
	Transaction
	BlockHeader
	Block
	ChainInventory
	BlockInventory
	Inventory
	Items
	DynamicProperties
	DisconnectMessage
	HelloMessage
	InventoryItems
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountCreateContract struct {
	Type         AccountType `protobuf:"varint,1,opt,name=type,enum=protocol.AccountType" json:"type,omitempty"`
	AccountName  []byte      `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	OwnerAddress []byte      `protobuf:"bytes,3,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *AccountCreateContract) Reset()                    { *m = AccountCreateContract{} }
func (m *AccountCreateContract) String() string            { return proto.CompactTextString(m) }
func (*AccountCreateContract) ProtoMessage()               {}
func (*AccountCreateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountCreateContract) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_Normal
}

func (m *AccountCreateContract) GetAccountName() []byte {
	if m != nil {
		return m.AccountName
	}
	return nil
}

func (m *AccountCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

// update account name if the account has no name.
type AccountUpdateContract struct {
	AccountName  []byte `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *AccountUpdateContract) Reset()                    { *m = AccountUpdateContract{} }
func (m *AccountUpdateContract) String() string            { return proto.CompactTextString(m) }
func (*AccountUpdateContract) ProtoMessage()               {}
func (*AccountUpdateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccountUpdateContract) GetAccountName() []byte {
	if m != nil {
		return m.AccountName
	}
	return nil
}

func (m *AccountUpdateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type TransferContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferContract) Reset()                    { *m = TransferContract{} }
func (m *TransferContract) String() string            { return proto.CompactTextString(m) }
func (*TransferContract) ProtoMessage()               {}
func (*TransferContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransferContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransferAssetContract struct {
	AssetName    []byte `protobuf:"bytes,1,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferAssetContract) Reset()                    { *m = TransferAssetContract{} }
func (m *TransferAssetContract) String() string            { return proto.CompactTextString(m) }
func (*TransferAssetContract) ProtoMessage()               {}
func (*TransferAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransferAssetContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *TransferAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferAssetContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferAssetContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteAssetContract struct {
	OwnerAddress []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	VoteAddress  [][]byte `protobuf:"bytes,2,rep,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	Support      bool     `protobuf:"varint,3,opt,name=support" json:"support,omitempty"`
	Count        int32    `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
}

func (m *VoteAssetContract) Reset()                    { *m = VoteAssetContract{} }
func (m *VoteAssetContract) String() string            { return proto.CompactTextString(m) }
func (*VoteAssetContract) ProtoMessage()               {}
func (*VoteAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VoteAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *VoteAssetContract) GetVoteAddress() [][]byte {
	if m != nil {
		return m.VoteAddress
	}
	return nil
}

func (m *VoteAssetContract) GetSupport() bool {
	if m != nil {
		return m.Support
	}
	return false
}

func (m *VoteAssetContract) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VoteWitnessContract struct {
	OwnerAddress []byte                      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Votes        []*VoteWitnessContract_Vote `protobuf:"bytes,2,rep,name=votes" json:"votes,omitempty"`
	Support      bool                        `protobuf:"varint,3,opt,name=support" json:"support,omitempty"`
}

func (m *VoteWitnessContract) Reset()                    { *m = VoteWitnessContract{} }
func (m *VoteWitnessContract) String() string            { return proto.CompactTextString(m) }
func (*VoteWitnessContract) ProtoMessage()               {}
func (*VoteWitnessContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VoteWitnessContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *VoteWitnessContract) GetVotes() []*VoteWitnessContract_Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *VoteWitnessContract) GetSupport() bool {
	if m != nil {
		return m.Support
	}
	return false
}

type VoteWitnessContract_Vote struct {
	VoteAddress []byte `protobuf:"bytes,1,opt,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	VoteCount   int64  `protobuf:"varint,2,opt,name=vote_count,json=voteCount" json:"vote_count,omitempty"`
}

func (m *VoteWitnessContract_Vote) Reset()                    { *m = VoteWitnessContract_Vote{} }
func (m *VoteWitnessContract_Vote) String() string            { return proto.CompactTextString(m) }
func (*VoteWitnessContract_Vote) ProtoMessage()               {}
func (*VoteWitnessContract_Vote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *VoteWitnessContract_Vote) GetVoteAddress() []byte {
	if m != nil {
		return m.VoteAddress
	}
	return nil
}

func (m *VoteWitnessContract_Vote) GetVoteCount() int64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

type WitnessCreateContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Url          []byte `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *WitnessCreateContract) Reset()                    { *m = WitnessCreateContract{} }
func (m *WitnessCreateContract) String() string            { return proto.CompactTextString(m) }
func (*WitnessCreateContract) ProtoMessage()               {}
func (*WitnessCreateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WitnessCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessCreateContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

type WitnessUpdateContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	UpdateUrl    []byte `protobuf:"bytes,12,opt,name=update_url,json=updateUrl,proto3" json:"update_url,omitempty"`
}

func (m *WitnessUpdateContract) Reset()                    { *m = WitnessUpdateContract{} }
func (m *WitnessUpdateContract) String() string            { return proto.CompactTextString(m) }
func (*WitnessUpdateContract) ProtoMessage()               {}
func (*WitnessUpdateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WitnessUpdateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessUpdateContract) GetUpdateUrl() []byte {
	if m != nil {
		return m.UpdateUrl
	}
	return nil
}

type AssetIssueContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Name         []byte `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TotalSupply  int64  `protobuf:"varint,4,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	TrxNum       int32  `protobuf:"varint,6,opt,name=trx_num,json=trxNum" json:"trx_num,omitempty"`
	Num          int32  `protobuf:"varint,8,opt,name=num" json:"num,omitempty"`
	StartTime    int64  `protobuf:"varint,9,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime      int64  `protobuf:"varint,10,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	DecayRatio   int32  `protobuf:"varint,15,opt,name=decay_ratio,json=decayRatio" json:"decay_ratio,omitempty"`
	VoteScore    int32  `protobuf:"varint,16,opt,name=vote_score,json=voteScore" json:"vote_score,omitempty"`
	Description  []byte `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	Url          []byte `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *AssetIssueContract) Reset()                    { *m = AssetIssueContract{} }
func (m *AssetIssueContract) String() string            { return proto.CompactTextString(m) }
func (*AssetIssueContract) ProtoMessage()               {}
func (*AssetIssueContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *AssetIssueContract) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AssetIssueContract) GetTotalSupply() int64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *AssetIssueContract) GetTrxNum() int32 {
	if m != nil {
		return m.TrxNum
	}
	return 0
}

func (m *AssetIssueContract) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *AssetIssueContract) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AssetIssueContract) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AssetIssueContract) GetDecayRatio() int32 {
	if m != nil {
		return m.DecayRatio
	}
	return 0
}

func (m *AssetIssueContract) GetVoteScore() int32 {
	if m != nil {
		return m.VoteScore
	}
	return 0
}

func (m *AssetIssueContract) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *AssetIssueContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

type ParticipateAssetIssueContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetName    []byte `protobuf:"bytes,3,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	Amount       int64  `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *ParticipateAssetIssueContract) Reset()                    { *m = ParticipateAssetIssueContract{} }
func (m *ParticipateAssetIssueContract) String() string            { return proto.CompactTextString(m) }
func (*ParticipateAssetIssueContract) ProtoMessage()               {}
func (*ParticipateAssetIssueContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ParticipateAssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type DeployContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Script       []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
}

func (m *DeployContract) Reset()                    { *m = DeployContract{} }
func (m *DeployContract) String() string            { return proto.CompactTextString(m) }
func (*DeployContract) ProtoMessage()               {}
func (*DeployContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeployContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *DeployContract) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountCreateContract)(nil), "protocol.AccountCreateContract")
	proto.RegisterType((*AccountUpdateContract)(nil), "protocol.AccountUpdateContract")
	proto.RegisterType((*TransferContract)(nil), "protocol.TransferContract")
	proto.RegisterType((*TransferAssetContract)(nil), "protocol.TransferAssetContract")
	proto.RegisterType((*VoteAssetContract)(nil), "protocol.VoteAssetContract")
	proto.RegisterType((*VoteWitnessContract)(nil), "protocol.VoteWitnessContract")
	proto.RegisterType((*VoteWitnessContract_Vote)(nil), "protocol.VoteWitnessContract.Vote")
	proto.RegisterType((*WitnessCreateContract)(nil), "protocol.WitnessCreateContract")
	proto.RegisterType((*WitnessUpdateContract)(nil), "protocol.WitnessUpdateContract")
	proto.RegisterType((*AssetIssueContract)(nil), "protocol.AssetIssueContract")
	proto.RegisterType((*ParticipateAssetIssueContract)(nil), "protocol.ParticipateAssetIssueContract")
	proto.RegisterType((*DeployContract)(nil), "protocol.DeployContract")
}

func init() { proto.RegisterFile("core/Contract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0x96, 0x73, 0x6b, 0x72, 0x92, 0xbf, 0xed, 0xef, 0x36, 0xc5, 0x54, 0x8a, 0x48, 0xcd, 0xa6,
	0x5d, 0x90, 0x48, 0x65, 0xc3, 0xb6, 0x2d, 0x42, 0xb0, 0xa0, 0x42, 0x6e, 0x0a, 0x12, 0x2c, 0xac,
	0xa9, 0x33, 0x04, 0x4b, 0xf1, 0x8c, 0x35, 0x73, 0x4c, 0x9b, 0x57, 0x80, 0x1d, 0x4b, 0xde, 0x88,
	0x37, 0xe0, 0x71, 0xd0, 0x9c, 0xb1, 0x53, 0x27, 0x6d, 0x45, 0xc2, 0xca, 0x33, 0xdf, 0xb9, 0x7d,
	0xe7, 0x36, 0x86, 0x9d, 0x48, 0x2a, 0x3e, 0x3c, 0x93, 0x02, 0x15, 0x8b, 0x70, 0x90, 0x2a, 0x89,
	0xd2, 0x6d, 0xd2, 0x27, 0x92, 0xd3, 0xfd, 0x2d, 0x12, 0x8f, 0x94, 0x14, 0x56, 0xe4, 0x7f, 0x73,
	0xa0, 0x7b, 0x12, 0x45, 0x32, 0x13, 0x78, 0xa6, 0x38, 0x43, 0x5e, 0x98, 0xba, 0x47, 0x50, 0xc3,
	0x59, 0xca, 0x3d, 0xa7, 0xef, 0x1c, 0x6e, 0x1e, 0x77, 0x07, 0x85, 0x8f, 0x41, 0xae, 0x3e, 0x9a,
	0xa5, 0x3c, 0x20, 0x15, 0xf7, 0x00, 0x3a, 0xcc, 0x82, 0xa1, 0x60, 0x09, 0xf7, 0x2a, 0x7d, 0xe7,
	0xb0, 0x13, 0xb4, 0x73, 0xec, 0x9c, 0x25, 0xdc, 0x7d, 0x0a, 0xff, 0xc9, 0x6b, 0xc1, 0x55, 0xc8,
	0xc6, 0x63, 0xc5, 0xb5, 0xf6, 0xaa, 0xa4, 0xd3, 0x21, 0xf0, 0xc4, 0x62, 0x7e, 0x38, 0xe7, 0x72,
	0x99, 0x8e, 0xcb, 0x5c, 0x96, 0x03, 0x38, 0x2b, 0x04, 0xa8, 0xdc, 0x13, 0x40, 0xc0, 0xf6, 0x48,
	0x31, 0xa1, 0x3f, 0x73, 0x35, 0xf7, 0x7d, 0xc7, 0xd0, 0xb9, 0x6b, 0xe8, 0xf6, 0x00, 0x50, 0x2e,
	0xb9, 0x6e, 0xa1, 0x2c, 0xc4, 0x7b, 0xd0, 0x60, 0x89, 0xa1, 0x42, 0x69, 0x55, 0x83, 0xfc, 0xe6,
	0xff, 0x70, 0xa0, 0x5b, 0x04, 0x3c, 0xd1, 0x9a, 0xe3, 0x3c, 0x6a, 0x0f, 0x80, 0x19, 0xa0, 0x9c,
	0x4f, 0x8b, 0x90, 0x95, 0xb3, 0x59, 0x22, 0x55, 0x7d, 0x98, 0x54, 0x6d, 0x81, 0xd4, 0x77, 0x07,
	0xfe, 0x7f, 0x2f, 0x91, 0x2f, 0x12, 0x5a, 0xa9, 0x0c, 0x07, 0xd0, 0xf9, 0x2a, 0x91, 0x97, 0x58,
	0x55, 0x4d, 0x1f, 0x0c, 0x56, 0xa8, 0x78, 0xb0, 0xa1, 0xb3, 0x34, 0x95, 0xca, 0xd6, 0xa2, 0x19,
	0x14, 0x57, 0x77, 0x17, 0xea, 0xd4, 0x2e, 0xaf, 0xde, 0x77, 0x0e, 0xeb, 0x81, 0xbd, 0xf8, 0xbf,
	0x1d, 0xd8, 0x31, 0x6c, 0x3e, 0xc4, 0x28, 0xb8, 0xd6, 0xeb, 0xf1, 0x79, 0x01, 0x75, 0x13, 0xdb,
	0x12, 0x69, 0x1f, 0xfb, 0xb7, 0x43, 0x7a, 0x8f, 0x4b, 0xc2, 0x02, 0x6b, 0xf0, 0x30, 0xcd, 0xfd,
	0xd7, 0x50, 0x33, 0x8a, 0x77, 0x72, 0xcd, 0x67, 0xae, 0x9c, 0x6b, 0x0f, 0x80, 0x54, 0x6c, 0x5a,
	0x15, 0xaa, 0x72, 0xcb, 0x20, 0x67, 0x94, 0xda, 0x39, 0x74, 0x0b, 0x0a, 0x8b, 0xab, 0xb5, 0x52,
	0x6e, 0xdb, 0x50, 0xcd, 0xd4, 0x34, 0x6f, 0xbc, 0x39, 0xfa, 0x9f, 0xe6, 0xfe, 0x96, 0xd6, 0x63,
	0xd5, 0x11, 0xce, 0xc8, 0x2c, 0x34, 0x6e, 0x3b, 0x76, 0x5a, 0x2c, 0x72, 0xa9, 0xa6, 0xfe, 0xaf,
	0x0a, 0xb8, 0x34, 0x11, 0x6f, 0xb4, 0xce, 0xd6, 0x74, 0xed, 0x42, 0xad, 0xb4, 0xf7, 0x74, 0x36,
	0xe5, 0x43, 0x89, 0x6c, 0x1a, 0x9a, 0xba, 0x4e, 0x67, 0xf9, 0x0c, 0xb6, 0x09, 0xbb, 0x20, 0xc8,
	0x7d, 0x04, 0x1b, 0xa8, 0x6e, 0x42, 0x91, 0x25, 0x5e, 0x83, 0x46, 0xa2, 0x81, 0xea, 0xe6, 0x3c,
	0x4b, 0x4c, 0xea, 0x06, 0x6c, 0x12, 0x68, 0x8e, 0x86, 0xbc, 0x46, 0xa6, 0x30, 0xc4, 0x38, 0xe1,
	0x5e, 0xcb, 0x56, 0x9a, 0x90, 0x51, 0x9c, 0x70, 0xf7, 0x31, 0x34, 0xb9, 0x18, 0x5b, 0x21, 0x90,
	0x70, 0x83, 0x8b, 0x31, 0x89, 0x9e, 0x40, 0x7b, 0xcc, 0x23, 0x36, 0x0b, 0x15, 0xc3, 0x58, 0x7a,
	0x5b, 0xe4, 0x13, 0x08, 0x0a, 0x0c, 0x32, 0x6f, 0xa2, 0x36, 0x4f, 0xa3, 0xb7, 0x4d, 0x72, 0x6a,
	0xe2, 0x85, 0x01, 0xdc, 0xbe, 0xb1, 0xd7, 0x91, 0x8a, 0x53, 0x8c, 0xa5, 0xf0, 0x76, 0xed, 0x14,
	0x94, 0xa0, 0xa2, 0x51, 0xdd, 0xdb, 0x46, 0xfd, 0x74, 0xa0, 0xf7, 0x8e, 0x29, 0x8c, 0xa3, 0x38,
	0x65, 0xf9, 0xa2, 0xfd, 0x43, 0x59, 0xff, 0xf2, 0xe8, 0x2c, 0x3e, 0x21, 0xd5, 0xe5, 0x27, 0xe4,
	0xa1, 0xf5, 0x7f, 0x0b, 0x9b, 0x2f, 0x79, 0x3a, 0x95, 0xb3, 0xf5, 0xc8, 0xec, 0x41, 0xc3, 0xa6,
	0x9c, 0x13, 0xc9, 0x6f, 0xa7, 0xaf, 0x60, 0x4b, 0xaa, 0xc9, 0x00, 0xe7, 0xbf, 0x14, 0x7d, 0xda,
	0x2c, 0x3c, 0x7f, 0x3c, 0x9a, 0xc4, 0xf8, 0x25, 0xbb, 0x1a, 0x44, 0x32, 0x19, 0x1a, 0x8d, 0x62,
	0x3d, 0x87, 0x13, 0x95, 0x46, 0xcf, 0x26, 0x0c, 0xf9, 0x35, 0x9b, 0x0d, 0x4d, 0x9d, 0xaf, 0x1a,
	0x24, 0x7b, 0xfe, 0x27, 0x00, 0x00, 0xff, 0xff, 0x42, 0xab, 0xca, 0xa1, 0xc1, 0x06, 0x00, 0x00,
}