/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/spf13/cobra"
)

var (
	reqFilePath string
)

func sendRequest(incomingReq lib.Request, payload []byte) (string, error) {
	req, err := http.NewRequest(string(incomingReq.Method), incomingReq.GetFormatedURL(), bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	var client *http.Client = &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	stringfiedBody := string(body)
	return stringfiedBody, nil
}

// sendCmd represents the get command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send a request to remote server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var req lib.Request
		if err := req.UnmarshalJson(reqFilePath); err != nil {
			log.Fatal(err)

			return
		}

		fmt.Println(req.GetFormatedURL())

		payload, err := req.MarshalBody()
		if err != nil {
			log.Fatal(err)

			return
		}

		resBody, err := sendRequest(req, payload)
		if err != nil {
			log.Fatal(err)

			return
		}

		fmt.Println("# Response body: \n" + resBody)
	},
}

func setGetFlags() {
	sendCmd.Flags().StringVarP(&reqFilePath, "filepath", "f", "", "The path to the file")

	if err := sendCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(sendCmd)

	setGetFlags()
}
