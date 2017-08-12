package main

import "fmt"

func main() {
	name := "fred"
	fmt.Println("Hello %s", name)
	// one
	// two
	fmt.Println("Hello %s", name)
	fmt.Println("Hello %s", name) // another
}

// OneUnused ..
func OneUnused(a, b string) string { return a }
