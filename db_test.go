package main

import (
	"os"
	"testing"
)

const test_src = `
productA:
  name: Product A
  cost: 1400
  sagio-price: 1500
  nayax-price: 1500
  mobilepay-price: 1600
  machine: manatee
  tags:
    - soft-drink

productB:
  name: Product B
  cost: 1403
  sagio-price: 1503
  nayax-price: 1503
  mobilepay-price: 1603
  machine: ven
  tags:
    - beer

`

func SetupTest() error {
	os.Setenv("KANTINE_DB", "./test.yaml")
	file, err := os.Open("./test.yaml")
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(test_src))
	if err != nil {
		return err
	}
	return nil
}

func TestRead(t *testing.T) {
	if SetupTest() != nil {
		t.Failed()
	}

	db, err := readDb()
	if err != nil {
		t.Failed()
	}

	a, ok := db["productA"]
	if !ok {
		t.Failed()
	}
	if a.Name != "Product A" {
		t.Failed()
	}
	if a.Cost != 1400 {
		t.Failed()
	}
	if a.SagioPrice != 1500 {
		t.Failed()
	}
	if a.NayaxPrice != 1500 {
		t.Failed()
	}
	if a.MobilepayPrice != 1600 {
		t.Failed()
	}
	if a.Machine != "manatee" {
		t.Failed()
	}
	if len(a.Tags) != 1 || a.Tags[0] != "soft-drink" {
		t.Failed()
	}

	b, ok := db["productB"]
	if !ok {
		t.Failed()
	}
	if b.Name != "Product B" {
		t.Failed()
	}
	if b.Cost != 1403 {
		t.Failed()
	}
	if b.SagioPrice != 1503 {
		t.Failed()
	}
	if b.NayaxPrice != 1503 {
		t.Failed()
	}
	if b.MobilepayPrice != 1603 {
		t.Failed()
	}
	if b.Machine != "ven" {
		t.Failed()
	}
	if len(b.Tags) != 1 || b.Tags[0] != "beer" {
		t.Failed()
	}
}

func TestReadWrite(t *testing.T) {
	if SetupTest() != nil {
		t.Failed()
	}

	db, err := readDb()
	if err != nil {
		t.Failed()
	}

	err = writeDb(db)
	if err != nil {
		t.Failed()
	}
	TestRead(t)
}

func TestSetProductPrice(t *testing.T) {
	if SetupTest() != nil {
		t.Failed()
	}

	setProductPrice("productA", 1000)
	db, err := readDb()
	if err != nil {
		t.Failed()
	}

	a, ok := db["productA"]
	if !ok {
		t.Failed()
	}
	if a.Cost != 1000 {
		t.Failed()
	}
}

func TestGetProductPrice(t *testing.T) {
	if SetupTest() != nil {
		t.Failed()
	}

	a, err := getProductPrice("productA")
	if err != nil {
		t.Failed()
	}

	if a.Cost != 1400 {
		t.Failed()
	}
	if a.SagioPrice != 1500 {
		t.Failed()
	}
	if a.NayaxPrice != 1500 {
		t.Failed()
	}
	if a.MobilepayPrice != 1600 {
		t.Failed()
	}
}

func TestInvalidPrice(t *testing.T) {
	if SetupTest() != nil {
		t.Failed()
	}

	setProductPrice("productA", 37)
	db, err := readDb()
	if err.Error() != "Invalid price:" {
		t.Failed()
	}

	a, ok := db["productA"]
	if !ok {
		t.Failed()
	}
	if a.Cost != 1000 {
		t.Failed()
	}
}
