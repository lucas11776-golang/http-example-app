package main

import (
	"context"
	"fmt"
	"server/bootstrap"
	"server/env"
	"server/workers/agents/scraper"
)

func main() {
	_ = bootstrap.Boot(".env")

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	// article, err := analyst.ResearchArticle(
	// 	context.Background(),
	// 	"https://www.news24.com/sport/soccer/psl/im-not-a-small-coach-nabi-defiant-in-his-process-as-chiefs-end-season-on-sombre-note-20250526-1042",
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("ARTICLES -> ", article)

	news := &scraper.WebSearch{}

	links, err := news.All(context.Background())

	if err != nil {
		panic(err)
	}

	fmt.Println(links)

	// server.Listen()
}
