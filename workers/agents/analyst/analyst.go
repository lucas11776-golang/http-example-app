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
func (ctx *SeniorAnalyst) ValidateArticle(context context.Context, article *models.ArticleCaputure) (*models.Article, error) {
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
	What to expect from interest rates, and major blow to Trump's tariffs plan

	Published At:
	28 May 2025

	Category:
	Business

	Description:
	The South African Reserve Bank is set to announce its monetary policy decision on May 29, with economists divided on a rate cut, while a US federal court has blocked most of Donald Trump's global tariffs.

	Article Content:
	The South African rand remained steady on Wednesday, just one day before the central bank's interest rate decision. It was trading at 17.9375 against the U.S. dollar, showing little change from Tuesday's closing level. The US dollar was approximately 0.3% stronger against a basket of currencies as markets awaited the release of the minutes from the Federal Reserve's latest policy meeting and upcoming economic data for clues about the US interest rate outlook. The South African Reserve Bank (SARB) will announce its monetary policy decision today at around 13h00 on May 29. On Thursday, 29 May, the rand was trading at R17.95 to the dollar, R24.13 to the pound and R20.19 to the euro. Oil was trading slightly lower at $65.77 a barrel. Here are five other important things happening in and affecting South Africa today: What to expect from interest rates: The majority of economists polled by Reuters and Bloomberg expect the bank to cut its main lending rate by 25 basis points. However, a significant minority, including local economists, think the rate could be left unchanged. Tariffs ruled illegal: The US trade court ruled most of President Donald Trump's global tariffs illegal, undermining a key part of his economic agenda. A three-judge panel at the US Court of International Trade unanimously sided with Democratic-led states and small businesses, stating that Trump wrongfully invoked an emergency law to justify his tariffs. 30% electricity price hike a reality: The Electricity Resellers Association of South Africa (ERASA) will determine a response this week to a looming 30% electricity tariff increase affecting end-users in Eskom areas. Recent bills show that a building in eastern Pretoria saw its costs rise from R357,921 to R464,081. Resellers must pass on these costs, despite users expecting only a 12.74% increase approved by Nersa. ERASA argues this average does not reflect larger increases for lower electricity users.
 
	
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

type ArticeImageDescription struct {
	Title       string
	Description string
}

// Comment
func (ctx *SeniorAnalyst) DescribeArticle(context context.Context, article any) (*ArticeImageDescription, error) {

	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := `
	You are a senior analyst working for company that analyze web for news for clients, you are working with a senior graphic designer.
	You want a news article image for article to be publish to our company news site.
	You must describe to the senior graphic designer what the image should be like.
	Only state the concept.
	Please place your description in <result><result>.

	Below are are bullet points of what a good article image should contain:

	✅ 1. Relevance
	Directly relates to the article's topic or theme.

	Helps explain or reinforce the main message.

	✅ 2. Visual Clarity
	High-resolution and sharp.

	Simple composition with a clear subject.

	✅ 3. Attention-Grabbing
	Uses strong colors, contrast, or movement.

	Faces or human emotion often improve appeal.

	✅ 4. Brand Consistency
	Aligns with the visual identity of your brand or publication.

	Consistent style, color palette, and tone.

	✅ 5. Emotionally Evocative
	Triggers curiosity or emotion.

	Strengthens the reader's connection to the content.

	✅ 6. Originality
	Unique visuals stand out better than generic stock photos.

	Custom illustrations or AI-generated images can be effective.

	✅ 7. Complementary (Not Distracting)
	Enhances the message without overwhelming it.

	Avoids clutter or overly complex scenes.

	✅ 8. Text-Friendly (if needed)
	Leaves space for titles or captions.

	Ensures readability with good contrast.

	✅ 9. Correct Format & Size
	Optimized for the web (e.g., JPEG, WebP).

	Compressed for fast loading, but not at the expense of quality.

	✅ 10. Legal and Ethical
	Proper licensing or original content.

	Respects copyrights and usage rights.


	Below is an article to be published:

	Title:
	What to expect from interest rates, and major blow to Trump's tariffs plan

	Category:
	Finance

	Content:
	The South African rand remained steady on Wednesday, just one day before the central bank's interest rate decision. It was trading at 17.9375 against the U.S. dollar, showing little change from Tuesday's closing level. The US dollar was approximately 0.3% stronger against a basket of currencies as markets awaited the release of the minutes from the Federal Reserve's latest policy meeting and upcoming economic data for clues about the US interest rate outlook. The South African Reserve Bank (SARB) will announce its monetary policy decision today at around 13h00 on May 29. On Thursday, 29 May, the rand was trading at R17.95 to the dollar, R24.13 to the pound and R20.19 to the euro. Oil was trading slightly lower at $65.77 a barrel. Here are five other important things happening in and affecting South Africa today: What to expect from interest rates: The majority of economists polled by Reuters and Bloomberg expect the bank to cut its main lending rate by 25 basis points. However, a significant minority, including local economists, think the rate could be left unchanged. Tariffs ruled illegal: The US trade court ruled most of President Donald Trump's global tariffs illegal, undermining a key part of his economic agenda. A three-judge panel at the US Court of International Trade unanimously sided with Democratic-led states and small businesses, stating that Trump wrongfully invoked an emergency law to justify his tariffs. 30% electricity price hike a reality: The Electricity Resellers Association of South Africa (ERASA) will determine a response this week to a looming 30% electricity tariff increase affecting end-users in Eskom areas. Recent bills show that a building in eastern Pretoria saw its costs rise from R357,921 to R464,081. Resellers must pass on these costs, despite users expecting only a 12.74% increase approved by Nersa. ERASA argues this average does not reflect larger increases for lower electricity users
	`

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

	file, _ := os.Create("article-description.txt")

	file.Write([]byte(response.Text()))

	if len(matches) == 0 {
		return nil, nil
	}

	var articles []models.ArticleCaputure

	err = json.Unmarshal([]byte(matches[0]), &articles)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
