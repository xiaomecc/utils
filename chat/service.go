package chat

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/sashabaranov/go-openai"
)

type Api struct {
	config OpenAIConfig
	cli    *openai.Client
}

func NewApi(config *OpenAIConfig) *Api {
	return &Api{
		config: *config,
		cli:    NewClient(config.ApiKey, config.Proxy),
	}
}

func NewClient(authToken, proxy string) *openai.Client {
	config := openai.DefaultConfig(authToken)
	if proxy != "" {
		uri, _ := url.Parse(proxy)
		config.HTTPClient = &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				// 设置代理
				Proxy: http.ProxyURL(uri),
			},
		}
	}
	return openai.NewClientWithConfig(config)
}

func (api *Api) GetChatMessage(ctx context.Context, reqMsgs []openai.ChatCompletionMessage) (chatResponse openai.ChatCompletionResponse, response openai.CompletionResponse, err error) {
	switch api.config.Model {
	case openai.GPT3Dot5Turbo0301, openai.GPT3Dot5Turbo, openai.GPT4, openai.GPT40314, openai.GPT432K0314, openai.GPT432K:
		req := openai.ChatCompletionRequest{
			Model:            api.config.Model,
			MaxTokens:        api.config.MaxLength,
			Temperature:      1.0,
			Messages:         reqMsgs,
			Stream:           false,
			TopP:             1,
			FrequencyPenalty: 0.1,
			PresencePenalty:  0.1,
		}
		chatResponse, err = api.cli.CreateChatCompletion(ctx, req)
		return
	case openai.GPT3TextDavinci003, openai.GPT3TextDavinci002, openai.GPT3TextCurie001, openai.GPT3TextBabbage001, openai.GPT3TextAda001, openai.GPT3TextDavinci001, openai.GPT3DavinciInstructBeta, openai.GPT3Davinci, openai.GPT3CurieInstructBeta, openai.GPT3Curie, openai.GPT3Ada, openai.GPT3Babbage:
		prompt := reqMsgs[len(reqMsgs)-1].Content
		req := openai.CompletionRequest{
			Model:       api.config.Model,
			MaxTokens:   api.config.MaxLength,
			Temperature: 0.6,
			Prompt:      prompt,
			Stream:      true,
			//Stop:             []string{"\n\n\n"},
			TopP:             1,
			FrequencyPenalty: 0.1,
			PresencePenalty:  0.1,
		}
		response, err = api.cli.CreateCompletion(ctx, req)
		return
	default:
		err = fmt.Errorf("模型不存在")
		return
	}
}

func (api *Api) GetChatMessageStream(ctx context.Context, reqMsgs []openai.ChatCompletionMessage) (stream *openai.ChatCompletionStream, err error) {
	req := openai.ChatCompletionRequest{
		Model:            api.config.Model,
		MaxTokens:        api.config.MaxLength,
		Temperature:      1.0,
		Messages:         reqMsgs,
		Stream:           true,
		TopP:             1,
		FrequencyPenalty: 0.1,
		PresencePenalty:  0.1,
	}
	stream, err = api.cli.CreateChatCompletionStream(ctx, req)
	if err != nil {
		err = fmt.Errorf("[ERROR] ChatCompletionStream error: %s", err.Error())
		return
	}
	return
}

func (api *Api) GetImageMessage(ctx context.Context, requestMsg string) {
	var err error

	prompt := strings.TrimPrefix(requestMsg, "/image ")
	req := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	resp, err := api.cli.CreateImage(ctx, req)
	if err != nil {
		err = fmt.Errorf("[ERROR] generate image error: %s", err.Error())
		return
	}

	imgBytes, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		err = fmt.Errorf("[ERROR] image base64 decode error: %s", err.Error())
		return
	}

	date := time.Now().Format("2006-01-02")
	imageDir := fmt.Sprintf("assets/images/%s", date)
	err = os.MkdirAll(imageDir, 0700)
	if err != nil {
		err = fmt.Errorf("[ERROR] create image directory error: %s", err.Error())
		return
	}

	imageFileName := fmt.Sprintf("%s.png", RandomString(16))
	err = os.WriteFile(fmt.Sprintf("%s/%s", imageDir, imageFileName), imgBytes, 0600)
	if err != nil {
		err = fmt.Errorf("[ERROR] write png image error: %s", err.Error())
		return
	}

	//msg := fmt.Sprintf("api/%s/%s", imageDir, imageFileName)
	return
}
