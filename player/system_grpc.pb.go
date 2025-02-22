// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: player/system.proto

package player

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Matches_GetPlayer_FullMethodName           = "/player.Matches/GetPlayer"
	Matches_GetMatchHistory_FullMethodName     = "/player.Matches/GetMatchHistory"
	Matches_SendUpcomingMatches_FullMethodName = "/player.Matches/SendUpcomingMatches"
	Matches_ProcessMatchResults_FullMethodName = "/player.Matches/ProcessMatchResults"
)

// MatchesClient is the client API for Matches service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchesClient interface {
	// Simple unary RPC
	//
	// Given a name of a server, get its corresponding Status
	// If the name doesn't exist, return a Status with an empty name
	GetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*PlayerSummary, error)
	GetMatchHistory(ctx context.Context, in *Player, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MatchResult], error)
	SendUpcomingMatches(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpcomingMatch, MatchSummary], error)
	ProcessMatchResults(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MatchResult, MatchResult], error)
}

type matchesClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchesClient(cc grpc.ClientConnInterface) MatchesClient {
	return &matchesClient{cc}
}

func (c *matchesClient) GetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*PlayerSummary, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlayerSummary)
	err := c.cc.Invoke(ctx, Matches_GetPlayer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchesClient) GetMatchHistory(ctx context.Context, in *Player, opts ...grpc.CallOption) (grpc.ServerStreamingClient[MatchResult], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Matches_ServiceDesc.Streams[0], Matches_GetMatchHistory_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Player, MatchResult]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_GetMatchHistoryClient = grpc.ServerStreamingClient[MatchResult]

func (c *matchesClient) SendUpcomingMatches(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UpcomingMatch, MatchSummary], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Matches_ServiceDesc.Streams[1], Matches_SendUpcomingMatches_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UpcomingMatch, MatchSummary]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_SendUpcomingMatchesClient = grpc.ClientStreamingClient[UpcomingMatch, MatchSummary]

func (c *matchesClient) ProcessMatchResults(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MatchResult, MatchResult], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Matches_ServiceDesc.Streams[2], Matches_ProcessMatchResults_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MatchResult, MatchResult]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_ProcessMatchResultsClient = grpc.BidiStreamingClient[MatchResult, MatchResult]

// MatchesServer is the server API for Matches service.
// All implementations must embed UnimplementedMatchesServer
// for forward compatibility.
type MatchesServer interface {
	// Simple unary RPC
	//
	// Given a name of a server, get its corresponding Status
	// If the name doesn't exist, return a Status with an empty name
	GetPlayer(context.Context, *Player) (*PlayerSummary, error)
	GetMatchHistory(*Player, grpc.ServerStreamingServer[MatchResult]) error
	SendUpcomingMatches(grpc.ClientStreamingServer[UpcomingMatch, MatchSummary]) error
	ProcessMatchResults(grpc.BidiStreamingServer[MatchResult, MatchResult]) error
	mustEmbedUnimplementedMatchesServer()
}

// UnimplementedMatchesServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMatchesServer struct{}

func (UnimplementedMatchesServer) GetPlayer(context.Context, *Player) (*PlayerSummary, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (UnimplementedMatchesServer) GetMatchHistory(*Player, grpc.ServerStreamingServer[MatchResult]) error {
	return status.Errorf(codes.Unimplemented, "method GetMatchHistory not implemented")
}
func (UnimplementedMatchesServer) SendUpcomingMatches(grpc.ClientStreamingServer[UpcomingMatch, MatchSummary]) error {
	return status.Errorf(codes.Unimplemented, "method SendUpcomingMatches not implemented")
}
func (UnimplementedMatchesServer) ProcessMatchResults(grpc.BidiStreamingServer[MatchResult, MatchResult]) error {
	return status.Errorf(codes.Unimplemented, "method ProcessMatchResults not implemented")
}
func (UnimplementedMatchesServer) mustEmbedUnimplementedMatchesServer() {}
func (UnimplementedMatchesServer) testEmbeddedByValue()                 {}

// UnsafeMatchesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchesServer will
// result in compilation errors.
type UnsafeMatchesServer interface {
	mustEmbedUnimplementedMatchesServer()
}

func RegisterMatchesServer(s grpc.ServiceRegistrar, srv MatchesServer) {
	// If the following call pancis, it indicates UnimplementedMatchesServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Matches_ServiceDesc, srv)
}

func _Matches_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchesServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Matches_GetPlayer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchesServer).GetPlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Matches_GetMatchHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Player)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MatchesServer).GetMatchHistory(m, &grpc.GenericServerStream[Player, MatchResult]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_GetMatchHistoryServer = grpc.ServerStreamingServer[MatchResult]

func _Matches_SendUpcomingMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MatchesServer).SendUpcomingMatches(&grpc.GenericServerStream[UpcomingMatch, MatchSummary]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_SendUpcomingMatchesServer = grpc.ClientStreamingServer[UpcomingMatch, MatchSummary]

func _Matches_ProcessMatchResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MatchesServer).ProcessMatchResults(&grpc.GenericServerStream[MatchResult, MatchResult]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Matches_ProcessMatchResultsServer = grpc.BidiStreamingServer[MatchResult, MatchResult]

// Matches_ServiceDesc is the grpc.ServiceDesc for Matches service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Matches_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "player.Matches",
	HandlerType: (*MatchesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPlayer",
			Handler:    _Matches_GetPlayer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetMatchHistory",
			Handler:       _Matches_GetMatchHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendUpcomingMatches",
			Handler:       _Matches_SendUpcomingMatches_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ProcessMatchResults",
			Handler:       _Matches_ProcessMatchResults_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "player/system.proto",
}
