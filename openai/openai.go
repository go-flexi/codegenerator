package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Config is a struct that holds the configuration for the API
type Config struct {
	url              string
	model            string
	temperature      float64
	maxToken         int
	topP             float64
	frequencyPenalty float64
	presencePenalty  float64
}

// WithURL sets the URL for the API
func (c Config) WithURL(url string) Config {
	c.url = url
	return c
}

// WithModel sets the model for the API
func (c Config) WithModel(model string) Config {
	c.model = model
	return c
}

// WithTemperature sets the temperature for the API
func (c Config) WithTemperature(temperature float64) Config {
	c.temperature = temperature
	return c
}

// WithMaxToken sets the maxToken for the API
func (c Config) WithMaxToken(maxToken int) Config {
	c.maxToken = maxToken
	return c
}

// WithTopP sets the topP for the API
func (c Config) WithTopP(topP float64) Config {
	c.topP = topP
	return c
}

// WithFrequencyPenalty sets the frequencyPenalty for the API
func (c Config) WithFrequencyPenalty(frequencyPenalty float64) Config {
	c.frequencyPenalty = frequencyPenalty
	return c
}

// WithPresencePenalty sets the presencePenalty for the API
func (c Config) WithPresencePenalty(presencePenalty float64) Config {
	c.presencePenalty = presencePenalty
	return c
}

// DefaultConfig returns a Config with the default values
func DefaultConfig() Config {
	return Config{
		url:              "https://api.openai.com/v1/chat/completions",
		model:            "gpt-4-0125-preview",
		temperature:      1,
		maxToken:         3968,
		topP:             1,
		frequencyPenalty: 0,
		presencePenalty:  0,
	}
}

// API is the main struct for interacting with the OpenAI API
type API struct {
	apiKey     string
	httpClient *http.Client
}

// NewAPI creates a new API instance
func NewAPI(apiKey string, config Config) *API {
	return &API{
		apiKey:     apiKey,
		httpClient: http.DefaultClient,
	}
}

// Role is a type that represents the role of a message
type Message struct {
	Role    Role
	Content string
}

func convertMessageToPayload(message []Message) []map[string]string {
	payload := []map[string]string{}
	for _, m := range message {
		payload = append(payload, map[string]string{
			"role":    m.Role.name,
			"content": m.Content,
		})
	}
	return payload
}

// Send sends a list of messages to the OpenAI API and returns the response
func (api *API) Send(ctx context.Context, cfg Config, messages []Message) (Respoinse, error) {
	payload := map[string]interface{}{
		"model":             cfg.model,
		"temperature":       cfg.temperature,
		"max_tokens":        cfg.maxToken,
		"top_p":             cfg.topP,
		"frequency_penalty": cfg.frequencyPenalty,
		"presence_penalty":  cfg.presencePenalty,
		"messages":          convertMessageToPayload(messages),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Respoinse{}, fmt.Errorf("json.Marshal[%v]: %w", payload, err)
	}

	req, err := http.NewRequest(http.MethodPost, cfg.url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return Respoinse{}, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+api.apiKey)

	res, err := api.httpClient.Do(req)
	if err != nil {
		return Respoinse{}, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer res.Body.Close()

	response := Respoinse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return Respoinse{}, fmt.Errorf("json decode: %w", err)
	}

	return response, nil
}

type Respoinse struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"message"`
}
