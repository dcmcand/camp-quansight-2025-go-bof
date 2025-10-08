# Go Tooling Demo - Birds of a Feather Session

A simple Go API application demonstrating Go tooling, build systems, and CI/CD practices.

## What This Project Demonstrates

- **Basic Go API**: Simple HTTP server with JSON endpoints
- **Go Tooling**: `go mod`, `go test`, `go vet`, `go build`
- **Multi-platform Builds**: GoReleaser for Windows, macOS, and Linux
- **Docker**: Multistage builds with minimal container images
- **CI/CD**: GitHub Actions for testing and releases

## API Endpoints

### POST `/`

Accepts a JSON payload with a number and returns whether it's even.

**Request:**

```bash
curl -X POST http://localhost:8080/ \
    -H "Content-Type: application/json" \
    -d '{"number": 42}'
```

**Response:**

```json
{
  "is_even": true
}
```

### GET `/health`

Health check endpoint.

**Response:**

```json
{
  "status": "ok"
}
```

## Prerequisites

- Go 1.23 or later
- Docker (optional, for container builds)
- GoReleaser (optional, for releases)

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/dcmcand/camp-quansight-2025-go-bof.git
cd camp-quansight-2025-go-bof
```

### 2. Download Dependencies

```bash
go mod download
```

### 3. Run the Application

```bash
go run main.go
```

The server will start on port 8080 (configurable via `PORT` environment variable).

### 4. Test the API

In another terminal:

```bash
# Test the even number endpoint
curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{"number": 42}'

# Expected: {"is_even":true}

curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{"number": 17}'

# Expected: {"is_even":false}

# Test the health endpoint
curl http://localhost:8080/health

# Expected: {"status":"ok"}
```

## Go Tooling Commands

### Code Quality and Testing

```bash
# Format your code (automatically fixes formatting issues)
go fmt ./...

# Run the vet tool (catches common mistakes)
go vet ./...

# Run tests (add -v for verbose output)
go test ./...
go test -v ./...

# Run tests with coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out  # View coverage in browser
```

### Dependency Management

```bash
# Download dependencies
go mod download

# Add missing and remove unused modules
go mod tidy

# Verify dependencies
go mod verify

# View dependency graph
go mod graph

# Upgrade dependencies
go get -u ./...
```

### Building

```bash
# Build the binary (output: camp-quansight-2025-go-bof or .exe on Windows)
go build

# Build with custom output name
go build -o myapp

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o app-linux-amd64
GOOS=windows GOARCH=amd64 go build -o app-windows-amd64.exe
GOOS=darwin GOARCH=arm64 go build -o app-darwin-arm64

# Build with optimizations (smaller binary)
go build -ldflags="-s -w" -o app
```

## Docker

### Build the Docker Image

```bash
docker build -t go-bof-app .
```

The multistage Dockerfile:

1. Uses `golang:1.23-alpine` to build the binary
2. Copies the binary to a `scratch` container (minimal, ~10MB final image)

### Run with Docker Compose (Recommended)

The easiest way to run the application:

```bash
# Build and run
docker-compose up

# Run in background
docker-compose up -d

# View logs
docker-compose logs -f

# Stop
docker-compose down
```

**Note:** The `EXPOSE 8080` directive in the Dockerfile is just documentation. Port publishing must be configured at runtime using the `-p` flag or in `docker-compose.yml`.

### Run the Docker Container Directly

```bash
# Run on port 8080 (the -p flag is required to publish the port)
docker run -p 8080:8080 go-bof-app

# Run on custom port
docker run -p 3000:3000 -e PORT=3000 go-bof-app

# Run in background
docker run -d -p 8080:8080 --name go-bof go-bof-app

# View logs
docker logs go-bof

# Stop container
docker stop go-bof
docker rm go-bof
```

## GoReleaser

GoReleaser automates building binaries for multiple platforms.

### Install GoReleaser

```bash
# macOS
brew install goreleaser

# Linux
go install github.com/goreleaser/goreleaser@latest

# Or download from https://github.com/goreleaser/goreleaser/releases
```

### Build Snapshot (without git tag)

```bash
goreleaser release --snapshot --clean
```

This creates binaries in the `dist/` folder for:

- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

### Create a Release

```bash
# Create and push a tag
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0

# Build release locally
goreleaser release --clean
```

When you push a tag to GitHub, the release workflow automatically runs and creates a GitHub release with all binaries.

## GitHub Actions Workflows

### CI Workflow (`.github/workflows/ci.yaml`)

Runs on every push and pull request:

- Runs `go vet ./...`
- Runs `go test ./...`
- On pushes to `main`: Builds snapshot release with GoReleaser

### Release Workflow (`.github/workflows/release.yaml`)

Runs when you push a version tag (e.g., `v1.0.0`):

- Builds binaries for all platforms
- Creates a GitHub release
- Uploads all binaries and checksums

## Exercises for Participants

Try these hands-on exercises:

### 1. Add a New Endpoint

Add a POST `/fibonacci` endpoint that calculates the nth Fibonacci number.

### 2. Add Tests

Create `main_test.go` and add unit tests:

```bash
go test -v
```

### 3. Add Input Validation

Modify the `/` endpoint to return an error for negative numbers.

### 4. Build for Your Platform

Build a binary for your specific OS and architecture:

```bash
go build -o myapp
./myapp
```

### 5. Try GoReleaser

Build binaries for all platforms:

```bash
goreleaser release --snapshot --clean
ls dist/
```

### 6. Modify the Dockerfile

Try changing the base image from `scratch` to `alpine` and compare sizes:

```bash
docker images go-bof-app
```

### 7. Explore Go Modules

Add a new dependency and see how `go.mod` changes:

```bash
go get github.com/gorilla/mux
go mod tidy
```

## Project Structure

```
.
├── main.go                      # Main application entrypoint
├── go.mod                       # Go module definition
├── pkg/
│   ├── even/                    # Even number package
│   │   └── even.go              # Package implementation
│   └── server/                  # HTTP server package
│       └── server.go            # Server implementation
├── Dockerfile                   # Multistage Docker build
├── docker-compose.yml           # Docker Compose configuration
├── .dockerignore                # Docker build exclusions
├── .goreleaser.yaml             # GoReleaser configuration
├── .github/
│   └── workflows/
│       ├── ci.yaml              # CI/CD workflow
│       └── release.yaml         # Release workflow
└── README.md                    # This file
```

## Useful Resources

- [A Tour of Go](https://go.dev/tour/) - Interactive introduction to Go
- [Go Official Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)
- [GoReleaser Documentation](https://goreleaser.com/)
- [Docker Go Best Practices](https://docs.docker.com/language/golang/)

## License

MIT License - see LICENSE file for details
