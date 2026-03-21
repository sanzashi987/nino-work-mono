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

Located in
- `apps/`:
  - `user/` - SSO/authentication service (sso.nino.work)
  - `canvix/` - Canvas/design application (canvix.nino.work)
  - `storage/` - File storage service (storage.nino.work)
  - `chat/` - Chat service (chat.nino.work)
  - `config-center/` - Configuration management
- `pkg/` - Shared Go packages
- `proto/` - Protobuf definitions

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

### Database

Uses SQLite with GORM. Database files are created at repository root with names defined in `config.ini` (e.g., `nino-mono.db`, `sso.nino.work.db`).

### Testing

Run Go tests:
```bash
go test ./apps/user/...
go test ./pkg/...
```

### Debugging

Install Delve debugger:
```bash
go env -w GOARCH=amd64
go install github.com/go-delve/delve/cmd/dlv@latest
```

### Infrastructure Requirements

**Required Services:**
- **etcd**: Service discovery for microservices mode
  - Download from https://github.com/etcd-io/etcd/releases/
  - Start with: `./etcd` (default: localhost:2379)
- **Typesense** (optional): Vector database for search features

### Key Patterns

#### Service Bootstrap

Each service follows the same pattern in `bootstrap.go`:
1. Parse configuration with `bootstrap.ParseConfig(serviceName)`
2. Connect to database with `dao.ConnectDB()`
3. Create HTTP router with `http.NewRouter()`
4. Run on configured WebPort

Services can optionally initialize RPC clients for inter-service communication.

#### Microservices Communication

Services communicate via gRPC when running in microservice mode. Proto definitions in `proto/` define service interfaces. Clients are created using proto-generated code and go-micro client.

#### Backend-Frontend Communication

**Development:**
- Frontend dev server proxies `/backend/*` to backend services
- Each frontend app's `infra.config.js` defines proxy rules
- Example: main app → port 8081 (user service), canvix app → port 8111 (canvix service)

**Production:**
- Static files served by backend service
- Backend API on same origin (no CORS needed)
- Public URL defined in frontend package.json `homepage` field

## Frontend (React/TypeScript)

### Architecture

Frontend is a pnpm monorepo with workspace structure. Uses custom rsbuild-based infrastructure defined in `@nino-work/infra` for building and serving applications.

**Key Technologies:**
- **Build Tool**: Rsbuild (migrated from webpack)
- **Module Federation**: For micro-frontend architecture
- **Single-spa**: Micro-frontend framework
- **Plugin System**: React, Sass, SVGR, Module Federation

**Infrastructure Documentation:**
See `@frontend/@nino-work/infra/README.md` for complete documentation on:
- Configuration (`infra.config.js`)
- Running modes (standalone/micro-host/micro-app)
- Entry file patterns (index.tsx vs index.micro.tsx)
- Environment variables
- Module Federation setup
- Build and deployment

### Workspace Structure

Located in `frontend/`:
- `apps/` - Application entry points
  - `main/` - Main portal (nino.work, port 3000)
  - `canvix/` - Canvas application (canvix.nino.work, port 3003)
  - `storage/` - Storage management UI
  - `root/` - Root config for single-spa micro-frontend
- `@nino-work/` - Shared workspace packages
  - `infra/` - Custom rsbuild build tools and configuration
  - `requester/` - API client utilities
  - `shared/` - Shared utilities and types
  - `form/` - Form components and utilities
  - `ui-components/` - Shared UI component library
  - `mf/` - Micro-frontend utilities (import maps, module federation)
  - `assets/` - Shared assets
  - `converter/` - Data conversion utilities
- `@canvix/` - Canvas-specific packages
  - `sdk/` - Canvas SDK
  - `txt/` - Text rendering utilities

### Running Frontend Apps

Each app supports **two modes**:

```bash
# Install dependencies (from frontend/ directory)
pnpm install

# Standalone mode - runs as independent app
cd apps/main && pnpm start        # Port 3000
cd apps/canvix && pnpm start      # Port 3003

# Micro-frontend mode - runs with single-spa
cd apps/main && pnpm start-mf     # Mode: micro-host
cd apps/canvix && pnpm start-mf   # Mode: micro-app
```

**Mode Types:**
- `standalone` (default): Independent app with HTML entry
- `micro-host`: Main container app, loads micro-apps via Module Federation
- `micro-app`: Child app exposed via Module Federation, no HTML generated

### Building

```bash
cd apps/main && pnpm build     # Outputs to apps/main/build/
cd apps/canvix && pnpm build   # Outputs to apps/canvix/build/
```

### Code Style

ESLint and Prettier configurations are in `frontend/.eslintrc` and `frontend/.prettierrc`. Linting is enforced across the workspace.

### Key Patterns

#### App Bootstrap

Each frontend app follows the same structure:

**1. Configuration (infra.config.js):**
- Define port, mode, and build-time constants
- Configure proxy for backend API
- Setup Tailwind CSS
- See `@frontend/@nino-work/infra/README.md` for detailed configuration options

**2. Entry Files:**
- `index.tsx` - Standalone mode entry
- `index.micro.tsx` - Micro-frontend mode entry
- See `@frontend/@nino-work/infra/README.md` for entry file patterns

**3. App Component:**
- Receives `importMapPromise` prop (micro-frontend config)
- Uses `usePromise` hook to load import map
- Wraps with ThemeProvider and RouterProvider

#### App Structure

Each frontend app uses:
- React 18 with TypeScript
- React Router for navigation
- React Hook Form for forms
- Material-UI for components
- Emotion for styling
- Single-spa for micro-frontend support (optional)

#### Micro-frontend Architecture

**Host App (main):**
- Mode: `micro-host`
- Loads import map from `@nino-work/mf`
- Registers child applications via `single-spa.registerApplication()`
- Provides container DOM element (`#nino-sub-app`)

**Child Apps (canvix, storage, etc.):**
- Mode: `micro-app`
- Exposed via Module Federation plugin
- Shared dependencies: react, react-dom, single-spa
- No HTML generated (pure JS bundle)

**Architecture Details:**
See `@frontend/@nino-work/infra/README.md` for Module Federation configuration, import map flow, and implementation examples.

#### Dependency Management

**Workspace Dependencies:**
- Use `workspace:*` protocol for internal packages
- Shared dependencies: react, react-dom, material-ui, emotion
- Build dependencies: @rsbuild/core, @rsbuild/plugin-*

**Peer Dependencies:**
- Tailwind CSS in `@nino-work/infra` (apps install their own version)
- Ensures flexibility and version control

#### Code Organization

**Frontend Structure:**
- `frontend/apps/[app]/` - Application entry and pages
- `frontend/@nino-work/[package]/` - Shared packages
- `frontend/@canvix/[package]/` - Canvas-specific packages
- `infra.config.js` - Per-app configuration
- `public/index.html` - Custom HTML template (optional)
