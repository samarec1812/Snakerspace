FROM golang:1.20-alpine as builder
WORKDIR /

COPY ./ ./

RUN go mod download
RUN go build -o bin/app ./cmd/snaker/main.go

FROM alpine:3.15
WORKDIR /
COPY --from=builder / /
EXPOSE 80

CMD ["./bin/app"]