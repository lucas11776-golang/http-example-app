package designer

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"server/env"
	"server/models"

	"google.golang.org/genai"
)

var (
	RESULT_REGEX = regexp.MustCompile(`<result>(.*?)</result>`)
)

type GraphicDesigner struct {
}

// TODO: Need to improve prompt and find out the is image gen that has thinking if not designer has to ask senior analyst for key points.
// Comment
func (ctx *GraphicDesigner) DesignArticleImage(context context.Context, intrests []string) ([]models.ArticleCaputure, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := `
	You are a senior graphic designer working for a analyst company that analyses the web for news,
	You are working close with a senior analyst that requires you to design article image base on this concept:

	Below concept you should use:
	A striking visual centered on Parks Tau, depicted under intense scrutiny with a serious expression. A dynamic, stylized hand, predominantly in a strong red (symbolizing the EFF), is positioned to convey pressure and accusation towards him. A subtle, blurred background hints at the National Lotteries Commission.
	`

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				{Text: ask},
			},
			Role: genai.RoleUser,
		},
	}

	result, err := client.Models.GenerateContent(context, env.Env("AI_MODEL_IMAGE"), content, &genai.GenerateContentConfig{
		ResponseModalities: []string{"TEXT", "IMAGE"},
	})

	if err != nil {
		return nil, err
	}

	for _, part := range result.Candidates[0].Content.Parts {
		if part.Text != "" {
			fmt.Println(part.Text)
		} else if part.InlineData != nil {
			imageBytes := part.InlineData.Data
			outputFilename := "./temp/gemini_generated_image.png"
			_ = os.WriteFile(outputFilename, imageBytes, 0644)
		}
	}

	return nil, nil
}
