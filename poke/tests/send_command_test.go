package tests

import (
	"testing"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
)

func TestSendRequest(t *testing.T) {
	request := lib.Request{
		Url: "http://localhost:3000/api/services/posts/only/clnxegjme0004voigngacw3ye",
		Method: "GET",
		Header: {
			"Accept": "*/*",
			"Content-Type": "application/json",
			"X-API-Key": "wnLYD9F2gZbjckqpONOL4EI0dM7tf7fTeNh+pKrKZgY="
		}
	}

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
	
	write := net.SendRequest()
}
