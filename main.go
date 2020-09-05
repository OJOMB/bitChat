package main

import (
	"flag"
	"net/http"

	"github.com/OJOMB/bitChat/bitlog"
	"github.com/OJOMB/bitChat/bitrepo"
	"github.com/OJOMB/bitChat/bitserver"
	"github.com/OJOMB/bitChat/bitconfig"
)

flag.String("env", "dev", "Set the environment in which to run the server")

func main() {
	flag.Parse()
	config = bitconfig.ConfigMap(*env)

	if config.Logger == "local" {
		logger = bitlog.NewLocalLogger()
	} else {
		panic(fmt.Errorf("missing logger config"))
	}

	// intitalize repo
	if config.DBType == "mongodb" {
		client, err := bitrepo.ConnectToMongoDB(config.DBHost, config.DBPort)
		if err != nil {
			logger.Fatal(err)
		}	
		repo := bitrepo.GetMongodb()
	} else {
		panic(fmt.Errorf("missing repo config"))
	}

	// initialize an http.Server instance
	server := bitserver.NewBitServer(
		config,
		http.NewServeMux(),
		repo,
		logger
	)
	server.ListenAndServe()
}
