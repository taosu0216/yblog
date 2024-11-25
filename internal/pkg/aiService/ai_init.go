package aiService

import (
	"github.com/sashabaranov/go-openai"
)

var (
	openaiClient *openai.Client
	//schemaClient *instructor.InstructorOpenAI
	model string
)

func InitAiservice(baseurl, apikey, modelstr string) {
	config := openai.DefaultConfig(apikey)
	config.BaseURL = baseurl
	openaiClient = openai.NewClientWithConfig(config)

	model = modelstr

	//schemaClient = instructor.FromOpenAI(
	//	openaiClient,
	//	instructor.WithMode(instructor.ModeJSON),
	//	instructor.WithMaxRetries(3),
	//)
}
