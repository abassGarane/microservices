package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type HelloHandler struct{
    l *log.Logger
}

func NewHelloHandler (l *log.Logger) *HelloHandler{
    return &HelloHandler{l}
}

func (h *HelloHandler) ServeHTTP( w http.ResponseWriter, r *http.Request)  {
    d,err := ioutil.ReadAll(r.Body)
    if err != nil{
    //   w.WriteHeader(http.StatusBadRequest)
    //   w.Write([]byte("Oops! Could not parse request !"))
      http.Error(w, "Could not parse request", http.StatusBadRequest)
      return
    }
    h.l.Printf("Call to %s was made", r.URL)
    fmt.Fprintf(w,"Response is %s\n", d)

}
