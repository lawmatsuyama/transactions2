# pismo-transactions

API to create and list transactions and accounts.
This API was built using:
- Go language
- Mongodb as nosql database to record transactions
- API REST

It is using hexagonal architecture:
- domain: contains all domain models and ports/interfaces
- infra: contains all adapters to external services
  - apimanager: contains request/response of apis, routes, handlers and setup of http server
  - containerhelper: contains a helper to up testcontainer-go
  - repository: contains setup for mongodb client and implementations for database operations
- usecases: contains all orchestrations of business logic to domain.
- docs: contains files to swagger
- db: contains script to initialize database
- pkg: contains helpers functions to general purpose

## How to start API

You need to have `docker` and `docker-compose` installed before execute `pismo-transactions` api

In a terminal shell:

- Clone this repository
```
git clone https://github.com/lawmatsuyama/pismo-transactions.git
cd pismo-transactions
```

- In pismo-transactions directory, run the command below to start the API
```
docker-compose up -d
```
It will build the API and create containers to mongodb and pismo-transactions API

Observation: It will use the net ports 27017 and 8080. So be attention to don't lock those ports before run the command above.

## How to test API

With the API up and running, you can access the swagger in some browser:
```
http://localhost:8080/swagger/index.html
```

It contains four operations:
- POST /accounts: Receives account data and registed it in application. 

Request example:
```json
{
  "document_number": "string"
}
```

`document_number` must be a valid CPF

`operation_type_id` field accept the follow integer values:

1 COMPRA A VISTA

2 COMPRA PARCELADA

3 SAQUE

4 PAGAMENTO

- `GET` `/accounts/{accountID}`: returns account info by given accountID.
Request:
`http://localhost:8080/accounts/74211f47-9c6f-4648-88a3-cb7d0614b5fe`

- `POST` `/transactions`: create transaction for given account_id and returns the transaction_id created.
Request:

```json
{
  "account_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
  "amount": 22222.999,
  "description":"teste2", // optional
  "operation_type_id": 4
}
```

`account_id` must exists on database account collection accounts (mongodb). For this project, there are 3 default account_id:

74211f47-9c6f-4648-88a3-cb7d0614b5fe

52814c2d-657b-4e7b-be5c-9f28e59253f88

355daea3-bfdc-41d5-8ecf-c9bcd21f4dbf

- `POST` `/transactions/query`: List transactions by giving filter. All fields you send will be considered for filtering transactions. For example if you send `account_id` and `event_date_from`, it will return transactions by account_id and from the event_date you sent. For the fields `event_date_from` and `event_date_tp`, you have to send in RFC3339 format, like `2022-12-31T23:59:59-03:00`. Also, you can omit or just send zero values in the fields that you want to ignore in filter.

Request example:
```json
{
  "_id": "",
  "amount_greater": 0,
  "amount_less": 0,
  "description": "",
  "operation_type_id": "",  
  "date_from":"2022-12-10T23:19:42-03:00",
  "account_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe"
}
```

It will return 20 transactions for each page. So if you want to view the next page of transactions, you can check response object `paging`, get the `next_page` and then set `page` in the new request.
If response don't return `next_page` in `paging`, it means there are no more transactions to return.


```json
{
  "_id": "",
  "amount_greater": 0,
  "amount_less": 0,
  "description": "",
  "operation_type": "",
  "origin": "",
  "date_from":"2022-12-10T23:19:42-03:00",
  "user_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
  "paging":{
        "page": 3
    }
}
```





## Check transactions and accounts in mongodb
Mongodb is running on the port 27017. It's a replica set with only one container to deal with transaction and session. You can access using any mongodb client passing URI: 
- mongodb://usercore:usercore@mongodb:27017
- mongodb://usercore:usercore@localhost:27017

There is an `account` database with `transactions` and `accounts` collections
