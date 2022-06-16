# payment-system-go
Эмулятор платежного сервиса, написанный на Go.

## Запустить сервис
`docker run --rm -itp 8080:8080 --name payment-system-go adiletelf/payment-system-go`

## REST API
Описание API.

### GetAllTransactions
`GET /transactions?userId=xxx&email=xxx`
```
curl -i -H 'Accept: application/json' -d 'userId=1&email=first@gmail.com' http://localhost:8080/transactions
```

### CreateTransaction
`POST /transaction`

### GetTransactionStatus
`GET /transaction/:id`
```
curl -i -H 'Accept: application/json' http://localhost:8080/transaction/1
```
