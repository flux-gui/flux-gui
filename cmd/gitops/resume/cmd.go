package resume

import (
	"github.com/spf13/cobra"

	"github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/cmd/gitops/resume/terraform"
)

func Command(opts *config.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "resume",
		Short: "Resume a resource",
		Example: `
# Suspend a Terraform object from the "flux-system" namespace
gitops resume terraform --namespace flux-system my-resource
`,
	}

	cmd.AddCommand(terraform.Command(opts))

	return cmd
}
