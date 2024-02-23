/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "All things information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {}
