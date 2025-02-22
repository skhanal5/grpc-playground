// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: player/system.proto

package player

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rank int32

const (
	Rank_RANK_UNSPECIFIED Rank = 0
	Rank_RANK_BRONZE      Rank = 1
	Rank_RANK_SILVER      Rank = 2
	Rank_RANK_GOLD        Rank = 3
	Rank_RANK_PLATINUM    Rank = 4
	Rank_RANK_DIAMOND     Rank = 5
	Rank_RANK_MASTER      Rank = 6
	Rank_RANK_FINAL       Rank = 7
)

// Enum value maps for Rank.
var (
	Rank_name = map[int32]string{
		0: "RANK_UNSPECIFIED",
		1: "RANK_BRONZE",
		2: "RANK_SILVER",
		3: "RANK_GOLD",
		4: "RANK_PLATINUM",
		5: "RANK_DIAMOND",
		6: "RANK_MASTER",
		7: "RANK_FINAL",
	}
	Rank_value = map[string]int32{
		"RANK_UNSPECIFIED": 0,
		"RANK_BRONZE":      1,
		"RANK_SILVER":      2,
		"RANK_GOLD":        3,
		"RANK_PLATINUM":    4,
		"RANK_DIAMOND":     5,
		"RANK_MASTER":      6,
		"RANK_FINAL":       7,
	}
)

func (x Rank) Enum() *Rank {
	p := new(Rank)
	*p = x
	return p
}

func (x Rank) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Rank) Descriptor() protoreflect.EnumDescriptor {
	return file_player_system_proto_enumTypes[0].Descriptor()
}

func (Rank) Type() protoreflect.EnumType {
	return &file_player_system_proto_enumTypes[0]
}

func (x Rank) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Rank.Descriptor instead.
func (Rank) EnumDescriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{0}
}

type Result int32

const (
	Result_RESULT_UNSPECIFIED Result = 0
	Result_RESULT_WIN         Result = 1
	Result_RESULT_LOSS        Result = 2
	Result_RESULT_DRAW        Result = 3
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "RESULT_UNSPECIFIED",
		1: "RESULT_WIN",
		2: "RESULT_LOSS",
		3: "RESULT_DRAW",
	}
	Result_value = map[string]int32{
		"RESULT_UNSPECIFIED": 0,
		"RESULT_WIN":         1,
		"RESULT_LOSS":        2,
		"RESULT_DRAW":        3,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_player_system_proto_enumTypes[1].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_player_system_proto_enumTypes[1]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{1}
}

type Player struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerId      int32                  `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Player) Reset() {
	*x = Player{}
	mi := &file_player_system_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_system_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

type PlayerSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerId      int32                  `protobuf:"varint,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Rank          Rank                   `protobuf:"varint,3,opt,name=rank,proto3,enum=player.Rank" json:"rank,omitempty"`
	Elo           int32                  `protobuf:"varint,4,opt,name=elo,proto3" json:"elo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayerSummary) Reset() {
	*x = PlayerSummary{}
	mi := &file_player_system_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSummary) ProtoMessage() {}

func (x *PlayerSummary) ProtoReflect() protoreflect.Message {
	mi := &file_player_system_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSummary.ProtoReflect.Descriptor instead.
func (*PlayerSummary) Descriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{1}
}

func (x *PlayerSummary) GetPlayerId() int32 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *PlayerSummary) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlayerSummary) GetRank() Rank {
	if x != nil {
		return x.Rank
	}
	return Rank_RANK_UNSPECIFIED
}

func (x *PlayerSummary) GetElo() int32 {
	if x != nil {
		return x.Elo
	}
	return 0
}

type MatchResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Player        int32                  `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"` // could be a Player instead
	Opponent      int32                  `protobuf:"varint,2,opt,name=opponent,proto3" json:"opponent,omitempty"`
	Result        Result                 `protobuf:"varint,3,opt,name=result,proto3,enum=player.Result" json:"result,omitempty"`
	Timestamp     string                 `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EloGain       int32                  `protobuf:"varint,5,opt,name=elo_gain,json=eloGain,proto3" json:"elo_gain,omitempty"`
	EloLoss       int32                  `protobuf:"varint,6,opt,name=elo_loss,json=eloLoss,proto3" json:"elo_loss,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchResult) Reset() {
	*x = MatchResult{}
	mi := &file_player_system_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResult) ProtoMessage() {}

