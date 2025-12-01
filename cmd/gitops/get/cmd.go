package get

import (
	"github.com/spf13/cobra"

	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/cmd/gitops/get/bcrypt"
	configCmd "github.com/flux-gui/flux-gui/cmd/gitops/get/config"
)

func GetCommand(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Display one or many Flux-GUI resources",
		Example: `
# Get the CLI configuration for Flux-GUI
gitops get config

# Generate a hashed secret
PASSWORD="<your password>"
echo -n $PASSWORD | gitops get bcrypt-hash`,
	}

	cmd.AddCommand(bcrypt.HashCommand(opts))
	cmd.AddCommand(configCmd.ConfigCommand(opts))

	return cmd
}
