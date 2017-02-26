// Code generated by protoc-gen-go.
// source: shogi.proto
// DO NOT EDIT!

/*
Package shogi is a generated protocol buffer package.

It is generated from these files:
	shogi.proto

It has these top-level messages:
	Player
	Record
	HandshakeRequest
	HandshakeResult
	LoginRequest
	LoginResponce
	LogoutRequest
	LogoutResponce
	MatchRequest
	MatchResult
	AckRequest
	AckResult
	GameAction
	GameEvent
	PingRequest
	PingResponce
*/
package shogi

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

type Side int32

const (
	Side_BLACK Side = 0
	Side_WHITE Side = 1
)

var Side_name = map[int32]string{
	0: "BLACK",
	1: "WHITE",
}
var Side_value = map[string]int32{
	"BLACK": 0,
	"WHITE": 1,
}

func (x Side) String() string {
	return proto.EnumName(Side_name, int32(x))
}
func (Side) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Record_WellKnown_Type int32

const (
	Record_WellKnown_CSA Record_WellKnown_Type = 0
	Record_WellKnown_KI2 Record_WellKnown_Type = 1
)

var Record_WellKnown_Type_name = map[int32]string{
	0: "CSA",
	1: "KI2",
}
var Record_WellKnown_Type_value = map[string]int32{
	"CSA": 0,
	"KI2": 1,
}

func (x Record_WellKnown_Type) String() string {
	return proto.EnumName(Record_WellKnown_Type_name, int32(x))
}
func (Record_WellKnown_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0, 0} }

type Player struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Record struct {
	// Types that are valid to be assigned to Format:
	//	*Record_Builtin_
	//	*Record_WellKnown_
	Format isRecord_Format `protobuf_oneof:"format"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isRecord_Format interface {
	isRecord_Format()
}

type Record_Builtin_ struct {
	Builtin *Record_Builtin `protobuf:"bytes,1,opt,name=builtin,oneof"`
}
type Record_WellKnown_ struct {
	WellKnown *Record_WellKnown `protobuf:"bytes,2,opt,name=well_known,json=wellKnown,oneof"`
}

func (*Record_Builtin_) isRecord_Format()   {}
func (*Record_WellKnown_) isRecord_Format() {}

func (m *Record) GetFormat() isRecord_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

func (m *Record) GetBuiltin() *Record_Builtin {
	if x, ok := m.GetFormat().(*Record_Builtin_); ok {
		return x.Builtin
	}
	return nil
}

func (m *Record) GetWellKnown() *Record_WellKnown {
	if x, ok := m.GetFormat().(*Record_WellKnown_); ok {
		return x.WellKnown
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Record) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Record_OneofMarshaler, _Record_OneofUnmarshaler, _Record_OneofSizer, []interface{}{
		(*Record_Builtin_)(nil),
		(*Record_WellKnown_)(nil),
	}
}

func _Record_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Record)
	// format
	switch x := m.Format.(type) {
	case *Record_Builtin_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Builtin); err != nil {
			return err
		}
	case *Record_WellKnown_:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WellKnown); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Record.Format has unexpected type %T", x)
	}
	return nil
}

func _Record_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Record)
	switch tag {
	case 1: // format.builtin
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Record_Builtin)
		err := b.DecodeMessage(msg)
		m.Format = &Record_Builtin_{msg}
		return true, err
	case 2: // format.well_known
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Record_WellKnown)
		err := b.DecodeMessage(msg)
		m.Format = &Record_WellKnown_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Record_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Record)
	// format
	switch x := m.Format.(type) {
	case *Record_Builtin_:
		s := proto.Size(x.Builtin)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Record_WellKnown_:
		s := proto.Size(x.WellKnown)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Record format constructed by text data.
type Record_WellKnown struct {
	Type Record_WellKnown_Type `protobuf:"varint,1,opt,name=type,enum=shogi.Record_WellKnown_Type" json:"type,omitempty"`
	Data []byte                `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Record_WellKnown) Reset()                    { *m = Record_WellKnown{} }
