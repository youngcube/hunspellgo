// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/spellcheck/spellcheck.proto

package spellcheck

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

type StemRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StemRequest) Reset()         { *m = StemRequest{} }
func (m *StemRequest) String() string { return proto.CompactTextString(m) }
func (*StemRequest) ProtoMessage()    {}
func (*StemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce4f05f433ca2a7, []int{0}
}

func (m *StemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StemRequest.Unmarshal(m, b)
}
func (m *StemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StemRequest.Marshal(b, m, deterministic)
}
func (m *StemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StemRequest.Merge(m, src)
}
func (m *StemRequest) XXX_Size() int {
	return xxx_messageInfo_StemRequest.Size(m)
}
func (m *StemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StemRequest proto.InternalMessageInfo

func (m *StemRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *StemRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type SuggestionRequest struct {
	Word                 string   `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Count                int32    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuggestionRequest) Reset()         { *m = SuggestionRequest{} }
func (m *SuggestionRequest) String() string { return proto.CompactTextString(m) }
func (*SuggestionRequest) ProtoMessage()    {}
func (*SuggestionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce4f05f433ca2a7, []int{1}
}

func (m *SuggestionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuggestionRequest.Unmarshal(m, b)
}
func (m *SuggestionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuggestionRequest.Marshal(b, m, deterministic)
}
func (m *SuggestionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuggestionRequest.Merge(m, src)
}
func (m *SuggestionRequest) XXX_Size() int {
	return xxx_messageInfo_SuggestionRequest.Size(m)
}
func (m *SuggestionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SuggestionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SuggestionRequest proto.InternalMessageInfo

func (m *SuggestionRequest) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *SuggestionRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *SuggestionRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type WordListReply struct {
	WordList             []string `protobuf:"bytes,1,rep,name=wordList,proto3" json:"wordList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WordListReply) Reset()         { *m = WordListReply{} }
func (m *WordListReply) String() string { return proto.CompactTextString(m) }
func (*WordListReply) ProtoMessage()    {}
func (*WordListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce4f05f433ca2a7, []int{2}
}

func (m *WordListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WordListReply.Unmarshal(m, b)
}
func (m *WordListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WordListReply.Marshal(b, m, deterministic)
}
func (m *WordListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WordListReply.Merge(m, src)
}
func (m *WordListReply) XXX_Size() int {
	return xxx_messageInfo_WordListReply.Size(m)
}
func (m *WordListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_WordListReply.DiscardUnknown(m)
}

var xxx_messageInfo_WordListReply proto.InternalMessageInfo

func (m *WordListReply) GetWordList() []string {
	if m != nil {
		return m.WordList
	}
	return nil
}

type StemLineReply struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StemLineReply) Reset()         { *m = StemLineReply{} }
func (m *StemLineReply) String() string { return proto.CompactTextString(m) }
func (*StemLineReply) ProtoMessage()    {}
func (*StemLineReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ce4f05f433ca2a7, []int{3}
}

func (m *StemLineReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StemLineReply.Unmarshal(m, b)
}
func (m *StemLineReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StemLineReply.Marshal(b, m, deterministic)
}
func (m *StemLineReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StemLineReply.Merge(m, src)
}
func (m *StemLineReply) XXX_Size() int {
	return xxx_messageInfo_StemLineReply.Size(m)
}
func (m *StemLineReply) XXX_DiscardUnknown() {
	xxx_messageInfo_StemLineReply.DiscardUnknown(m)
}

var xxx_messageInfo_StemLineReply proto.InternalMessageInfo

func (m *StemLineReply) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*StemRequest)(nil), "spellcheck.StemRequest")
	proto.RegisterType((*SuggestionRequest)(nil), "spellcheck.SuggestionRequest")
	proto.RegisterType((*WordListReply)(nil), "spellcheck.WordListReply")
	proto.RegisterType((*StemLineReply)(nil), "spellcheck.StemLineReply")
}

func init() { proto.RegisterFile("src/spellcheck/spellcheck.proto", fileDescriptor_3ce4f05f433ca2a7) }

