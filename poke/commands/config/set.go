package config

import (
	cfg "github.com/Gabriel-Spinola/PokeGelo-CLI/cmd_config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ConfigCmd represents the info command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set specified settings",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var config cfg.CMDConfig

		if config.OutputPath != "" {
			viper.Set(cfg.OutputPathField, config.OutputPath)
		}
	},
}

// / Set default output path
// / Set
func init() {
	SetCmd.Flags().StringVarP(&config.OutputPath, "output-path", "o", "", "Set the output path for any writting related command")
}
