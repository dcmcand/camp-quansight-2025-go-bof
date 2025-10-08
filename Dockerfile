# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /build

# Copy go mod files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w' -o app .

# Final stage
FROM scratch

# Copy the binary from builder
COPY --from=builder /build/app /app

# Expose port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/app"]
