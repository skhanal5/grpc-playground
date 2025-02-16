package main

import (
	"context"
	"errors"
	pb "grpc-playground/player"
)

type matchServer struct {

}

func (s *matchServer) GetPlayer(_ context.Context, player *pb.Player) (*pb.PlayerSummary, error) {
	return nil, errors.New("foo")
}

func (s *matchServer) GetMatchHistory(player *pb.Player, stream *pb.Matches_GetMatchHistoryServer) error {
	return errors.New("foo")
}


func (s *matchServer) SendUpcomingMatches(_ context.Context, strea *pb.Matches_SendUpcomingMatchesServer) error {
	return errors.New("foo")
}


func (s *matchServer) ProcessMatchResults(_ context.Context, strea *pb.Matches_SendUpcomingMatchesServer) error {
	return errors.New("foo")
}
