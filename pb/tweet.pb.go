// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/tweet.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Tweet struct {
	TweetData            string   `protobuf:"bytes,1,opt,name=tweetData,proto3" json:"tweetData,omitempty"`
	TweetID              string   `protobuf:"bytes,2,opt,name=tweetID,proto3" json:"tweetID,omitempty"`
	Retweets             int64    `protobuf:"varint,3,opt,name=retweets,proto3" json:"retweets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tweet) Reset()         { *m = Tweet{} }
func (m *Tweet) String() string { return proto.CompactTextString(m) }
func (*Tweet) ProtoMessage()    {}
func (*Tweet) Descriptor() ([]byte, []int) {
	return fileDescriptor_tweet_df9085b732c9fb82, []int{0}
}
func (m *Tweet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tweet.Unmarshal(m, b)
}
func (m *Tweet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tweet.Marshal(b, m, deterministic)
}
func (dst *Tweet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tweet.Merge(dst, src)
}
func (m *Tweet) XXX_Size() int {
	return xxx_messageInfo_Tweet.Size(m)
}
func (m *Tweet) XXX_DiscardUnknown() {
	xxx_messageInfo_Tweet.DiscardUnknown(m)
}

var xxx_messageInfo_Tweet proto.InternalMessageInfo

func (m *Tweet) GetTweetData() string {
	if m != nil {
		return m.TweetData
	}
	return ""
}

func (m *Tweet) GetTweetID() string {
	if m != nil {
		return m.TweetID
	}
	return ""
}

func (m *Tweet) GetRetweets() int64 {
	if m != nil {
		return m.Retweets
	}
	return 0
}

type TweetRequest struct {
	TweetData            string   `protobuf:"bytes,1,opt,name=tweetData,proto3" json:"tweetData,omitempty"`
	TweetID              string   `protobuf:"bytes,2,opt,name=tweetID,proto3" json:"tweetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TweetRequest) Reset()         { *m = TweetRequest{} }
