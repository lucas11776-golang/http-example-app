package main

import (
	"context"
	"fmt"
	"server/bootstrap"
	"server/env"
	"server/workers/agents/analyst"
)

func main() {
	_ = bootstrap.Boot(".env")

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	article, err := analyst.ResearchArticle(
		context.Background(),
		"https://iol.co.za/news/south-africa/2025-05-27-tensions-rise-in-zulu-royal-family-as-supreme-court-appeal-looms/",
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("ARTICLES -> ", article)

	// server.Listen()
}
