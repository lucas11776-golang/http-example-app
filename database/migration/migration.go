package main

import (
	"fmt"
	"server/env"
	"server/jobs/workspace/paperwork/analyst"

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
	return append(presention(), paperwork()...)

}

// Comment
func presention() orm.Models {
	return orm.Models{
		// models.Article{},
		// models.NewsQuery{},
		// models.ArticleCaputure{},
	}
}

// Comment
func paperwork() orm.Models {
	return orm.Models{
		analyst.ArticleCapture{},
		analyst.ArticleVerified{},
		analyst.ArticleVerifiedSource{},
	}
}
