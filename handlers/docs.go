package handlers

import "github.com/abassGarane/microservices/data"

// A list of products returned as a response
// swagger:response productsRespose
type productsRespose struct {
	// All products in the Datastore
	// in:body
	Body []data.Product
}

// swagger:response  productResponse
type productResponseWrapper struct {
	// Product struct to Add Product
	// in:body
	// requred:true
	Body data.Product
}

// swagger:parameters updateProduct addProduct
type productParamsWrapper struct {
	// Product struct to update Product  or Add Product
	// in:body
	// requred:true
	Body data.Product
}

// swagger:parameters deleteProduct updateProduct
type ProductIDParameterWrapper struct {
	// ID of product to delete from dataStrore
	// in: path
	// required: true
	ID int `json:"id"`
}

//swagger:response  noContent
type ProductsNoContentWrapper struct {
}
