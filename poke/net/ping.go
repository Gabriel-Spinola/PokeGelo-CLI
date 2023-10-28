/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote URL and returns the response",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := ping(urlPath)

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(res)
	},
}

func ping(domain string) (int, error) {
	var url string = "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return 0, err
	}

	res, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	res.Body.Close()

	return res.StatusCode, nil
}

func setPingFlags() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(pingCmd)

	setPingFlags()
}
