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
func (ctx *SeniorAnalyst) createVerifiedArticleSource(article *analyst.ArticleVerified, source map[string]interface{}, trusted bool) (*analyst.ArticleVerifiedSource, error) {
	return orm.Model(analyst.ArticleVerifiedSource{}).
		Insert(orm.Values{
			"article_verified_id": article.ID,
			"name":                source["name"],
			"website":             source["website"],
			"trusted":             trusted,
		})
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

	for _, trusted := range article["trusted"].([]interface{}) {
		source, err := ctx.createVerifiedArticleSource(verified, trusted.(map[string]interface{}), true)

		if err != nil {
			// Do something
			continue
		}

		verified.Trusted = append(verified.Trusted, source)
	}

	for _, untrusted := range article["untrusted"].([]interface{}) {
		source, err := ctx.createVerifiedArticleSource(verified, untrusted.(map[string]interface{}), false)

		if err != nil {
			// Do something
			continue
		}

		verified.Trusted = append(verified.Trusted, source)
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
		return verified, nil
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

	result := utils.ResultFromPaperword(response.Text())

	if result == "" {
		return nil, errors.New("something went wrong when tring to validate article")
	}

	var verification map[string]interface{}

	if err := json.Unmarshal([]byte(result), &verification); err != nil {
		return nil, err
	}

	return verification, nil
}
