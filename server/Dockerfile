FROM golang:1.14 AS builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .
RUN go build -o app .

FROM alpine:latest

COPY --from=builder /app/app .
CMD [ "./app" ]
