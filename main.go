package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abassGarane/microservices/handlers"
	"github.com/gorilla/mux"
)


func main()  {

  // Define a new router / servemux
  // mux := http.NewServeMux()
  
  mux := mux.NewRouter()

  l := log.New(os.Stdout, "Product Api ::", log.LstdFlags)

  h := handlers.NewProducts(l)

  //GET
  GetRouter := mux.Methods(http.MethodGet).Subrouter()
  GetRouter.HandleFunc("/", h.GetProducts)

  //PUT
  PutRouter := mux.Methods(http.MethodPut).Subrouter()
  PutRouter.Use(h.MiddlewareProductValidator)
  PutRouter.HandleFunc("/{id:[0-9]+}", h.UpdateProduct)

  //POST 
  PostRouter := mux.Methods(http.MethodPost).Subrouter()
  PostRouter.Use(h.MiddlewareProductValidator)
  PostRouter.HandleFunc("/", h.AddProduct)
	// mux.Handle("/products", h).Methods("GET")


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
