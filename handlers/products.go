package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/abassGarane/microservices/data"
	"github.com/gorilla/mux"
)

type ProductHandler struct{
  l *log.Logger
}

func NewProducts (l *log.Logger) *ProductHandler{
  return &ProductHandler{l}
}


func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
}

func (p *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  err := lp.ToJson(w)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
}

func (p *ProductHandler)AddProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)

	prod := &data.Product{}
	err:= prod.FromJSON(r.Body)
	if err != nil{
		http.Error(w,"Could not unmarshal object", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *ProductHandler)UpdateProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	vars := mux.Vars(r)
	id , err:= strconv.Atoi(vars["id"])

	if err != nil{
		http.Error(w,"Could not parse id", http.StatusBadRequest)
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil{
		http.Error(w,"Could not unmarshal object", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id,prod)
	if err == data.ErrorProductNotFound{
		http.Error(w,"Could not find product", http.StatusBadRequest)
		return
	}
}






