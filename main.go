package main

import (
	"net/http"

	"github.com/OJOMB/bitChat/bitServer"
)

func main() {

	// initialize an http.Server instance
	server := bitServer.NewBitServer(
		"0.0.0.0:8080",
		http.NewServeMux(),
	)
	server.ListenAndServe()
}
