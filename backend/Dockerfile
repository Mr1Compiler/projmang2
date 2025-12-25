# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install git for fetching dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

# Install golang-migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /api cmd/api/main.go

# Run stage
FROM alpine:latest

WORKDIR /

# Install certificates for SSL DB connections
RUN apk --no-cache add ca-certificates

COPY --from=builder /api /api
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/migrations /migrations
COPY start.sh /start.sh

RUN chmod +x /start.sh

# Expose the port (Railway will map this)
EXPOSE 8080

CMD ["/start.sh"]
