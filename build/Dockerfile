FROM golang:1.20-alpine as builder
WORKDIR /app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o snaker ./cmd/snaker/main.go

FROM alpine:3.18

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/snaker ./

EXPOSE 8080

CMD ["./snaker", "--config", "/app/config/config.yml"]