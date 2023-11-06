package tests

import (
	"net/http"
	"testing"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/commands/net"
	"github.com/Gabriel-Spinola/PokeGelo-CLI/lib"
	"github.com/stretchr/testify/assert"
)

var requestPaths = map[lib.HttpMethod]string{
	lib.GET:    "./requests/get.json",
	lib.POST:   "./requests/post.json",
	lib.PATCH:  "./requests/patch.json",
	lib.PUT:    "./requests/put.json",
	lib.DELETE: "./requests/delete.json",
}

func TestSendRequest(t *testing.T) {
	possibleRequests := []string{
		requestPaths[lib.GET],
		requestPaths[lib.POST],
		requestPaths[lib.PATCH],
		requestPaths[lib.PUT],
		requestPaths[lib.DELETE],
	}

	for _, request := range possibleRequests {
		var req lib.Request
		if err := req.UnmarshalJson(request); err != nil {
			t.Errorf("Failed to unmarshal the json sent. In: %v, err: %v", request, err)

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
}

func TestWriteResponseFile(t *testing.T) {

}
