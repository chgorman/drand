// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dkg/dkg.proto

/*
Package dkg is a generated protocol buffer package.

It is generated from these files:
	dkg/dkg.proto

It has these top-level messages:
	DKGPacket
	DKGResponse
	Deal
	Response
	Justification
*/
package dkg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import vss "github.com/dedis/drand/protobuf/crypto/share/vss"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// DKGPacket is used by the nodes to run the dkg protocol before being able to
// generate randomness beacons.
type DKGPacket struct {
	Deal          *Deal          `protobuf:"bytes,1,opt,name=deal" json:"deal,omitempty"`
	Response      *Response      `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
	Justification *Justification `protobuf:"bytes,3,opt,name=justification" json:"justification,omitempty"`
}

func (m *DKGPacket) Reset()                    { *m = DKGPacket{} }
func (m *DKGPacket) String() string            { return proto.CompactTextString(m) }
func (*DKGPacket) ProtoMessage()               {}
func (*DKGPacket) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DKGPacket) GetDeal() *Deal {
	if m != nil {
		return m.Deal
	}
	return nil
}

func (m *DKGPacket) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *DKGPacket) GetJustification() *Justification {
	if m != nil {
		return m.Justification
	}
	return nil
}

type DKGResponse struct {
}

func (m *DKGResponse) Reset()                    { *m = DKGResponse{} }
func (m *DKGResponse) String() string            { return proto.CompactTextString(m) }
func (*DKGResponse) ProtoMessage()               {}
func (*DKGResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Deal contains a share for a participant.
type Deal struct {
	// index of the dealer, the issuer of the share
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// encrypted version of the deal
	Deal *vss.EncryptedDeal `protobuf:"bytes,2,opt,name=deal" json:"deal,omitempty"`
	// signature of the whole deal
	// NOTE: this is almost duplicated data, since the vss deal already includes
	// a signature. However it does not include the index of the dealer that
	// issue this deal, so another one is required. Best would be to merge vss
	// and dkg so we could use only one field of signature. For future work...
	// :)
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Deal) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Deal) GetDeal() *vss.EncryptedDeal {
	if m != nil {
		return m.Deal
	}
	return nil
}

func (m *Deal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// Response holds the response that a participant broadcast after having
// received a deal.
type Response struct {
	// index of the dealer for which this response is for
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// response from the participant which received a deal
	Response *vss.Response `protobuf:"bytes,2,opt,name=response" json:"response,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response) GetResponse() *vss.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

// Justification holds the justification from a dealer after a participant
// issued a complaint response because of a supposedly invalid deal.
type Justification struct {
	// index of the dealer who is issuing this justification
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// justification from the dealer
	Justification *vss.Justification `protobuf:"bytes,2,opt,name=justification" json:"justification,omitempty"`
}

func (m *Justification) Reset()                    { *m = Justification{} }
func (m *Justification) String() string            { return proto.CompactTextString(m) }
func (*Justification) ProtoMessage()               {}
func (*Justification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Justification) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Justification) GetJustification() *vss.Justification {
	if m != nil {
		return m.Justification
	}
	return nil
}

func init() {
	proto.RegisterType((*DKGPacket)(nil), "dkg.DKGPacket")
	proto.RegisterType((*DKGResponse)(nil), "dkg.DKGResponse")
	proto.RegisterType((*Deal)(nil), "dkg.Deal")
	proto.RegisterType((*Response)(nil), "dkg.Response")
	proto.RegisterType((*Justification)(nil), "dkg.Justification")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Dkg service

type DkgClient interface {
	Setup(ctx context.Context, in *DKGPacket, opts ...grpc.CallOption) (*DKGResponse, error)
}

type dkgClient struct {
	cc *grpc.ClientConn
}

func NewDkgClient(cc *grpc.ClientConn) DkgClient {
	return &dkgClient{cc}
}

func (c *dkgClient) Setup(ctx context.Context, in *DKGPacket, opts ...grpc.CallOption) (*DKGResponse, error) {
	out := new(DKGResponse)
	err := grpc.Invoke(ctx, "/dkg.Dkg/Setup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dkg service

type DkgServer interface {
	Setup(context.Context, *DKGPacket) (*DKGResponse, error)
}

func RegisterDkgServer(s *grpc.Server, srv DkgServer) {
	s.RegisterService(&_Dkg_serviceDesc, srv)
}

func _Dkg_Setup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DKGPacket)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DkgServer).Setup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkg.Dkg/Setup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DkgServer).Setup(ctx, req.(*DKGPacket))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dkg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dkg.Dkg",
	HandlerType: (*DkgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Setup",
			Handler:    _Dkg_Setup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dkg/dkg.proto",
}

