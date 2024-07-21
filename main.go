package main

import (
	"fmt"
	"log"
	"os"

	"github.com/esteanes/expense-manager/datafetcher"
)

func main() {
	fmt.Println("Starting server on port 8080...")
	logger := log.New(os.Stdout, "datafetcher:", log.Default().Flags())
	datafetcher.HandleRequests(os.Getenv("UP_BANK_TOKEN"), logger)
}
