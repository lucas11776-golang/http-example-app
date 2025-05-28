package senior

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

	1. Check the Source's Reputation
	Summary: Use well-known, credible outlets with a history of accurate reporting.
	Reference: Trusted institutions (e.g., Columbia Journalism School) stress credibility through editorial standards.

	2. Verify the Author's Credentials
	Summary: Ensure the author is an expert or a professional journalist.
	Reference: Journalism education encourages sourcing from qualified, identifiable individuals.

	3. Look for Citations and Evidence
	Summary: Reliable sources back claims with data, documents, or expert quotes.
	Reference: The Elements of Journalism highlights evidence-based reporting.

	4. Evaluate the Date of Publication
	Summary: Use up-to-date information relevant to current events.
	Reference: News accuracy depends on timeliness, as noted in AP Stylebook and news writing guides.

	5. Detect Bias or Objectivity
	Summary: Trust sources with balanced, fact-based reporting over opinion-heavy content.
	Reference: Journalism ethics (e.g., SPJ Code of Ethics) emphasize fairness and impartiality.

	6. Cross-Check with Other Reliable Sources
	Summary: Confirm information by comparing it with other credible reports.
	Reference: Fact-checking is a core skill taught in journalism schools.

	7. Inspect the Domain and Site Design
	Summary: Professional outlets usually have clean, ad-limited websites with proper URLs (.edu, .org, .gov) or can be trusted media hows like (dailymaverick,iol etc).
	Reference: Media literacy curricula advise scrutiny of web sources’ authenticity.

	Below is the article you analyze and give feedback:

	Title: 
	Wits Researchers Identify New Biomarker for Early Diabetes Detection

	Article Content:
	JOHANNESBURG – In a significant step forward for public health, researchers at the University of the Witwatersrand (Wits) have announced the discovery of a new biomarker that shows great promise for the early detection of Type 2 diabetes. This groundbreaking finding could revolutionize how the disease is diagnosed, allowing for interventions before irreversible complications set in.\n\nLed by Professor Lerato Ndlovu from the Wits Medical Research Council Developmental Pathways for Health Research Unit (DPHRU), the team has identified a specific metabolic signature in blood samples that appears to be present years before traditional symptoms of Type 2 diabetes manifest. 'Current diagnostic methods often identify diabetes once significant damage has already occurred,' stated Prof. Ndlovu. 'Our newly identified biomarker offers a window of opportunity for early diagnosis and preventative measures, which is crucial for managing this escalating global health crisis.'\n\nThe research involved longitudinal studies of large cohorts of South African populations, examining various health markers over several years. The biomarker is detectable through a simple blood test, making it potentially scalable for widespread screening. Early detection means patients can make lifestyle changes, receive early medication, or participate in prevention programs that can halt or even reverse the progression of the disease.\n\nDiabetes and its complications, such as heart disease, kidney failure, and blindness, place a heavy burden on South Africa's healthcare system. This research, published in the prestigious journal 'Nature Medicine', positions Wits as a leader in metabolic disease research and offers tangible hope for improving the health outcomes of millions. Further clinical trials are planned to validate the biomarker's efficacy in diverse populations.

	Our system whats the result in JSON object containing the following interface and
	place the data inside <result><result> also do not include ` + "```json ``` in results." + `

	interface Article {
		trusted: []string;      // Trusted website/source reported on article.
		untrusted: []string;    // Untrusted website/source reported on article.
		title: string;          // Update article title base on all trusted website/source.
		content: string;        // Update article content base on all trusted website/source.
		rating: number;         // Article rating out of 10 not float only decimals.
	}`

	content := []*genai.Content{
		{
			Parts: []*genai.Part{
				// {Text: fmt.Sprintf(strings.Join(ask, "\r\n"), news.Finance)},
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
