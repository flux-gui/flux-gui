package logs

import (
	"github.com/spf13/cobra"

	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/cmd/gitops/logs/terraform"
)

func GetCommand(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logs",
		Short: "Get logs for a resource",
	}

	cmd.AddCommand(terraform.Command(opts))

	return cmd
}
