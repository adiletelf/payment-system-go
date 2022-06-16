FROM golang:1.18.3-bullseye AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o build/main cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./build/main"]
