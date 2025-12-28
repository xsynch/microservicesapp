package data

import "testing"

func TestChecksValidation(t *testing.T){
	p := &Product{
		Name: "people",
		Price: 1.00,
		SKU: "abs-jon-str",
	}
	err := p.Validate() 
	if err != nil {
		t.Fatal(err)
	}
}