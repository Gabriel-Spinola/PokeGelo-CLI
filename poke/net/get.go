/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
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

type Header struct {
	APIKey      string `json:"X-API-KEY"`
	ContentType string `json:"Content-Type"`
}

func (req *Request) ValidateRequest() error {
	if req.Method == "" {
		return errors.New("Missing method in request")
	}

	if req.Url == "" {
		return errors.New("missing url in request")
	}

	return nil
}

func (req *Request) UnmarshalJson(path string) error {
	jsonFile, err := os.Open(reqFilePath)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &req)

	if err := req.ValidateRequest(); err != nil {
		return err
	}

	return nil
}

func handleGetRequest(req Request) (string, error) {
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

		fmt.Println("worked: " + req.Method)
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
