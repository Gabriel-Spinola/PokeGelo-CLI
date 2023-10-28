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

type Methods uint8

const (
	Created Methods = iota
	GET
	POST
	PUT
	PATCH
)

var (
	reqFilePath string
)

func (method Methods) String() string {
	switch method {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	}

	return "unknown"
}

type Request struct {
	Url    string `json:"url"`
	Method string `json:"method"`
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

	byteValue, _ := io.ReadAll(jsonFile)

	var request Request
	json.Unmarshal(byteValue, &request)

	fmt.Println(request.Url)
	fmt.Println(request.Method)

	defer jsonFile.Close()

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
