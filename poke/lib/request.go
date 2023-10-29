package lib

import (
	"encoding/json"
	"errors"
	"io"
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
	// Add more methods as needed
)

type Request struct {
	Url            string                 `json:"url"`    // Required
	Params         map[string]interface{} `json:"params"` // Required
	Method         HttpMethod             `json:"method"`
	Headers        map[string]string      `json:"headers"`
	Body           map[string]interface{} `json:"body"`
	Cookies        map[string]interface{} `json:"cookies"`
	Timeout        float32                `json:"timeout"`
	AllowRedirects bool                   `json:"allow_redirects"`
	Proxies        map[string]string      `json:"proxies"`
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

// TODO - accept yml
func (req *Request) UnmarshalYml(path string) error {
	return nil
}
