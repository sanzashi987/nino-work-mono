# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a full-stack monorepo containing Go microservices and React/TypeScript frontend applications. The backend uses gRPC with go-micro framework, while the frontend uses a pnpm workspace with multiple apps and shared packages.

## Backend (Go)

### Architecture

The backend is structured as microservices using the `go-micro` framework with gRPC/Protobuf for inter-service communication and etcd for service discovery. Services can run in two modes:
- **Microservice mode**: Uses etcd for service discovery and RPC for communication
- **Standalone mode**: Runs independently without etcd dependency

### Services

Located in `apps/`:
- `user/` - SSO/authentication service (sso.nino.work)
- `canvix/` - Canvas/design application (canvix.nino.work)
- `storage/` - File storage service (storage.nino.work)
- `chat/` - Chat service (chat.nino.work)
- `config-center/` - Configuration management

### Running Services

Each service can be started using `fresh` for hot-reload development:
```bash
fresh -c ./apps/user/runner.conf
fresh -c ./apps/canvix/runner.conf
```

Or run directly:
```bash
go run ./apps/user/bootstrap.go
```

Services read configuration from `config.ini` at the repository root. Each service has its own section defining Host, Port, WebPort, and DbName.

### Protocol Buffers

Proto files are in `proto/`. To regenerate Go code from proto files:
```bash
protoc --micro_out=. --go_out=. -I=./proto proto/user.proto
protoc --micro_out=. --go_out=. -I=./proto proto/storage.proto
```

Required tools (install once):
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Testing

Run Go tests:
```bash
go test ./apps/user/...
go test ./pkg/...
```

### Database

Uses SQLite with GORM. Database files are created at repository root with names defined in `config.ini` (e.g., `nino-mono.db`, `sso.nino.work.db`).

## Frontend (React/TypeScript)

### Architecture

Frontend is a pnpm monorepo with workspace structure. Uses custom webpack infrastructure defined in `@nino-work/infra` for building and serving applications.

### Workspace Structure

Located in `frontend/`:
- `apps/` - Application entry points
  - `main/` - Main portal (nino.work)
  - `canvix/` - Canvas application (canvix.nino.work)
  - `storage/` - Storage management UI
  - `root/` - Root config for single-spa micro-frontend
- `@nino-work/` - Shared workspace packages
  - `infra/` - Custom webpack build tools
  - `requester/` - API client utilities
  - `shared/` - Shared utilities and types
  - `form/` - Form components and utilities
  - `ui-components/` - Shared UI component library
  - `mf/` - Micro-frontend utilities
  - `assets/` - Shared assets
  - `converter/` - Data conversion utilities
- `@canvix/` - Canvas-specific packages
  - `sdk/` - Canvas SDK
  - `txt/` - Text rendering utilities

### Running Frontend Apps

From `frontend/` directory:

```bash
# Install dependencies
pnpm install

# Run main app
cd apps/main && pnpm start

# Run canvix app
cd apps/canvix && pnpm start

# Run in micro-frontend mode
cd apps/main && pnpm start-mf
cd apps/canvix && pnpm start-mf
```

### Building

```bash
cd apps/main && pnpm build
cd apps/canvix && pnpm build
```

### Code Style

ESLint and Prettier configurations are in `frontend/.eslintrc` and `frontend/.prettierrc`. Linting is enforced across the workspace.

## Infrastructure Requirements

### Required Services
- **etcd**: Service discovery for microservices mode
  - Download from https://github.com/etcd-io/etcd/releases/
  - Start with: `./etcd` (default: localhost:2379)
- **Typesense** (optional): Vector database for search features

### Debugging Go

Install Delve debugger:
```bash
go env -w GOARCH=amd64
go install github.com/go-delve/delve/cmd/dlv@latest
```

## Key Patterns

### Backend Service Bootstrap

Each service follows the same pattern in `bootstrap.go`:
1. Parse configuration with `bootstrap.ParseConfig(serviceName)`
2. Connect to database with `dao.ConnectDB()`
3. Create HTTP router with `http.NewRouter()`
4. Run on configured WebPort

Services can optionally initialize RPC clients for inter-service communication.

### Frontend App Structure

Each frontend app uses:
- React 18 with TypeScript
- React Router for navigation
- React Hook Form for forms
- Material-UI for components
- Emotion for styling
- Single-spa for micro-frontend support (optional)

### Microservices Communication

Services communicate via gRPC when running in microservice mode. Proto definitions in `proto/` define service interfaces. Clients are created using proto-generated code and go-micro client.
