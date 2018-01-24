package main

import (
	"fmt"

	"github.com/codermarcel/go_test_travis/company"
)

func main() {
	c := company.New("TheLions", "https://www.thelions.com")

	fmt.Printf("%+v\n", c)
}
