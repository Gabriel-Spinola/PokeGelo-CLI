/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/commands/config"
	"github.com/Gabriel-Spinola/PokeGelo-CLI/commands/info"
	"github.com/Gabriel-Spinola/PokeGelo-CLI/commands/net"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "poke",
	Short: "Utilities toolset",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func addSubCommadsPalletes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
	rootCmd.AddCommand(config.ConfigCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addSubCommadsPalletes()
}
