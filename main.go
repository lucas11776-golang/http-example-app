package main

import (
	"context"
	"fmt"
	"server/bootstrap"
	"server/workers/agents/analyst"
	"server/workers/agents/capture"
	"server/workers/agents/designer"
	"server/workers/agents/studio/producer"
)

// Comment
func juniorAnalyst() {
	juniorAnalyst := &capture.JuniorAnalyst{}

	links, err := juniorAnalyst.ResearchArticle(context.Background(), []string{})

	if err != nil {
		panic(err)
	}

	fmt.Println(links)
}

// Comment
func seniorAnalyst() {
	seniorAnalyst := &analyst.SeniorAnalyst{}

	article, err := seniorAnalyst.ValidateArticle(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(article)
}

// Comment
func seniorAnalystArticleDescription() {
	seniorAnalyst := &analyst.SeniorAnalyst{}

	description, err := seniorAnalyst.DescribeArticle(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(description)
}

// Comment
func seniorGraphicDesign() {
	graphicDesigner := &designer.GraphicDesigner{}

	image, err := graphicDesigner.DesignArticleImage(context.Background(), []string{})

	if err != nil {
		panic(err)
	}

	fmt.Println(image)
}

// Comment
func studioProducer() {
	graphicDesigner := &producer.Producer{}

	image, err := graphicDesigner.PrepareNewsScript(context.Background(), []string{}, nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(image)
}

func main() {
	_ = bootstrap.Boot(".env")

	// fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	fmt.Println("Working Hard")

	// juniorAnalyst()
	// seniorAnalyst()
	// seniorAnalystArticleDescription()
	studioProducer()
}
