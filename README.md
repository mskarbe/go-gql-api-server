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

## GraphQL

The GraphQL server was created and is developed with use of [gqlgen](https://github.com/99designs/gqlgen) package.
