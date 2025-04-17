# Build stage
FROM golang:1.24.2 as builder

WORKDIR /app

# Copy only mod files to leverage cache
COPY inventory/go.mod inventory/go.sum ./
COPY cloud_commons/go.mod cloud_commons/go.sum ./cloud_commons/

RUN go mod download

# Copy full source
COPY inventory/ .
COPY cloud_commons/ ./cloud_commons/

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/.env .env

RUN mkdir -p /app/logs

EXPOSE 50051 

CMD ["./main"]
