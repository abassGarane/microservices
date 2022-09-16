// Package classification Product Api.
//
// Documentation for Product Api
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
// License: BSD-3
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
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

// A list of products returned as a response
// swagger:response productsRespose
type productsRespose struct{
	// All products in the Datastore
	// in:body
	Body []data.Product
}
// swagger:response  productResponse
type productResponseWrapper struct{
	// Product struct to Add Product
	// in:body
	// requred:true
	Body data.Product
}
// swagger:parameters updateProduct addProduct
type productParamsWrapper struct{
	// Product struct to update Product  or Add Product
	// in:body
	// requred:true
	Body data.Product
}



func NewProducts (l *log.Logger) *ProductHandler{
  return &ProductHandler{l}
}

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsRespose

// GetProducts returns all Products in the API
func (p ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  err := lp.ToJson(w)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
}
// swagger:route POST /products products addProduct
// Adds a product into dataStore
// responses:
//  200: productResponse

// AddProduct adds a new product into the dataStore 
func (p ProductHandler)AddProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	prod := r.Context().Value(ProductKey{}).(data.Product)	
	data.AddProduct(&prod)
}
// swagger:route PUT /products/{id} products updateProduct
// Updates a product in the dataStore
// responses:
//  201: noContent

// UpdateProduct updates a product in database
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

// swagger:route DELETE /products/{id} products deleteProduct
// deletes a product from Datastore
// responses:
//  201: noContent

// DeleteProduct deletes a product from database
func (p ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request )  {
	p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	vars := mux.Vars(r)
	id , err:= strconv.Atoi(vars["id"])
	p.l.Printf("ID is :: %d", id)
	if err != nil{
		http.Error(w,"Could not parse id", http.StatusBadRequest)
		return
	}
	err = data.DeleteProduct(id)
	if err != nil{
		http.Error(w,fmt.Sprintf("Could not delete product:: %s", err), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
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






