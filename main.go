package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"optimistic-ui-demo/middleware"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// force failure of requests?
	var teapot bool
	flag.BoolVar(&teapot, "teapot", false, "Is the server a teapot?")
	flag.Parse()

	// initialise the pgx middleware
	pgxMiddleware, err := middleware.NewPgxMiddleware()
	if err != nil {
		log.Fatal(err)
	}

	// setup router
	router := SetupRouter(pgxMiddleware, teapot)

	// handle graceful shutdown
	srv := &http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown the server
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatal("Forced Server Shutdown:", err)
	}

	// server is shutdown so close the pgx connection pool
	pgxMiddleware.Close()

	log.Println("Server exiting")
}
