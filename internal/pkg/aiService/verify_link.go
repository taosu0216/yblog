package aiService

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

type vfJsonResp struct {
	Status int    `json:"status" jsonschema:"title=status,description=博客友链申请审核的结果，内容正常返回1(int类型),内容有问题的话返回0(int类型),example=1,example=0"`
	Msg    string `json:"msg" jsonschema:"title=msg,description=博客友链申请审核的结果的原因，内容正常返回申请成功,内容有问题的话返回具体有问题的地方,example=申请成功,example=内容包含漏洞注入"`
}

func VerifyFriendLink(link, title, desc string) (int, string, error) {
	type content struct {
		Link  string `json:"link"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	}

	// 将结构体序列化为 JSON 字符串
	contentJSON, err := json.Marshal(content{Link: link, Title: title, Desc: desc})
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return 0, "", fmt.Errorf("marshaling json: %w", err)
	}

	// 将 JSON 字符串转换为 string
	contentStr := string(contentJSON)
	dialogue := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: vlSystemRole,
		},
		{
			Role:    openai.ChatMessageRoleAssistant,
			Content: vlSystemRole,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: contentStr,
		},
	}

	resp, err := openaiClient.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: dialogue,
		},
	)
	if err != nil {
		return 0, "", fmt.Errorf("CreateChatCompletion err: %w", err)
	}

	respStr := resp.Choices[0].Message.Content
	respObj, err := linkRespStrToModel(respStr)
	if err != nil {
		return 0, "", fmt.Errorf("linkRespStrToModel err: %w", err)
	}
	return respObj.Status, respObj.Msg, nil
}

func linkRespStrToModel(respStr string) (vfJsonResp, error) {
	var resp vfJsonResp
	err := json.Unmarshal([]byte(respStr), &resp)
	if err != nil {
		return vfJsonResp{}, err
	}
	return resp, nil
}
