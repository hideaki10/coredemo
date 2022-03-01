package main

import (
	"net/http"

	"github.com/hideaki10/coredemo/framework"
)

func main() {
	// core
	core := framework.NewCore()

	// router
	registerRouter(core)

	//
	server := &http.Server{
		Addr:    ":8080",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
