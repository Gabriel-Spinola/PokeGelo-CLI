/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var (
	reqFilePath string
)

type Request struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get request",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := readJson(reqFilePath)

		if err != nil {
			fmt.Println(err)

			return
		}

		fmt.Println(data)
	},
}

func readJson(pathToRead string) (string, error) {
	jsonFile, err := os.Open(reqFilePath)

	if err != nil {
		return "", err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	// NOTE - Decode json into the result map
	var request Request
	json.Unmarshal(byteValue, &request)

	fmt.Println(request.Url)
	fmt.Println(request.Method)

	fmt.Println("Header Data:")
	for key, value := range request.Headers {
		fmt.Printf("\tKey: %s, Value: %s\n", key, value)
	}

	return "Successfuly opened json", nil
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
