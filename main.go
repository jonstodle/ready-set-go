package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	symbol := flag.String("symbol", "!", "The trailing symbol of the greeting")
	flag.Parse()

	fmt.Printf("Hello %v%v", strings.Join(flag.Args(), ", "), *symbol)
}