func (x *MatchResult) ProtoReflect() protoreflect.Message {
	mi := &file_player_system_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResult.ProtoReflect.Descriptor instead.
func (*MatchResult) Descriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{2}
}

func (x *MatchResult) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *MatchResult) GetOpponent() int32 {
	if x != nil {
		return x.Opponent
	}
	return 0
}

func (x *MatchResult) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_RESULT_UNSPECIFIED
}

func (x *MatchResult) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *MatchResult) GetEloGain() int32 {
	if x != nil {
		return x.EloGain
	}
	return 0
}

func (x *MatchResult) GetEloLoss() int32 {
	if x != nil {
		return x.EloLoss
	}
	return 0
}

type UpcomingMatch struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Player        int32                  `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"` // could be a Player instead
	Opponent      int32                  `protobuf:"varint,2,opt,name=opponent,proto3" json:"opponent,omitempty"`
	Timestamp     string                 `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	EloGain       int32                  `protobuf:"varint,4,opt,name=elo_gain,json=eloGain,proto3" json:"elo_gain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpcomingMatch) Reset() {
	*x = UpcomingMatch{}
	mi := &file_player_system_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpcomingMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpcomingMatch) ProtoMessage() {}

func (x *UpcomingMatch) ProtoReflect() protoreflect.Message {
	mi := &file_player_system_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpcomingMatch.ProtoReflect.Descriptor instead.
func (*UpcomingMatch) Descriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{3}
}

func (x *UpcomingMatch) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *UpcomingMatch) GetOpponent() int32 {
	if x != nil {
		return x.Opponent
	}
	return 0
}

func (x *UpcomingMatch) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *UpcomingMatch) GetEloGain() int32 {
	if x != nil {
		return x.EloGain
	}
	return 0
}

