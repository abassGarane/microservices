package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main()  {

  // Define a new router / servemux
  mux := http.NewServeMux()

  // map urls to routes
  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    d,err := ioutil.ReadAll(r.Body)
    if err != nil{
    //   w.WriteHeader(http.StatusBadRequest)
    //   w.Write([]byte("Oops! Could not parse request !"))
      http.Error(w, "Could not parse request", http.StatusBadRequest)
      return
    }
    log.Println("Hello from page")
    fmt.Fprintf(w,"Response is %s\n", d)
  })

  http.ListenAndServe(":8080", mux)
}
