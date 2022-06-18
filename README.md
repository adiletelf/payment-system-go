# payment-system-go
Эмулятор платежного сервиса, написанный на Go.

## Запустить сервис
`docker run --rm -itp 8080:8080 --name payment-system-go adiletelf/payment-system-go`

## REST API
Описание API.

### GetAllTransactions
`GET /transactions?userId=xxx&email=xxx`

```
curl --request GET --url 'http://localhost:8080/transactions?userId=1'
```
```
[
	{
		"id": 1,
		"userId": 1,
		"email": "first@gmail.com",
		"amount": 100,
		"currency": "ruble",
		"status": "created",
		"createdAt": "2022-06-18T17:26:55.302262334+06:00",
		"updatedAt": "2022-06-18T17:26:55.302262445+06:00"
	},
	{
		"id": 2,
		"userId": 1,
		"email": "first@gmail.com",
		"amount": 125,
		"currency": "ruble",
		"status": "created",
		"createdAt": "2022-06-18T17:26:55.302282363+06:00",
		"updatedAt": "2022-06-18T17:26:55.302282523+06:00"
	},
	{
		"id": 3,
		"userId": 1,
		"email": "first@gmail.com",
		"amount": 126,
		"currency": "ruble",
		"status": "created",
		"createdAt": "2022-06-18T17:26:55.302297+06:00",
		"updatedAt": "2022-06-18T17:26:55.302297071+06:00"
	}
]
```

### CreateTransaction
`POST /transaction`

```
curl --request POST \
  --url http://localhost:8080/transaction \
  --header 'Content-Type: application/json' \
  --data '{
	"userId": 4,
	"email": "fourth@gmail.com",
	"amount": 99.99,
	"currency": "ruble"
}'
```
```
{
   "amount" : 99.99,
   "createdAt" : "2022-06-18T17:35:50.986302371+06:00",
   "currency" : "ruble",
   "email" : "fourth@gmail.com",
   "id" : 7,
   "status" : "failed",
   "updatedAt" : "2022-06-18T17:35:50.986302632+06:00",
   "userId" : 4
}
```

### GetTransactionStatus
`GET /transaction/:id`
```
curl --request GET \
  --url 'http://localhost:8080/transaction/1?userId=1&email=first%40gmail.com'
```
```
{
   "status" : "created"
}
```

### UpdateTransactionStatus
`PUT /transaction/:id`
```
curl --request PUT \
  --url http://localhost:8080/transaction/3 \
  --header 'Content-Type: application/json' \
  --data '{
	"status": "succeed"
}'
```
```
{
	"id": 3,
	"userId": 1,
	"email": "first@gmail.com",
	"amount": 126,
	"currency": "ruble",
	"status": "succeed",
	"createdAt": "2022-06-18T17:35:38.903364565+06:00",
	"updatedAt": "2022-06-18T17:35:38.903364615+06:00"
}
```

### CancelTransaction
`GET /transaction/cancel/:id`
```
curl --request GET \
  --url http://localhost:8080/transaction/cancel/1
```
```
{
	"id": 1,
	"userId": 1,
	"email": "first@gmail.com",
	"amount": 100,
	"currency": "ruble",
	"status": "canceled",
	"createdAt": "2022-06-18T17:35:38.903344827+06:00",
	"updatedAt": "2022-06-18T17:35:38.903344868+06:00"
}
```
