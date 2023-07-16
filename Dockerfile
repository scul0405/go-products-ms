#Build stage
FROM golang:1.20-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o service cmd/main.go

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/service .
COPY --from=builder /app/config/config.yaml ./config/config.yaml

#EXPOSE 5000
CMD ["/app/service"]