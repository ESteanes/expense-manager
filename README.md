# expense-manager
A Go based expense manager and visulisation with a variety of input sources

Phase 1:
* Use the [Up Bank API](https://github.com/up-banking/api) to fetch existing transactions and store them in a SQLite database
* Create webhooks so that it also fetches real-time transactions
* Create a simple web UI to manually interact with the system using [templ](https://github.com/a-h/templ)
    * Have some basic static charts using [go-echarts](https://github.com/go-echarts/go-echarts) to summarise spending

Phase 2:
* Implement some level of authentication
* Implement this in a runnable docker container

Phase 3:
* Interactable UI
    * Have a UI in which you can interact with the transactions - give them additional metadata such as tags, photos (receipts/invoices)
    * Mutate the graphs - apply filters to change what the graphs output    

## Modules

* expense-manager
    * data-fetcher
    * data-storer
    * data-manager
    * visualisation

### data-fetcher
Fetching spending data from UpBank using their open api spec

### data-storer
Storing the data in a local sqlite database

### data-manager
interface for updating the transactions/metadata

### visualisation
endpoitns to drive the UI


## How to develop

1. Install the OpenAPI generator
`brew install openapi-generator`
2. Fetch the open-api definition JSON file
`curl https://raw.githubusercontent.com/up-banking/api/master/v1/openapi.json`
3. Generate the REST client code
```
openapi-generator-cli generate   -i openapi.json   -g go   -o ./datafetcher/upclient   --additional-properties packageName=upclient   --git-user-id esteanes   --git-repo-id expense-manager/datafetcher/upclient
```
Its recommended to have a file called `.openapi-generator-ignore` inside the /upclient with the following contents:
```
# OpenAPI Generator Ignore

go.mod
go.sum
LICENSE
```
This will stop the generator from generating those files (which will mess up the Go compilation)

4. Generate the TEMPL components
`templ generate`

    * Note if you're finding it hard to run `templ`, it might be because `$GOPATH` is not on your `$PATH` so your computer can't find the executable. In that case either update your path or run `./$GOPATH/templ`

4. Format code
`gofmt -s -w .`

5. Fix up any import changes
`go mod tidy`

6. Make sure that you've set your up bank token in the environment variables in your `.bashrc` or `.zshrc` file - typically located at `~/.<editor>rc`
```
export UP_BANK_TOKEN='<your token here>' 
```
6. Build an executable
```
go build -o expense-manager.exe #windows
go build -o expense-manager #linux
```

7. Run the executable
```
.\expense-manager.exe #windows
./expense-manager #linux
```

## Helpful Commands
Add to your `.bashrc` file.
```
alias expenseManagerBuild="templ generate && gofmt -s -w . && go mod tidy && go build -o expense-manager"
alias expenseManagerBuildRun="templ generate && gofmt -s -w . && go mod tidy && go build -o expense-manager && ./expense-manager"
```
