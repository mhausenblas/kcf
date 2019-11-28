package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mhausenblas/kcf/pkg/fleet"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	// KubernetesConfigFlags makes the config flags available globally
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

// RootCmd runs the fleet root command
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "fleet",
		Short:         "",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := fleet.RunFleetCommand(KubernetesConfigFlags); err != nil {
				return errors.Cause(err)
			}
			return nil
		},
	}
	cobra.OnInitialize(initConfig)
	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

// InitAndExecute sets up and executes the root command
func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}