var fileDescriptor_3ce4f05f433ca2a7 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0x4a, 0xd6,
	0x2f, 0x2e, 0x48, 0xcd, 0xc9, 0x49, 0xce, 0x48, 0x4d, 0xce, 0x46, 0x62, 0xea, 0x15, 0x14, 0xe5,
	0x97, 0xe4, 0x0b, 0x71, 0x21, 0x44, 0x94, 0x4c, 0xb9, 0xb8, 0x83, 0x4b, 0x52, 0x73, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0xca, 0xf3, 0x8b, 0x52, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x4e, 0x62, 0x5e, 0xba, 0x04, 0x13, 0x44, 0x0c,
	0xc4, 0x56, 0x0a, 0xe4, 0x12, 0x0c, 0x2e, 0x4d, 0x4f, 0x4f, 0x2d, 0x2e, 0xc9, 0xcc, 0xcf, 0x23,
	0x51, 0xb3, 0x90, 0x08, 0x17, 0x6b, 0x72, 0x7e, 0x69, 0x5e, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06,
	0x6b, 0x10, 0x84, 0xa3, 0xa4, 0xcd, 0xc5, 0x1b, 0x9e, 0x5f, 0x94, 0xe2, 0x93, 0x59, 0x5c, 0x12,
	0x94, 0x5a, 0x90, 0x53, 0x29, 0x24, 0xc5, 0xc5, 0x51, 0x0e, 0x15, 0x90, 0x60, 0x54, 0x60, 0xd6,
	0xe0, 0x0c, 0x82, 0xf3, 0x95, 0xd4, 0xb9, 0x78, 0x41, 0xce, 0xf6, 0xc9, 0xcc, 0x4b, 0x85, 0x28,
	0x16, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x81, 0xda, 0x0e, 0xe5, 0x19, 0xdd, 0x60,
	0xe4, 0xe2, 0xf0, 0x28, 0xcd, 0x03, 0xfb, 0x58, 0xc8, 0x81, 0x8b, 0x03, 0xa4, 0x0b, 0x64, 0x8d,
	0x90, 0xb8, 0x1e, 0x52, 0xb8, 0x20, 0x05, 0x81, 0x94, 0x24, 0xb2, 0x04, 0x8a, 0x8b, 0x94, 0x18,
	0x60, 0x26, 0x80, 0xec, 0x25, 0xd2, 0x04, 0x14, 0x67, 0x2a, 0x31, 0x08, 0x79, 0x70, 0x71, 0x21,
	0x42, 0x4e, 0x48, 0x16, 0x45, 0x29, 0x7a, 0x88, 0xe2, 0x75, 0x4b, 0x12, 0x1b, 0x38, 0x36, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xda, 0x25, 0x79, 0xe6, 0xf0, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HunspellClient is the client API for Hunspell service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HunspellClient interface {
	StemWord(ctx context.Context, in *StemRequest, opts ...grpc.CallOption) (*WordListReply, error)
	StemLine(ctx context.Context, in *StemRequest, opts ...grpc.CallOption) (*StemLineReply, error)
	Suggestion(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*WordListReply, error)
}

type hunspellClient struct {
	cc *grpc.ClientConn
}

func NewHunspellClient(cc *grpc.ClientConn) HunspellClient {
	return &hunspellClient{cc}
}

func (c *hunspellClient) StemWord(ctx context.Context, in *StemRequest, opts ...grpc.CallOption) (*WordListReply, error) {
	out := new(WordListReply)
	err := c.cc.Invoke(ctx, "/spellcheck.Hunspell/StemWord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hunspellClient) StemLine(ctx context.Context, in *StemRequest, opts ...grpc.CallOption) (*StemLineReply, error) {
	out := new(StemLineReply)
	err := c.cc.Invoke(ctx, "/spellcheck.Hunspell/StemLine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hunspellClient) Suggestion(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*WordListReply, error) {
	out := new(WordListReply)
	err := c.cc.Invoke(ctx, "/spellcheck.Hunspell/Suggestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HunspellServer is the server API for Hunspell service.
type HunspellServer interface {
	StemWord(context.Context, *StemRequest) (*WordListReply, error)
	StemLine(context.Context, *StemRequest) (*StemLineReply, error)
	Suggestion(context.Context, *SuggestionRequest) (*WordListReply, error)
}

// UnimplementedHunspellServer can be embedded to have forward compatible implementations.
type UnimplementedHunspellServer struct {
}

func (*UnimplementedHunspellServer) StemWord(ctx context.Context, req *StemRequest) (*WordListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StemWord not implemented")
}
func (*UnimplementedHunspellServer) StemLine(ctx context.Context, req *StemRequest) (*StemLineReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StemLine not implemented")
}
func (*UnimplementedHunspellServer) Suggestion(ctx context.Context, req *SuggestionRequest) (*WordListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Suggestion not implemented")
}

func RegisterHunspellServer(s *grpc.Server, srv HunspellServer) {
	s.RegisterService(&_Hunspell_serviceDesc, srv)
}

func _Hunspell_StemWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HunspellServer).StemWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spellcheck.Hunspell/StemWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HunspellServer).StemWord(ctx, req.(*StemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hunspell_StemLine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HunspellServer).StemLine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spellcheck.Hunspell/StemLine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HunspellServer).StemLine(ctx, req.(*StemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Hunspell_Suggestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HunspellServer).Suggestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spellcheck.Hunspell/Suggestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HunspellServer).Suggestion(ctx, req.(*SuggestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hunspell_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spellcheck.Hunspell",
	HandlerType: (*HunspellServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StemWord",
			Handler:    _Hunspell_StemWord_Handler,
		},
		{
			MethodName: "StemLine",
			Handler:    _Hunspell_StemLine_Handler,
		},
		{
			MethodName: "Suggestion",
			Handler:    _Hunspell_Suggestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/spellcheck/spellcheck.proto",
}
