package core

import (
	"github.com/atotto/clipboard"

	"github.com/go-flexi/codegenerator/generator/backend"
	"github.com/go-flexi/codegenerator/ui"
	"github.com/rivo/tview"
)

// Core is the core of the application.
type Core struct {
	app           *tview.Application
	model         *ui.MultiLineEditor
	userText      *ui.MultiLineEditor
	generatedCode *ui.MultiLineEditor
	generator     *backend.Generator
}

// NewCore creates a new Core.
func NewCore(generator *backend.Generator) *Core {
	c := Core{}
	c.generator = generator
	c.app = tview.NewApplication()
	c.model = ui.NewMultiLineEditor(c.app, "Write Model", c.hanldeModleEvent)
	c.userText = ui.NewMultiLineEditor(c.app, "Add Text to Modify Response", c.handleUserTextEvent)
	c.generatedCode = ui.NewMultiLineEditor(c.app, "Generated Code", c.handleGeneratedCodeEvent)

	return &c
}

// View shows the application.
func (c *Core) View() {
	flex := tview.NewFlex().
		AddItem(c.model.View(), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(c.userText.View(), 0, 1, false).
			AddItem(c.generatedCode.View(), 0, 3, false), 0, 2, false)
	if err := c.app.SetRoot(flex, true).SetFocus(c.model.View()).Run(); err != nil {
		panic(err)
	}
}

func (c *Core) hanldeModleEvent(e ui.Event, content string) {
	if e != ui.SubmitEvent {
		return
	}

	c.app.SetFocus(c.userText.View())

	code, err := c.generator.FirstCall(content)
	if err != nil {
		c.generatedCode.Reset(err.Error())
		return
	}

	c.generatedCode.Reset(code)
}

func (c *Core) handleUserTextEvent(e ui.Event, content string) {
	switch e {
	case ui.SubmitEvent:
		c.userText.Clear()

		code, err := c.generator.UserMessage(content)
		if err != nil {
			c.generatedCode.Reset(err.Error())
			return
		}

		c.generatedCode.Reset(code)
	}
	if e == ui.SubmitEvent {

		return
	}

	if e == ui.NextEvent {
		c.app.SetFocus(c.generatedCode.View())
	}
}

func (c *Core) handleGeneratedCodeEvent(e ui.Event, content string) {
	switch e {
	case ui.NextEvent:
		c.app.SetFocus(c.userText.View())
	case ui.CopyEvent:
		clipboard.WriteAll(content)
	}
}
