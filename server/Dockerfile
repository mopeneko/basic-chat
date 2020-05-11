FROM golang:1.14 AS builder

WORKDIR /app

COPY go.mod go.sum /app/
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o app .

FROM alpine:latest

COPY --from=builder /app/app .
CMD [ "./app" ]
