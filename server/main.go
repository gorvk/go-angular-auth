package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorvk/angular-go-auth/server/routes"
)

func main() {
	routes.Setup()
	configureListenAndServe()
}

func configureListenAndServe() {
	server := &http.Server{
		Addr:         ":9090",
		Handler:      nil,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
		}
	}()

	osSignalChan := make(chan os.Signal)
	signal.Notify(osSignalChan, os.Interrupt, os.Kill)
	osSignal := <-osSignalChan
	fmt.Println("gracefully shutting down server due to :", osSignal)
	contextWithTimeout, cancelContext := context.WithTimeout(context.Background(), 30*time.Second)
	cancelContext()
	server.Shutdown(contextWithTimeout)
}
