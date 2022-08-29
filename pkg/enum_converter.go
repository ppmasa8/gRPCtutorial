package pkg

import "github.com/ppmasa8/grpctutorial/pb"

func EncodeHandShapes(n int32) pb.HandShapes {
	switch n {
	case 1:
		return pb.HandShapes_ROCK
	case 2:
		return pb.HandShapes_SCISSORS
	case 3 :
		return pb.HandShapes_PAPER
	default:
		return pb.HandShapes_UNKNOWN
	}
}