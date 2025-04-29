# Build stage
FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dimbus

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/dimbus .

EXPOSE 8080

CMD ["./dimbus"]
