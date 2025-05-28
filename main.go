package main

import (
	"context"
	"fmt"
	"server/bootstrap"
	"server/env"
	"server/workers/agents/analyst"
	"server/workers/agents/capture"
)

func juniorAnalyst() {
	news := &capture.JuniorAnalyst{}

	links, err := news.ResearchArticle(context.Background(), []string{})

	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}

func seniorAnalyst() {
	news, err := analyst.ValidateArticle(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(news)
}

func main() {
	_ = bootstrap.Boot(".env")

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	// juniorAnalyst()
	seniorAnalyst()
}
