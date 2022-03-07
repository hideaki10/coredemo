package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
	go func() {
		server.ListenAndServe()
	}()

	quit := make(chan os.Signal)
	//signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal("server shutdown error: ", err)
	}
}
