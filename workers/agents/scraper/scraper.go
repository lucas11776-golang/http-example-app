package scraper

import (
	"context"
	"fmt"
	"server/app/services/news"
	"server/env"
	"server/models"
	"strings"

	"google.golang.org/genai"
)

// General,Business,Politics,Science,Health,Entertainment,Sport,Technology

type Scraper interface {
	All() ([]*models.Article, error)
	Category(category news.Category) ([]*models.Article, error)
}

type WebSearch struct{}

// TODO: Improve the agent promp...
func (ctx *WebSearch) All(context context.Context) ([]*models.ArticleResearch, error) {

	// TODO: Look all categories...
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := []string{
		"You are a analyst you have to analyze the must relevant and up to data news about this category: %s I need more then 5 article,",
		"You need to analyze all articles and responsed with a JSON object containing the following interface and place the response in <result></result>",
		`interface Article {
			title: string;        // Article tile.
			category: string;     // Article category pick on based on article - (General,Business,Politics,Science,Health,Entertainment,Sport,Technology)
			url: string;          // Article url.
			description: string;  // Short description of article.
			image: string;        // Article URL image.
			publisher: string;    // Article publisher.
			published_at: string; // Article date.
			content: string;      // Article content (must be text).
		}? please do not include` + "```json ``` in results.",
	}

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				{Text: fmt.Sprintf(strings.Join(ask, "\r\n"), news.Politics)},
			},
			Role: genai.RoleUser,
		},
	}

	response, err := client.Models.GenerateContent(context, env.Env("AI_MODEL"), content, &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			IncludeThoughts: true,
		},
	})

	if err != nil {
		return nil, err
	}

	fmt.Println(response.Text())

	return nil, nil
}

func (ctx *WebSearch) Category(category news.Category) ([]*models.ArticleResearch, error) {
	return nil, nil
}
