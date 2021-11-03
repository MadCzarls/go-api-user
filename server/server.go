package server

import (
	"context"
	"github.com/mad-czarls/go-api-user/container"
	"github.com/mad-czarls/go-api-user/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	redisDataSource := container.GetRedisDataSource()

	//setting up the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.SetUpRouter(),
	}

	//starting server; it is done in goroutine since we want code below (shutdown handling) to be executed
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", server.Addr)

	//create a channel to listen to interrupt signals from OS (e.g. SIGINT = ctrl+c on Linux)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//listening on the channel - blocks further (below code's) execution until signal won't be passed to channel
	<-quit

	//HANDLING SHUTDOWN BELOW
	//context is for notifying the server that it has 5 seconds to finish tasks
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//cleanup called AFTER shutdown below will be executed
	defer func() {
		// add custom processes cleanup here: close database, handle queues etc.
		err := redisDataSource.Close()
		if err != nil {
			log.Fatalf("Error when shutting down Redis connection: %v\n", err)
		}
		// canceling all built-in processes
		cancel()
	}()

	log.Println("Shutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown with error: %v\n", err)
	}
	log.Println("Server has been shut down")
}
