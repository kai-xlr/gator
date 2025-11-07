# Gator

A CLI application for managing users with PostgreSQL persistence. Built with Go, using sqlc for type-safe database queries and goose for migrations.

## Prerequisites

- Go 1.25+
- PostgreSQL
- [sqlc](https://sqlc.dev/) (for code generation)
- [goose](https://github.com/pressly/goose) (for migrations)

## Setup

1. Create a PostgreSQL database
2. Set up your configuration file at `~/.gatorconfig.json`:
   ```json
   {
     "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
     "current_user_name": ""
   }
   ```
3. Run database migrations:
   ```bash
   goose -dir sql/schema postgres "$GATOR_DB_URL" up
   ```
4. Build the application:
   ```bash
   go build -o gator
   ```

## Usage

```bash
# Register a new user and set as current
./gator register <name>

# Switch to an existing user
./gator login <name>

# Delete all users from database
./gator reset
```

## Development

### Adding New Database Operations

1. Write SQL queries in `sql/queries/*.sql` with sqlc annotations
2. Generate Go code:
   ```bash
   sqlc generate
   ```
3. Use the generated methods via `state.db.*` in your handlers

### Adding New Commands

1. Create a handler function in a `handler_*.go` file:
   ```go
   func handlerYourCommand(s *state, cmd command) error {
       // Implementation
   }
   ```
2. Register it in `main()`:
   ```go
   cmds.register("yourcommand", handlerYourCommand)
   ```

## Architecture

- **Command Registry**: Commands are registered and dispatched via a registry pattern
- **State Management**: A shared `state` struct provides access to database and config
- **Type-Safe SQL**: sqlc generates Go code from SQL queries in `sql/queries/`
- **Migrations**: goose manages schema changes in `sql/schema/`
