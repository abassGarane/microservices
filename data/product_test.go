package data

import "testing"

func TestIsUnderValidation(t *testing.T) {
	p := &Product{
		ID:          3,
		Price:       33.55,
		Description: "A wild stuff",
		Name:        "Rice",
		SKU:         "abc-def-123",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
