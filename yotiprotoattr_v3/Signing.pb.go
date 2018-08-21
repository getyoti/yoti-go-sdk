// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Signing.proto

package yotiprotoattr_v3

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

type AttributeSigning struct {
	Name                 string      `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value                []byte      `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	ContentType          ContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,enum=Yoti.Protobuf.attrpubapi_v3.ContentType" json:"content_type,omitempty"`
	ArtifactSignature    []byte      `protobuf:"bytes,4,opt,name=artifact_signature,json=artifactSignature,proto3" json:"artifact_signature,omitempty"`
	SubType              string      `protobuf:"bytes,5,opt,name=sub_type,json=subType" json:"sub_type,omitempty"`
	SignedTimeStamp      []byte      `protobuf:"bytes,6,opt,name=signed_time_stamp,json=signedTimeStamp,proto3" json:"signed_time_stamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttributeSigning) Reset()         { *m = AttributeSigning{} }
func (m *AttributeSigning) String() string { return proto.CompactTextString(m) }
func (*AttributeSigning) ProtoMessage()    {}
func (*AttributeSigning) Descriptor() ([]byte, []int) {
	return fileDescriptor_Signing_d68001faa0a6803c, []int{0}
}
func (m *AttributeSigning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeSigning.Unmarshal(m, b)
}
func (m *AttributeSigning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeSigning.Marshal(b, m, deterministic)
}
func (dst *AttributeSigning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeSigning.Merge(dst, src)
}
func (m *AttributeSigning) XXX_Size() int {
	return xxx_messageInfo_AttributeSigning.Size(m)
}
func (m *AttributeSigning) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeSigning.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeSigning proto.InternalMessageInfo

func (m *AttributeSigning) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttributeSigning) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttributeSigning) GetContentType() ContentType {
	if m != nil {
		return m.ContentType
	}
	return ContentType_UNDEFINED
}

func (m *AttributeSigning) GetArtifactSignature() []byte {
	if m != nil {
		return m.ArtifactSignature
	}
	return nil
}

func (m *AttributeSigning) GetSubType() string {
	if m != nil {
		return m.SubType
	}
	return ""
}

func (m *AttributeSigning) GetSignedTimeStamp() []byte {
	if m != nil {
		return m.SignedTimeStamp
	}
	return nil
}

func init() {
	proto.RegisterType((*AttributeSigning)(nil), "Yoti.Protobuf.attrpubapi_v3.AttributeSigning")
}

func init() { proto.RegisterFile("Signing.proto", fileDescriptor_Signing_d68001faa0a6803c) }

var fileDescriptor_Signing_d68001faa0a6803c = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x68, 0x0b, 0x35, 0x85, 0xb6, 0x16, 0x42, 0x01, 0x96, 0x88, 0x29, 0x42, 0xc2,
	0x03, 0x11, 0x3f, 0x80, 0x32, 0xb2, 0xa0, 0xa4, 0x0b, 0x2c, 0x96, 0x1d, 0xae, 0x91, 0x25, 0xe2,
	0x58, 0xce, 0x39, 0x52, 0x7e, 0x3d, 0xc8, 0x36, 0x05, 0xb1, 0xb0, 0xdd, 0xbb, 0x77, 0x7e, 0xf7,
	0xf9, 0xc8, 0x69, 0xa5, 0x1a, 0xad, 0x74, 0xc3, 0x8c, 0xed, 0xb0, 0xa3, 0xd7, 0xaf, 0x1d, 0x2a,
	0xf6, 0xe2, 0x6b, 0xe9, 0x76, 0x4c, 0x20, 0x5a, 0xe3, 0xa4, 0x30, 0x8a, 0x0f, 0xc5, 0xd5, 0xf2,
	0x11, 0xd1, 0x2a, 0xe9, 0x10, 0xe2, 0xf4, 0xcd, 0x67, 0x42, 0x56, 0x3f, 0xbd, 0xef, 0x20, 0x4a,
	0xc9, 0x44, 0x8b, 0x16, 0xd2, 0x24, 0x4b, 0xf2, 0x79, 0x19, 0x6a, 0x7a, 0x4e, 0xa6, 0x83, 0xf8,
	0x70, 0x90, 0x1e, 0x64, 0x49, 0xbe, 0x28, 0xa3, 0xa0, 0xcf, 0x64, 0x51, 0x77, 0x1a, 0x41, 0x23,
	0xc7, 0xd1, 0x40, 0x7a, 0x98, 0x25, 0xf9, 0xd9, 0x7d, 0xce, 0xfe, 0x61, 0x60, 0x4f, 0xf1, 0xc1,
	0x76, 0x34, 0x50, 0x9e, 0xd4, 0xbf, 0x82, 0xde, 0x11, 0x2a, 0x2c, 0xaa, 0x9d, 0xa8, 0x91, 0xf7,
	0xaa, 0xd1, 0x02, 0x9d, 0x85, 0x74, 0x12, 0xf6, 0xad, 0xf7, 0x4e, 0xb5, 0x37, 0xe8, 0x25, 0x39,
	0xee, 0x9d, 0x8c, 0x7b, 0xa7, 0x81, 0xf4, 0xa8, 0x77, 0x32, 0x24, 0xdd, 0x92, 0xb5, 0x0f, 0x80,
	0x77, 0x8e, 0xaa, 0x05, 0xde, 0xa3, 0x68, 0x4d, 0x3a, 0x0b, 0x41, 0xcb, 0x68, 0x6c, 0x55, 0x0b,
	0x95, 0x6f, 0x6f, 0x1e, 0xc8, 0x45, 0xdd, 0xb5, 0x6c, 0xf4, 0xc4, 0x7f, 0x40, 0x37, 0x73, 0x7f,
	0x98, 0xf0, 0x89, 0xb7, 0x95, 0xb7, 0xc3, 0xc5, 0xfc, 0x08, 0x1f, 0x0a, 0x39, 0x0b, 0xaa, 0xf8,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xce, 0x8b, 0x53, 0x3a, 0x7e, 0x01, 0x00, 0x00,
}
