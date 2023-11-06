package tests

import (
	"net/http"
	"testing"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/commands/net"
	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/stretchr/testify/assert"
)

var reqFilePath string = "/home/pingola/Desktop/pokegelo-cli/poke/tests/requests/request.json"

func TestSendRequest(t *testing.T) {
	var req lib.Request
	if err := req.UnmarshalJson(reqFilePath); err != nil {
		t.Errorf("Failed to unmarshal the json sent %v", err)

		return
	}

	payload, err := req.MarshalBody()
	if err != nil {
		t.Errorf("Failed to marshal payload body %v", err)

		return
	}

	writer, err := net.SendRequest(req, payload)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, writer.StatusCode)
}
