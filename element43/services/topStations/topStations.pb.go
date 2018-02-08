// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topStations.proto

/*
Package topStations is a generated protocol buffer package.

It is generated from these files:
	topStations.proto

It has these top-level messages:
	GetTopStationsRequest
	GetTopStationsResponse
	StationStats
*/
package topStations

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GetTopStationsRequest struct {
}

func (m *GetTopStationsRequest) Reset()                    { *m = GetTopStationsRequest{} }
func (m *GetTopStationsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTopStationsRequest) ProtoMessage()               {}
func (*GetTopStationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetTopStationsResponse struct {
	// Station list
	Stations []*StationStats `protobuf:"bytes,1,rep,name=stations" json:"stations,omitempty"`
}

func (m *GetTopStationsResponse) Reset()                    { *m = GetTopStationsResponse{} }
func (m *GetTopStationsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetTopStationsResponse) ProtoMessage()               {}
func (*GetTopStationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTopStationsResponse) GetStations() []*StationStats {
	if m != nil {
		return m.Stations
	}
	return nil
}

type StationStats struct {
	// The station's ID
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The station's ask orders total volume
	AskVolume float64 `protobuf:"fixed64,2,opt,name=ask_volume,json=askVolume" json:"ask_volume,omitempty"`
	// The station's bid orders total volume
	BidVolume float64 `protobuf:"fixed64,3,opt,name=bid_volume,json=bidVolume" json:"bid_volume,omitempty"`
	// The station's orders total volume
	TotalVolume float64 `protobuf:"fixed64,4,opt,name=total_volume,json=totalVolume" json:"total_volume,omitempty"`
	// The station's total number of orders
	TotalOrders float64 `protobuf:"fixed64,5,opt,name=total_orders,json=totalOrders" json:"total_orders,omitempty"`
}

func (m *StationStats) Reset()                    { *m = StationStats{} }
func (m *StationStats) String() string            { return proto.CompactTextString(m) }
func (*StationStats) ProtoMessage()               {}
func (*StationStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *StationStats) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StationStats) GetAskVolume() float64 {
	if m != nil {
		return m.AskVolume
	}
	return 0
}

func (m *StationStats) GetBidVolume() float64 {
	if m != nil {
		return m.BidVolume
	}
	return 0
}

func (m *StationStats) GetTotalVolume() float64 {
	if m != nil {
		return m.TotalVolume
	}
	return 0
}

func (m *StationStats) GetTotalOrders() float64 {
	if m != nil {
		return m.TotalOrders
	}
	return 0
}

