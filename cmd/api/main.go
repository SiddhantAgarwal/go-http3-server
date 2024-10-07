package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/SiddhantAgarwal/go-http3-server/internal/ping/handler"
	"github.com/gorilla/mux"
	"github.com/quic-go/quic-go/http3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handler.Ping).Methods("GET")

	go func() {
		srv, err := NewHTTP3Server(r)
		if err != nil {
			log.Fatal(err)
		}

		log.Println("http/3 server starting at port :8080")
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("server interrupted, exiting")
}

func NewHTTP3Server(handler http.Handler) (*http3.Server, error) {
	cert, err := tls.LoadX509KeyPair(
		os.Getenv("TLS_CERT"),
		os.Getenv("TLS_KEY"),
	)
	if err != nil {
		return nil, err
	}

	return &http3.Server{
		Addr: "127.0.0.1:8080",
		Port: 8080,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		Handler: handler,
	}, nil
}
