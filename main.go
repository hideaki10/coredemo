package main

import (
	"net/http"

	"github.com/hideaki10/coredemo/framework"
	"github.com/hideaki10/coredemo/router"
)

func main() {
	// core
	core := framework.NewCore()

	// // router
	router.RegisterRouter(core)

	//
	server := &http.Server{
		Addr:    ":8888",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