func (m *TweetRequest) String() string { return proto.CompactTextString(m) }
func (*TweetRequest) ProtoMessage()    {}
func (*TweetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tweet_df9085b732c9fb82, []int{1}
}
func (m *TweetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TweetRequest.Unmarshal(m, b)
}
func (m *TweetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TweetRequest.Marshal(b, m, deterministic)
}
func (dst *TweetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TweetRequest.Merge(dst, src)
}
func (m *TweetRequest) XXX_Size() int {
	return xxx_messageInfo_TweetRequest.Size(m)
}
func (m *TweetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TweetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TweetRequest proto.InternalMessageInfo

func (m *TweetRequest) GetTweetData() string {
	if m != nil {
		return m.TweetData
	}
	return ""
}

func (m *TweetRequest) GetTweetID() string {
	if m != nil {
		return m.TweetID
	}
	return ""
}

type TweetResponse struct {
	Tweet                *Tweet   `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TweetResponse) Reset()         { *m = TweetResponse{} }
func (m *TweetResponse) String() string { return proto.CompactTextString(m) }
func (*TweetResponse) ProtoMessage()    {}
func (*TweetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tweet_df9085b732c9fb82, []int{2}
}
func (m *TweetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TweetResponse.Unmarshal(m, b)
}
func (m *TweetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TweetResponse.Marshal(b, m, deterministic)
}
func (dst *TweetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TweetResponse.Merge(dst, src)
}
func (m *TweetResponse) XXX_Size() int {
	return xxx_messageInfo_TweetResponse.Size(m)
}
func (m *TweetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TweetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TweetResponse proto.InternalMessageInfo

func (m *TweetResponse) GetTweet() *Tweet {
	if m != nil {
		return m.Tweet
	}
	return nil
}

type TopRetweetsResponse struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopRetweetsResponse) Reset()         { *m = TopRetweetsResponse{} }
func (m *TopRetweetsResponse) String() string { return proto.CompactTextString(m) }
func (*TopRetweetsResponse) ProtoMessage()    {}
func (*TopRetweetsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_tweet_df9085b732c9fb82, []int{3}
}
func (m *TopRetweetsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopRetweetsResponse.Unmarshal(m, b)
}
func (m *TopRetweetsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopRetweetsResponse.Marshal(b, m, deterministic)
}
func (dst *TopRetweetsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopRetweetsResponse.Merge(dst, src)
}
func (m *TopRetweetsResponse) XXX_Size() int {
	return xxx_messageInfo_TopRetweetsResponse.Size(m)
}
func (m *TopRetweetsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TopRetweetsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TopRetweetsResponse proto.InternalMessageInfo

func (m *TopRetweetsResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type TopRetweetsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopRetweetsRequest) Reset()         { *m = TopRetweetsRequest{} }
func (m *TopRetweetsRequest) String() string { return proto.CompactTextString(m) }
func (*TopRetweetsRequest) ProtoMessage()    {}
func (*TopRetweetsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_tweet_df9085b732c9fb82, []int{4}
}
func (m *TopRetweetsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopRetweetsRequest.Unmarshal(m, b)
}
func (m *TopRetweetsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopRetweetsRequest.Marshal(b, m, deterministic)
}
func (dst *TopRetweetsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopRetweetsRequest.Merge(dst, src)
}
func (m *TopRetweetsRequest) XXX_Size() int {
	return xxx_messageInfo_TopRetweetsRequest.Size(m)
}
func (m *TopRetweetsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TopRetweetsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TopRetweetsRequest proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Tweet)(nil), "pb.Tweet")
	proto.RegisterType((*TweetRequest)(nil), "pb.TweetRequest")
	proto.RegisterType((*TweetResponse)(nil), "pb.TweetResponse")
	proto.RegisterType((*TopRetweetsResponse)(nil), "pb.TopRetweetsResponse")
	proto.RegisterType((*TopRetweetsRequest)(nil), "pb.TopRetweetsRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TweetServiceClient is the client API for TweetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TweetServiceClient interface {
	Tweet(ctx context.Context, in *TweetRequest, opts ...grpc.CallOption) (*TweetResponse, error)
	TopRetweets(ctx context.Context, in *TopRetweetsRequest, opts ...grpc.CallOption) (*TopRetweetsResponse, error)
}

type tweetServiceClient struct {
	cc *grpc.ClientConn
}

func NewTweetServiceClient(cc *grpc.ClientConn) TweetServiceClient {
	return &tweetServiceClient{cc}
}

func (c *tweetServiceClient) Tweet(ctx context.Context, in *TweetRequest, opts ...grpc.CallOption) (*TweetResponse, error) {
	out := new(TweetResponse)
	err := c.cc.Invoke(ctx, "/pb.TweetService/Tweet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tweetServiceClient) TopRetweets(ctx context.Context, in *TopRetweetsRequest, opts ...grpc.CallOption) (*TopRetweetsResponse, error) {
	out := new(TopRetweetsResponse)
	err := c.cc.Invoke(ctx, "/pb.TweetService/TopRetweets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TweetServiceServer is the server API for TweetService service.
type TweetServiceServer interface {
	Tweet(context.Context, *TweetRequest) (*TweetResponse, error)
	TopRetweets(context.Context, *TopRetweetsRequest) (*TopRetweetsResponse, error)
}

func RegisterTweetServiceServer(s *grpc.Server, srv TweetServiceServer) {
	s.RegisterService(&_TweetService_serviceDesc, srv)
}

func _TweetService_Tweet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TweetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).Tweet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TweetService/Tweet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).Tweet(ctx, req.(*TweetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TweetService_TopRetweets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopRetweetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TweetServiceServer).TopRetweets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TweetService/TopRetweets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TweetServiceServer).TopRetweets(ctx, req.(*TopRetweetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TweetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TweetService",
	HandlerType: (*TweetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tweet",
			Handler:    _TweetService_Tweet_Handler,
		},
		{
			MethodName: "TopRetweets",
			Handler:    _TweetService_TopRetweets_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/tweet.proto",
}

func init() { proto.RegisterFile("pb/tweet.proto", fileDescriptor_tweet_df9085b732c9fb82) }

var fileDescriptor_tweet_df9085b732c9fb82 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x2f,
	0x29, 0x4f, 0x4d, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x8a,
	0xe6, 0x62, 0x0d, 0x01, 0x09, 0x09, 0xc9, 0x70, 0x71, 0x82, 0xe5, 0x5c, 0x12, 0x4b, 0x12, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10, 0x02, 0x42, 0x12, 0x5c, 0xec, 0x60, 0x8e, 0xa7, 0x8b,
	0x04, 0x13, 0x58, 0x0e, 0xc6, 0x15, 0x92, 0xe2, 0xe2, 0x28, 0x4a, 0x05, 0x73, 0x8a, 0x25, 0x98,
	0x15, 0x18, 0x35, 0x98, 0x83, 0xe0, 0x7c, 0x25, 0x37, 0x2e, 0x1e, 0xb0, 0xe1, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x64, 0xdb, 0xa1, 0x64, 0xc0, 0xc5, 0x0b, 0x35, 0xa7, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0x48, 0x9e, 0x8b, 0x15, 0x2c, 0x07, 0x36, 0x84, 0xdb, 0x88, 0x53, 0xaf, 0x20, 0x49,
	0x0f, 0xa2, 0x02, 0x22, 0xae, 0xa4, 0xc9, 0x25, 0x1c, 0x92, 0x5f, 0x10, 0x04, 0x75, 0x08, 0x5c,
	0x9f, 0x10, 0x17, 0x4b, 0x0a, 0xcc, 0x6e, 0x9e, 0x20, 0x30, 0x5b, 0x49, 0x84, 0x4b, 0x08, 0x45,
	0x29, 0xd8, 0xa9, 0x46, 0x0d, 0x8c, 0x50, 0xb7, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a,
	0xe9, 0xc1, 0x02, 0x4a, 0x00, 0x61, 0x19, 0x44, 0xad, 0x94, 0x20, 0x92, 0x08, 0xc4, 0x22, 0x25,
	0x06, 0x21, 0x07, 0x2e, 0x6e, 0x24, 0x63, 0x85, 0xc4, 0xc0, 0x6a, 0x30, 0xec, 0x91, 0x12, 0xc7,
	0x10, 0x87, 0x99, 0x90, 0xc4, 0x06, 0x8e, 0x25, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2f,
	0x2f, 0x95, 0x84, 0xb7, 0x01, 0x00, 0x00,
}
