package main

import (
	"fmt"
	"os"
)

// csv_transformer - CSV processing toolkit
func csv_transformer(path string) {
	fmt.Println("========================================")
	fmt.Println("  CSV-Transformer")
	fmt.Println("  CSV processing toolkit")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	csv_transformer(path)
}