func (m *Record_WellKnown) String() string            { return proto.CompactTextString(m) }
func (*Record_WellKnown) ProtoMessage()               {}
func (*Record_WellKnown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

func (m *Record_WellKnown) GetType() Record_WellKnown_Type {
	if m != nil {
		return m.Type
	}
	return Record_WellKnown_CSA
}

func (m *Record_WellKnown) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// Built-in record format.
type Record_Builtin struct {
}

func (m *Record_Builtin) Reset()                    { *m = Record_Builtin{} }
func (m *Record_Builtin) String() string            { return proto.CompactTextString(m) }
func (*Record_Builtin) ProtoMessage()               {}
func (*Record_Builtin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 1} }

type HandshakeRequest struct {
	Version    uint32 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ClientName string `protobuf:"bytes,2,opt,name=client_name,json=clientName" json:"client_name,omitempty"`
}

func (m *HandshakeRequest) Reset()                    { *m = HandshakeRequest{} }
func (m *HandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*HandshakeRequest) ProtoMessage()               {}
func (*HandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HandshakeRequest) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *HandshakeRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type HandshakeResult struct {
}

func (m *HandshakeResult) Reset()                    { *m = HandshakeResult{} }
func (m *HandshakeResult) String() string            { return proto.CompactTextString(m) }
func (*HandshakeResult) ProtoMessage()               {}
func (*HandshakeResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type LoginRequest struct {
	Username string  `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string  `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Player   *Player `protobuf:"bytes,3,opt,name=player" json:"player,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type LoginResponce struct {
}

func (m *LoginResponce) Reset()                    { *m = LoginResponce{} }
func (m *LoginResponce) String() string            { return proto.CompactTextString(m) }
func (*LoginResponce) ProtoMessage()               {}
func (*LoginResponce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type LogoutRequest struct {
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LogoutResponce struct {
}

func (m *LogoutResponce) Reset()                    { *m = LogoutResponce{} }
func (m *LogoutResponce) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponce) ProtoMessage()               {}
func (*LogoutResponce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type MatchRequest struct {
	// Types that are valid to be assigned to Condition:
	//	*MatchRequest_Any_
	Condition isMatchRequest_Condition `protobuf_oneof:"condition"`
}

func (m *MatchRequest) Reset()                    { *m = MatchRequest{} }
func (m *MatchRequest) String() string            { return proto.CompactTextString(m) }
func (*MatchRequest) ProtoMessage()               {}
func (*MatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type isMatchRequest_Condition interface {
	isMatchRequest_Condition()
}

type MatchRequest_Any_ struct {
	Any *MatchRequest_Any `protobuf:"bytes,1,opt,name=any,oneof"`
}

func (*MatchRequest_Any_) isMatchRequest_Condition() {}

func (m *MatchRequest) GetCondition() isMatchRequest_Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func (m *MatchRequest) GetAny() *MatchRequest_Any {
	if x, ok := m.GetCondition().(*MatchRequest_Any_); ok {
		return x.Any
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MatchRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MatchRequest_OneofMarshaler, _MatchRequest_OneofUnmarshaler, _MatchRequest_OneofSizer, []interface{}{
		(*MatchRequest_Any_)(nil),
	}
}

func _MatchRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MatchRequest)
	// condition
	switch x := m.Condition.(type) {
	case *MatchRequest_Any_:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Any); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MatchRequest.Condition has unexpected type %T", x)
	}
	return nil
}

func _MatchRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MatchRequest)
	switch tag {
	case 1: // condition.any
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MatchRequest_Any)
		err := b.DecodeMessage(msg)
		m.Condition = &MatchRequest_Any_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MatchRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MatchRequest)
	// condition
	switch x := m.Condition.(type) {
	case *MatchRequest_Any_:
		s := proto.Size(x.Any)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type MatchRequest_Any struct {
}

func (m *MatchRequest_Any) Reset()                    { *m = MatchRequest_Any{} }
func (m *MatchRequest_Any) String() string            { return proto.CompactTextString(m) }
func (*MatchRequest_Any) ProtoMessage()               {}
func (*MatchRequest_Any) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

type MatchResult struct {
	MatchId   uint32  `protobuf:"varint,1,opt,name=match_id,json=matchId" json:"match_id,omitempty"`
	OtherSide *Player `protobuf:"bytes,2,opt,name=other_side,json=otherSide" json:"other_side,omitempty"`
	YourSide  Side    `protobuf:"varint,3,opt,name=your_side,json=yourSide,enum=shogi.Side" json:"your_side,omitempty"`
	Setup     *Record `protobuf:"bytes,4,opt,name=setup" json:"setup,omitempty"`
}

func (m *MatchResult) Reset()                    { *m = MatchResult{} }
func (m *MatchResult) String() string            { return proto.CompactTextString(m) }
func (*MatchResult) ProtoMessage()               {}
func (*MatchResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MatchResult) GetMatchId() uint32 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *MatchResult) GetOtherSide() *Player {
	if m != nil {
		return m.OtherSide
	}
	return nil
}

func (m *MatchResult) GetYourSide() Side {
	if m != nil {
		return m.YourSide
	}
	return Side_BLACK
}

func (m *MatchResult) GetSetup() *Record {
	if m != nil {
		return m.Setup
	}
	return nil
}

type AckRequest struct {
	MatchId uint32 `protobuf:"varint,1,opt,name=match_id,json=matchId" json:"match_id,omitempty"`
	Accept  bool   `protobuf:"varint,2,opt,name=accept" json:"accept,omitempty"`
}

func (m *AckRequest) Reset()                    { *m = AckRequest{} }
func (m *AckRequest) String() string            { return proto.CompactTextString(m) }
func (*AckRequest) ProtoMessage()               {}
func (*AckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AckRequest) GetMatchId() uint32 {
	if m != nil {
		return m.MatchId
	}
	return 0
}

func (m *AckRequest) GetAccept() bool {
	if m != nil {
		return m.Accept
	}
	return false
}

type AckResult struct {
	Accepted bool `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
}

func (m *AckResult) Reset()                    { *m = AckResult{} }
func (m *AckResult) String() string            { return proto.CompactTextString(m) }
func (*AckResult) ProtoMessage()               {}
func (*AckResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *AckResult) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type GameAction struct {
}

func (m *GameAction) Reset()                    { *m = GameAction{} }
func (m *GameAction) String() string            { return proto.CompactTextString(m) }
func (*GameAction) ProtoMessage()               {}
func (*GameAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type GameEvent struct {
}

func (m *GameEvent) Reset()                    { *m = GameEvent{} }
func (m *GameEvent) String() string            { return proto.CompactTextString(m) }
func (*GameEvent) ProtoMessage()               {}
func (*GameEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type PingRequest struct {
	// Current client's epoch milliseconds.
	// Server may/can measure the network latency by this value.
	EpochMs uint64 `protobuf:"varint,1,opt,name=epoch_ms,json=epochMs" json:"epoch_ms,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PingRequest) GetEpochMs() uint64 {
	if m != nil {
		return m.EpochMs
	}
	return 0
}

type PingResponce struct {
	// Current server's epoch milliseconds.
	EpochMs uint64 `protobuf:"varint,1,opt,name=epoch_ms,json=epochMs" json:"epoch_ms,omitempty"`
}

func (m *PingResponce) Reset()                    { *m = PingResponce{} }
func (m *PingResponce) String() string            { return proto.CompactTextString(m) }
func (*PingResponce) ProtoMessage()               {}
func (*PingResponce) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *PingResponce) GetEpochMs() uint64 {
	if m != nil {
		return m.EpochMs
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "shogi.Player")
	proto.RegisterType((*Record)(nil), "shogi.Record")
	proto.RegisterType((*Record_WellKnown)(nil), "shogi.Record.WellKnown")
	proto.RegisterType((*Record_Builtin)(nil), "shogi.Record.Builtin")
	proto.RegisterType((*HandshakeRequest)(nil), "shogi.HandshakeRequest")
	proto.RegisterType((*HandshakeResult)(nil), "shogi.HandshakeResult")
	proto.RegisterType((*LoginRequest)(nil), "shogi.LoginRequest")
	proto.RegisterType((*LoginResponce)(nil), "shogi.LoginResponce")
	proto.RegisterType((*LogoutRequest)(nil), "shogi.LogoutRequest")
	proto.RegisterType((*LogoutResponce)(nil), "shogi.LogoutResponce")
	proto.RegisterType((*MatchRequest)(nil), "shogi.MatchRequest")
	proto.RegisterType((*MatchRequest_Any)(nil), "shogi.MatchRequest.Any")
	proto.RegisterType((*MatchResult)(nil), "shogi.MatchResult")
	proto.RegisterType((*AckRequest)(nil), "shogi.AckRequest")
	proto.RegisterType((*AckResult)(nil), "shogi.AckResult")
	proto.RegisterType((*GameAction)(nil), "shogi.GameAction")
	proto.RegisterType((*GameEvent)(nil), "shogi.GameEvent")
	proto.RegisterType((*PingRequest)(nil), "shogi.PingRequest")
	proto.RegisterType((*PingResponce)(nil), "shogi.PingResponce")
	proto.RegisterEnum("shogi.Side", Side_name, Side_value)
	proto.RegisterEnum("shogi.Record_WellKnown_Type", Record_WellKnown_Type_name, Record_WellKnown_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProtocolService service

type ProtocolServiceClient interface {
	Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResult, error)
}

type protocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewProtocolServiceClient(cc *grpc.ClientConn) ProtocolServiceClient {
	return &protocolServiceClient{cc}
}

func (c *protocolServiceClient) Handshake(ctx context.Context, in *HandshakeRequest, opts ...grpc.CallOption) (*HandshakeResult, error) {
	out := new(HandshakeResult)
	err := grpc.Invoke(ctx, "/shogi.ProtocolService/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProtocolService service

type ProtocolServiceServer interface {
	Handshake(context.Context, *HandshakeRequest) (*HandshakeResult, error)
}

func RegisterProtocolServiceServer(s *grpc.Server, srv ProtocolServiceServer) {
	s.RegisterService(&_ProtocolService_serviceDesc, srv)
}

func _ProtocolService_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolServiceServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.ProtocolService/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolServiceServer).Handshake(ctx, req.(*HandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.ProtocolService",
	HandlerType: (*ProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _ProtocolService_Handshake_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shogi.proto",
}

// Client API for SessionService service

type SessionServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponce, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponce, error)
}

type sessionServiceClient struct {
	cc *grpc.ClientConn
}

func NewSessionServiceClient(cc *grpc.ClientConn) SessionServiceClient {
	return &sessionServiceClient{cc}
}

func (c *sessionServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponce, error) {
	out := new(LoginResponce)
	err := grpc.Invoke(ctx, "/shogi.SessionService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionServiceClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponce, error) {
	out := new(LogoutResponce)
	err := grpc.Invoke(ctx, "/shogi.SessionService/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionService service

type SessionServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponce, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponce, error)
}

func RegisterSessionServiceServer(s *grpc.Server, srv SessionServiceServer) {
	s.RegisterService(&_SessionService_serviceDesc, srv)
}

func _SessionService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.SessionService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.SessionService/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionServiceServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.SessionService",
	HandlerType: (*SessionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _SessionService_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _SessionService_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shogi.proto",
}

// Client API for MatchService service

type MatchServiceClient interface {
	Request(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResult, error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResult, error)
}

type matchServiceClient struct {
	cc *grpc.ClientConn
}

func NewMatchServiceClient(cc *grpc.ClientConn) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) Request(ctx context.Context, in *MatchRequest, opts ...grpc.CallOption) (*MatchResult, error) {
	out := new(MatchResult)
	err := grpc.Invoke(ctx, "/shogi.MatchService/Request", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*AckResult, error) {
	out := new(AckResult)
	err := grpc.Invoke(ctx, "/shogi.MatchService/Ack", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MatchService service

type MatchServiceServer interface {
	Request(context.Context, *MatchRequest) (*MatchResult, error)
	Ack(context.Context, *AckRequest) (*AckResult, error)
}

func RegisterMatchServiceServer(s *grpc.Server, srv MatchServiceServer) {
	s.RegisterService(&_MatchService_serviceDesc, srv)
}

func _MatchService_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.MatchService/Request",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).Request(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shogi.MatchService/Ack",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MatchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _MatchService_Request_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _MatchService_Ack_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shogi.proto",
}

// Client API for GameService service

type GameServiceClient interface {
	Start(ctx context.Context, opts ...grpc.CallOption) (GameService_StartClient, error)
}

type gameServiceClient struct {
	cc *grpc.ClientConn
}

func NewGameServiceClient(cc *grpc.ClientConn) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Start(ctx context.Context, opts ...grpc.CallOption) (GameService_StartClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GameService_serviceDesc.Streams[0], c.cc, "/shogi.GameService/Start", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceStartClient{stream}
	return x, nil
}

type GameService_StartClient interface {
	Send(*GameAction) error
	Recv() (*GameEvent, error)
	grpc.ClientStream
}

type gameServiceStartClient struct {
	grpc.ClientStream
}

func (x *gameServiceStartClient) Send(m *GameAction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceStartClient) Recv() (*GameEvent, error) {
	m := new(GameEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for GameService service

type GameServiceServer interface {
	Start(GameService_StartServer) error
}

func RegisterGameServiceServer(s *grpc.Server, srv GameServiceServer) {
	s.RegisterService(&_GameService_serviceDesc, srv)
}

func _GameService_Start_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).Start(&gameServiceStartServer{stream})
}

type GameService_StartServer interface {
	Send(*GameEvent) error
	Recv() (*GameAction, error)
	grpc.ServerStream
}

type gameServiceStartServer struct {
	grpc.ServerStream
}

func (x *gameServiceStartServer) Send(m *GameEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceStartServer) Recv() (*GameAction, error) {
	m := new(GameAction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GameService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Start",
			Handler:       _GameService_Start_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "shogi.proto",
}

// Client API for PingService service

type PingServiceClient interface {
	Ping(ctx context.Context, opts ...grpc.CallOption) (PingService_PingClient, error)
}

type pingServiceClient struct {
	cc *grpc.ClientConn
}

func NewPingServiceClient(cc *grpc.ClientConn) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) Ping(ctx context.Context, opts ...grpc.CallOption) (PingService_PingClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PingService_serviceDesc.Streams[0], c.cc, "/shogi.PingService/Ping", opts...)
	if err != nil {
		return nil, err
	}
	x := &pingServicePingClient{stream}
	return x, nil
}

type PingService_PingClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponce, error)
	grpc.ClientStream
}

type pingServicePingClient struct {
	grpc.ClientStream
}

func (x *pingServicePingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pingServicePingClient) Recv() (*PingResponce, error) {
	m := new(PingResponce)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PingService service

type PingServiceServer interface {
	Ping(PingService_PingServer) error
}

func RegisterPingServiceServer(s *grpc.Server, srv PingServiceServer) {
	s.RegisterService(&_PingService_serviceDesc, srv)
}

func _PingService_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PingServiceServer).Ping(&pingServicePingServer{stream})
}

type PingService_PingServer interface {
	Send(*PingResponce) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type pingServicePingServer struct {
	grpc.ServerStream
}

func (x *pingServicePingServer) Send(m *PingResponce) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pingServicePingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shogi.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Ping",
			Handler:       _PingService_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "shogi.proto",
}

func init() { proto.RegisterFile("shogi.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x4f, 0xdb, 0x48,
	0x14, 0x8d, 0xc9, 0xa7, 0xaf, 0x03, 0x84, 0x61, 0x61, 0xb3, 0x16, 0xd2, 0xae, 0xbc, 0x5a, 0x6d,
	0x96, 0x5d, 0x45, 0xac, 0x51, 0xa5, 0x3e, 0x54, 0x42, 0x09, 0x42, 0x0d, 0x02, 0x0a, 0x9a, 0x20,
	0xf1, 0x18, 0x19, 0x7b, 0x4a, 0xac, 0x38, 0x33, 0xae, 0x3d, 0x21, 0x8a, 0xfa, 0x7b, 0xfa, 0x1b,
	0xfb, 0x5a, 0xcd, 0x97, 0xe3, 0xd0, 0x96, 0xb7, 0x39, 0xf7, 0x9c, 0xb9, 0x73, 0xef, 0x3d, 0xd7,
	0x06, 0x27, 0x9f, 0xb2, 0xa7, 0xb8, 0x9f, 0x66, 0x8c, 0x33, 0x54, 0x97, 0xc0, 0x3b, 0x82, 0xc6,
	0x5d, 0x12, 0xac, 0x48, 0x86, 0x10, 0xd4, 0x68, 0x30, 0x27, 0x5d, 0xeb, 0x0f, 0xab, 0x67, 0x63,
	0x79, 0xf6, 0xbe, 0x5a, 0xd0, 0xc0, 0x24, 0x64, 0x59, 0x84, 0xfe, 0x87, 0xe6, 0xe3, 0x22, 0x4e,
	0x78, 0x4c, 0xa5, 0xc2, 0xf1, 0x0f, 0xfa, 0x2a, 0x9d, 0xe2, 0xfb, 0x43, 0x45, 0x8e, 0x2a, 0xd8,
	0xe8, 0xd0, 0x5b, 0x80, 0x25, 0x49, 0x92, 0xc9, 0x8c, 0xb2, 0x25, 0xed, 0x6e, 0xc9, 0x5b, 0xbf,
	0x6e, 0xde, 0x7a, 0x20, 0x49, 0x72, 0x25, 0xe8, 0x51, 0x05, 0xdb, 0x4b, 0x03, 0xdc, 0x19, 0xd8,
	0x05, 0x83, 0x4e, 0xa0, 0xc6, 0x57, 0xa9, 0x2a, 0x6c, 0xc7, 0x3f, 0xfa, 0x49, 0x82, 0xfe, 0xfd,
	0x2a, 0x25, 0x58, 0x2a, 0x45, 0x2b, 0x51, 0xc0, 0x03, 0xf9, 0x64, 0x1b, 0xcb, 0xb3, 0xd7, 0x85,
	0x9a, 0x50, 0xa0, 0x26, 0x54, 0xcf, 0xc7, 0x83, 0x4e, 0x45, 0x1c, 0xae, 0x2e, 0xfd, 0x8e, 0xe5,
	0xda, 0xd0, 0xd4, 0xc5, 0x0f, 0x5b, 0xd0, 0xf8, 0xc8, 0xb2, 0x79, 0xc0, 0xbd, 0x1b, 0xe8, 0x8c,
	0x02, 0x1a, 0xe5, 0xd3, 0x60, 0x46, 0x30, 0xf9, 0xb4, 0x20, 0x39, 0x47, 0x5d, 0x68, 0x3e, 0x93,
	0x2c, 0x8f, 0x99, 0x1a, 0xc1, 0x36, 0x36, 0x10, 0xfd, 0x0e, 0x4e, 0x98, 0xc4, 0x84, 0xf2, 0x89,
	0x1c, 0xe1, 0x96, 0x1c, 0x21, 0xa8, 0xd0, 0x07, 0x31, 0xc8, 0x3d, 0xd8, 0x2d, 0xa5, 0xcb, 0x17,
	0x09, 0xf7, 0xe6, 0xd0, 0xbe, 0x66, 0x4f, 0x31, 0x35, 0xd9, 0x5d, 0x68, 0x2d, 0x72, 0x92, 0x95,
	0x3c, 0x28, 0xb0, 0xe0, 0xd2, 0x20, 0xcf, 0x97, 0x2c, 0x8b, 0x74, 0xf2, 0x02, 0xa3, 0xbf, 0xa0,
	0x91, 0x4a, 0x07, 0xbb, 0x55, 0x39, 0xe1, 0x6d, 0x3d, 0x20, 0x65, 0x2b, 0xd6, 0xa4, 0xb7, 0x0b,
	0xdb, 0xfa, 0xb9, 0x3c, 0x65, 0x34, 0x24, 0x3a, 0xc0, 0x16, 0x5c, 0x17, 0xe0, 0x75, 0x60, 0xc7,
	0x04, 0xb4, 0xe4, 0x16, 0xda, 0x37, 0x01, 0x0f, 0xa7, 0xa6, 0xc4, 0x7f, 0xa1, 0x1a, 0xd0, 0x95,
	0xf6, 0xdf, 0x38, 0x59, 0x56, 0xf4, 0x07, 0x74, 0x35, 0xaa, 0x60, 0xa1, 0x72, 0xeb, 0x50, 0x1d,
	0xd0, 0xd5, 0xd0, 0x01, 0x3b, 0x64, 0x34, 0x8a, 0x79, 0xcc, 0xa8, 0xf7, 0xc5, 0x02, 0x47, 0xeb,
	0xc5, 0x0c, 0xd0, 0x6f, 0xd0, 0x9a, 0x0b, 0x38, 0x89, 0x23, 0x33, 0x52, 0x89, 0x2f, 0x23, 0xf4,
	0x1f, 0x00, 0xe3, 0x53, 0x92, 0x4d, 0xf2, 0x38, 0x22, 0x7a, 0x79, 0x5e, 0xb4, 0x66, 0x4b, 0xc1,
	0x38, 0x8e, 0x08, 0xea, 0x81, 0xbd, 0x62, 0x0b, 0x2d, 0xae, 0xca, 0x45, 0x71, 0xb4, 0x58, 0xf0,
	0xb8, 0x25, 0x58, 0xa9, 0xfc, 0x13, 0xea, 0x39, 0xe1, 0x8b, 0xb4, 0x5b, 0xdb, 0x48, 0xa9, 0xd6,
	0x09, 0x2b, 0xce, 0x3b, 0x03, 0x18, 0x84, 0x33, 0xd3, 0xf6, 0x2b, 0x55, 0x1e, 0x42, 0x23, 0x08,
	0x43, 0x92, 0x72, 0x59, 0x61, 0x0b, 0x6b, 0xe4, 0xfd, 0x0d, 0xb6, 0x4c, 0x20, 0xbb, 0x74, 0xa1,
	0xa5, 0xc2, 0x44, 0xdd, 0x6f, 0xe1, 0x02, 0x7b, 0x6d, 0x80, 0xf7, 0xc1, 0x9c, 0x0c, 0x42, 0x39,
	0x1f, 0x07, 0x6c, 0x81, 0x2e, 0x9e, 0x09, 0xe5, 0x5e, 0x0f, 0x9c, 0xbb, 0x98, 0x3e, 0x95, 0xaa,
	0x20, 0x29, 0x0b, 0xa7, 0x93, 0x79, 0x2e, 0xb3, 0xd4, 0x70, 0x53, 0xe2, 0x9b, 0xdc, 0xfb, 0x07,
	0xda, 0x4a, 0xa9, 0x7c, 0x7b, 0x45, 0x7a, 0x7c, 0x04, 0x35, 0x39, 0x06, 0x1b, 0xea, 0xc3, 0xeb,
	0xc1, 0xf9, 0x55, 0xa7, 0x22, 0x8e, 0x0f, 0xa3, 0xcb, 0xfb, 0x8b, 0x8e, 0xe5, 0xdf, 0xc2, 0xee,
	0x9d, 0xf8, 0x3b, 0x84, 0x2c, 0x19, 0x93, 0xec, 0x39, 0x0e, 0x09, 0x7a, 0x07, 0x76, 0xb1, 0xb9,
	0xc8, 0x78, 0xfe, 0xf2, 0xd3, 0x70, 0x0f, 0xbf, 0x27, 0x44, 0xeb, 0xfe, 0x67, 0xd8, 0x19, 0x93,
	0x5c, 0x7c, 0x23, 0x26, 0x9f, 0x0f, 0x75, 0xb9, 0x87, 0x68, 0x5f, 0x5f, 0x29, 0x7f, 0x04, 0xee,
	0x2f, 0x9b, 0x41, 0xdd, 0xcf, 0x1b, 0x68, 0xa8, 0xcd, 0x44, 0x25, 0x7e, 0xbd, 0xb9, 0xee, 0xc1,
	0x8b, 0xa8, 0xba, 0xe6, 0x53, 0xbd, 0xbe, 0xeb, 0xa7, 0x9b, 0x66, 0x98, 0xfb, 0x3f, 0x58, 0x5e,
	0x17, 0x6d, 0x06, 0xa5, 0x77, 0xc7, 0x50, 0x1d, 0x84, 0x33, 0xb4, 0xa7, 0xa9, 0xf5, 0x56, 0xb8,
	0x9d, 0x72, 0x48, 0x36, 0x7b, 0x06, 0x8e, 0x70, 0xcf, 0x3c, 0x77, 0x02, 0xf5, 0x31, 0x0f, 0x32,
	0x5e, 0x5c, 0x5e, 0x1b, 0x5d, 0x5c, 0x2e, 0xdc, 0xee, 0x59, 0x27, 0x96, 0x3f, 0x54, 0x8e, 0x9b,
	0x04, 0xa7, 0x50, 0x13, 0x10, 0x99, 0xba, 0x4a, 0xdb, 0xe0, 0xee, 0x6f, 0xc4, 0x54, 0xc3, 0x22,
	0xc7, 0x63, 0x43, 0xfe, 0xde, 0x4f, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x26, 0xfb, 0xf1,
	0xed, 0x05, 0x00, 0x00,
}