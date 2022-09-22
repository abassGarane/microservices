package handlers

import (
	"context"
	"github.com/abassGarane/microservices/grpc/protos"
	"net/http"

	"github.com/abassGarane/microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//  200: productsRespose

// GetProducts returns all Products in the API
func (p ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
	// Get exchange rate
	rr := &protos.RateRequest{
		Base:        protos.Currencies_KSH,
		Destination: protos.Currencies_TSH,
	}
	resp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		p.l.Println("Error connecting to grpc server", err)
	}
	lp := data.GetProducts()
	for i, _ := range lp {
		lp[i].Price = lp[i].Price * resp.Rate
	}
	//Convert to json
	err := lp.ToJson(w)
	if err != nil {
		http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
	}
}
