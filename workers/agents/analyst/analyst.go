package analyst

import (
	"context"
	"fmt"
	"server/env"
	"server/models"
	"strings"

	"google.golang.org/genai"
)

// Comment
// Ask is week promp just for test workflow of Scraper and Analyst.
func ResearchArticle(context context.Context, url string) (*models.Article, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := []string{
		"Read the the following URL as an analyst: https://www.sowetanlive.co.za/news/2025-05-27-fires-leave-12-dead-in-gauteng-in-just-a-week/",
		"Response with a JSON object cotaining the following this object below and place it in <result></result>",
		`interface Article {
			title: string;        // Article tile.
			url: string;          // Article url.
			description: string;  // Short description of article.
			image: string;        // Article URL image.
			publisher: string;    // Article publisher.
			published_at: string; // Article date.
			content: string;      // Article content (must be html div class name must be article-single).
			locations: []string   // Address affected by article or where the article occurred.
		}?`,
	}

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				{Text: strings.Join(ask, "\r\n")},
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
