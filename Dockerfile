FROM golang:1.18.3-alpine3.16 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGP_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]