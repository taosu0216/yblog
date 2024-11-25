package aiService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"log"
)

type gdJsonResp struct {
	Msg string `json:"msg" jsonschema:"title=msg,description=生成的对应的大约20字左右的文章摘要"`
}

func GenDesc(info string) (string, error) {
	type content struct {
		Msg string `json:"msg"`
	}

	// 将结构体序列化为 JSON 字符串
	contentJSON, err := json.Marshal(content{Msg: info})
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "", fmt.Errorf("marshaling json: %w", err)
	}

	// 将 JSON 字符串转换为 string
	contentStr := string(contentJSON)
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: gdSystem,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: gdSystem,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: contentStr,
		},
	}

	log.Println("开始生成文章摘要")

	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: dialogue,
		},
	)
	if err != nil {
		return "", fmt.Errorf("CreateChatCompletion err: %w", err)
	}
	log.Println("文章摘要生成成功")

	respStr := resp.Choices[0].Message.Content
	respObj, err := articleRespStrToModel(respStr)
	if err != nil {
		return "", fmt.Errorf("linkRespStrToModel err: %w", err)
	}

	return respObj.Msg, nil
}

func articleRespStrToModel(respStr string) (gdJsonResp, error) {
	var resp gdJsonResp
	err := json.Unmarshal([]byte(respStr), &resp)
	log.Println("@@@@@@@@@@@@@@@", respStr)
	if err != nil {
		return gdJsonResp{}, err
	}
	return resp, nil
}
