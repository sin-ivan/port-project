package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	router "github.com/port-project/api-service/port/router"
)

func setupServer() {
	r := router.NewRouter()
	srv := &http.Server{
		Handler:      r.Handler,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Graceful Shutdown
	waitForShutdown(srv)

	log.Println("Starting service")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}

func setupLogging() {
	// Configure Logging
	logFileLocation := os.Getenv("LOG_FILE_LOCATION")
	if logFileLocation != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   logFileLocation,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

func main() {
	setupLogging()
	setupServer()
}

func waitForShutdown(srv *http.Server) {

	go func() {
		interruptChan := make(chan os.Signal, 1)
		signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal.
		<-interruptChan

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		srv.Shutdown(ctx)

		log.Println("Shutting down")
		println("Shut down server")
		os.Exit(0)
	}()
}