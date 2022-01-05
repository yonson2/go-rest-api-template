package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yonson2/go-rest-api-template/pkg/server"
	"github.com/yonson2/go-rest-api-template/pkg/storage/sql"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	storage, err := sql.NewSqliteDB("file::memory:?cache=shared")
	if err != nil {
		log.Fatalln("Error starting database: ", err)
	}

	err = storage.Migrate()
	if err != nil {
		log.Fatalln("Error migrating database schema: ", err)
	}

	srv := server.New(storage)

	// channel to listen for errors coming from the listener.
	serverErrors := make(chan error, 1)

	// channel to listen for an interrupt or terminate signal from the OS.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Println("[MAIN] API listening...")
		serverErrors <- srv.StartServer()
	}()

	// blocking run and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return fmt.Errorf("error: starting server: %s", err)

	case <-shutdown:
		log.Println("[MAIN] Start shutdown")

		//give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// asking listener to shutdown
		err := srv.Shutdown(ctx)
		if err != nil {
			log.Printf("[MAIN] Graceful shutdown did not complete in %v : %v", timeout, err)
			err = srv.Close()
		}

		if err != nil {
			return fmt.Errorf("[MAIN] could not stop server gracefully : %v", err)
		}
	}
	return nil
}
