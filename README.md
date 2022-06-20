# REST API Payment System
Эмулятор платежного сервиса, написанный на Go.

## Запустить сервис
```
docker run --rm -itp 8080:8080 --name payment-system-go adiletelf/payment-system-go
```

## REST API
Описание API представлено ниже.


### Register
`POST /api/admin/register`
```
curl --request POST \
  --url 'http://localhost:8080/api/admin/register' \
  --header "Content-Type: application/json" \
  --data '{"username": "root", "password": "root"}'
```

### Login
`POST /api/admin/login`
```
curl --request POST \
  --url 'http://localhost:8080/api/admin/login' \
  --header "Content-Type: application/json" \
  --data '{"username": "root", "password": "root"}'
```
```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6MSwiYXV0aG9yaXplZCI6dHJ1ZSwiZXhwIjoxNjU1NzE3MjM2fQ.SUOo4bnIbLDhSLPWMKPeBTFezHNbmYbXEn6ryioIvFg"
}
```


### GetAllTransactions
`GET /api/transactions?userId=xxx&email=xxx`

```
curl --request GET --url 'http://localhost:8080/api/transactions' \
  -d 'userId=1' \
  -d 'token={{TOKEN}}'
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
`POST /api/transaction`

```
curl --request POST \
  --url http://localhost:8080/api/transaction \
  --header 'Content-Type: application/json' \
  --data '{
    "userId": 4,
    "email": "fourth@gmail.com",
    "amount": 99.99,
    "currency": "ruble"
  }'
  -d 'token={{TOKEN}}'
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
`GET /api/transaction/:id`
```
curl --request GET \
  --url 'http://localhost:8080/api/transaction/1?userId=1&email=first%40gmail.com' \
  -d 'token={{TOKEN}}'
```
```
{
  "status" : "created"
}
```

### UpdateTransactionStatus
`PUT /api/transaction/:id`
```
curl --request PUT \
  --url 'http://localhost:8080/api/transaction/3' \
  --header 'Content-Type: application/json' \
  --data '{ "status": "succeed" }' \
  -d 'token={{TOKEN}}'
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
`GET /api/transaction/cancel/:id`
```
curl --request GET \
  --url 'http://localhost:8080/api/transaction/cancel/1' \
  -d 'token={{TOKEN}}'
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

### CurrentAdmin
`GET /api/admin`
```
curl --request GET \
  --url 'http://localhost:8080/api/admin' \
  -d 'token={{TOKEN}}'
```
```
{
  "data": {
    "id": 1,
    "username": "root",
    "password": ""
  },
  "message": "success"
}
```
