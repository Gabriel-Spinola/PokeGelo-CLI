package lib

import (
	"encoding/json"
	"errors"
	"io"
	"os"
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
