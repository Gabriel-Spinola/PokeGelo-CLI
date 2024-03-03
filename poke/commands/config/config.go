package config

import (
	cfg "github.com/Gabriel-Spinola/PokeGelo-CLI/cmd_config"
	"github.com/spf13/cobra"
)

var (
	config cfg.CMDConfig
)

// ConfigCmd represents the info command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration pallete",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ConfigCmd.AddCommand(NewCmd)
	ConfigCmd.AddCommand(LinkCmd)
	ConfigCmd.AddCommand(SetCmd)
}
