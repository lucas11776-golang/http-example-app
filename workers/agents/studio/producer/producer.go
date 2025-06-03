package producer

import (
	"context"
	"encoding/json"
	"os"
	"regexp"
	"server/env"
	"server/models"
	"server/workers/agents/studio/manager"

	"google.golang.org/genai"
)

var (
	RESULT_REGEX = regexp.MustCompile(`<result>(.*?)</result>`)
)

type Producer struct {
}

type NewsScript struct {
}

// Comment
// TODO: producer can edit studio effect base on article.
func (ctx *Producer) PrepareNewsScript(context context.Context, article []string, studio *manager.Studio) (*NewsScript, error) {

	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  env.Env("AI_KEY_AI"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		return nil, err
	}

	ask := `
	You are a news producer, working with reporter and you must write a news article to something a reporter can read on air,
	You should describe the studio scene and what the report should wear place you work in <result></result> tag.
	
	Details:

	Reporter's Name:
	Jane

	Anchor's Name:
	Bob

	News Article:

	Title:
	Critical Economic Junctions: SARB's Rate Decision & A Legal Blow to Trump's Tariff Legacy

	Category:
	Finance

	Content:
	The global financial landscape is abuzz with anticipated economic shifts, notably the South African Reserve Bank's (SARB) impending monetary policy decision and a significant legal challenge that has undermined a key aspect of former US President Donald Trump's trade agenda.\n\nFinancial markets are closely watching the SARB, which is poised to announce its latest interest rate decision on May 29. While the South African rand has shown relative stability against major currencies, trading around R17.93 against the U.S. dollar in recent sessions, economists remain divided on the SARB's next move. A majority of analysts, as polled by reputable financial news agencies like Reuters and Bloomberg, anticipate a 25-basis point cut to the main lending rate, aiming to stimulate economic growth. However, a notable contingent, including local economists, suggests the central bank might opt to keep rates unchanged, prioritizing inflation control amidst global uncertainties. The decision comes as markets also digest minutes from the U.S. Federal Reserve's latest policy meeting, seeking clues on the direction of global interest rates.\n\nConcurrently, a landmark ruling by the US Court of International Trade has delivered a major setback to the tariff policies implemented during the Trump administration. A unanimous three-judge panel sided with various Democratic-led states and small businesses, determining that President Trump had wrongfully invoked a national security law (Section 232) to justify imposing broad global tariffs on a range of imports, including steel and aluminum. This ruling significantly challenges the legal foundation of a central pillar of his past economic strategy, potentially impacting future trade policy discussions.\n\nDomestically in South Africa, concerns continue to mount over electricity tariffs. The Electricity Resellers Association of South Africa (ERASA) is deliberating its response to a looming 30% electricity tariff increase that is set to impact end-users in Eskom-supplied areas. Despite the National Energy Regulator of South Africa (NERSA) having approved an average increase of 12.74%, the actual costs passed on to consumers, particularly by resellers, are significantly higher due to various charges and complexities. This disparity has led to substantial jumps in recent electricity bills, underscoring ongoing challenges within the country's power sector and its impact on household and business finances.
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

	file, _ := os.Create("reporter-script.txt")

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
