// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ghostcloud/ghostcloud/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateDeploymentRequest struct {
	Meta    *Meta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Payload *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgCreateDeploymentRequest) Reset()         { *m = MsgCreateDeploymentRequest{} }
func (m *MsgCreateDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDeploymentRequest) ProtoMessage()    {}
func (*MsgCreateDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad6ede0eb448cbc, []int{0}
}
func (m *MsgCreateDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDeploymentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDeploymentRequest.Merge(m, src)
}
func (m *MsgCreateDeploymentRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDeploymentRequest proto.InternalMessageInfo

func (m *MsgCreateDeploymentRequest) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *MsgCreateDeploymentRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgCreateDeploymentResponse struct {
}

func (m *MsgCreateDeploymentResponse) Reset()         { *m = MsgCreateDeploymentResponse{} }
func (m *MsgCreateDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDeploymentResponse) ProtoMessage()    {}
func (*MsgCreateDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad6ede0eb448cbc, []int{1}
}
func (m *MsgCreateDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDeploymentResponse.Merge(m, src)
}
func (m *MsgCreateDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDeploymentResponse proto.InternalMessageInfo

type MsgUpdateDeploymentRequest struct {
	Meta    *Meta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Payload *Payload `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgUpdateDeploymentRequest) Reset()         { *m = MsgUpdateDeploymentRequest{} }
func (m *MsgUpdateDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDeploymentRequest) ProtoMessage()    {}
func (*MsgUpdateDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad6ede0eb448cbc, []int{2}
}
func (m *MsgUpdateDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDeploymentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDeploymentRequest.Merge(m, src)
}
func (m *MsgUpdateDeploymentRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDeploymentRequest proto.InternalMessageInfo

func (m *MsgUpdateDeploymentRequest) GetMeta() *Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *MsgUpdateDeploymentRequest) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgUpdateDeploymentResponse struct {
}

func (m *MsgUpdateDeploymentResponse) Reset()         { *m = MsgUpdateDeploymentResponse{} }
func (m *MsgUpdateDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateDeploymentResponse) ProtoMessage()    {}
func (*MsgUpdateDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dad6ede0eb448cbc, []int{3}
}
func (m *MsgUpdateDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateDeploymentResponse.Merge(m, src)
}
func (m *MsgUpdateDeploymentResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateDeploymentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateDeploymentRequest)(nil), "ghostcloud.ghostcloud.MsgCreateDeploymentRequest")
	proto.RegisterType((*MsgCreateDeploymentResponse)(nil), "ghostcloud.ghostcloud.MsgCreateDeploymentResponse")
	proto.RegisterType((*MsgUpdateDeploymentRequest)(nil), "ghostcloud.ghostcloud.MsgUpdateDeploymentRequest")
	proto.RegisterType((*MsgUpdateDeploymentResponse)(nil), "ghostcloud.ghostcloud.MsgUpdateDeploymentResponse")
}

func init() { proto.RegisterFile("ghostcloud/ghostcloud/tx.proto", fileDescriptor_dad6ede0eb448cbc) }

var fileDescriptor_dad6ede0eb448cbc = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xc8, 0x2f,
	0x2e, 0x49, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x47, 0x62, 0x96, 0x54, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x89, 0x22, 0x04, 0xf5, 0x10, 0x4c, 0x29, 0x05, 0xec, 0xda, 0x72, 0x53, 0x4b, 0x12,
	0x21, 0x1a, 0xa5, 0x94, 0xb1, 0xab, 0x28, 0x48, 0xac, 0xcc, 0xc9, 0x4f, 0x4c, 0x81, 0x28, 0x52,
	0x6a, 0x67, 0xe4, 0x92, 0xf2, 0x2d, 0x4e, 0x77, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x49, 0x2d,
	0xc8, 0xc9, 0xaf, 0xcc, 0x4d, 0xcd, 0x2b, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xd2,
	0xe7, 0x62, 0x01, 0x99, 0x28, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xad, 0x87, 0xd5, 0x2d,
	0x7a, 0xbe, 0xa9, 0x25, 0x89, 0x41, 0x60, 0x85, 0x42, 0x16, 0x5c, 0xec, 0x50, 0x0b, 0x24, 0x98,
	0xc0, 0x7a, 0xe4, 0x70, 0xe8, 0x09, 0x80, 0xa8, 0x0a, 0x82, 0x29, 0x57, 0x92, 0xe5, 0x92, 0xc6,
	0xea, 0x90, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x98, 0x43, 0x43, 0x0b, 0x52, 0x06, 0x87, 0x43,
	0x31, 0x1d, 0x02, 0x71, 0xa8, 0xd1, 0x6f, 0x46, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x4a, 0x2e,
	0x01, 0x74, 0xcf, 0x08, 0x19, 0xe2, 0x72, 0x17, 0xce, 0x18, 0x90, 0x32, 0x22, 0x45, 0x0b, 0xc4,
	0x09, 0x20, 0xab, 0xd1, 0x9d, 0x87, 0xcf, 0x6a, 0x1c, 0x61, 0x8a, 0xcf, 0x6a, 0x5c, 0xbe, 0x77,
	0x32, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x59, 0xa4, 0x34, 0x58,
	0x81, 0x92, 0xd2, 0x2b, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xe9, 0xd1, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xe5, 0x7c, 0x41, 0xce, 0x0f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateDeployment(ctx context.Context, in *MsgCreateDeploymentRequest, opts ...grpc.CallOption) (*MsgCreateDeploymentResponse, error)
	UpdateDeployment(ctx context.Context, in *MsgUpdateDeploymentRequest, opts ...grpc.CallOption) (*MsgUpdateDeploymentResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDeployment(ctx context.Context, in *MsgCreateDeploymentRequest, opts ...grpc.CallOption) (*MsgCreateDeploymentResponse, error) {
	out := new(MsgCreateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/ghostcloud.ghostcloud.Msg/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDeployment(ctx context.Context, in *MsgUpdateDeploymentRequest, opts ...grpc.CallOption) (*MsgUpdateDeploymentResponse, error) {
	out := new(MsgUpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/ghostcloud.ghostcloud.Msg/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateDeployment(context.Context, *MsgCreateDeploymentRequest) (*MsgCreateDeploymentResponse, error)
	UpdateDeployment(context.Context, *MsgUpdateDeploymentRequest) (*MsgUpdateDeploymentResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateDeployment(ctx context.Context, req *MsgCreateDeploymentRequest) (*MsgCreateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployment not implemented")
}
func (*UnimplementedMsgServer) UpdateDeployment(ctx context.Context, req *MsgUpdateDeploymentRequest) (*MsgUpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployment not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghostcloud.ghostcloud.Msg/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDeployment(ctx, req.(*MsgCreateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ghostcloud.ghostcloud.Msg/UpdateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDeployment(ctx, req.(*MsgUpdateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ghostcloud.ghostcloud.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeployment",
			Handler:    _Msg_CreateDeployment_Handler,
		},
		{
			MethodName: "UpdateDeployment",
			Handler:    _Msg_UpdateDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ghostcloud/ghostcloud/tx.proto",
}

func (m *MsgCreateDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDeploymentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDeploymentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Payload != nil {
		{
			size, err := m.Payload.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Meta != nil {
		{
			size, err := m.Meta.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateDeploymentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateDeploymentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Meta != nil {
		l = m.Meta.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Payload != nil {
		l = m.Payload.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Payload{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = &Meta{}
			}
			if err := m.Meta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payload == nil {
				m.Payload = &Payload{}
			}
			if err := m.Payload.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
