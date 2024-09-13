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

### Dependencies
1. Install the OpenAPI generator
`sudo npm install openapi-generator`
2. Install Templ and air
* [Air](https://github.com/air-verse/air?tab=readme-ov-file#installation)
* [Templ](https://templ.guide/quick-start/installation)
3. Install Tailwind CSS
Install using npm as we're already using NPM for the API client generation.
[Tailwind Installation](https://tailwindcss.com/docs/installation)
4. Fetch the open-api definition JSON file
`curl https://raw.githubusercontent.com/up-banking/api/master/v1/openapi.json`
5. Make sure that you've set your up bank token in the environment variables in your `.bashrc` or `.zshrc` file - typically located at `~/.<editor>rc`
[Get your token from Up Bank](https://api.up.com.au/getting_started)
```
export UP_BANK_TOKEN='<your token here>' 
```
### Automatically
Most of the comands have been simplified down to a really simple make file

Doing active development, run:
```bash
make watch  
```
as this will create a hot-reloading environment which will automatically update CSS, Templ generated HTML and Go logic while editing.

If you need to regenerate the upbank API client run:
```
make upclient-generate
```

If you want to generate a hard executable file then run:
```
make build
```

### Manually
1. Generate the REST client code
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

1. Generate the TEMPL components
`templ generate`

    * Note if you're finding it hard to run `templ`, it might be because `$GOPATH` is not on your `$PATH` so your computer can't find the executable. In that case either update your path or run `./$GOPATH/templ`

1. Format code
`gofmt -s -w .`

1. Fix up any import changes
`go mod tidy`


1. Build an executable
```
go build -o expense-manager
```

1. Run the executable
```
./expense-manager
```

## Helpful Commands
Add to your `.bashrc` file.
```
alias expenseManagerBuild="templ generate && gofmt -s -w . && go mod tidy && go build -o expense-manager"
alias expenseManagerBuildRun="templ generate && gofmt -s -w . && go mod tidy && go build -o expense-manager && ./expense-manager"
```
