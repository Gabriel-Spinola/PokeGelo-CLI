/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	reqFilePath string
)

func handleGetRequest(req request.Request) (string, error) {
	resp, err := http.Get(req.Url)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	stringfiedBody := string(body)
	return stringfiedBody, nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var req Request
		err := req.UnmarshalJson(reqFilePath)

		if err != nil {
			log.Fatal(err)

			return
		}

		resp, err := handleGetRequest(req)

		if err != nil {
			log.Fatal(err)

			return
		}

		fmt.Println("worked: " + resp)
	},
}

func setGetFlags() {
	getCmd.Flags().StringVarP(&reqFilePath, "filepath", "f", "", "The path to the file")

	if err := getCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(getCmd)

	setGetFlags()
}
