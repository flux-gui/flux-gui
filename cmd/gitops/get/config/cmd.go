package config

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	cfg "github.com/flux-gui/flux-gui/cmd/gitops/config"
	"github.com/flux-gui/flux-gui/pkg/config"
	"github.com/flux-gui/flux-gui/pkg/logger"
)

func ConfigCommand(opts *cfg.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Prints out the CLI configuration for Flux-GUI",
		Example: `
# Prints out the CLI configuration for Flux-GUI
gitops get config`,
		SilenceUsage:      true,
		SilenceErrors:     true,
		RunE:              getConfigCommandRunE(opts),
		DisableAutoGenTag: true,
	}

	return cmd
}

func getConfigCommandRunE(opts *cfg.Options) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var err error

		log := logger.NewCLILogger(os.Stdout)

		gitopsConfig, err := config.GetConfig(false)
		if err != nil {
			log.Warningf(config.WrongConfigFormatMsg)
			return err
		}

		log.Successf("Your CLI configuration for Flux-GUI:")

		cfgStr := gitopsConfig.String()
		fmt.Println(cfgStr)

		return nil
	}
}
