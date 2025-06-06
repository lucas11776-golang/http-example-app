package main

import (
	"context"
	"fmt"
	"os"
	"server/bootstrap"
	"server/jobs"
	"server/workers/agents/analyst"
	"server/workers/agents/capture"
	"server/workers/agents/designer"
	"server/workers/agents/studio/artist"
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

// Comment
func voiceArtist() {
	graphicDesigner := &artist.VoiceArtist{}

	image, err := graphicDesigner.RecordArticle(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(image)
}

func main() {

	_ = bootstrap.Boot(".env")

	// fmt.Printf("\r\nRunning Server %s:%d\r\n", env.Env("HOST"), env.EnvInt("PORT"))

	if len(os.Args[1]) > 1 && os.Args[1] == "jobs" {
		jobs.Setup()

		return
	}

	fmt.Println()

	fmt.Println("Working Hard")

	// juniorAnalyst()
	// seniorAnalyst()
	// seniorAnalystArticleDescription()
	// studioProducer()
	// voiceArtist()
}
