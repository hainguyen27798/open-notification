FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/server

FROM scratch

COPY --from=builder /build/backend /
COPY --from=builder /build/templates /templates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./backend"]
