package handlers

import (
	"context"
	"fmt"
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



func (p ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  err := lp.ToJson(w)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
}

func (p ProductHandler)AddProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	prod := r.Context().Value(ProductKey{}).(data.Product)	
	data.AddProduct(&prod)
}

func (p ProductHandler)UpdateProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	vars := mux.Vars(r)
	id , err:= strconv.Atoi(vars["id"])

	if err != nil{
		http.Error(w,"Could not parse id", http.StatusBadRequest)
		return
	}
	prod := r.Context().Value(ProductKey{}).(data.Product)
	err = data.UpdateProduct(id,&prod)
	if err == data.ErrorProductNotFound{
		http.Error(w,"Could not find product", http.StatusBadRequest)
		return
	}
}

type ProductKey struct{}

func (p ProductHandler) MiddlewareProductValidator( next http.Handler)http.Handler  {
  p.l.Printf("Reached the product middleware")
	middleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil{
  		p.l.Printf("Error deserializing product")
			http.Error(w,"Error Reading product", http.StatusBadRequest)
			return
  	}
  	err = prod.Validate()
  	if err != nil{
  		p.l.Println("Error Validating product", err)
			http.Error(w,fmt.Sprintf("Error Validating product :: %s", err), http.StatusBadRequest)
			return
  	}
  	ctx := context.WithValue(r.Context(), ProductKey{}, prod)
  	req := r.WithContext(ctx)
  	next.ServeHTTP(w,req)
	})

	return middleware
}






