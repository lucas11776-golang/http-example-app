package analyst

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

type SeniorAnalyst struct {
}

// Comment
// Ask is week promp just for test workflow of Scraper and Analyst.
func ValidateArticle(context context.Context, article *models.ArticleCaputure) (*models.Article, error) {
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := `
	You are a senior analyst working for company that analyze web for news for clients, you are working with junior/mid level analyst that
	scrap and analyze news in the web as a senior you supports to validate the article that is valid or not by checking the points below
	in the web:

	1. ✅ Check the Source's Reputation
	Use established news outlets (e.g., BBC, Reuters, The New York Times).

	Why it matters: Trusted organizations follow strict editorial standards.
	📚 Taught by: Columbia Journalism School, NYU Journalism

	2. 👤 Verify the Author's Credentials
	Look for named journalists or experts with relevant experience.

	Why it matters: Credible authors can be held accountable.
	📚 Referenced in: SPJ Code of Ethics

	3. 📊 Look for Citations and Evidence
	Choose sources that provide facts, documents, and expert quotes.

	Why it matters: Evidence-based reporting builds trust.
	📚 From: The Elements of Journalism

	4. 🕒 Evaluate the Date of Publication
	Make sure the source is current and relevant.

	Why it matters: Outdated information can mislead your reporting.
	📚 Taught in: News writing & editing courses

	5. ⚖️ Detect Bias or Objectivity
	Watch for emotionally charged or one-sided language.

	Why it matters: Reliable sources strive for fairness and neutrality.
	📚 From: SPJ Code of Ethics

	6. 🔍 Cross-Check with Other Reliable Sources
	Confirm information with at least 2-3 other reputable outlets.

	Why it matters: Truth stands up to independent verification.
	📚 Core skill in: Fact-checking workshops

	7. 🌐 Inspect the Domain and Site Design
	Use sites with professional design and trusted domains (e.g., .gov, .edu).

	Why it matters: Real sources don't mimic clickbait aesthetics.
	📚 Taught in: Media literacy programs


	Below is the article you analyze and give feedback one more key point the article content maybe empty if it empty
	get content from trust source and build the article content please the content must atleast be 120 words:

	Title:
	Standard Bank slammed in court after repossessing and auctioning Soweto home for R200

	Published At:
	16 May 2025

	Category:
	Finance

	Description:
	1992 property loan sparks 30-year fight.

	Article Content:
	Standard Bank faced legal action after controversially repossessing and auctioning a Soweto home for a mere R200, a move stemming from a property loan dispute initiated in 1992 that has now escalated into a three-decade legal battle.


	Our system whats the result in JSON object in array containing the following interface and
	place the data inside <result><result> also do not include` + "```json ``` in results." + `

	interface Article {
		rating: number;         // Article rating out of 10 not float only decimals.
		trusted: []string;      // Trusted source reported on article only the name.
		untrusted: []string;    // Untrusted source reported on article.
		title: string;          // Get the article a attractive title that will want the use to read it.
		content: string;        // Update article content base on all trusted source.
		html: string;        	// Update article base on all trusted source - in html please make the content easy readable here by apply br,b etc.
		description: string;    // Short summary of article within 20-50 words.
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
	})

	if err != nil {
		return nil, err
	}

	matches := RESULT_REGEX.FindStringSubmatch(response.Text())

	file, _ := os.Create("article-report.txt")

	file.Write([]byte(response.Text()))

	if len(matches) == 0 {
		return nil, nil
	}

	var articles []models.ArticleCaputure

	err = json.Unmarshal([]byte(matches[0]), &articles)

	if err != nil {
		return nil, err
	}

	// return articles, nil

	return nil, nil
}
