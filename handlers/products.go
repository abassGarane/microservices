package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/abassGarane/microservices/data"
)

type Products struct{
  l *log.Logger
}

func NewProducts (l *log.Logger) *Products{
  return &Products{l}
}


func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a Read Request for %s :: %s\n", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  d, err := json.Marshal(lp)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
  // Write as Response
  w.Write(d)
}
