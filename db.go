package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

const PRODUCT_DB_FILE = "products.yaml"

const (
	Ven     = "ven"
	Manatee = "Manatee"
	Heino   = "heino"
	Dupont  = "dupont"
	Dupond  = "dupond"
)

const (
	PRODUCT_MISSING = "Product missing"
)

type ProductError struct {
	name  string
	cause string
}

func (e ProductError) Error() string {
	return fmt.Sprintf("%s: %s", e.cause, e.name)
}

type Product struct {
	Name string `yaml:"name"`
	// NOTE: Prices are given as integeres of øres
	Cost           int      `yaml:"cost"`
	SagioPrice     int      `yaml:"sagio-price"`
	NayaxPrice     int      `yaml:"yanax-price"`
	MobilepayPrice int      `yaml:"mobilepay-price"`
	Machine        string   `yaml:"machine"`
	Tags           []string `yaml:"tags"`
}

type Costs struct {
	Cost           int
	SagioPrice     int
	NayaxPrice     int
	MobilepayPrice int
}

type Db map[string]Product

func readDb() (Db, error) {
	file_path := path.Join(os.Getenv("KANTINE_DB"), PRODUCT_DB_FILE)
	file, err := os.Open(file_path)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	products := make(map[string]Product)
	err = yaml.Unmarshal([]byte(data), &products)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func writeDb(products Db) error {
	file_path := path.Join(os.Getenv("KANTINE_DB"), PRODUCT_DB_FILE)
	file, err := os.Open(file_path)

	if err != nil {
		return err
	}

	data, err := yaml.Marshal(&products)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func getProductPrice(name string) (Costs, error) {
	products, err := readDb()
	if err != nil {
		return Costs{}, err
	}

	product, ok := products[name]
	if !ok {
		return Costs{}, ProductError{name: name, cause: PRODUCT_MISSING}
	}

	costs := Costs{
		Cost:           product.Cost,
		SagioPrice:     product.SagioPrice,
		NayaxPrice:     product.NayaxPrice,
		MobilepayPrice: product.MobilepayPrice,
	}

	return costs, nil
}

func setProductPrice(name string, cost int) error {
	products, err := readDb()
	if err != nil {
		return err
	}

	product, ok := products[name]
	if !ok {
		return ProductError{name: name, cause: PRODUCT_MISSING}
	}

	// TODO: Calculate other prices
	product.Cost = cost
	return writeDb(products)
}
