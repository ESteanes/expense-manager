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


## Notes

* To generate the open-api interface run the following command: 
```
sudo docker run --rm   -v ${PWD}:/local openapitools/openapi-generator-cli generate   -i /local/openapi.json   -g go   -o /local/expense-manager/up-bank-interface
```

