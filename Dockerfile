# Stage 1: build
FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /app
RUN apk add --virtual build-dependencies build-base gcc wget git

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main cmd/server/main.go

# Stage 2: base image
FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /app/main /app/.env /app/
ENV USE_AUTHENTICATION=true
CMD ["/app/main"]
