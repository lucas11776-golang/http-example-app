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
func (ctx *SeniorAnalyst) ResearchArticles(context context.Context, interest []string) ([]*analyst.ArticleVerified, error) {
	unverifiedArticles, err := ctx.workspace.JuniorAnalyst.ResearchArticles(context, interest)

	if err != nil {
		return nil, err
	}

	verifiedArticles := []*analyst.ArticleVerified{}

	for _, unverifiedArticle := range unverifiedArticles {
		verifiedArticle, err := ctx.VerifiedArticle(context, unverifiedArticle)

		if err != nil {
			continue
		}

		verifiedArticles = append(verifiedArticles, verifiedArticle)
	}

	return verifiedArticles, nil
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
func (ctx *SeniorAnalyst) createVerifiedArticle(unverifiedArticle *analyst.ArticleCapture, verifiedArticle orm.Values) (*analyst.ArticleVerified, error) {
	verifiedArticle["article_capture_id"] = unverifiedArticle.ID

	verified, err := orm.Model(analyst.ArticleVerified{}).Insert(verifiedArticle)

	if err != nil {
		return nil, err
	}

	if trusted, ok := verifiedArticle["trusted"].([]interface{}); ok {
		verified.Trusted, _ = ctx.createSources(verified, trusted, true)
	}

	if trusted, ok := verifiedArticle["untrusted"].([]interface{}); ok {
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

	verified, err = ctx.createVerifiedArticle(article, verification)

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

	var verification orm.Values

	if err := json.Unmarshal([]byte(result), &verification); err != nil {
		return nil, err
	}

	return verification, nil
}
