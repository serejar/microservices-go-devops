FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/server

FROM alpine:3.18

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /app/user-service .

EXPOSE 8081

CMD ["./user-service"]