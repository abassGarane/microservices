package handlers

import (
	"net/http"

	"github.com/abassGarane/microservices/data"
)

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

