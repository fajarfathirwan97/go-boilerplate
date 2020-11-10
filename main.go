package main

import (
	"context"
	"fmt"
	"go-docker/config"
	"go-docker/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Warningln("ENV LOAD FAILED, LOAD FROM ENV")
	}
}

func loadRouter(r *mux.Router) {
	router.API(r)
}
func startServer(r *mux.Router, srv *http.Server) {
	log.Infoln("Starting Server at port", config.GetEnv().AppPort)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	loadEnv()
	r := mux.NewRouter()
	loadRouter(r)
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%v", config.GetEnv().AppPort),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go startServer(r, srv)
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
