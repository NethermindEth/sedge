# Test Configuration Guide

This document provides guidance on setting up and running tests for the Sedge project.

## Prerequisites

### Required Dependencies
- Go 1.21 or later
- Docker and Docker Compose (for E2E tests)
- mockgen (for generating mocks)

### Installing Dependencies

1. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

2. **Install mockgen:**
   ```bash
   go install github.com/golang/mock/mockgen@latest
   export PATH=$PATH:~/go/bin
   ```

3. **Generate mocks:**
   ```bash
   go generate ./...
   ```

4. **Install Docker (if not already installed):**
   ```bash
   ./scripts/setup-test-docker.sh
   ```

## Running Tests

### Unit Tests
```bash
# Run all unit tests
go test ./... -v

# Run specific package tests
go test ./cli -v
go test ./internal/pkg/dependencies -v
```

### Network-Dependent Tests
Some tests require network connectivity to external services (IPFS, Ethereum nodes, etc.). These tests will be skipped in CI environments or when network connectivity is unavailable.

```bash
# Run tests with network timeout
go test ./internal/lido/contracts/csfeedistributor -v
go test ./internal/lido/contracts/mevboostrelaylist -v
```

### E2E Tests
E2E tests require Docker to be running:

```bash
# Setup Docker first
./scripts/setup-test-docker.sh

# Run E2E tests
go test ./e2e -v
```

## Test Categories

### 1. Unit Tests
- **Location**: `./cli`, `./internal/pkg/*`
- **Dependencies**: None
- **Network**: No
- **Docker**: No

### 2. Integration Tests
- **Location**: `./internal/lido/contracts/*`
- **Dependencies**: Ethereum client libraries
- **Network**: Yes (external APIs)
- **Docker**: No

### 3. E2E Tests
- **Location**: `./e2e`
- **Dependencies**: Docker, Docker Compose
- **Network**: Yes
- **Docker**: Yes

## Troubleshooting

### Mock Generation Issues
If you see errors about missing mock packages:
```bash
# Regenerate all mocks
go generate ./...
```

### Docker Issues
If Docker tests are failing:
```bash
# Check Docker status
sudo docker ps

# Restart Docker
sudo systemctl restart docker

# Or start manually
sudo dockerd &
```

### Network Timeout Issues
If network-dependent tests are timing out:
- Check your internet connection
- Tests will be skipped automatically if network is unavailable
- You can run tests with increased timeout: `go test -timeout 30s ./...`

### Permission Issues
If you get permission errors with Docker:
```bash
# Add user to docker group
sudo usermod -aG docker $USER

# Reload group membership
newgrp docker
```

## CI/CD Considerations

### GitHub Actions
The CI environment automatically:
- Installs dependencies
- Generates mocks
- Runs unit tests
- Skips network-dependent tests
- Runs E2E tests if Docker is available

### Local Development
For local development:
1. Run `./scripts/setup-test-docker.sh` to set up Docker
2. Run `go generate ./...` to generate mocks
3. Run `go test ./... -v` to run all tests

## Test Environment Variables

The following environment variables affect test behavior:

- `CI=true`: Skips network-dependent tests
- `GITHUB_ACTIONS=true`: Skips network-dependent tests
- `DOCKER_HOST`: Docker daemon socket (default: unix:///var/run/docker.sock)
- `TEST_TIMEOUT`: Test timeout (default: 10s for network tests)

## Adding New Tests

### Unit Tests
1. Create test file: `package_test.go`
2. Use standard Go testing patterns
3. No special setup required

### Integration Tests
1. Add timeout handling
2. Check for network errors
3. Skip tests in CI environment

### E2E Tests
1. Use Docker Compose for setup
2. Clean up resources after tests
3. Handle Docker availability gracefully