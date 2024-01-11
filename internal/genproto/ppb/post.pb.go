// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post.proto

package ppb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ID struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{0}
}

func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (m *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(m, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type PostResponse struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId               int32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body                 string   `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostResponse) Reset()         { *m = PostResponse{} }
func (m *PostResponse) String() string { return proto.CompactTextString(m) }
func (*PostResponse) ProtoMessage()    {}
func (*PostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{1}
}

func (m *PostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostResponse.Unmarshal(m, b)
}
func (m *PostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostResponse.Marshal(b, m, deterministic)
}
func (m *PostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostResponse.Merge(m, src)
}
func (m *PostResponse) XXX_Size() int {
	return xxx_messageInfo_PostResponse.Size(m)
}
func (m *PostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostResponse proto.InternalMessageInfo

func (m *PostResponse) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *PostResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *PostResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PostResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*ID)(nil), "post.ID")
	proto.RegisterType((*PostResponse)(nil), "post.PostResponse")
}

func init() {
	proto.RegisterFile("proto/post.proto", fileDescriptor_5e57bba97c118b83)
}

var fileDescriptor_5e57bba97c118b83 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0x2f, 0x2e, 0xd1, 0x03, 0x33, 0x85, 0x58, 0x40, 0x6c, 0x25, 0x11, 0x2e,
	0x26, 0x4f, 0x17, 0x21, 0x3e, 0x10, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0xe4, 0xe9,
	0xa2, 0x94, 0xc8, 0xc5, 0x13, 0x90, 0x5f, 0x5c, 0x12, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x8a, 0x2e, 0x2f, 0x24, 0xce, 0xc5, 0x5e, 0x5a, 0x9c, 0x5a, 0x14, 0x9f, 0x99, 0x22, 0xc1, 0x04,
	0x16, 0x64, 0x03, 0x71, 0x3d, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0x32, 0x4b, 0x72, 0x52, 0x25,
	0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x21, 0x2e, 0x96, 0xa4, 0xfc, 0x94, 0x4a,
	0x09, 0x16, 0xb0, 0x20, 0x98, 0x6d, 0x64, 0xc3, 0xc5, 0x0d, 0xb2, 0x22, 0x38, 0xb5, 0xa8, 0x2c,
	0x33, 0x39, 0x55, 0x48, 0x97, 0x8b, 0xdb, 0x3d, 0xb5, 0x04, 0x24, 0xe2, 0x54, 0xe9, 0xe9, 0x22,
	0xc4, 0xa1, 0x07, 0x76, 0xa9, 0xa7, 0x8b, 0x94, 0x10, 0x84, 0x85, 0xec, 0x1c, 0x25, 0x06, 0x27,
	0xb6, 0x28, 0x16, 0xfd, 0x82, 0x82, 0xa4, 0x24, 0x36, 0xb0, 0x5f, 0x8c, 0x01, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x4a, 0x76, 0x90, 0x3d, 0xdf, 0x00, 0x00, 0x00,
}