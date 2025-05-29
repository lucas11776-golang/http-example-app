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
	A visually striking concept featuring a stylized, assertive profile of Donald Trump, perhaps with an outstretched hand gesture as if presenting or imposing something. Emanating from his hand or as an extension of his rhetorical stance, a bold, distinct label or speech bubble shape prominently displays the text "AFRIKANER REFUGEES." However, this label is visually depicted as fragmented, cracked, or actively disintegrating into digital dust or pixels, symbolizing its falsity and the article's "unpacking" of the claim. In the background, subtly rendered and slightly out of focus, are elements suggestive of South Africa, such as the vibrant colors of its flag, a faint outline of its map, or a hint of an iconic natural landscape. The overall mood should be serious and analytical, highlighting the political opportunism and the fabricated nature of the claim.
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
