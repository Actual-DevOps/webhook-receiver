ARG GOLANG_VERSION

FROM golang:${GOLANG_VERSION} AS builder

ARG BUILD_CMD

WORKDIR /app

COPY . .

RUN go mod download && sh -c "${BUILD_CMD}"

FROM alpine:3.21.2

WORKDIR /app

RUN apk --no-cache add ca-certificates \
    && update-ca-certificates \
    && addgroup -g 9999 app \
    && adduser -s /dev/false -u 9999 -D -G app app \
    && mkdir -p /app/config \
    && chown -R app:app /app \
    && rm -rf /var/cache/apk/*

COPY --from=builder /app/webhook-receiver .

EXPOSE 8081

VOLUME ["/app/config"]

USER app

CMD ["./webhook-receiver"]
