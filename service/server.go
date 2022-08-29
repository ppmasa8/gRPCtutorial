package service

import (
	"context"
	"math/rand"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ppmasa8/grpctutorial/pb"
	"github.com/ppmasa8/grpctutorial/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type RockScissorsPaperService struct {
	numberOfGames int32
	numberOfWins  int32
	matchResults  []*pb.MatchResult
}

func NewRockScissorsPaperService() *RockScissorsPaperService {
	return &RockScissorsPaperService{
		numberOfGames: 0,
		numberOfWins: 0,
		matchResults: make([]*pb.MatchResult, 0),
	}
}

func (s *RockScissorsPaperService) PlayGame(ctx context.Context, req *pb.PlayRequest) (*pb.PlayResponse, error) {
	if req.HandShapes == pb.HandShapes_UNKNOWN {
		return nil, status.Errorf(codes.InvalidArgument, "Choose Rock, Paper, or Scissors.")
	}

	opponentHandShapes := pkg.EncodeHandShapes(int32(rand.Intn(3) + 1))

	var result pb.Result
	if req.HandShapes == opponentHandShapes {
		result = pb.Result_DRAW
	} else if (req.HandShapes.Number()-opponentHandShapes.Number()+3)%3 == 1 {
		result = pb.Result_WIN
	} else {
		result = pb.Result_LOSE
	}

	now := time.Now()
	matchResult := &pb.MatchResult{
		YourHandShapes:     req.HandShapes,
		OpponentHandShapes: opponentHandShapes,
		Result:             result,
		CreateTime: &timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		},
	}

	s.numberOfGames = s.numberOfGames + 1
	if result == pb.Result_WIN {
		s.numberOfWins = s.numberOfWins + 1
	}
	s.matchResults = append(s.matchResults, matchResult)

	return &pb.PlayResponse{
		MatchResult: matchResult,
	}, nil
}

func (s *RockScissorsPaperService) ReportMatchResults(ctx context.Context, req *pb.ReportRequest) (*pb.ReportResponse, error) {
	return &pb.ReportResponse{
		Report: &pb.Report{
			NumberOfGames: s.numberOfGames,
			NumberOfWins:  s.numberOfWins,
			MatchResults:  s.matchResults,
		},
	}, nil
}