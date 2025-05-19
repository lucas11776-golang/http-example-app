# HTTP Example APP


## Getting Started


### Prerequisites

HTTP requests [Go](https://go.dev) version [1.23](https://go.dev/doc/devel/release#go1.22.0) or above

**Setup**


First you have to register to [NewsApi](https://newsapi.org/) to get `NEWS_API_KEY` after you have the key
paste it in the `.env` at `NEWS_API_KEY="PasteKeyHere"` then.

Run `go run database/migration/migration.go` command to run database migrations.

```sh
go run database/migration/migration.go
```

After you can run `go run main.go` to start the application

```sh
go run main.go
```

by default address is `127.0.0.1:8080` you can change then enviroment variable is `.env` to suit your needs.
