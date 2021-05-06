# go-gql-api-server

_GraphQL API for PostgreSQL database written in Go._

<!-- TOC -->

- [go-gql-api-server](#go-gql-api-server)
  - [Database](#database)
  - [GraphQL](#graphql)

<!-- /TOC -->

## Database

The connection and sql queries are done with [database/sql](https://golang.org/pkg/database/sql/) package. Exclusively parametrized queries are used.

The schema represents a simple book catalogue.

The migrations are performed with [golang-migrate](https://github.com/golang-migrate/migrate). On the startup of the server there is always performed a migration to the newest version present in `./database/migrations` folder.

As the Postgres is used it is important to install it with correct driver:

```
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

```

Then each migration can be created and applied with:

```
migrate create -ext sql -dir db/migrations -seq <migration_name>
migrate -database "postgres://<username>:<pass>@localhost:5432/<dbname>?sslmode=disable&search_path=public" -path db/migrations up
```

## GraphQL

The GraphQL server was created and is developed with use of [gqlgen](https://github.com/99designs/gqlgen) package.
