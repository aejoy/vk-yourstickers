FROM golang:1.23.4-alpine3.20 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY .. .
RUN go build -v -ldflags "-s -w" -o /usr/local/bin/app ./cmd/main.go

FROM alpine:3.20

WORKDIR /usr/local/bin

COPY --from=builder /usr/local/bin/app /usr/local/bin/app
COPY --from=builder /usr/src/app/migrations /usr/local/bin/migrations

CMD ["app"]