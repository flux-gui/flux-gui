package set

import (
	"github.com/spf13/cobra"

	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	configCmd "github.com/flux-gui/flux-gui/cmd/gitops/set/config"
)

func SetCommand(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets one or many Flux-GUI CLI configs or resources",
		Example: `
# Enables analytics in the current user's CLI configuration for Flux-GUI
gitops set config analytics true`,
	}

	cmd.AddCommand(configCmd.ConfigCommand(opts))

	return cmd
}
