package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	validator "github.com/go-playground/validator/v10"
)

// Product defines the structure for an API Product
// swagger:model
type Product struct {

	// the id for the product
	//
	//requred: false
	// min: 1
	ID int `json:"id"`

	// The name for the product
	//
	// requred: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// The description of the product
	//
	// required: true
	//max length: 10000
	Description string `json:"description" validate:"required"`

	// The price of the product

	// requred: true
	// min:0.01
	Price float32 `json:"price" validate:"required,gte=0,lte=10000"`

	// The sku for the product
	//
	// requred:true
	// pattern:[a-z]+-[a-z]+-[1-9]
	SKU string `json:"sku" validate:"required,sku"`

	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}
func validateSKU(f validator.FieldLevel) bool {

	// Sku format abc-dcf-324
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[0-9]+`)
	matches := re.FindAllString(f.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}
func (p *Product) Validate() error {

	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

var ErrorProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	// for i := 0; i < len(productList); i++ {
	// if productList[i].ID == id{
	// return &productList[i],nil
	// }
	// }
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrorProductNotFound
}
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)
	if i == -1 {
		return ErrorProductNotFound
	}
	productList = append(productList[:i], productList[:i+1]...)
	return nil
}

func GetNextID() int {
	id := productList[len(productList)-1].ID
	return id + 1
}

type Products []*Product

func (pp *Products) ToJson(w io.Writer) error {
	en := json.NewEncoder(w)
	return en.Encode(pp)
}

func GetProducts() Products {
	return productList
}

// Simple in-memory products storage
var productList = []*Product{
	{
		ID:          1,
		Name:        "Rice",
		Description: "A bag of rice",
		Price:       2200,
		SKU:         "cdcd34",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Beef",
		Description: "A real pound of beef",
		Price:       800.90,
		SKU:         "cdcd66",
		CreatedAt:   time.Now().UTC().String(),
		UpdatedAt:   time.Now().UTC().String(),
		DeletedAt:   time.Now().UTC().String(),
	},
}
