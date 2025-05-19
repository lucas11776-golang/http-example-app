package database

import (
	"server/env"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
)

// Comment
func Setup() {
	orm.DB.Add("sqlite", sqlite.Connect(env.Env("DATABASE"))) // SQLite database config...
}
