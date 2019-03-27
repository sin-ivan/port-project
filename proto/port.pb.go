// Code generated by protoc-gen-go. DO NOT EDIT.
// source: port.proto

package proto_grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type TextResponse struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextResponse) Reset()         { *m = TextResponse{} }
func (m *TextResponse) String() string { return proto.CompactTextString(m) }
func (*TextResponse) ProtoMessage()    {}
func (*TextResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{1}
}

func (m *TextResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextResponse.Unmarshal(m, b)
}
func (m *TextResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextResponse.Marshal(b, m, deterministic)
}
func (m *TextResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextResponse.Merge(m, src)
}
func (m *TextResponse) XXX_Size() int {
	return xxx_messageInfo_TextResponse.Size(m)
}
func (m *TextResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TextResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TextResponse proto.InternalMessageInfo

func (m *TextResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TextResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DeleteResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{2}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *DeleteResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type ListPort struct {
	Ports                []*Port  `protobuf:"bytes,1,rep,name=ports,proto3" json:"ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPort) Reset()         { *m = ListPort{} }
func (m *ListPort) String() string { return proto.CompactTextString(m) }
func (*ListPort) ProtoMessage()    {}
func (*ListPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{3}
}

func (m *ListPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPort.Unmarshal(m, b)
}
func (m *ListPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPort.Marshal(b, m, deterministic)
}
func (m *ListPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPort.Merge(m, src)
}
func (m *ListPort) XXX_Size() int {
	return xxx_messageInfo_ListPort.Size(m)
}
func (m *ListPort) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPort.DiscardUnknown(m)
}

var xxx_messageInfo_ListPort proto.InternalMessageInfo

func (m *ListPort) GetPorts() []*Port {
	if m != nil {
		return m.Ports
	}
	return nil
}

type SingleRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SingleRequest) Reset()         { *m = SingleRequest{} }
func (m *SingleRequest) String() string { return proto.CompactTextString(m) }
func (*SingleRequest) ProtoMessage()    {}
func (*SingleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{4}
}

func (m *SingleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SingleRequest.Unmarshal(m, b)
}
func (m *SingleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SingleRequest.Marshal(b, m, deterministic)
}
func (m *SingleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SingleRequest.Merge(m, src)
}
func (m *SingleRequest) XXX_Size() int {
	return xxx_messageInfo_SingleRequest.Size(m)
}
func (m *SingleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SingleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SingleRequest proto.InternalMessageInfo

func (m *SingleRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Port struct {
	ID                   string    `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string    `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string    `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Alias                []string  `protobuf:"bytes,5,rep,name=alias,proto3" json:"alias,omitempty"`
	Regions              []string  `protobuf:"bytes,6,rep,name=regions,proto3" json:"regions,omitempty"`
	Coordinates          []float32 `protobuf:"fixed32,7,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
	Province             string    `protobuf:"bytes,8,opt,name=province,proto3" json:"province,omitempty"`
	Timezone             string    `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Unlocs               []string  `protobuf:"bytes,10,rep,name=unlocs,proto3" json:"unlocs,omitempty"`
	Code                 string    `protobuf:"bytes,11,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Port) Reset()         { *m = Port{} }
func (m *Port) String() string { return proto.CompactTextString(m) }
func (*Port) ProtoMessage()    {}
func (*Port) Descriptor() ([]byte, []int) {
	return fileDescriptor_729c3d36e9010a8e, []int{5}
}

func (m *Port) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Port.Unmarshal(m, b)
}
func (m *Port) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Port.Marshal(b, m, deterministic)
}
func (m *Port) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Port.Merge(m, src)
}
func (m *Port) XXX_Size() int {
	return xxx_messageInfo_Port.Size(m)
}
func (m *Port) XXX_DiscardUnknown() {
	xxx_messageInfo_Port.DiscardUnknown(m)
}

var xxx_messageInfo_Port proto.InternalMessageInfo

func (m *Port) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Port) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Port) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Port) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Port) GetAlias() []string {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Port) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *Port) GetCoordinates() []float32 {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

func (m *Port) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *Port) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Port) GetUnlocs() []string {
	if m != nil {
		return m.Unlocs
	}
	return nil
}

func (m *Port) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "proto_grpc.EmptyRequest")
	proto.RegisterType((*TextResponse)(nil), "proto_grpc.TextResponse")
	proto.RegisterType((*DeleteResponse)(nil), "proto_grpc.DeleteResponse")
	proto.RegisterType((*ListPort)(nil), "proto_grpc.ListPort")
	proto.RegisterType((*SingleRequest)(nil), "proto_grpc.SingleRequest")
	proto.RegisterType((*Port)(nil), "proto_grpc.Port")
}

func init() { proto.RegisterFile("port.proto", fileDescriptor_729c3d36e9010a8e) }

var fileDescriptor_729c3d36e9010a8e = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0xdb, 0x30,
	0x10, 0xdd, 0x38, 0xb1, 0x93, 0x4c, 0xb6, 0xa1, 0x88, 0xa5, 0xa8, 0xb9, 0xd4, 0xf8, 0x50, 0x7c,
	0x69, 0x0e, 0x29, 0x94, 0x1e, 0xf6, 0xb2, 0x90, 0xb2, 0x5d, 0xe8, 0xa1, 0x68, 0xdb, 0x73, 0x51,
	0xed, 0x21, 0x08, 0x1c, 0xc9, 0x95, 0xc6, 0xa5, 0xe9, 0x8f, 0xe8, 0x7f, 0xe8, 0x3f, 0x2d, 0x92,
	0x3f, 0xd6, 0x4b, 0xa0, 0xec, 0xc9, 0xf3, 0xe6, 0x8d, 0x67, 0xe6, 0x3d, 0x0d, 0x40, 0x6d, 0x2c,
	0x6d, 0x6b, 0x6b, 0xc8, 0x30, 0x08, 0x9f, 0x6f, 0x07, 0x5b, 0x17, 0xd9, 0x1a, 0x2e, 0x3f, 0x1c,
	0x6b, 0x3a, 0x09, 0xfc, 0xd1, 0xa0, 0xa3, 0xec, 0x1a, 0x2e, 0xbf, 0xe0, 0x2f, 0x12, 0xe8, 0x6a,
	0xa3, 0x1d, 0x32, 0x06, 0xb3, 0xc2, 0x94, 0xc8, 0x27, 0xe9, 0x24, 0x8f, 0x45, 0x88, 0x19, 0x87,
	0xf9, 0x11, 0x9d, 0x93, 0x07, 0xe4, 0x51, 0x3a, 0xc9, 0x97, 0xa2, 0x87, 0xd9, 0x35, 0xac, 0xf7,
	0x58, 0x21, 0xe1, 0xf0, 0xff, 0x0b, 0x48, 0x1c, 0x49, 0x6a, 0x5c, 0xe8, 0xb0, 0x14, 0x1d, 0x1a,
	0xfa, 0x46, 0x0f, 0x7d, 0xb3, 0x1d, 0x2c, 0x3e, 0x29, 0x47, 0x9f, 0x8d, 0x25, 0xf6, 0x1a, 0x62,
	0xbf, 0xb1, 0xff, 0x6d, 0x9a, 0xaf, 0x76, 0xcf, 0xb7, 0x0f, 0x3b, 0x6f, 0x7d, 0x81, 0x68, 0xe9,
	0xec, 0x15, 0x3c, 0xbb, 0x57, 0xfa, 0x50, 0x61, 0x27, 0x80, 0xad, 0x21, 0x52, 0x65, 0x37, 0x2c,
	0x52, 0x65, 0xf6, 0x27, 0x82, 0x59, 0xe8, 0xb8, 0x86, 0xe8, 0x6e, 0xdf, 0x13, 0x77, 0x7b, 0xbf,
	0x81, 0x96, 0xc7, 0x5e, 0x42, 0x88, 0xc3, 0x56, 0x8a, 0x4e, 0x7c, 0xda, 0xe6, 0x7c, 0xec, 0xd5,
	0x16, 0xa6, 0xd1, 0x64, 0x4f, 0x7c, 0xd6, 0xaa, 0xed, 0x20, 0xbb, 0x82, 0x58, 0x56, 0x4a, 0x3a,
	0x1e, 0xa7, 0xd3, 0x7c, 0x29, 0x5a, 0xe0, 0xeb, 0x2d, 0x1e, 0x94, 0xd1, 0x8e, 0x27, 0x21, 0xdf,
	0x43, 0x96, 0xc2, 0xaa, 0x30, 0xc6, 0x96, 0x4a, 0x4b, 0x42, 0xc7, 0xe7, 0xe9, 0x34, 0x8f, 0xc4,
	0x38, 0xc5, 0x36, 0xb0, 0xa8, 0xad, 0xf9, 0xa9, 0x74, 0x81, 0x7c, 0x11, 0x86, 0x0d, 0xd8, 0x73,
	0xa4, 0x8e, 0xf8, 0xdb, 0x68, 0xe4, 0xcb, 0x96, 0xeb, 0xb1, 0x77, 0xb9, 0xd1, 0x95, 0x29, 0x1c,
	0x87, 0x30, 0xb2, 0x43, 0x83, 0xcb, 0xab, 0x4e, 0x8f, 0x29, 0x71, 0xf7, 0x37, 0x82, 0x95, 0x37,
	0xe4, 0xa3, 0xd4, 0x65, 0x85, 0x96, 0xbd, 0x83, 0xf9, 0x2d, 0xb6, 0xa6, 0xbf, 0x1c, 0xbb, 0xfc,
	0xc8, 0xd6, 0xcd, 0xd9, 0x03, 0xb0, 0xf7, 0x90, 0xdc, 0x22, 0xdd, 0x54, 0x15, 0xe3, 0x63, 0x6e,
	0x7c, 0x4d, 0x9b, 0xab, 0x31, 0x33, 0xbc, 0xed, 0x1b, 0x88, 0xef, 0xc9, 0x58, 0x64, 0x67, 0x4d,
	0xcf, 0xc7, 0x64, 0x17, 0x6c, 0x0b, 0xc9, 0xd7, 0xba, 0x94, 0xf4, 0xd4, 0xfa, 0x1b, 0x48, 0xda,
	0x23, 0xfc, 0x9f, 0x9e, 0xcd, 0x98, 0x7a, 0x7c, 0xb3, 0xd9, 0xc5, 0xf7, 0x24, 0x90, 0x6f, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x36, 0xa4, 0xf7, 0x63, 0x36, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PortHandlerClient is the client API for PortHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PortHandlerClient interface {
	// Sends info about the port
	GetPort(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Port, error)
	GetAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListPort, error)
	Store(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
	Update(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error)
	Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type portHandlerClient struct {
	cc *grpc.ClientConn
}

func NewPortHandlerClient(cc *grpc.ClientConn) PortHandlerClient {
	return &portHandlerClient{cc}
}

func (c *portHandlerClient) GetPort(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/proto_grpc.PortHandler/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) GetAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListPort, error) {
	out := new(ListPort)
	err := c.cc.Invoke(ctx, "/proto_grpc.PortHandler/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Store(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/proto_grpc.PortHandler/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Update(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/proto_grpc.PortHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/proto_grpc.PortHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortHandlerServer is the server API for PortHandler service.
type PortHandlerServer interface {
	// Sends info about the port
	GetPort(context.Context, *SingleRequest) (*Port, error)
	GetAll(context.Context, *EmptyRequest) (*ListPort, error)
	Store(context.Context, *Port) (*Port, error)
	Update(context.Context, *Port) (*Port, error)
	Delete(context.Context, *SingleRequest) (*DeleteResponse, error)
}

// UnimplementedPortHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedPortHandlerServer struct {
}

func (*UnimplementedPortHandlerServer) GetPort(ctx context.Context, req *SingleRequest) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPort not implemented")
}
func (*UnimplementedPortHandlerServer) GetAll(ctx context.Context, req *EmptyRequest) (*ListPort, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedPortHandlerServer) Store(ctx context.Context, req *Port) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Store not implemented")
}
func (*UnimplementedPortHandlerServer) Update(ctx context.Context, req *Port) (*Port, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedPortHandlerServer) Delete(ctx context.Context, req *SingleRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterPortHandlerServer(s *grpc.Server, srv PortHandlerServer) {
	s.RegisterService(&_PortHandler_serviceDesc, srv)
}

func _PortHandler_GetPort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortHandlerServer).GetPort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_grpc.PortHandler/GetPort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).GetPort(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortHandler_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortHandlerServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_grpc.PortHandler/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).GetAll(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortHandler_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortHandlerServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_grpc.PortHandler/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).Store(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Port)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_grpc.PortHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).Update(ctx, req.(*Port))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto_grpc.PortHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).Delete(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto_grpc.PortHandler",
	HandlerType: (*PortHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPort",
			Handler:    _PortHandler_GetPort_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _PortHandler_GetAll_Handler,
		},
		{
			MethodName: "Store",
			Handler:    _PortHandler_Store_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PortHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PortHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "port.proto",
}