package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	fmt.Printf("Hello %v!", strings.Join(args, ", "))
}
