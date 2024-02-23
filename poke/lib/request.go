package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	PATCH   HttpMethod = "PATCH"
	DELETE  HttpMethod = "DELETE"
	HEADER  HttpMethod = "HEADER"
	OPTIONS HttpMethod = "OPTIONS"
)

type Request struct {
	Url            string                 `json:"url"`          // Required
	QueryParams    map[string]string      `json:"query-params"` // Required
	Method         HttpMethod             `json:"method"`
	Headers        map[string]string      `json:"headers"`
	Body           map[string]interface{} `json:"body"`
	Cookies        map[string]http.Cookie `json:"cookies"`
	Timeout        uint                   `json:"timeout"`
	AllowRedirects bool                   `json:"allow-redirects"`
	Proxies        map[string]string      `json:"proxies"`
	VerifyTLS      bool                   `json:"verify"`
}

type Header struct {
	APIKey      string `json:"X-API-KEY"`
	ContentType string `json:"Content-Type"`
}

func (req *Request) ValidateRequest() error {
	validMethods := map[HttpMethod]bool{
		GET:     true,
		POST:    true,
		PUT:     true,
		PATCH:   true,
		HEADER:  true,
		DELETE:  true,
		OPTIONS: true,
	}

	if !validMethods[req.Method] {
		return errors.New("Invalid HTTP method")
	}

	if req.Method == "" {
		return errors.New("Missing method in request")
	}

	if req.Url == "" {
		return errors.New("missing url in request")
	}

	return nil
}

func (req *Request) UnmarshalJson(path string) error {
	jsonFile, err := os.Open(path)

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

func (req *Request) MarshalBody() ([]byte, error) {
	if len(req.Body) <= 0 {
		return nil, nil
	}

	bytesValue, err := json.Marshal(req.Body)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return bytesValue, nil

}

// TODO - accept yml
func (req *Request) UnmarshalYml(path string) error {
	return nil
}

func (req *Request) GetFormatedURL() string {
	if len(req.QueryParams) <= 0 {
		return req.Url
	}

	var index uint = 0

	for key, value := range req.QueryParams {
		if index == 0 {
			req.Url = req.Url + fmt.Sprintf("?%s=%s", key, value)

			index++
			continue
		}

		req.Url = req.Url + fmt.Sprintf("&%s=%s", key, value)

		index++
	}

	return req.Url
}
