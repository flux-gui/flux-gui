package server_test

import (
	"testing"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"

	"github.com/flux-gui/flux-gui/core/clustersmngr"
	"github.com/flux-gui/flux-gui/core/clustersmngr/cluster"
	"github.com/flux-gui/flux-gui/core/clustersmngr/fetcher"
	"github.com/flux-gui/flux-gui/core/server"
	pb "github.com/flux-gui/flux-gui/pkg/api/core"
	"github.com/flux-gui/flux-gui/pkg/featureflags"
	"github.com/flux-gui/flux-gui/pkg/health"
	"github.com/flux-gui/flux-gui/pkg/kube"
)

func TestGetFeatureFlags(t *testing.T) {
	RegisterFailHandler(Fail)

	ctx := t.Context()

	featureflags.Set("this is a flag", "you won't find it anywhere else")

	scheme, err := kube.CreateScheme()
	if err != nil {
		t.Fatal(err)
	}

	cluster, err := cluster.NewSingleCluster("Default", k8sEnv.Rest, scheme, kube.UserPrefixes{})
	if err != nil {
		t.Fatal(err)
	}

	clustersManager := clustersmngr.NewClustersManager([]clustersmngr.ClusterFetcher{
		fetcher.NewSingleClusterFetcher(cluster),
	}, &nsChecker, logr.Discard())

	hc := health.NewHealthChecker()

	cfg, err := server.NewCoreConfig(logr.Discard(), &rest.Config{}, "test", clustersManager, hc)
	Expect(err).NotTo(HaveOccurred())
	coreSrv, err := server.NewCoreServer(ctx, cfg)
	Expect(err).NotTo(HaveOccurred())

	resp, err := coreSrv.GetFeatureFlags(t.Context(), &pb.GetFeatureFlagsRequest{})
	Expect(err).NotTo(HaveOccurred())
	Expect(resp.Flags).To(HaveKeyWithValue("this is a flag", "you won't find it anywhere else"))
}