func init() { proto.RegisterFile("dkg/dkg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4b, 0x33, 0x31,
	0x10, 0xa6, 0x5f, 0x2f, 0xed, 0xf4, 0x8d, 0x48, 0xf0, 0x50, 0x8a, 0x82, 0xac, 0x28, 0xf6, 0xd2,
	0x48, 0xbd, 0x78, 0x96, 0x95, 0x82, 0xbd, 0x48, 0xbc, 0x79, 0x91, 0xec, 0x26, 0x4d, 0xe3, 0xd6,
	0x64, 0x49, 0xb2, 0xa2, 0xbf, 0xc2, 0xbf, 0x2c, 0x49, 0xbb, 0xda, 0x52, 0xeb, 0x21, 0x87, 0x99,
	0x3c, 0x3c, 0x1f, 0x33, 0x03, 0x88, 0x17, 0x92, 0xf0, 0x42, 0x8e, 0x4b, 0x6b, 0xbc, 0xc1, 0x2d,
	0x5e, 0xc8, 0xe1, 0x30, 0xb7, 0x1f, 0xa5, 0x37, 0xc4, 0x2d, 0x98, 0x15, 0xe4, 0xcd, 0xb9, 0xf0,
	0x56, 0x80, 0xe4, 0xb3, 0x01, 0xbd, 0x74, 0x36, 0x7d, 0x60, 0x79, 0x21, 0x3c, 0x3e, 0x81, 0x36,
	0x17, 0x6c, 0x39, 0x68, 0x9c, 0x36, 0x2e, 0xfb, 0x93, 0xde, 0x38, 0x10, 0xa5, 0x82, 0x2d, 0x69,
	0x6c, 0xe3, 0x11, 0x74, 0xad, 0x70, 0xa5, 0xd1, 0x4e, 0x0c, 0x9a, 0x11, 0x82, 0x22, 0x84, 0xae,
	0x9b, 0xf4, 0xfb, 0x1b, 0xdf, 0x00, 0x7a, 0xa9, 0x9c, 0x57, 0x73, 0x95, 0x33, 0xaf, 0x8c, 0x1e,
	0xb4, 0x22, 0x1e, 0x47, 0xfc, 0xfd, 0xe6, 0x0f, 0xdd, 0x06, 0x26, 0x08, 0xfa, 0xe9, 0x6c, 0x5a,
	0x53, 0x26, 0x19, 0xb4, 0x83, 0x03, 0x7c, 0x04, 0x1d, 0xa5, 0xb9, 0x78, 0x8f, 0xde, 0x10, 0x5d,
	0x15, 0xf8, 0x62, 0x6d, 0xb8, 0xb9, 0x66, 0x0f, 0xc1, 0xee, 0x74, 0xcc, 0x2b, 0xf8, 0x86, 0xf3,
	0x63, 0xe8, 0x39, 0x25, 0x35, 0xf3, 0x95, 0x15, 0xd1, 0xca, 0x7f, 0xfa, 0xd3, 0x48, 0x66, 0xd0,
	0xad, 0xf5, 0xf6, 0xe8, 0xfc, 0x96, 0x3c, 0x68, 0xed, 0x26, 0x4f, 0x9e, 0x01, 0x6d, 0xe5, 0xdb,
	0xc3, 0xb8, 0x33, 0xa0, 0xcd, 0x08, 0x7f, 0x0d, 0x68, 0x72, 0x05, 0xad, 0xb4, 0x90, 0x78, 0x04,
	0x9d, 0x47, 0xe1, 0xab, 0x12, 0x1f, 0xac, 0xd6, 0x54, 0x2f, 0x71, 0x78, 0x58, 0xd7, 0xb5, 0xb9,
	0xdb, 0xf3, 0xa7, 0x33, 0xa9, 0xfc, 0xa2, 0xca, 0xc6, 0xb9, 0x79, 0x25, 0x5c, 0x70, 0xe5, 0x08,
	0xb7, 0x4c, 0x73, 0x12, 0x8f, 0x20, 0xab, 0xe6, 0xe1, 0x64, 0xb2, 0x7f, 0xb1, 0xba, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0xb2, 0xeb, 0xb8, 0x54, 0x44, 0x02, 0x00, 0x00,
}