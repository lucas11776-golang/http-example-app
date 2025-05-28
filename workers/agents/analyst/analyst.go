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
	State measures in place to cushion poor amid rising living costs — Ramaphosa

	Published At:
	22 May 2025

	Category:
	Politics

	Description:
	President Cyril Ramaphosa stated that his government has implemented measures to protect low-income South Africans from rising living costs, including macroeconomic policies and social welfare initiatives, despite a significant increase in the household food basket.
	
	Article Content:
	President Cyril Ramaphosa says his government, the Government of National Unity (GNU), has implemented adequate measures to protect low-income South Africans from the rising cost of necessities. This comes as the average household food basket has increased by nearly 40% relative to inflation. Ramaphosa emphasised that the government was fully aware of the financial pressures facing citizens and remains committed to supporting the most vulnerable through targeted relief measures. He was responding to oral questions in Parliament, Cape Town on Tuesday. “Government recognises the high cost of living facing South Africans. “Tackling poverty and the cost of living is one of the three strategic priorities of the GNU and forms a central pillar of the Medium TermDevelopment Plan. “South Africa's macroeconomic policy framework has been a key lever for shielding the poor from the high cost of living,” he said. Ramaphosa mentioned that the framework included an inflation target, which has helped to keep prices low and stable and has been important in reducing average prices. “Food price inflation has fallen quite significantly from 12.7 percent at the end of 2022 to 2.2 percent in March 2025. “Headline inflation, which is a measure of the general cost of living, has also declined, averaging 4.4 percent in 2024 and inflation has even moderated further to 2.7 percent in March 2025,” he said. “Food staples, such as maize meal, brown bread, rice, samp, milk, eggs, and other basic foodstuffs remain exempt from VAT, to help to cushion lower-income households in our country. “Our fiscal policy has been redistributive, prioritising poor and low-income households. The government spends around 60 percent of its revenue on the social wage, which includes spending on social grants on education and health.” Last week, Finance Minister Enoch Godongwana, presented his 2025 national budget, without Value-Added-Tax (VAT) but social grants were increased at a rate higher than inflation. The provision of free basic services, such as water and electricity, for indigent households is an essential measure in reducing the high cost of living. “This package of free municipal services continues to be a key tool for reducing poverty and inequality, and raising living standards and facilitating access to greater economic opportunities for many of our people.


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
