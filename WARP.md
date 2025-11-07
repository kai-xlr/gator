# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project Overview

Gator is a CLI application written in Go that manages users with PostgreSQL persistence. It uses sqlc for type-safe database queries and goose for migrations.

## Common Commands

### Building and Running
```bash
# Build the application
go build -o gator

# Run the CLI (after building)
./gator <command> [args...]
```

### Database Operations
```bash
# Run migrations (requires GATOR_DB_URL environment variable)
goose -dir sql/schema postgres "$GATOR_DB_URL" up

# Rollback migrations
goose -dir sql/schema postgres "$GATOR_DB_URL" down

# Check migration status
goose -dir sql/schema postgres "$GATOR_DB_URL" status

# Generate Go code from SQL queries (run after modifying sql/queries/*.sql)
sqlc generate
```

### Available CLI Commands
```bash
./gator register <name>    # Create a new user and set as current
./gator login <name>        # Switch to an existing user
./gator reset              # Delete all users from database
```

## Architecture

### Configuration System
- Config file: `~/.gatorconfig.json` (stored in user's home directory)
- Contains: `db_url` (PostgreSQL connection string) and `current_user_name`
- Package: `internal/config/config.go`
- The config is read on startup and updated when users register or login

### Command Registration Pattern
The application uses a command registry pattern:
1. Commands are registered in `main()` using `cmds.register(name, handlerFunc)`
2. Each handler is a function with signature: `func(*state, command) error`
3. Handlers are in `handler_*.go` files (e.g., `handler_user.go`, `handler_reset.go`)
4. The `commands` struct dispatches to the appropriate handler based on CLI args

### Database Layer
- **sqlc** generates type-safe Go code from SQL queries
- SQL queries: `sql/queries/*.sql` (annotated with sqlc directives)
- Database migrations: `sql/schema/*.sql` (goose format with `-- +goose Up/Down`)
- Generated code: `internal/database/` (never edit these files directly)
- To add new database operations:
  1. Write SQL in `sql/queries/`
  2. Run `sqlc generate`
  3. Use generated methods via `state.db.*`

### State Management
The `state` struct is passed to all command handlers and contains:
- `db`: sqlc-generated `*database.Queries` for database operations
- `cfg`: `*config.Config` for accessing configuration

### Adding New Commands
1. Create a handler function: `func handler<Name>(s *state, cmd command) error`
2. Put it in an appropriate `handler_*.go` file
3. Register it in `main()`: `cmds.register("commandname", handler<Name>)`
4. If database operations are needed, add SQL to `sql/queries/` and run `sqlc generate`

## Dependencies
- `github.com/lib/pq`: PostgreSQL driver
- `github.com/google/uuid`: UUID generation
- `sqlc`: Type-safe SQL code generation (must be installed)
- `goose`: Database migrations (must be installed)

## Environment
Requires `GATOR_DB_URL` environment variable for database connection (read from config file at `~/.gatorconfig.json`).
