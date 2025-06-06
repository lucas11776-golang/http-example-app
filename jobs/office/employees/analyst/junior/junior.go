package junior

import (
	"context"
	"encoding/json"
	"server/env"
	"server/jobs/office/utils"
	"server/jobs/workspace"
	"server/jobs/workspace/paperwork/analyst"
	"server/utils/prompt"
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
func (ctx *JuniorAnalyst) createArticleCapture(capture orm.Values) (*analyst.ArticleCapture, error) {
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
func (ctx *JuniorAnalyst) ResearchArticles(context context.Context, interest []string) ([]*analyst.ArticleCapture, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	prompt, err := ctx.workspace.Prompt.Generate("analyst.junior.research-article", prompt.PromptData{"interest": &interest})

	if err != nil {
		return nil, err
	}

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				{Text: prompt},
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

	result := utils.PaperworkResult(string(response.Text()))

	if result == "" {
		return []*analyst.ArticleCapture{}, nil
	}

	var captures []orm.Values

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