func init() {
	proto.RegisterType((*GetTopStationsRequest)(nil), "topStations.GetTopStationsRequest")
	proto.RegisterType((*GetTopStationsResponse)(nil), "topStations.GetTopStationsResponse")
	proto.RegisterType((*StationStats)(nil), "topStations.StationStats")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TopStations service

type TopStationsClient interface {
	GetTopStations(ctx context.Context, in *GetTopStationsRequest, opts ...grpc.CallOption) (*GetTopStationsResponse, error)
}

type topStationsClient struct {
	cc *grpc.ClientConn
}

func NewTopStationsClient(cc *grpc.ClientConn) TopStationsClient {
	return &topStationsClient{cc}
}

func (c *topStationsClient) GetTopStations(ctx context.Context, in *GetTopStationsRequest, opts ...grpc.CallOption) (*GetTopStationsResponse, error) {
	out := new(GetTopStationsResponse)
	err := grpc.Invoke(ctx, "/topStations.TopStations/GetTopStations", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TopStations service

type TopStationsServer interface {
	GetTopStations(context.Context, *GetTopStationsRequest) (*GetTopStationsResponse, error)
}

func RegisterTopStationsServer(s *grpc.Server, srv TopStationsServer) {
	s.RegisterService(&_TopStations_serviceDesc, srv)
}

func _TopStations_GetTopStations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopStationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopStationsServer).GetTopStations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/topStations.TopStations/GetTopStations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopStationsServer).GetTopStations(ctx, req.(*GetTopStationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopStations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "topStations.TopStations",
	HandlerType: (*TopStationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTopStations",
			Handler:    _TopStations_GetTopStations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topStations.proto",
}

func init() { proto.RegisterFile("topStations.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x99, 0x54, 0x45, 0x5f, 0x4a, 0xc1, 0x01, 0x35, 0x16, 0x8b, 0x31, 0x22, 0x64, 0xd5,
	0xc1, 0x8a, 0x67, 0x70, 0x59, 0x88, 0xe2, 0x56, 0x26, 0x64, 0x28, 0x43, 0xd2, 0xbc, 0x98, 0xf7,
	0xda, 0xb5, 0x78, 0x02, 0xc1, 0x1b, 0x78, 0x25, 0xaf, 0xe0, 0x41, 0x24, 0x69, 0xa2, 0x49, 0x11,
	0x57, 0x33, 0xfc, 0xdf, 0xb7, 0x98, 0xff, 0x1f, 0x38, 0x64, 0x2c, 0xee, 0x59, 0xb3, 0xc5, 0x9c,
	0xa6, 0x45, 0x89, 0x8c, 0xd2, 0xed, 0x44, 0xe3, 0xb3, 0x05, 0xe2, 0x22, 0x33, 0x4a, 0x17, 0x56,
	0xe9, 0x3c, 0xc7, 0x9e, 0x1a, 0x9c, 0xc0, 0xd1, 0x9d, 0xe1, 0x87, 0x5f, 0x3f, 0x32, 0xcf, 0x2b,
	0x43, 0x1c, 0xcc, 0xe1, 0x78, 0x1b, 0x50, 0x81, 0x39, 0x19, 0x79, 0x0b, 0xfb, 0xd4, 0x64, 0x9e,
	0xf0, 0x07, 0xa1, 0x3b, 0x3b, 0x9d, 0x76, 0xdf, 0xd0, 0x5c, 0xaa, 0x83, 0xa2, 0x1f, 0x35, 0xf8,
	0x10, 0x30, 0xec, 0x22, 0x39, 0x02, 0xc7, 0x26, 0x9e, 0xf0, 0x45, 0x38, 0x88, 0x1c, 0x9b, 0xc8,
	0x09, 0x80, 0xa6, 0xf4, 0x69, 0x8d, 0xd9, 0x6a, 0x69, 0x3c, 0xc7, 0x17, 0xa1, 0x88, 0x0e, 0x34,
	0xa5, 0x8f, 0x75, 0x50, 0xe1, 0xd8, 0x26, 0x2d, 0x1e, 0x6c, 0x70, 0x6c, 0x93, 0x06, 0x5f, 0xc0,
	0x90, 0x91, 0x75, 0xd6, 0x0a, 0x3b, 0xb5, 0xe0, 0xd6, 0xd9, 0xb6, 0x82, 0x65, 0x62, 0x4a, 0xf2,
	0x76, 0x3b, 0xca, 0xbc, 0x8e, 0x66, 0x6f, 0x02, 0xdc, 0x4e, 0x67, 0xf9, 0x22, 0x60, 0xd4, 0x9f,
	0x41, 0x06, 0xbd, 0xb2, 0x7f, 0x8e, 0x37, 0xbe, 0xfc, 0xd7, 0xd9, 0xec, 0x18, 0x5c, 0xbd, 0x7e,
	0x7e, 0xbd, 0x3b, 0xe7, 0x72, 0xa2, 0xd6, 0xd7, 0x6a, 0xa9, 0xcb, 0xd4, 0xb0, 0xaa, 0xd6, 0x22,
	0xd5, 0x6e, 0xa6, 0x18, 0x8b, 0x78, 0xaf, 0xfe, 0xa8, 0x9b, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf5, 0xd8, 0x9e, 0x9d, 0xe8, 0x01, 0x00, 0x00,
}
