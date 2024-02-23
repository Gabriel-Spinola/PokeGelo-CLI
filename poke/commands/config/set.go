package config

import (
	"github.com/spf13/cobra"
)

// ConfigCmd represents the info command
var SetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set specified settings",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
