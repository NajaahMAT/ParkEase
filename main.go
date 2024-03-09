package main

import (
	"ParkEase/helper"
	"ParkEase/router"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Base context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Initialize dependencies
	deps := router.InitializeDependencies()

	//router
	routes := router.NewRouter(deps)

	server := &http.Server{
		Addr:    ":8081",
		Handler: routes,
	}

	done := make(chan bool)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v", err)
			helper.ErrorPanic(err)
		}
		close(done)
	}()

	// Listening for shutdown signal in a goroutine
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint

		// Shutdown signal received, cancel the context
		log.Println("Signal received, initiating graceful shutdown")
		cancel()

		// Graceful shutdown
		shutdownCtx, cancelShutdown := context.WithTimeout(ctx, 30*time.Second)
		defer cancelShutdown()
		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
			helper.ErrorPanic(err)
		}

		close(done)
	}()

	// Wait here until the done channel is closed
	<-done
	log.Println("Server stopped")
}
