# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is **sqlc-gen-go**, a Go code generator plugin for [sqlc](https://sqlc.dev/) that generates type-safe Go code from SQL queries. It was extracted from the main sqlc repository to allow customization. The project builds both native Go binaries and WASM plugins for sqlc v2.

## Key Commands

### Build Commands
```bash
# Build both Go binary and WASM plugin
make all

# Build only the Go binary
make build

# Build only the WASM plugin  
make bin/sqlc-gen-go.wasm

# Run tests (requires WASM build first)
make test
```

### Development Commands
```bash
# Run Go tests directly
go test ./...

# Run specific test
go test -run TestName ./internal/...

# Check formatting
go fmt ./...

# Run linting (if golangci-lint is installed)
golangci-lint run

# Update dependencies
go mod tidy
```

## Architecture & Structure

### Core Flow
1. **Entry Point**: `plugin/main.go` calls `codegen.Run(golang.Generate)`
2. **Generation Engine**: `internal/gen.go::Generate()` orchestrates the entire process
3. **Template System**: Uses Go's `text/template` with driver-specific templates in `internal/templates/`

### Key Components

- **internal/gen.go**: Main code generation logic, orchestrates parsing, validation, and output
- **internal/opts/**: Configuration parsing and validation, driver detection
- **internal/templates/**: Driver-specific code generation templates
  - `template.tmpl`: Shared template definitions
  - `stdlib/`, `pgx/`, `go-sql-driver-mysql/`, `gofr/`: Driver-specific templates
- **internal/driver.go**: SQL driver detection and configuration
- **internal/imports.go**: Import statement generation and management

### Template System
The project uses a modular template system where each SQL driver has its own template set:
- Templates are embedded using `embed.FS`
- Driver selection happens in `parseOptions()` based on `sql_package` config
- Templates generate structs, interfaces, and query methods

### Supported SQL Drivers
- PostgreSQL: `pgx/v4`, `pgx/v5`, `database/sql` with `lib/pq`
- MySQL: `database/sql` with `go-sql-driver/mysql`
- SQLite: `database/sql` with various SQLite drivers
- GoFr: Custom framework support (recently added)

## Current Branch Context

You're on branch `nikolubbe/add-gofr-sql-support` which is adding GoFr framework support. Key changes include:
- New GoFr templates in `internal/templates/gofr/`
- Driver enum updates in `internal/opts/enum.go`
- Import handling for GoFr in `internal/imports.go`

## Testing Approach

- End-to-end tests in `internal/endtoend/` using testdata fixtures
- Limited unit test coverage (only 2 test files currently)
- Tests focus on query result handling and options parsing
- Run with `make test` after building WASM

## Important Conventions

### When Adding New SQL Drivers
1. Add driver enum in `internal/opts/enum.go`
2. Create template directory under `internal/templates/`
3. Update driver detection in `internal/driver.go`
4. Add import handling in `internal/imports.go`
5. Update template embedding in `internal/gen.go`

### Code Generation Patterns
- All generated code includes a header comment indicating it's generated
- Structs are generated for each table/query result
- Interface `Querier` contains all query methods
- `New()` function creates a new Queries struct with the database connection

### Template Variables
Key variables available in templates:
- `$.Package`: Package name for generated code
- `$.Structs`: Generated struct definitions
- `$.Queries`: Query definitions with SQL and methods
- `$.SourceName`: Source file that triggered generation
- `$.EmitInterface`: Whether to generate the Querier interface