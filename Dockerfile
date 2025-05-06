# Build stage
FROM golang:1.23 as builder

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate templ files
RUN templ generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o dimbus ./cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/dimbus .

EXPOSE 8080

CMD ["./dimbus"] 