package openai

import (
	"context"
	"errors"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"go-chatgpt-api/cache"
	"go-chatgpt-api/config"
	"go-chatgpt-api/models"
	services "go-chatgpt-api/service"
	"go-chatgpt-api/utils/ip"
	"io"
	"log"
	"time"
)

var askLogService services.AskLogService

// ChatGPTResponseBody 请求体
type ChatGPTResponseBody struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int                      `json:"created"`
	Model   string                   `json:"model"`
	Choices []map[string]interface{} `json:"choices"`
	Usage   map[string]interface{}   `json:"usage"`
}

// ChatGPTRequestBody 响应体
type ChatGPTRequestBody struct {
	Model            string  `json:"model"`
	Prompt           string  `json:"prompt"`
	MaxTokens        int     `json:"max_tokens"`
	Temperature      float32 `json:"temperature"`
	TopP             int     `json:"top_p"`
	FrequencyPenalty int     `json:"frequency_penalty"`
	PresencePenalty  int     `json:"presence_penalty"`
}

func Ask(msg, ipParam string) (string, error) {
	log.Printf("ask request content:%s", msg)

	openAiClient := GetOpenAIClientFromCache(ipParam)

	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4000,
		Prompt:    msg,
	}
	resp, err := openAiClient.CreateCompletion(ctx, req)
	if err != nil {
		go askLogService.Record(&models.AskLog{
			UserId:     0,
			Request:    msg,
			Content:    err.Error(),
			Method:     "Ask",
			CreateTime: time.Now(),
			RequestIp:  ipParam,
			Address:    *ip.GetIpAddress(ipParam),
		})
		return "", err
	}
	result := resp.Choices[0].Text
	go askLogService.Record(&models.AskLog{
		UserId:     0,
		Request:    msg,
		Content:    result,
		Method:     "Ask",
		CreateTime: time.Now(),
		RequestIp:  ipParam,
		Address:    *ip.GetIpAddress(ipParam),
	})
	fmt.Println(result)
	return result, nil
}

func GetOpenAIClientFromCache(ip string) *gogpt.Client {
	var openAiClient *gogpt.Client
	if x, found := cache.OpenAiClientCache.Get(ip); found {
		openAiClient = x.(*gogpt.Client)
	} else {
		apiKey := config.GetOpenAiApiKey()
		if apiKey == nil {
			panic("未配置apiKey")
		}
		openAiClient = gogpt.NewClient(*apiKey)
		cache.OpenAiClientCache.Set(ip, openAiClient, 600*time.Second)
	}
	return openAiClient
}

func AskStream(msg, ip string) {
	ctx := context.Background()
	openAiClient := GetOpenAIClientFromCache(ip)
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 4096,
		Prompt:    msg,
		Stream:    true,
	}
	stream, err := openAiClient.CreateCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return
		}

		fmt.Printf("Stream response: %v\n", response)
	}
}

func GenerateImg(msg, ipParam string) (string, error) {
	openAiClient := GetOpenAIClientFromCache(ipParam)
	ctx := context.Background()
	req := gogpt.ImageRequest{
		Prompt:         msg,
		ResponseFormat: "url",
		Size:           "512x512",
	}
	resp, err := openAiClient.CreateImage(ctx, req)
	if err != nil {
		go askLogService.Record(&models.AskLog{
			UserId:     0,
			Request:    msg,
			Method:     "GenerateImg",
			Content:    err.Error(),
			CreateTime: time.Now(),
			RequestIp:  ipParam,
			Address:    *ip.GetIpAddress(ipParam),
		})
		return "", err
	}

	go askLogService.Record(&models.AskLog{
		UserId:     0,
		Request:    msg,
		Method:     "GenerateImg",
		Content:    resp.Data[0].URL,
		CreateTime: time.Now(),
		RequestIp:  ipParam,
		Address:    *ip.GetIpAddress(ipParam),
	})

	return resp.Data[0].URL, nil
}
