package main

import (
	"context"
	"github.com/abassGarane/microservices/grpc/protos"
	"google.golang.org/grpc"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/abassGarane/microservices/handlers"
	"github.com/go-openapi/runtime/middleware"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Define a new router / servemux
	// mux := http.NewServeMux()

	sm := mux.NewRouter()

	l := log.New(os.Stdout, "Product Api ::", log.LstdFlags)
	// Create a connection
	conn, err := grpc.Dial("localhost:9092")
	if err != nil {
		l.Fatal(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			l.Fatal(err)
		}
	}(conn)
	// Create client
	cc := protos.NewCurrencyClient(conn)

	h := handlers.NewProducts(l, &cc)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	//GET
	GetRouter := sm.Methods(http.MethodGet).Subrouter()
	GetRouter.HandleFunc("/", h.GetProducts)
	GetRouter.Handle("/docs", sh)
	// insecure -> all files on ./ directory are assessible
	GetRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	//PUT
	PutRouter := sm.Methods(http.MethodPut).Subrouter()
	PutRouter.Use(h.MiddlewareProductValidator)
	PutRouter.HandleFunc("/{id:[0-9]+}", h.UpdateProduct)

	//POST
	PostRouter := sm.Methods(http.MethodPost).Subrouter()
	PostRouter.Use(h.MiddlewareProductValidator)
	PostRouter.HandleFunc("/", h.AddProduct)
	// mux.Handle("/products", h).Methods("GET")

	DeleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	DeleteRouter.HandleFunc("/{id:[0-9]+}", h.DeleteProduct)

	//CORS
	cors := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Create a server
	s := &http.Server{
		IdleTimeout:  time.Second * 60,
		Addr:         ":8080",
		Handler:      cors(sm),
		ReadTimeout:  time.Second * 60,
		WriteTimeout: time.Second * 60,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Printf("Commencing graceful shutdown %s", sig)
	ctx, Cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer Cancel()
	_ = s.Shutdown(ctx)
}
