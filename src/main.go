package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/divo1/psapi"
)

func main() {
	s := psapi.NewService("https://cattaneo.pl/api", "5QINNHUJE64QS4T6AB2G7GPQ3E9KA8AK")

	products := s.GetProducts(1)

	fmt.Println(len(products))
	spew.Dump(products)
	//s.AddProduct(products[0])
}