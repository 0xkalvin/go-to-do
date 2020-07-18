FROM golang:1.14-alpine as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build src/server/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080
