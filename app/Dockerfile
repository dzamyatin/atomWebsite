FROM golang:1.24 AS builder

WORKDIR /app

COPY . .

FROM alpine:3.21.3 AS production

EXPOSE 8502
EXPOSE 8503

WORKDIR /app

COPY --from=builder /exec /app