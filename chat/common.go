package chat

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sashabaranov/go-openai"
)

func RandomString(n int) string {
	var letter []rune
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	numberChars := "0123456789"
	chars := fmt.Sprintf("%s%s", lowerChars, numberChars)
	letter = []rune(chars)
	var str string
	b := make([]rune, n)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letter[seededRand.Intn(len(letter))]
	}
	str = string(b)
	return str
}

var (
	GPTModels = []string{
		openai.GPT432K0314,
		openai.GPT432K,
		openai.GPT40314,
		openai.GPT4,
		openai.GPT3Dot5Turbo0301,
		openai.GPT3Dot5Turbo,
		openai.GPT3TextDavinci003,
		openai.GPT3TextDavinci002,
		openai.GPT3TextCurie001,
		openai.GPT3TextBabbage001,
		openai.GPT3TextAda001,
		openai.GPT3TextDavinci001,
		openai.GPT3DavinciInstructBeta,
		openai.GPT3Davinci,
		openai.GPT3CurieInstructBeta,
		openai.GPT3Curie,
		openai.GPT3Ada,
		openai.GPT3Babbage,
	}
)
