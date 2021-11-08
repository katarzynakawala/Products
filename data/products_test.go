package data

import "testing"

func TestChecksValidation(t*testing.T) {
	p := &Product{
		Name: "Kacper",
		Price: 10.00,
		SKU: "abs-abss-etet",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}