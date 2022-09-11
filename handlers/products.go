package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/abassGarane/microservices/data"
)

type ProductHandler struct{
  l *log.Logger
}

func NewProducts (l *log.Logger) *ProductHandler{
  return &ProductHandler{l}
}


func (p *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
  if r.Method == http.MethodGet{
    p.getProducts(w,r)
    return
  }

  if r.Method == http.MethodPost{
  	p.addProduct(w,r)
  	return
  }
  if r.Method == http.MethodPut{
  	p := r.URL.Path
  	// re, _ := regexp.Compile("/(.*)")
  	// values := re.FindStringSubmatch(p)
  	reg := regexp.MustCompile(`/([0-9]+)`)
  	values := reg.FindAllStringSubmatch(p, -1)
  	if len(values) !=1 || len(values[0][1]) != 1 {
  		http.Error(w,"Error parsing url", http.StatusBadRequest)
  		return
  	}
  	id,_ := strconv.Atoi(values[0][1])
		fmt.Printf("Parsed id :: %d", id)

  }
	// Catch all router
  w.WriteHeader(http.StatusNotImplemented)
  p.l.Printf("Recieved a non-get Request for %s :: %s\n", r.Method, r.URL)
}

func (p *ProductHandler) getProducts(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)
  lp := data.GetProducts() 
  //Convert to json
  err := lp.ToJson(w)
  if err != nil{
    http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
  }
}

func (p *ProductHandler)addProduct(w http.ResponseWriter, r *http.Request)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)

	prod := &data.Product{}
	err:= prod.FromJSON(r.Body)
	if err != nil{
		http.Error(w,"Could not unmarshal object", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}

func (p *ProductHandler)updateProduct(w http.ResponseWriter, r *http.Request, id int)  {
  p.l.Printf("Recieved a %s request from :: %s", r.Method, r.URL)

	prod := &data.Product{}
	err:= prod.FromJSON(r.Body)
	if err != nil{
		http.Error(w,"Could not unmarshal object", http.StatusBadRequest)
	}
	data.AddProduct(prod)
}






