package server // this is your APP or MAIN package

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mad-czarls/go-api-user/config"
	datasource "github.com/mad-czarls/go-api-user/datasource/redis" // bad nameing
	"github.com/mad-czarls/go-api-user/handler"
	"github.com/mad-czarls/go-api-user/repository/redis"
	"github.com/mad-czarls/go-api-user/router"
)

func Run() (err error) {

	cfg := config.Config{} //! TODO parse from env

	ds := datasource.NewDataSource(cfg)

	repository := redis.NewUserRepository(ds)

	handler := handler.NewHandler(repository)

	//setting up the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.SetUpRouter(cfg, *handler),
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
	//! should be buffered?
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	//listening on the channel - blocks further (code below it) execution until signal won't be passed to channel
	<-quit

	//HANDLING SHUTDOWN BELOW
	//context is for notifying the server that it has 5 seconds to finish tasks
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	//cleanup called AFTER shutdown below will be executed
	defer func() {
		// add custom processes cleanup here: close database, handle queues etc.
		dserr := ds.Close()
		if dserr != nil {
			log.Fatalf("Error when shutting down Redis connection: %v\n", err)
		}

		err = dserr

		// canceling all built-in processes
		cancel()
	}()

	log.Println("Shutting down server...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown with error: %v\n", err)
	}
	log.Println("Server has been shut down")

	return nil
}
