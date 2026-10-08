# ShopFlow

ShopFlow is an e-commerce backend written in Go and backed by PostgreSQL.

## Tech Stack

- **Language:** Go 1.26
- **Database:** PostgreSQL 18 (via `pgx` / `pgxpool`)
- **Config:** `godotenv`
- **Local infrastructure:** Docker Compose

## Project Structure

```
shop_flow/
├── cmd/
│   └── server/          # application entrypoint
│       └── main.go
├── config/              # environment/config loading
├── database/            # database connection (PostgreSQL)
├── internal/            # domain packages
│   ├── authentication/  # users, login, tokens
│   ├── cart/            # shopping cart
│   ├── categories/      # product categories
│   ├── orders/          # orders and checkout
│   ├── products/        # product catalog
│   └── users/           # user accounts
├── middleware/          # HTTP middleware
├── migrations/          # SQL schema migrations
├── docker-compose.yml   # local PostgreSQL
├── go.mod
└── .env
```

## Prerequisites

- Go 1.26+
- Docker (for the local PostgreSQL instance)

## Getting Started

1. Start PostgreSQL:

   ```sh
   docker compose up -d
   ```

2. Create a `.env` file in the project root:

   ```env
   DATABASE_URL=postgresql://shopflow:shopflow123@localhost:5432/shopflow
   ```

3. Run the server:

   ```sh
   go run ./cmd/server
   ```

   You should see `Connected to database`.

## Configuration

| Variable       | Description                     | Example                                                      |
| -------------- | ------------------------------- | ------------------------------------------------------------ |
| `DATABASE_URL` | PostgreSQL connection string    | `postgresql://shopflow:shopflow123@localhost:5432/shopflow`  |

The default PostgreSQL credentials are defined in `docker-compose.yml`:

- Database: `shopflow`
- User: `shopflow`
- Password: `shopflow123`
- Port: `5432`

## Roadmap

- [ ] HTTP server and routing
- [ ] Database migrations for users, products, categories, carts, orders
- [ ] Authentication (registration, login, JWT)
- [ ] Product and category management
- [ ] Shopping cart
- [ ] Orders and checkout
