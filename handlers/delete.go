package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abassGarane/microservices/data"
	"github.com/gorilla/mux"
)

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


