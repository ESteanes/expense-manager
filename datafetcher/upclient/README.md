# Go API client for upclient

The Up API gives you programmatic access to your balances and
transaction data. You can request past transactions or set up
webhooks to receive real-time events when new transactions hit your
account. It’s new, it’s exciting and it’s just the beginning.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Generator version: 7.7.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://github.com/up-banking/api](https://github.com/up-banking/api)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import upclient "github.com/esteanes/expense-manager/datafetcher/upclient"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `upclient.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), upclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `upclient.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), upclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `upclient.ContextOperationServerIndices` and `upclient.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), upclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), upclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.up.com.au/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsAPI* | [**AccountsGet**](docs/AccountsAPI.md#accountsget) | **Get** /accounts | List accounts
*AccountsAPI* | [**AccountsIdGet**](docs/AccountsAPI.md#accountsidget) | **Get** /accounts/{id} | Retrieve account
*AttachmentsAPI* | [**AttachmentsGet**](docs/AttachmentsAPI.md#attachmentsget) | **Get** /attachments | List attachments
*AttachmentsAPI* | [**AttachmentsIdGet**](docs/AttachmentsAPI.md#attachmentsidget) | **Get** /attachments/{id} | Retrieve attachment
*CategoriesAPI* | [**CategoriesGet**](docs/CategoriesAPI.md#categoriesget) | **Get** /categories | List categories
*CategoriesAPI* | [**CategoriesIdGet**](docs/CategoriesAPI.md#categoriesidget) | **Get** /categories/{id} | Retrieve category
*CategoriesAPI* | [**TransactionsTransactionIdRelationshipsCategoryPatch**](docs/CategoriesAPI.md#transactionstransactionidrelationshipscategorypatch) | **Patch** /transactions/{transactionId}/relationships/category | Categorize transaction
*TagsAPI* | [**TagsGet**](docs/TagsAPI.md#tagsget) | **Get** /tags | List tags
*TagsAPI* | [**TransactionsTransactionIdRelationshipsTagsDelete**](docs/TagsAPI.md#transactionstransactionidrelationshipstagsdelete) | **Delete** /transactions/{transactionId}/relationships/tags | Remove tags from transaction
*TagsAPI* | [**TransactionsTransactionIdRelationshipsTagsPost**](docs/TagsAPI.md#transactionstransactionidrelationshipstagspost) | **Post** /transactions/{transactionId}/relationships/tags | Add tags to transaction
*TransactionsAPI* | [**AccountsAccountIdTransactionsGet**](docs/TransactionsAPI.md#accountsaccountidtransactionsget) | **Get** /accounts/{accountId}/transactions | List transactions by account
*TransactionsAPI* | [**TransactionsGet**](docs/TransactionsAPI.md#transactionsget) | **Get** /transactions | List transactions
*TransactionsAPI* | [**TransactionsIdGet**](docs/TransactionsAPI.md#transactionsidget) | **Get** /transactions/{id} | Retrieve transaction
*UtilityEndpointsAPI* | [**UtilPingGet**](docs/UtilityEndpointsAPI.md#utilpingget) | **Get** /util/ping | Ping
*WebhooksAPI* | [**WebhooksGet**](docs/WebhooksAPI.md#webhooksget) | **Get** /webhooks | List webhooks
*WebhooksAPI* | [**WebhooksIdDelete**](docs/WebhooksAPI.md#webhooksiddelete) | **Delete** /webhooks/{id} | Delete webhook
*WebhooksAPI* | [**WebhooksIdGet**](docs/WebhooksAPI.md#webhooksidget) | **Get** /webhooks/{id} | Retrieve webhook
*WebhooksAPI* | [**WebhooksPost**](docs/WebhooksAPI.md#webhookspost) | **Post** /webhooks | Create webhook
*WebhooksAPI* | [**WebhooksWebhookIdLogsGet**](docs/WebhooksAPI.md#webhookswebhookidlogsget) | **Get** /webhooks/{webhookId}/logs | List webhook logs
*WebhooksAPI* | [**WebhooksWebhookIdPingPost**](docs/WebhooksAPI.md#webhookswebhookidpingpost) | **Post** /webhooks/{webhookId}/ping | Ping webhook


## Documentation For Models

 - [AccountResource](docs/AccountResource.md)
 - [AccountResourceAttributes](docs/AccountResourceAttributes.md)
 - [AccountResourceLinks](docs/AccountResourceLinks.md)
 - [AccountResourceRelationships](docs/AccountResourceRelationships.md)
 - [AccountResourceRelationshipsTransactions](docs/AccountResourceRelationshipsTransactions.md)
 - [AccountResourceRelationshipsTransactionsLinks](docs/AccountResourceRelationshipsTransactionsLinks.md)
 - [AccountTypeEnum](docs/AccountTypeEnum.md)
 - [AttachmentResource](docs/AttachmentResource.md)
 - [AttachmentResourceAttributes](docs/AttachmentResourceAttributes.md)
 - [AttachmentResourceRelationships](docs/AttachmentResourceRelationships.md)
 - [AttachmentResourceRelationshipsTransaction](docs/AttachmentResourceRelationshipsTransaction.md)
 - [AttachmentResourceRelationshipsTransactionData](docs/AttachmentResourceRelationshipsTransactionData.md)
 - [CardPurchaseMethodEnum](docs/CardPurchaseMethodEnum.md)
 - [CardPurchaseMethodObject](docs/CardPurchaseMethodObject.md)
 - [CashbackObject](docs/CashbackObject.md)
 - [CategoryInputResourceIdentifier](docs/CategoryInputResourceIdentifier.md)
 - [CategoryResource](docs/CategoryResource.md)
 - [CategoryResourceAttributes](docs/CategoryResourceAttributes.md)
 - [CategoryResourceRelationships](docs/CategoryResourceRelationships.md)
 - [CategoryResourceRelationshipsChildren](docs/CategoryResourceRelationshipsChildren.md)
 - [CategoryResourceRelationshipsChildrenDataInner](docs/CategoryResourceRelationshipsChildrenDataInner.md)
 - [CategoryResourceRelationshipsParent](docs/CategoryResourceRelationshipsParent.md)
 - [CategoryResourceRelationshipsParentData](docs/CategoryResourceRelationshipsParentData.md)
 - [CreateWebhookRequest](docs/CreateWebhookRequest.md)
 - [CreateWebhookResponse](docs/CreateWebhookResponse.md)
 - [CustomerObject](docs/CustomerObject.md)
 - [ErrorObject](docs/ErrorObject.md)
 - [ErrorObjectSource](docs/ErrorObjectSource.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [GetAccountResponse](docs/GetAccountResponse.md)
 - [GetAttachmentResponse](docs/GetAttachmentResponse.md)
 - [GetCategoryResponse](docs/GetCategoryResponse.md)
 - [GetTransactionResponse](docs/GetTransactionResponse.md)
 - [GetWebhookResponse](docs/GetWebhookResponse.md)
 - [HoldInfoObject](docs/HoldInfoObject.md)
 - [ListAccountsResponse](docs/ListAccountsResponse.md)
 - [ListAccountsResponseLinks](docs/ListAccountsResponseLinks.md)
 - [ListAttachmentsResponse](docs/ListAttachmentsResponse.md)
 - [ListCategoriesResponse](docs/ListCategoriesResponse.md)
 - [ListTagsResponse](docs/ListTagsResponse.md)
 - [ListTransactionsResponse](docs/ListTransactionsResponse.md)
 - [ListWebhookDeliveryLogsResponse](docs/ListWebhookDeliveryLogsResponse.md)
 - [ListWebhooksResponse](docs/ListWebhooksResponse.md)
 - [MoneyObject](docs/MoneyObject.md)
 - [NoteObject](docs/NoteObject.md)
 - [OwnershipTypeEnum](docs/OwnershipTypeEnum.md)
 - [PingResponse](docs/PingResponse.md)
 - [PingResponseMeta](docs/PingResponseMeta.md)
 - [RoundUpObject](docs/RoundUpObject.md)
 - [TagInputResourceIdentifier](docs/TagInputResourceIdentifier.md)
 - [TagResource](docs/TagResource.md)
 - [TransactionResource](docs/TransactionResource.md)
 - [TransactionResourceAttributes](docs/TransactionResourceAttributes.md)
 - [TransactionResourceRelationships](docs/TransactionResourceRelationships.md)
 - [TransactionResourceRelationshipsAccount](docs/TransactionResourceRelationshipsAccount.md)
 - [TransactionResourceRelationshipsAccountData](docs/TransactionResourceRelationshipsAccountData.md)
 - [TransactionResourceRelationshipsAttachment](docs/TransactionResourceRelationshipsAttachment.md)
 - [TransactionResourceRelationshipsAttachmentData](docs/TransactionResourceRelationshipsAttachmentData.md)
 - [TransactionResourceRelationshipsCategory](docs/TransactionResourceRelationshipsCategory.md)
 - [TransactionResourceRelationshipsCategoryLinks](docs/TransactionResourceRelationshipsCategoryLinks.md)
 - [TransactionResourceRelationshipsTags](docs/TransactionResourceRelationshipsTags.md)
 - [TransactionResourceRelationshipsTagsDataInner](docs/TransactionResourceRelationshipsTagsDataInner.md)
 - [TransactionResourceRelationshipsTagsLinks](docs/TransactionResourceRelationshipsTagsLinks.md)
 - [TransactionResourceRelationshipsTransferAccount](docs/TransactionResourceRelationshipsTransferAccount.md)
 - [TransactionResourceRelationshipsTransferAccountData](docs/TransactionResourceRelationshipsTransferAccountData.md)
 - [TransactionStatusEnum](docs/TransactionStatusEnum.md)
 - [UpdateTransactionCategoryRequest](docs/UpdateTransactionCategoryRequest.md)
 - [UpdateTransactionTagsRequest](docs/UpdateTransactionTagsRequest.md)
 - [WebhookDeliveryLogResource](docs/WebhookDeliveryLogResource.md)
 - [WebhookDeliveryLogResourceAttributes](docs/WebhookDeliveryLogResourceAttributes.md)
 - [WebhookDeliveryLogResourceAttributesRequest](docs/WebhookDeliveryLogResourceAttributesRequest.md)
 - [WebhookDeliveryLogResourceAttributesResponse](docs/WebhookDeliveryLogResourceAttributesResponse.md)
 - [WebhookDeliveryLogResourceRelationships](docs/WebhookDeliveryLogResourceRelationships.md)
 - [WebhookDeliveryLogResourceRelationshipsWebhookEvent](docs/WebhookDeliveryLogResourceRelationshipsWebhookEvent.md)
 - [WebhookDeliveryLogResourceRelationshipsWebhookEventData](docs/WebhookDeliveryLogResourceRelationshipsWebhookEventData.md)
 - [WebhookDeliveryStatusEnum](docs/WebhookDeliveryStatusEnum.md)
 - [WebhookEventCallback](docs/WebhookEventCallback.md)
 - [WebhookEventResource](docs/WebhookEventResource.md)
 - [WebhookEventResourceAttributes](docs/WebhookEventResourceAttributes.md)
 - [WebhookEventResourceRelationships](docs/WebhookEventResourceRelationships.md)
 - [WebhookEventResourceRelationshipsWebhook](docs/WebhookEventResourceRelationshipsWebhook.md)
 - [WebhookEventResourceRelationshipsWebhookData](docs/WebhookEventResourceRelationshipsWebhookData.md)
 - [WebhookEventTypeEnum](docs/WebhookEventTypeEnum.md)
 - [WebhookInputResource](docs/WebhookInputResource.md)
 - [WebhookInputResourceAttributes](docs/WebhookInputResourceAttributes.md)
 - [WebhookResource](docs/WebhookResource.md)
 - [WebhookResourceAttributes](docs/WebhookResourceAttributes.md)
 - [WebhookResourceRelationships](docs/WebhookResourceRelationships.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearer_auth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), upclient.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



