package main

import (
	"net/http"

	"github.com/hideaki10/coredemo/framework"
	"github.com/hideaki10/coredemo/middleware"
	"github.com/hideaki10/coredemo/router"
)

func main() {
	// core
	core := framework.NewCore()

	core.Use(
		middleware.Test1(),
		middleware.Test2(),
	)
	// // router
	router.RegisterRouter(core)

	subjectApi := core.Group("/subject")
	subjectApi.Use(middleware.Test3())

	//
	server := &http.Server{
		Addr:    ":8888",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
