/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/spf13/cobra"
)

var (
	reqFilePath         string
	shouldWriteResponse bool
)

func writeResponseFile(resp string) (string, error) {
	// Unmarshal the stringified JSON response body into a map
	var data interface{}

	err := json.Unmarshal([]byte(resp), &data)
	if err != nil {
		return "", err
	}

	// Marshal the map back to JSON with indentation and line breaks
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)

		return "", err
	}

	// Write the JSON data to a file
	err = os.WriteFile("output/output.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)

		return "", err
	}

	return "JSON data has been written to output.json", nil
}

func sendRequest(incomingReq lib.Request, payload []byte) (string, error) {
	req, err := http.NewRequest(string(incomingReq.Method), incomingReq.GetFormatedURL(), bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	for key, header := range incomingReq.Headers {
		req.Header.Set(key, header)
	}

	for _, cookie := range incomingReq.Cookies {
		req.AddCookie(&cookie)
	}

	var client *http.Client = &http.Client{
		Timeout: time.Duration(incomingReq.Timeout),
		Transport: &http.Transport{
			Proxy: func(*http.Request) (*url.URL, error) {
				for _, proxy := range incomingReq.Proxies {
					return &url.URL{Path: proxy}, nil
				}

				return nil, nil
			},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: incomingReq.VerifyTLS},
		},
	}

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
	if shouldWriteResponse {
		return writeResponseFile(stringfiedBody)
	}
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

func setSendFlags() {
	sendCmd.Flags().StringVarP(&reqFilePath, "filepath", "f", "", "The path to the file")
	sendCmd.Flags().BoolVarP(&shouldWriteResponse, "writeresponse", "w", false, "should write response file")

	if err := sendCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(sendCmd)

	setSendFlags()
}
