# Builder
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

# Dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy project
COPY . .

# Build static binary - THIS IS THE CORRECTED LINE
RUN CGO_ENABLED=0 GOOS=linux go build -o udhaar-server ./cmd

# Final image
FROM alpine:3.19

# Install curl for healthchecks
RUN apk add --no-cache curl

RUN adduser -D -g '' appuser
WORKDIR /app

COPY --from=builder /app/udhaar-server ./

# Optional: don't copy .env in production
COPY .env .env

USER appuser
EXPOSE 8080
EXPOSE 9090

CMD ["./udhaar-server"]