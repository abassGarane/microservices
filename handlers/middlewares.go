package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/abassGarane/microservices/data"
)
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


