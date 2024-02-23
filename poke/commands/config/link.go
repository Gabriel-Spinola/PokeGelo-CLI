package config

import (
	"github.com/spf13/cobra"
)

// ConfigCmd represents the info command
var LinkCmd = &cobra.Command{
	Use:   "link",
	Short: "Import specified config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
