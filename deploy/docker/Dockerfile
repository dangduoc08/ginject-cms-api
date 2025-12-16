FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app

RUN apk add --no-cache make

COPY go.mod go.sum Makefile ./

RUN make deps

COPY . .

RUN make build

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/bin .

EXPOSE 3005

CMD ["./ginject-cms-api"]
