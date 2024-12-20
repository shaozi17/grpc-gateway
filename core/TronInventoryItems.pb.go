// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/TronInventoryItems.proto

package core // import "github.com/shaozi17/grpc-gateway/core"

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

type InventoryItems struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Items                [][]byte `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InventoryItems) Reset()         { *m = InventoryItems{} }
func (m *InventoryItems) String() string { return proto.CompactTextString(m) }
func (*InventoryItems) ProtoMessage()    {}
func (*InventoryItems) Descriptor() ([]byte, []int) {
	return fileDescriptor_TronInventoryItems_0c712d51e1476087, []int{0}
}
func (m *InventoryItems) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InventoryItems.Unmarshal(m, b)
}
func (m *InventoryItems) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InventoryItems.Marshal(b, m, deterministic)
}
func (dst *InventoryItems) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InventoryItems.Merge(dst, src)
}
func (m *InventoryItems) XXX_Size() int {
	return xxx_messageInfo_InventoryItems.Size(m)
}
func (m *InventoryItems) XXX_DiscardUnknown() {
	xxx_messageInfo_InventoryItems.DiscardUnknown(m)
}

var xxx_messageInfo_InventoryItems proto.InternalMessageInfo

func (m *InventoryItems) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *InventoryItems) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*InventoryItems)(nil), "protocol.InventoryItems")
}

func init() {
	proto.RegisterFile("core/TronInventoryItems.proto", fileDescriptor_TronInventoryItems_0c712d51e1476087)
}

var fileDescriptor_TronInventoryItems_0c712d51e1476087 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2f, 0x4a,
	0xd5, 0x0f, 0x29, 0xca, 0xcf, 0xf3, 0xcc, 0x2b, 0x4b, 0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xf4, 0x2c,
	0x49, 0xcd, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39,
	0x4a, 0x56, 0x5c, 0x7c, 0xa8, 0x2a, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x4c, 0x90, 0xa4, 0x04, 0x93,
	0x02, 0xb3, 0x06, 0x4f, 0x10, 0x84, 0xe3, 0x14, 0xc0, 0xc5, 0x9f, 0x5f, 0x94, 0xae, 0x57, 0x52,
	0x94, 0x9f, 0x07, 0x31, 0xb7, 0xd8, 0x49, 0x08, 0xd3, 0xca, 0x28, 0xcd, 0xf4, 0xcc, 0x92, 0x8c,
	0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x90, 0x5a, 0x98, 0xdd, 0xfa, 0xe9, 0x45, 0x05, 0xc9,
	0xba, 0xe9, 0x89, 0x25, 0xa9, 0xe5, 0x89, 0x95, 0xfa, 0x20, 0x07, 0x27, 0xb1, 0x81, 0xe5, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0xdc, 0x30, 0x0d, 0xbf, 0x00, 0x00, 0x00,
}
