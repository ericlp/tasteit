FROM golang:alpine as builder

ENV GIN_MODE=release

RUN mkdir /app
WORKDIR /app

COPY ./go.mod /app/
COPY ./go.sum /app/
RUN go mod download

COPY . /app
RUN go build -o tasteit ./cmd/tasteit/main.go

FROM alpine

COPY --from=builder /app/tasteit /tasteit
COPY ./internal/db/migrations /internal/db/migrations

ENTRYPOINT ["/tasteit"]