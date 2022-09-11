package data

import (
	"encoding/json"
	"time"
	"io"
)

type Product struct{
  ID int    `json:"id"`
  Name string  `json:"name"`
  Description string `json:"description"`
  Price float32  `json:"price"`
  SKU string  `json:"sku"`
  CreatedAt string `json:"-"`
  UpdatedAt string `json:"-"`
  DeletedAt string `json:"-"`
}

type Products []Product

func (pp *Products)ToJson(w io.Writer) error  {
  en := json.NewEncoder(w)
  return en.Encode(pp)
}

func GetProducts()Products  {
 return productList 
}

// Simple in-memory products storage
var productList = Products{
  {
    ID:1,
    Name:"Rice",
    Description: "A bag of rice",
    Price:2200,
    SKU: "cdcd34",
    CreatedAt: time.Now().UTC().String(),
    UpdatedAt: time.Now().UTC().String(),
    DeletedAt: time.Now().UTC().String(),
  },
 {
    ID:2,
    Name:"Beef",
    Description: "A real pound of beef",
    Price:800.90,
    SKU: "cdcd66",
    CreatedAt: time.Now().UTC().String(),
    UpdatedAt: time.Now().UTC().String(),
    DeletedAt: time.Now().UTC().String(),
  },

}
