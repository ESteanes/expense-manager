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
	bearerToken := os.Getenv("UP_BANK_TOKEN")
	if bearerToken == "" {
		logger.Fatalln("No bearer token was detected with the key UP_BANK_TOKEN")
	}
	datafetcher.HandleRequests(bearerToken, logger)
}
