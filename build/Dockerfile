FROM golang:1.20-alpine as builder
WORKDIR /app

COPY ./ ./

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o bin ./cmd/snaker/main.go

FROM alpine:3.15
WORKDIR /
COPY --from=builder / /
EXPOSE 80

CMD ["./app/bin"]