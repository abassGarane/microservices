package handlers

import (
	"log"
	"net/http"

	"github.com/abassGarane/microservices/data"
)

type ProductHandler struct{
  l *log.Logger
}

func NewProducts (l *log.Logger) *ProductHandler{
  return &ProductHandler{l}
}


func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a Read Request for %s :: %s\n", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  err := lp.ToJson(w)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
}
