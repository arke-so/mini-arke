# Mini-Arke API

## Requirements

- Go 1.24+
- Docker & Docker Compose

## Quick Start

```bash
# Start the server
make start
```

API runs on `http://localhost:8080`

Test with Postman or:
```bash
curl http://localhost:8080/products
```

## Running Tests

```bash
make test
```

## Available Commands

- `make start` - Start server
- `make test` - Run tests
- `make down` - Stop databases