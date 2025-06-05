package junior

import (
	"context"
	"encoding/json"
	"server/env"
	"server/jobs/office/utils"
	"server/jobs/workspace"
	"server/jobs/workspace/paperwork/analyst"
	"time"

	"github.com/lucas11776-golang/orm"
	"google.golang.org/genai"
)

type JuniorAnalyst struct {
	workspace *workspace.Workspace
}

// Comment
func NewJuniorAnalyst(workspace *workspace.Workspace) *JuniorAnalyst {
	return &JuniorAnalyst{
		workspace: workspace,
	}
}

// Comment
func (ctx *JuniorAnalyst) Work(context context.Context) {

}

// TODO: We need embed to check if we have news article
// Comment
func (ctx *JuniorAnalyst) createArticleCapture(capture map[string]interface{}) (*analyst.ArticleCapture, error) {
	var err error

	capture["published_at"], err = time.Parse(time.DateOnly, capture["published_at"].(string))

	if err != nil {
		return nil, err
	}

	capture["published_at"] = capture["published_at"].(time.Time).Format(time.DateOnly)

	article, err := orm.Model(analyst.ArticleCapture{}).Where("title", "=", capture["title"]).
		AndWhere("published_at", "=", capture["published_at"]).
		First()

	if err != nil || article != nil {
		return article, nil
	}

	article, err = orm.Model(analyst.ArticleCapture{}).Insert(capture)

	if err != nil {
		return nil, err
	}

	return article, nil
}

// Comment
func (ctx *JuniorAnalyst) ResearchArticles(context context.Context, intrest []string) ([]*analyst.ArticleCapture, error) {
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

	- News must the the lastest current date is 05 June 2025.
	- News must be in South Africa.
	- Use news site from South Africa.
	- Get atleast 50 article but if the are not that intresting exclude them.

	After you are done analyzing the news article data please format the articles in JSON object in array containing the following interface and
	place the data inside <result><result> also do not include ` + "```json ``` in results." + `

	interface Article {
		title: string;        // Article tile.
		category: string;     // Article category pick on based on article - (General,Business,Politics,Science,Health,Entertainment,Sport,Technology,Finance)
		website: string;      // Article url/source only website host.
		description: string;  // Short description of article.
		image: string;        // Article image please do not make up one if you can not find it leave it empty.
		publisher: string;    // Article publisher
		published_at: string; // Article published at format YYYY-DD-MM.
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

	result := utils.ResultFromPaperword(string(response.Text()))

	if result == "" {
		return []*analyst.ArticleCapture{}, nil
	}

	var captures []map[string]interface{}

	if err := json.Unmarshal([]byte(result), &captures); err != nil {
		return nil, err
	}

	articles := []*analyst.ArticleCapture{}

	for _, capture := range captures {
		article, err := ctx.createArticleCapture(capture)

		if err != nil {
			continue
		}

		articles = append(articles, article)
	}

	return articles, nil
}

// Comment
func (ctx *JuniorAnalyst) UnverifiedArticles(context context.Context) ([]*analyst.ArticleCapture, error) {
	articles, err := orm.Model(analyst.ArticleCapture{}).
		Where("verified_at", "=", nil).
		AndWhereGroup(func(group orm.WhereGroupBuilder) {
			group.Where("verification_at", "=", nil).
				OrWhere("verification_at", "<", time.Now().Add(time.Minute*-5))
		}).
		Get()

	if err != nil {
		return nil, err
	}

	return articles, nil
}
