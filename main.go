package main

import (
	"fmt"
	"time"

	"github.com/bradleyfalzon/test/gopherci-dev1/internal/pkg"
	"github.com/pkg/errors"
)

func main() {
	// Importing an internal package ensures GopherCI clones the repository to
	// the correct locations.
	_, err := fmt.Println("Hello %s", pkg.InternalPkgName) // govet: possible formatting directive in Println call
	if err != nil {
		fmt.Println(errors.Wrap(err, "could not print")) // add external dependency
	}
}

func fn1() { // warning: func fn1 is unused (U1000) (unused)
	time.Parse("12345", "")    // warning: parsing time "12345" as "12345": cannot parse "" as "4" (SA1002) (staticcheck)
	time.Now().Sub(time.Now()) // warning: should use time.Since instead of time.Now().Sub (S1012) (gosimple)
}
