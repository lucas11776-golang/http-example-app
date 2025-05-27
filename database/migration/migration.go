package main

import (
	"fmt"
	"server/env"
	"server/models"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
)

func main() {
	env.Load(".env")

	// SQLite Connection
	sqlite := sqlite.Connect(env.Env("DATABASE"))

	// run migrations
	err := sqlite.Migration().Migrate(orm.Models{
		models.Article{},
		models.NewsQuery{},
		models.ArticleResearchLink{},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Successfully...")
}
