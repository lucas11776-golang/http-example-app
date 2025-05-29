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
	Trump's Afrikaners are South African opportunists, not refugees

	Published At:
	28 May 2025

	Category:
	Politics

	Description:
	Roger Southall argues that the claim of Afrikaners being refugees from genocide in South Africa is fabricated and serves the Trump administration's political agenda.

	Article Content:
	South Africans are wearily attuned to governments' Orwellian misuse of language. After all, South Africa is a country where a one-time government passed a law (the Natives Abolition of Passes and Coordination of Documents Act of 1952) which extended rather than abolishing the notorious pass system. This made it compulsory for black South Africans over the age of 16 to carry a passbook. And the same government passed the Extension of University Education Act of 1959 which made it more, not less, difficult for black students to register at “open” (or white) universities. So perhaps they should not be unduly surprised that the government of the US has imported 49 Afrikaners and labelled them as “refugees”. The claim is that they are escaping from the persecution of Afrikaners – and white people more broadly – in South Africa today. The Trump administration knows perfectly well this claim is a complete fabrication. As President Cyril Ramaphosa and his government have pointed out, there is no evidence whatsoever that Afrikaners or white people more generally are subject to genocide. True, South Africa has one of the highest murder rates in the world. But it is poor black South Africans – not whites – who are principal victims of such deadly violence. Nor are Afrikaners/whites subject to persecution. Along with all other South Africans, their human rights are protected by a constitution. This is no mere piece of paper. Its provisions are (albeit imperfectly, and unlike in the US these days) largely enforced by the courts. Furthermore, genocide implies the deliberate elimination of a people on racial, ethnic, or religious grounds. Therefore, if a genocide of whites and Afrikaners was taking place, we might assume that their numbers would be falling. In fact the reverse is true. The white population has continued to grow (albeit slowly) in absolute numbers since 1994. Worse, the characterisation of Afrikaners as refugees at a moment in time when the people of Gaza are daily subject to a regime of death, terror, and murder inflicted on them by the Israeli government is not merely an absurdity but a downright insult to those genuinely subject to genocide. So, what is really going on? The drivers Extensive commentary has correctly highlighted the motivations of the Trump administration. First, the administration has launched an attack on what it terms the “tyranny” of “diversity, equity and inclusion” policies across the entire spectrum of public and private institutions in America. Critics argue this is driven by an appeal to Trump's white Christian nationalist political base. Because post-apartheid South Africa, rightly or wrongly, has become the poster-country of diversity, equity and inclusion policies internationally, because of its constitutional commitment to non-racialism and diversity, it has been singled out for attack. Secondly, labelling Afrikaners as refugees plays to the insecurities of Trump's political base. This finds the idea of a white minority being ruled by a black majority government difficult to swallow. Third, characterising Afrikaners as subject to genocide is a very deliberate response to South Africa's charging of Israel as guilty of genocide against the Palestinian people before the International Court of Justice. But this is unacceptable to the US Christian nationalist right. For them the existence of Israel represents the realisation of Biblical truth – the return of Jews to the Holy Land. Trump is saying that the US can and will play the same game, using it to clobber South Africa regardless of the groundlessness of the charge. But, being Trump, he will balance pandering to his support base against what economic benefits he can extract from South Africa. The landscape But what of the 49 Afrikaners themselves? Why have they chosen to accept the opportunity offered to them by the US government? After all, extensive attention in the South African media has been given to Afrikaners who have defiantly stated that they are committed to staying in South Africa. The reasons they give are that it's their home. And they fully accept that, at least formally, South Africa has become a non-racial democracy. Likewise, as I have detailed in my book on Whites and Democracy in South Africa, Afrikaners and whites have not only survived in democratic South Africa but, generally, have prospered economically. Furthermore, whites as a “population group” (to use outdated apartheid-era terminology) have participated fully in South African democracy. They are more highly disposed to voting in elections than other racial groupings, and de facto, they are well represented in parliament and local government by the Democratic Alliance, which is a vigorous defender of their interests. But (there is always a but), if we want to guess the motivations of Trump's 49 “refugees”, we need to bear in mind the following. First, until we know more about the personal circumstances of the individuals involved, we cannot really know what has driven them to take the drastic step of leaving families and their personal history behind by moving to America. Second, most whites have responded to the arrival of democracy pragmatically. They have their numerous complaints, notably about equity employment (affirmative action policies in favour of blacks) which they view as discriminatory against whites. But they have continued to enjoy high rates of employment. Indeed they continue to occupy the higher ranks of employment in the private sector in disproportionate numbers. However, although many whites continue to live in a de facto overwhelmingly white world, both at work and at their homes in suburbia, there remains a minority which has remained wholly unreconciled to the changes which have taken place politically and economically since 1994. The armed opposers linked to the far-right have long been defeated. But we may presume the 49 belong to a broader category of passive resisters who have withdrawn into a white world as much as possible. Third, although most whites continue to do well economically, the changes which have taken place since 1994 have led to the re-appearance of a small class of largely uneducated poor whites who feel excluded from employment by equity employment legislation. And who generally feel the loss of their racial status under democracy. Opportunists, not refugees Having said all that, some interesting questions remain. Presumably the Afrikaner 49 belonged to that category of whites which, for one reason or another, is disposed to leave South Africa. However, emigrating requires jumping through numerous hoops; meeting educational and professional qualifications, getting a job offer, having sufficient financial resources to take with them to support themselves and their families before they can qualify for recipient countries' social security systems, and so on. Apart from the emotional costs involved, emigration is not always the easiest of options, even for those who wish to “escape”. The evidence suggests that the heads of household among the Afrikaner 49 are drawn not only from that minority of Afrikaners who are totally unreconciled to democracy, but who – quite simply – are opportunists who have availed themselves of a short cut to emigrate. The Conversation Roger Southall, Professor of Sociology, University of the Witwatersrand. This article is republished from The Conversation under a Creative Commons license. Read the original article.
  

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
func DescribeArticle(context context.Context, article any) (*ArticeImageDescription, error) {

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


	Below is an article to be published.

	Title:
	Unpacking Trump's 'Afrikaner Refugees' Claim: An Expert Analysis of Political Opportunism in South Africa

	Category:
	Politics

	Content:
	South Africans are wearily attuned to governments' Orwellian misuse of language. After all, South Africa is a country where a one-time government passed a law (the Natives Abolition of Passes and Coordination of Documents Act of 1952) which extended rather than abolishing the notorious pass system. This made it compulsory for black South Africans over the age of 16 to carry a passbook. And the same government passed the Extension of University Education Act of 1959 which made it more, not less, difficult for black students to register at “open” (or white) universities. So perhaps they should not be unduly surprised that the government of the US has imported 49 Afrikaners and labelled them as “refugees”. The claim is that they are escaping from the persecution of Afrikaners – and white people more broadly – in South Africa today. The Trump administration knows perfectly well this claim is a complete fabrication. As President Cyril Ramaphosa and his government have pointed out, there is no evidence whatsoever that Afrikaners or white people more generally are subject to genocide. True, South Africa has one of the highest murder rates in the world. But it is poor black South Africans – not whites – who are principal victims of such deadly violence. Nor are Afrikaners/whites subject to persecution. Along with all other South Africans, their human rights are protected by a constitution. This is no mere piece of paper. Its provisions are (albeit imperfectly, and unlike in the US these days) largely enforced by the courts. Furthermore, genocide implies the deliberate elimination of a people on racial, ethnic, or religious grounds. Therefore, if a genocide of whites and Afrikaners was taking place, we might assume that their numbers would be falling. In fact the reverse is true. The white population has continued to grow (albeit slowly) in absolute numbers since 1994. Worse, the characterisation of Afrikaners as refugees at a moment in time when the people of Gaza are daily subject to a regime of death, terror, and murder inflicted on them by the Israeli government is not merely an absurdity but a downright insult to those genuinely subject to genocide. So, what is really going on? The drivers Extensive commentary has correctly highlighted the motivations of the Trump administration. First, the administration has launched an attack on what it terms the “tyranny” of “diversity, equity and inclusion” policies across the entire spectrum of public and private institutions in America. Critics argue this is driven by an appeal to Trump's white Christian nationalist political base. Because post-apartheid South Africa, rightly or wrongly, has become the poster-country of diversity, equity and inclusion policies internationally, because of its constitutional commitment to non-racialism and diversity, it has been singled out for attack. Secondly, labelling Afrikaners as refugees plays to the insecurities of Trump's political base. This finds the idea of a white minority being ruled by a black majority government difficult to swallow. Third, characterising Afrikaners as subject to genocide is a very deliberate response to South Africa's charging of Israel as guilty of genocide against the Palestinian people before the International Court of Justice. But this is unacceptable to the US Christian nationalist right. For them the existence of Israel represents the realisation of Biblical truth – the return of Jews to the Holy Land. Trump is saying that the US can and will play the same game, using it to clobber South Africa regardless of the groundlessness of the charge. But, being Trump, he will balance pandering to his support base against what economic benefits he can extract from South Africa. The landscape But what of the 49 Afrikaners themselves? Why have they chosen to accept the opportunity offered to them by the US government? After all, extensive attention in the South African media has been given to Afrikaners who have defiantly stated that they are committed to staying in South Africa. The reasons they give are that it's their home. And they fully accept that, at least formally, South Africa has become a non-racial democracy. Likewise, as I have detailed in my book on Whites and Democracy in South Africa, Afrikaners and whites have not only survived in democratic South Africa but, generally, have prospered economically. Furthermore, whites as a “population group” (to use outdated apartheid-era terminology) have participated fully in South African democracy. They are more highly disposed to voting in elections than other racial groupings, and de facto, they are well represented in parliament and local government by the Democratic Alliance, which is a vigorous defender of their interests. But (there is always a but), if we want to guess the motivations of Trump's 49 “refugees”, we need to bear in mind the following. First, until we know more about the personal circumstances of the individuals involved, we cannot really know what has driven them to take the drastic step of leaving families and their personal history behind by moving to America. Second, most whites have responded to the arrival of democracy pragmatically. They have their numerous complaints, notably about equity employment (affirmative action policies in favour of blacks) which they view as discriminatory against whites. But they have continued to enjoy high rates of employment. Indeed they continue to occupy the higher ranks of employment in the private sector in disproportionate numbers. However, although many whites continue to live in a de facto overwhelmingly white world, both at work and at their homes in suburbia, there remains a minority which has remained wholly unreconciled to the changes which have taken place politically and economically since 1994. The armed opposers linked to the far-right have long been defeated. But we may presume the 49 belong to a broader category of passive resisters who have withdrawn into a white world as much as possible. Third, although most whites continue to do well economically, the changes which have taken place since 1994 have led to the re-appearance of a small class of largely uneducated poor whites who feel excluded from employment by equity employment legislation. And who generally feel the loss of their racial status under democracy. Opportunists, not refugees Having said all that, some interesting questions remain. Presumably the Afrikaner 49 belonged to that category of whites which, for one reason or another, is disposed to leave South Africa. However, emigrating requires jumping through numerous hoops; meeting educational and professional qualifications, getting a job offer, having sufficient financial resources to take with them to support themselves and their families before they can qualify for recipient countries' social security systems, and so on. Apart from the emotional costs involved, emigration is not always the easiest of options, even for those who wish to “escape”. The evidence suggests that the heads of household among the Afrikaner 49 are drawn not only from that minority of Afrikaners who are totally unreconciled to democracy, but who – quite simply – are opportunists who have availed themselves of a short cut to emigrate. The Conversation Roger Southall, Professor of Sociology, University of the Witwatersrand. This article is republished from The Conversation under a Creative Commons license. Read the original article.
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

	// return articles, nil

	return nil, nil
}
