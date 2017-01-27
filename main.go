package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	name := "fred"
	_, err := fmt.Println("Hello %s", name)
	if err != nil {
		fmt.Println(errors.Wrap(err, "could not print")) // add external dependency
	}
}
