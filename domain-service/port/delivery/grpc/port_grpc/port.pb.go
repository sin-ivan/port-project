// Code generated by protoc-gen-go. DO NOT EDIT.
// source: port.proto

package port_grpc

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
	proto.RegisterType((*EmptyRequest)(nil), "port_grpc.EmptyRequest")
	proto.RegisterType((*TextResponse)(nil), "port_grpc.TextResponse")
	proto.RegisterType((*DeleteResponse)(nil), "port_grpc.DeleteResponse")
	proto.RegisterType((*ListPort)(nil), "port_grpc.ListPort")
	proto.RegisterType((*SingleRequest)(nil), "port_grpc.SingleRequest")
	proto.RegisterType((*Port)(nil), "port_grpc.Port")
}

func init() { proto.RegisterFile("port.proto", fileDescriptor_729c3d36e9010a8e) }

var fileDescriptor_729c3d36e9010a8e = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xdd, 0xa6, 0x4d, 0xda, 0xdc, 0xae, 0x15, 0x46, 0xd1, 0xb1, 0x2f, 0x86, 0x80, 0x10, 0x50,
	0x0a, 0xae, 0xe2, 0xd3, 0x82, 0x08, 0x95, 0x75, 0xc1, 0x07, 0x99, 0xd5, 0x67, 0x19, 0x93, 0x4b,
	0x19, 0x48, 0x67, 0xe2, 0xcc, 0xad, 0x58, 0x7f, 0x84, 0xaf, 0xfe, 0x5d, 0x99, 0xc9, 0x87, 0xd9,
	0x0a, 0xe2, 0xdb, 0x3d, 0x73, 0x66, 0xee, 0x3d, 0xe7, 0xe4, 0x06, 0xa0, 0x31, 0x96, 0x36, 0x8d,
	0x35, 0x64, 0x58, 0xea, 0xeb, 0xcf, 0x3b, 0xdb, 0x94, 0xf9, 0x0a, 0xce, 0xdf, 0xee, 0x1b, 0x3a,
	0x0a, 0xfc, 0x7a, 0x40, 0x47, 0xf9, 0x25, 0x9c, 0x7f, 0xc4, 0xef, 0x24, 0xd0, 0x35, 0x46, 0x3b,
	0x64, 0x0c, 0x66, 0xa5, 0xa9, 0x90, 0x4f, 0xb2, 0x49, 0x11, 0x8b, 0x50, 0x33, 0x0e, 0xf3, 0x3d,
	0x3a, 0x27, 0x77, 0xc8, 0xa3, 0x6c, 0x52, 0xa4, 0xa2, 0x87, 0xf9, 0x25, 0xac, 0xb6, 0x58, 0x23,
	0xe1, 0xf0, 0xfe, 0x01, 0x24, 0x8e, 0x24, 0x1d, 0x5c, 0xe8, 0x90, 0x8a, 0x0e, 0x0d, 0x7d, 0xa3,
	0x3f, 0x7d, 0xf3, 0xe7, 0xb0, 0x78, 0xaf, 0x1c, 0x7d, 0x30, 0x96, 0xd8, 0x13, 0x88, 0xbd, 0x48,
	0xff, 0x6c, 0x5a, 0x2c, 0x2f, 0xee, 0x6e, 0x06, 0xc9, 0x1b, 0xcf, 0x8b, 0x96, 0xcd, 0x1f, 0xc3,
	0x9d, 0x1b, 0xa5, 0x77, 0x35, 0x76, 0xfa, 0xd9, 0x0a, 0x22, 0x55, 0x75, 0xb3, 0x22, 0x55, 0xe5,
	0x3f, 0x23, 0x98, 0x85, 0x86, 0x2b, 0x88, 0xae, 0xb7, 0x3d, 0x71, 0xbd, 0xf5, 0x02, 0xb4, 0xdc,
	0xf7, 0x0e, 0x42, 0x1d, 0x44, 0x29, 0x3a, 0xf2, 0x69, 0x7b, 0xe6, 0x6b, 0x6f, 0xb6, 0x34, 0x07,
	0x4d, 0xf6, 0xc8, 0x67, 0xad, 0xd9, 0x0e, 0xb2, 0xfb, 0x10, 0xcb, 0x5a, 0x49, 0xc7, 0xe3, 0x6c,
	0x5a, 0xa4, 0xa2, 0x05, 0xfe, 0xbe, 0xc5, 0x9d, 0x32, 0xda, 0xf1, 0x24, 0x9c, 0xf7, 0x90, 0x65,
	0xb0, 0x2c, 0x8d, 0xb1, 0x95, 0xd2, 0x92, 0xd0, 0xf1, 0x79, 0x36, 0x2d, 0x22, 0x31, 0x3e, 0x62,
	0x6b, 0x58, 0x34, 0xd6, 0x7c, 0x53, 0xba, 0x44, 0xbe, 0x08, 0xc3, 0x06, 0xec, 0x39, 0x52, 0x7b,
	0xfc, 0x61, 0x34, 0xf2, 0xb4, 0xe5, 0x7a, 0xec, 0x43, 0x3e, 0xe8, 0xda, 0x94, 0x8e, 0x43, 0x18,
	0xd9, 0xa1, 0x21, 0xe4, 0x65, 0xe7, 0xc7, 0x54, 0x78, 0xf1, 0x2b, 0x82, 0xa5, 0x0f, 0xe4, 0x9d,
	0xd4, 0x55, 0x8d, 0x96, 0xbd, 0x84, 0xf9, 0x15, 0xb6, 0x99, 0xf3, 0x51, 0xc8, 0xb7, 0x52, 0x5d,
	0x9f, 0xc6, 0xcf, 0x5e, 0x41, 0x72, 0x85, 0xf4, 0xa6, 0xae, 0xd9, 0xc3, 0x11, 0x35, 0xde, 0xa4,
	0xf5, 0xbd, 0x11, 0x31, 0x7c, 0xd6, 0xa7, 0x10, 0xdf, 0x90, 0xb1, 0xc8, 0x4e, 0x3b, 0xfe, 0x35,
	0x22, 0x3f, 0x63, 0xcf, 0x20, 0xf9, 0xd4, 0x54, 0x92, 0xfe, 0xef, 0xf6, 0x6b, 0x48, 0xda, 0xdd,
	0xfb, 0x87, 0x8f, 0x47, 0x23, 0xe6, 0xf6, 0xa2, 0xe6, 0x67, 0x5f, 0x92, 0xf0, 0x73, 0xbc, 0xf8,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x92, 0x78, 0xb6, 0x02, 0x2a, 0x03, 0x00, 0x00,
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
	err := c.cc.Invoke(ctx, "/port_grpc.PortHandler/GetPort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) GetAll(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ListPort, error) {
	out := new(ListPort)
	err := c.cc.Invoke(ctx, "/port_grpc.PortHandler/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Store(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/port_grpc.PortHandler/Store", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Update(ctx context.Context, in *Port, opts ...grpc.CallOption) (*Port, error) {
	out := new(Port)
	err := c.cc.Invoke(ctx, "/port_grpc.PortHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portHandlerClient) Delete(ctx context.Context, in *SingleRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/port_grpc.PortHandler/Delete", in, out, opts...)
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
		FullMethod: "/port_grpc.PortHandler/GetPort",
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
		FullMethod: "/port_grpc.PortHandler/GetAll",
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
		FullMethod: "/port_grpc.PortHandler/Store",
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
		FullMethod: "/port_grpc.PortHandler/Update",
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
		FullMethod: "/port_grpc.PortHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortHandlerServer).Delete(ctx, req.(*SingleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PortHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "port_grpc.PortHandler",
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
