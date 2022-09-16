package handlers

import (
	"net/http"
	"strconv"

	"github.com/abassGarane/microservices/data"
	"github.com/gorilla/mux"
)

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

