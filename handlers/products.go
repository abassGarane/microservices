package handlers

import "time"

type Product struct{
  ID int
  Name string
  Description string
  Price float32
  SKU string
  CreatedAt string
  UpdatedAt string
  DeletedAt string
}


// Simple in-memory products storage
var productList = []*Product{
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
