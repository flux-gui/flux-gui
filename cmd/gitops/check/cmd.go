package check

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/discovery"

	"github.com/flux-gui/flux-gui/cmd/gitops/check/oidcconfig"
	"github.com/flux-gui/flux-gui/cmd/gitops/cmderrors"
	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/pkg/run"
	"github.com/flux-gui/flux-gui/pkg/services/check"
)

func GetCommand(opts *config.Options) *cobra.Command {
	var kubeConfigArgs *genericclioptions.ConfigFlags

	cmd := &cobra.Command{
		Use:   "check",
		Short: "Validates flux compatibility",
		Example: `
# Validate flux and kubernetes compatibility
gitops check
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			kubeConfigArgs = run.GetKubeConfigArgs()
			kubeConfigArgs.AddFlags(cmd.Flags())

			cfg, err := kubeConfigArgs.ToRESTConfig()
			if err != nil {
				return err
			}

			c, err := discovery.NewDiscoveryClientForConfig(cfg)
			if err != nil {
				return cmderrors.ErrGetKubeClient
			}
			output, err := check.KubernetesVersion(c)
			if err != nil {
				return err
			}

			fmt.Println(output)

			return nil
		},
	}

	cmd.AddCommand(oidcconfig.OIDCConfigCommand(opts))

	return cmd
}
