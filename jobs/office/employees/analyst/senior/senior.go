package senior

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"server/env"
	"server/jobs/office/utils"
	"server/jobs/workspace"
	"server/jobs/workspace/paperwork/analyst"
	"server/utils/prompt"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/utils/cast"
	"google.golang.org/genai"
)

type SeniorAnalyst struct {
	workspace *workspace.Workspace
}

// Comment
func NewSeniorAnalyst(workspace *workspace.Workspace) *SeniorAnalyst {
	return &SeniorAnalyst{
		workspace: workspace,
	}
}

// Comment
func (ctx *SeniorAnalyst) Work(context context.Context) {
	go ctx.verifiedUnverifiedArticles(context)
}

// Comme
func (ctx *SeniorAnalyst) verifiedUnverifiedArticles(context context.Context) {
	articles, err := ctx.workspace.JuniorAnalyst.UnverifiedArticles(context)

	if err != nil {
		return
	}

	verifiedArticles := []*analyst.ArticleVerified{}

	for _, article := range articles {
		verified, err := ctx.VerifiedArticle(context, article)

		if err != nil {
			fmt.Println("ERROR", err)
			continue
		}

		verifiedArticles = append(verifiedArticles, verified)
	}

	fmt.Println(verifiedArticles)
}

// Comment
func (ctx *SeniorAnalyst) ResearchArticles(context context.Context, intrest []string) ([]*analyst.ArticleVerified, error) {
	return nil, nil
}

// Comment
func (ctx *SeniorAnalyst) createSources(article *analyst.ArticleVerified, sources []interface{}, trusted bool) ([]*analyst.ArticleVerifiedSource, error) {
	values := []orm.Values{}

	for _, source := range sources {
		site, ok := source.(map[string]interface{})

		if !ok {
			continue
		}

		values = append(values, orm.Values{
			"article_verified_id": article.ID,
			"name":                site["name"],
			"website":             site["website"],
			"trusted":             trusted,
		})
	}

	return orm.Model(analyst.ArticleVerifiedSource{}).InsertMany(values)
}

// Comment
func (ctx *SeniorAnalyst) createVerifiedArticle(article map[string]interface{}) (*analyst.ArticleVerified, error) {
	verified, err := orm.Model(analyst.ArticleVerified{}).
		Insert(orm.Values{
			"article_capture_id": article["article_capture_id"],
			"rating":             article["rating"],
			"title":              article["title"],
			"description":        article["description"],
			"content":            article["content"],
			"html":               article["html"],
		})

	if err != nil {
		return nil, err
	}

	if trusted, ok := article["trusted"].([]interface{}); ok {
		verified.Trusted, _ = ctx.createSources(verified, trusted, true)
	}

	if trusted, ok := article["untrusted"].([]interface{}); ok {
		verified.Untrusted, _ = ctx.createSources(verified, trusted, false)
	}

	return verified, nil
}

// Comment
func (ctx *SeniorAnalyst) VerifiedArticle(context context.Context, article *analyst.ArticleCapture) (*analyst.ArticleVerified, error) {
	verified, err := orm.Model(analyst.ArticleVerified{}).
		Where("article_capture_id", "=", article.ID).
		First()

	if err != nil || verified != nil {
		// TODO: Do something...
		return verified, nil // TODO: get sources...
	}

	if err := article.Verifying(time.Now()); err != nil {
		// TODO: Do something...
		return nil, err
	}

	verification, err := ctx.verifyArticle(context, article)

	if err != nil {
		return nil, err
	}

	rating := cast.Kind(reflect.Int, verification["rating"]).(int)

	if rating < 5 {
		reVerification, err := ctx.verifyArticle(context, article)

		if err != nil {
			return nil, err
		}

		reRating, ok := reVerification["rating"].(int)

		if ok && reRating > rating {
			verification = reVerification
		}
	}

	verification["article_capture_id"] = article.ID

	verified, err = ctx.createVerifiedArticle(verification)

	if err != nil {
		return nil, err
	}

	if err != article.Verified(time.Now()) {
		article.Verified(time.Now())
	}

	return verified, nil
}

// Comment
func (ctx *SeniorAnalyst) verifyArticle(context context.Context, article *analyst.ArticleCapture) (map[string]interface{}, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	prompt, err := ctx.workspace.Prompt.Generate("analyst.senior.verify-article", prompt.PromptData{"article": article})

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
	})

	if err != nil {
		return nil, err
	}

	result := utils.ResultFromPaperwork(response.Text())

	if result == "" {
		return nil, errors.New("something went wrong when trying to validate article")
	}

	var verification map[string]interface{}

	if err := json.Unmarshal([]byte(result), &verification); err != nil {
		return nil, err
	}

	return verification, nil
}
