package datafetcher

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/esteanes/expense-manager/datafetcher/upclient"
)

// homePage function to handle requests to the root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	pageSize := int32(30)                                                // int32 | The number of records to return in each page.  (optional)
	filterAccountType := upclient.AccountTypeEnum("SAVER")          // AccountTypeEnum | The type of account for which to return records. This can be used to filter Savers from spending accounts.  (optional)
	filterOwnershipType := upclient.OwnershipTypeEnum("INDIVIDUAL") // OwnershipTypeEnum | The account ownership structure for which to return records. This can be used to filter 2Up accounts from Up accounts.  (optional)
	auth := context.WithValue(context.Background(), upclient.ContextAccessToken, os.Getenv("up-bank-bearer-token"))

	configuration := upclient.NewConfiguration()
	apiClient := upclient.NewAPIClient(configuration)
	resp, r2, err := apiClient.AccountsAPI.AccountsGet(auth).PageSize(pageSize).FilterAccountType(filterAccountType).FilterOwnershipType(filterOwnershipType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r2)
	}
	// response from `AccountsGet`: ListAccountsResponse
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountsGet`: %v\n", resp)

	resp2, r2, err := apiClient.TransactionsAPI.TransactionsGet(auth).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionsAPI.TransactionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r2.Body)
	}
	// response from `TransactionsGet`: ListTransactionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TransactionsAPI.TransactionsGet`: %v\n", resp2)

}

// HandleRequests function to define the routes and start the server
func HandleRequests() {
	http.HandleFunc("/", homePage) // Set the root URL to call homePage function
	http.HandleFunc("/info", getInfo)
	log.Default().Println("Serving request at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080

}
