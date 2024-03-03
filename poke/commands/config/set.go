package config

import (
	"github.com/spf13/cobra"
)

var (
	defaultOutputPath string
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

// / Set default output path
// / Set
func init() {
	SetCmd.Flags().StringVarP(&defaultOutputPath, "default-output-path", "do", "", "Set the output path for any writting related command")
}
