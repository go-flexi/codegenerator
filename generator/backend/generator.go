package backend

import (
	"context"
	"fmt"

	"github.com/go-flexi/codegenerator/generator"
	"github.com/go-flexi/codegenerator/openai"
)

// Generator generates backend code using OpenAI API.
type Generator struct {
	api         *openai.API
	orgName     string
	projectName string

	messages generator.Messages
}

// NewGenerator creates a new Generator.
func NewGenerator(api *openai.API, orgName, projectName string) *Generator {
	return &Generator{
		api:         api,
		orgName:     orgName,
		projectName: projectName,
		messages:    generator.NewMessages(system),
	}
}

// Generate generates backend code and this function needs to be called at the beginning.
func (g *Generator) FirstCall(modelStruct string) (string, error) {
	code, err := g.UserMessage(modelStruct + "\n" + "write the code for model.go")
	if err != nil {
		return "", fmt.Errorf("UserMessage: %w", err)
	}

	g.messages.AddAssistantMessage(code)
	return code, nil
}

// UserMessage is used to receive user messages and generate backend code accordingly.
func (g *Generator) UserMessage(message string) (string, error) {
	if err := g.generateWithUserMessage(message); err != nil {
		return "", fmt.Errorf("generateWithUserMessage: %w", err)
	}
	return g.messages.LastAsistantMessage(), nil
}

func (g *Generator) generateWithUserMessage(message string) error {
	g.messages.AddUserMessage(message)
	if err := g.openaiCall(); err != nil {
		return fmt.Errorf("openaiCall: %w", err)
	}
	return nil
}

func (g *Generator) openaiCall() error {
	response, err := g.api.Send(context.Background(), openai.DefaultConfig(), g.messages.GetMessages())
	if err != nil {
		return fmt.Errorf("api.Send: %w", err)
	}

	if len(response.Choices) > 0 {
		g.messages.AddAssistantMessage(response.Choices[0].Message.Content)
	}

	return nil
}
