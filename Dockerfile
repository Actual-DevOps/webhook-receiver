FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o webhook-receiver

FROM alpine:3.21.2

COPY --from=builder /app/webhook-receiver /app/webhook-receiver

WORKDIR /app

EXPOSE 8081

CMD ["./webhook-receiver"]
