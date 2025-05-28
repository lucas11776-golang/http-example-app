package capture

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"server/env"
	"server/models"

	"google.golang.org/genai"
)

var (
	RESULT_REGEX = regexp.MustCompile(`<result>(.*?)</result>`)
)

type JuniorAnalyst struct {
}

// Comment
func (ctx *JuniorAnalyst) ResearchArticle(context context.Context, intrests []string) ([]models.ArticleCaputure, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := `
	You are an analyst working for a company that analyses the web for news based on client descriptions.
	Your job is to find those news articles on the web, and they will be submitted to your senior analyst for review and approval.
	Remember the client depends on those news articles for their daily operations.
	Below are are bullet points of what the client wants:

	- News must the the lastest current date is 28 May 2025.
	- News must be category of Finance.
	- News must be in South Africa.
	- Use news site from South Africa.
	- Get atleast 10 article but if the are not that intresting exclude them.

	After you are done analyzing the news article data please format the articles in JSON object in array containing the following interface and
	place the data inside <result><result> also do not include ` + "```json ``` in results." + `

	interface Article {
		title: string;        // Article tile.
		category: string;     // Article category pick on based on article - (General,Business,Politics,Science,Health,Entertainment,Sport,Technology,Finance)
		website: string;      // Article url/source only website hostname.
		description: string;  // Short description of article.
		image: string;        // Article image please do not make up one if you can not find it leave it empty.
		publisher: string;    // Article publisher data format YYYY-DD-MM
		published_at: string; // Article date.
		content: string;      // Article content (must be text).
	}`

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				{Text: ask},
			},
			Role: genai.RoleUser,
		},
	}

	response, err := client.Models.GenerateContent(context, env.Env("AI_MODEL"), content, &genai.GenerateContentConfig{
		ThinkingConfig: &genai.ThinkingConfig{
			IncludeThoughts: true,
		},
		Tools: []*genai.Tool{
			{
				GoogleSearch: &genai.GoogleSearch{},
			},
			{
				URLContext: &genai.URLContext{},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	matches := RESULT_REGEX.FindStringSubmatch(response.Text())

	file, _ := os.Create("cache.txt")

	file.Write([]byte(response.Text()))

	if len(matches) == 0 {
		return []models.ArticleCaputure{}, nil
	}

	var articles []models.ArticleCaputure

	err = json.Unmarshal([]byte(matches[0]), &articles)

	if err != nil {
		return nil, err
	}

	return articles, nil
}
