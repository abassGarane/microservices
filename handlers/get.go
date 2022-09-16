package handlers

import (
	"net/http"

	"github.com/abassGarane/microservices/data"
)

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

