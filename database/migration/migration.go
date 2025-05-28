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

	if err := sqlite.Connect(env.Env("DATABASE")).Migration().Migrate(Models()); err != nil {
		panic(err)
	}

	fmt.Println("Migration Successfully...")
}

// Comment
func Models() orm.Models {
	return orm.Models{
		models.Article{},
		models.NewsQuery{},
		models.ArticleCaputure{},
	}
}
