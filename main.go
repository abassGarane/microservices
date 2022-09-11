package main

import (
	"log"
	"net/http"
)


func main()  {

  // Define a new router / servemux
  mux := http.NewServeMux()

  // map urls to routes
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("Hello from page")
  })

  http.ListenAndServe(":8080", nil)
}
