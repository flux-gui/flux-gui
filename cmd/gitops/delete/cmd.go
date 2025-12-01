package delete

import (
	"github.com/spf13/cobra"

	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/cmd/gitops/delete/terraform"
)

func GetCommand(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a resource",
	}

	cmd.AddCommand(terraform.Command(opts))

	return cmd
}
