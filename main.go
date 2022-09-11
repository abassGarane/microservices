package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	go func ()  {
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Printf("Commencing graceful shutdown %s", sig)
	ctx,_:= context.WithTimeout(context.Background(), time.Second*30)
	s.Shutdown(ctx)
}
