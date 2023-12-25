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
	"path/filepath"
	"time"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/spf13/cobra"
)

var (
	reqFilePaths        []string
	shouldWriteResponse bool
)

type Response struct {
	StringifiedBody string
	StatusCode      int
}

func writeResponseFile(resp string, respFileName string) (string, error) {
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

	var outputPath = "output" + lib.PATH_SEPARATOR + "output_" + respFileName

	// Write the JSON data to a file
	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)

		return "", err
	}

	return "JSON data has been written to " + outputPath, nil
}

func SendRequest(incomingReq lib.Request, payload []byte, resFileName string) (Response, error) {
	req, err := http.NewRequest(string(incomingReq.Method), incomingReq.GetFormatedURL(), bytes.NewBuffer(payload))
	if err != nil {
		return Response{}, err
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
		log.Fatal("Failed to create client")

		return Response{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	stringfiedBody := string(body)
	if shouldWriteResponse {
		data, err := writeResponseFile(stringfiedBody, resFileName)
		if err != nil {
			return Response{}, err
		}

		return Response{StringifiedBody: data, StatusCode: resp.StatusCode}, nil
	}

	return Response{StringifiedBody: stringfiedBody, StatusCode: resp.StatusCode}, nil
}

func handleSend(filePath string) Response {
	var req lib.Request
	if err := req.UnmarshalJson(filePath); err != nil {
		log.Fatal("Failed to unmarshal file", err)

		return Response{}
	}

	payload, err := req.MarshalBody()
	if err != nil {
		log.Fatal("Failed to marshal body ", err)

		return Response{}
	}

	resBody, err := SendRequest(req, payload, filepath.Base(filePath))
	if err != nil {
		log.Fatal("Failed to send request: ", err)

		return Response{}
	}

	return resBody
}

// sendCmd represents the get command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "send a request to remote server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		fmt.Println(reqFilePaths)

		if len(reqFilePaths) == 1 {
			resBody := handleSend(reqFilePaths[0])

			fmt.Println(resBody.StringifiedBody)
			return
		}

		responseChannel := make(chan Response)
		go lib.ConcurrentFileProcessor(reqFilePaths, responseChannel, handleSend)

		for result := range responseChannel {
			fmt.Println(result)
		}

		<-responseChannel

		fmt.Println("TOOK: ", time.Since(start))
	},
}

func setSendFlags() {
	sendCmd.Flags().StringSliceVarP(&reqFilePaths, "filepath", "f", []string{}, "The path to the file")
	sendCmd.Flags().BoolVarP(&shouldWriteResponse, "writeresponse", "w", false, "should write response file")

	if err := sendCmd.MarkFlagRequired("filepath"); err != nil {
		fmt.Println(err)
	}
}

func init() {
	NetCmd.AddCommand(sendCmd)

	setSendFlags()
}
