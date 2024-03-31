package generator

import "github.com/go-flexi/codegenerator/openai"

// Messages is a struct to manage messages.
type Messages struct {
	messages []openai.Message
}

// NewMessages creates a new Messages with a system message.
func NewMessages(system string) Messages {
	return Messages{
		messages: []openai.Message{
			{Role: openai.SystemRole(), Content: system},
		},
	}
}

// AddUserMessage adds a user message to messages.
func (m *Messages) AddUserMessage(message string) {
	m.messages = append(
		m.messages,
		openai.Message{Role: openai.UserRole(), Content: message},
	)
}

// AddAssistantMessage adds an assistant message to messages.
func (m *Messages) AddAssistantMessage(message string) {
	m.messages = append(
		m.messages,
		openai.Message{Role: openai.AssistantRole(), Content: message},
	)
}

// GetMessages returns messages.
func (m *Messages) GetMessages() []openai.Message {
	return m.messages
}

// LastAsistantMessage returns the last assistant message.
func (m *Messages) LastAsistantMessage() string {
	for i := len(m.messages) - 1; i >= 0; i-- {
		if m.messages[i].Role == openai.AssistantRole() {
			return m.messages[i].Content
		}
	}
	return ""
}
