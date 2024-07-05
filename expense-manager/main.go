package main

import (
	"fmt"
	"os"
	"github.com/esteanes/expense-manager/datafetcher"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	datafetcher.HandleRequests()
}
