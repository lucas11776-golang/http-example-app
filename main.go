package main

import (
	"context"
	"fmt"
	"server/bootstrap"
	"server/env"
	"server/workers/agents/analyst"
	"server/workers/agents/capture"
	"server/workers/agents/designer"
)

func juniorAnalyst() {
	worker := &capture.JuniorAnalyst{}

	links, err := worker.ResearchArticle(context.Background(), []string{})

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

func seniorAnalystArticleDescription() {
	news, err := analyst.DescribeArticle(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(news)
}

func seniorGraphicDesign() {
	worker := &designer.GraphicDesigner{}

	links, err := worker.DesignArticleImage(context.Background(), []string{})

	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}

func main() {
	_ = bootstrap.Boot(".env")

	fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	// juniorAnalyst()
	// seniorAnalyst()
	// seniorAnalystArticleDescription()
	seniorGraphicDesign()
}