type MatchSummary struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Player        int32                  `protobuf:"varint,1,opt,name=player,proto3" json:"player,omitempty"`
	NumMatches    int32                  `protobuf:"varint,2,opt,name=num_matches,json=numMatches,proto3" json:"num_matches,omitempty"`
	AvgEloGain    int32                  `protobuf:"varint,3,opt,name=avg_elo_gain,json=avgEloGain,proto3" json:"avg_elo_gain,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MatchSummary) Reset() {
	*x = MatchSummary{}
	mi := &file_player_system_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MatchSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchSummary) ProtoMessage() {}

func (x *MatchSummary) ProtoReflect() protoreflect.Message {
	mi := &file_player_system_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchSummary.ProtoReflect.Descriptor instead.
func (*MatchSummary) Descriptor() ([]byte, []int) {
	return file_player_system_proto_rawDescGZIP(), []int{4}
}

func (x *MatchSummary) GetPlayer() int32 {
	if x != nil {
		return x.Player
	}
	return 0
}

func (x *MatchSummary) GetNumMatches() int32 {
	if x != nil {
		return x.NumMatches
	}
	return 0
}

func (x *MatchSummary) GetAvgEloGain() int32 {
	if x != nil {
		return x.AvgEloGain
	}
	return 0
}

var File_player_system_proto protoreflect.FileDescriptor

var file_player_system_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x25, 0x0a,
	0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x61,
	0x6e, 0x6b, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6c, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x6c, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0b, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e,
	0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6c, 0x6f, 0x5f, 0x67, 0x61, 0x69, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x6c, 0x6f, 0x47, 0x61, 0x69, 0x6e, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x6c, 0x6f, 0x5f, 0x6c, 0x6f, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x65, 0x6c, 0x6f, 0x4c, 0x6f, 0x73, 0x73, 0x22, 0x7c, 0x0a, 0x0d, 0x55, 0x70,
	0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x6c, 0x6f, 0x5f, 0x67, 0x61, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x65, 0x6c, 0x6f, 0x47, 0x61, 0x69, 0x6e, 0x22, 0x69, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x76, 0x67, 0x5f, 0x65, 0x6c, 0x6f, 0x5f, 0x67, 0x61, 0x69,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x76, 0x67, 0x45, 0x6c, 0x6f, 0x47,
	0x61, 0x69, 0x6e, 0x2a, 0x93, 0x01, 0x0a, 0x04, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x10,
	0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x42, 0x52, 0x4f, 0x4e, 0x5a,
	0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x53, 0x49, 0x4c, 0x56,
	0x45, 0x52, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x47, 0x4f, 0x4c,
	0x44, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x50, 0x4c, 0x41, 0x54,
	0x49, 0x4e, 0x55, 0x4d, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x41, 0x4e, 0x4b, 0x5f, 0x44,
	0x49, 0x41, 0x4d, 0x4f, 0x4e, 0x44, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x41, 0x4e, 0x4b,
	0x5f, 0x4d, 0x41, 0x53, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x41, 0x4e,
	0x4b, 0x5f, 0x46, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x07, 0x2a, 0x52, 0x0a, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x57, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x4c, 0x4f, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x44, 0x52, 0x41, 0x57, 0x10, 0x03, 0x32, 0x8a, 0x02,
	0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x1a, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x00, 0x12,
	0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x0e, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x1a, 0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46, 0x0a, 0x13, 0x53,
	0x65, 0x6e, 0x64, 0x55, 0x70, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x63, 0x6f,
	0x6d, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x1a, 0x14, 0x2e, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x45, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_player_system_proto_rawDescOnce sync.Once
	file_player_system_proto_rawDescData []byte
)

func file_player_system_proto_rawDescGZIP() []byte {
	file_player_system_proto_rawDescOnce.Do(func() {
		file_player_system_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_player_system_proto_rawDesc), len(file_player_system_proto_rawDesc)))
	})
	return file_player_system_proto_rawDescData
}

var file_player_system_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_player_system_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_player_system_proto_goTypes = []any{
	(Rank)(0),             // 0: player.Rank
	(Result)(0),           // 1: player.Result
	(*Player)(nil),        // 2: player.Player
	(*PlayerSummary)(nil), // 3: player.PlayerSummary
	(*MatchResult)(nil),   // 4: player.MatchResult
	(*UpcomingMatch)(nil), // 5: player.UpcomingMatch
	(*MatchSummary)(nil),  // 6: player.MatchSummary
}
var file_player_system_proto_depIdxs = []int32{
	0, // 0: player.PlayerSummary.rank:type_name -> player.Rank
	1, // 1: player.MatchResult.result:type_name -> player.Result
	2, // 2: player.Matches.GetPlayer:input_type -> player.Player
	2, // 3: player.Matches.GetMatchHistory:input_type -> player.Player
	5, // 4: player.Matches.SendUpcomingMatches:input_type -> player.UpcomingMatch
	4, // 5: player.Matches.ProcessMatchResults:input_type -> player.MatchResult
	3, // 6: player.Matches.GetPlayer:output_type -> player.PlayerSummary
	4, // 7: player.Matches.GetMatchHistory:output_type -> player.MatchResult
	6, // 8: player.Matches.SendUpcomingMatches:output_type -> player.MatchSummary
	4, // 9: player.Matches.ProcessMatchResults:output_type -> player.MatchResult
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_player_system_proto_init() }
func file_player_system_proto_init() {
	if File_player_system_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_player_system_proto_rawDesc), len(file_player_system_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_system_proto_goTypes,
		DependencyIndexes: file_player_system_proto_depIdxs,
		EnumInfos:         file_player_system_proto_enumTypes,
		MessageInfos:      file_player_system_proto_msgTypes,
	}.Build()
	File_player_system_proto = out.File
	file_player_system_proto_goTypes = nil
	file_player_system_proto_depIdxs = nil
}
