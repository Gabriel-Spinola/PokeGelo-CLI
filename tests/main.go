package main

import (
	"fmt"

	"github.com/Gabriel-Spinola/PokeGelo-CLI/tests/server"
)

func main() {
	fmt.Println("Starting test server")

	server.RunServer()
}
