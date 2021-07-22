package api

import (
	"context"
	"log"
	"net/http"
	"time"
)

var (
	BaseURL string = "https://df.dv/"
)

func Launch(readyCheck chan struct{}, shutdown chan struct{}, closeAnnounce chan struct{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten/", ShortenURLHandler)
	mux.HandleFunc("/", RetrieveFullURLHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("Server started")

	readyCheck <- struct{}{}

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	go func() {
		<-shutdown
		if err := srv.Shutdown(ctxShutDown); err != nil {
			log.Fatalf("server Shutdown Failed:%+s", err)
		}
		log.Printf("Server exited properly")
		close(closeAnnounce)
	}()
}
