package data

import "testing"

func TestCheckValiodation(t *testing.T) {
	p := &Product{
		Name:  "Mor",
		Price: 5,
		SKU:   "abc-def-ghi",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
