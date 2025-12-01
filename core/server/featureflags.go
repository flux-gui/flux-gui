package server

import (
	"context"

	pb "github.com/flux-gui/flux-gui/pkg/api/core"
	"github.com/flux-gui/flux-gui/pkg/featureflags"
)

func (cs *coreServer) GetFeatureFlags(ctx context.Context, msg *pb.GetFeatureFlagsRequest) (*pb.GetFeatureFlagsResponse, error) {
	return &pb.GetFeatureFlagsResponse{
		Flags: featureflags.GetFlags(),
	}, nil
}
