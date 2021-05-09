# go-gql-api-server

_GraphQL API for PostgreSQL database written in Go._

<!-- TOC -->

- [go-gql-api-server](#go-gql-api-server)
  - [Database](#database)
    - [Schema](#schema)
    - [Migrations](#migrations)
  - [GraphQL](#graphql)
  - [Todos](#todos)

<!-- /TOC -->

## Database

The connection and sql queries are done with [database/sql](https://golang.org/pkg/database/sql/) package. Exclusively parametrized queries are used.

### Schema

The schema represents a simple book catalogue.

![schema.png](/db/schema.png)

### Migrations

The migrations are performed with [golang-migrate](https://github.com/golang-migrate/migrate). On the startup of the server there is always performed a migration to the newest version present in `./database/migrations` folder.

As the Postgres is used it is important to install it with correct driver:

```
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

```

Then each migration can be created and applied with:

```
$ migrate create -ext sql -dir db/migrations -seq <migration_name>
$ migrate -database "postgres://<username>:<pass>@localhost:5432/<dbname>?sslmode=disable&search_path=public" -path db/migrations up
```

## GraphQL

The GraphQL server was created and is developed with use of [gqlgen](https://github.com/99designs/gqlgen) package.

The GrqphQL schema is defined in [schema.graphqls](graph/schema.graphqls).

Inside of [graph/resolver](graph/resolver) there are located `resolver_<resource>.go` files, which include functions to perform actual sql queries.

The [schema.resolvers](graph/resolver/schema.resolvers.go) file contains the Graphql queries and mutations.

After building and running the server:

```
$ go build server.go
$ ./server
```

Gqlgen GraphQL Playground is available at chosen port and the queries can be performed.

The API is reachable at `http://localhost:3000/query`:

```
$ curl \
  -H "Content-Type: application/json" \
  -d '{ "query": "query all_books { books { id title description authors { id full_name } } }"}' \
  http://localhost:3000/query
{"data":{"books":[{"id":"ckoembnr4000141069zel7zck","title":"Philosophy as a Mockery of Truth: Humour in Plato’s Charmides","description":"really cool article","authors":[{"id":"ckoembf9k00004106adegevg5","full_name":"Jan Skrobecki"}]},{"id":"ckoembrm000024106l291x19d","title":"Greek symposion as a space for philosophical discourse: Xenophanes and criticism of the poetic tradition","description":"difficult stuff","authors":[{"id":"ckoembf9k00004106adegevg5","full_name":"Jan Skrobecki"}]}]}}
```
