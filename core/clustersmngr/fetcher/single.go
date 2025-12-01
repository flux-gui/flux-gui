package fetcher

import (
	"context"

	mngr "github.com/flux-gui/flux-gui/core/clustersmngr"
	"github.com/flux-gui/flux-gui/core/clustersmngr/cluster"
)

type singleClusterFetcher struct {
	cluster cluster.Cluster
}

func NewSingleClusterFetcher(cluster cluster.Cluster) mngr.ClusterFetcher {
	return singleClusterFetcher{cluster}
}

func (cf singleClusterFetcher) Fetch(ctx context.Context) ([]cluster.Cluster, error) {
	return []cluster.Cluster{cf.cluster}, nil
}
