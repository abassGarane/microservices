package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abassGarane/microservices/handlers"
)


func main()  {

  // Define a new router / servemux
  mux := http.NewServeMux()

  l := log.New(os.Stdout, "Product Api ::", log.LstdFlags)

  h := handlers.NewHelloHandler(l)

	mux.Handle("/", h)

	// Create a server
	s := &http.Server{
		IdleTimeout: time.Second * 60,
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	s.ListenAndServe()
}
