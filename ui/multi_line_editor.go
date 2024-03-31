package ui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Event represents a command event
type Event string

// list of events
const (
	SubmitEvent Event = ":submit"
	NextEvent   Event = ":next"
	CopyEvent   Event = ":copy"
)

type OnEvent func(e Event, content string)

type MultiLineEditor struct {
	app       *tview.Application
	textView  *tview.TextView
	multiLine MultiLine

	onEvent OnEvent
}

func NewMultiLineEditor(app *tview.Application, title string, onEvent OnEvent) *MultiLineEditor {
	editor := MultiLineEditor{
		app: app,
		textView: tview.NewTextView().
			SetDynamicColors(true).
			SetWrap(true).
			SetRegions(true).SetScrollable(true),
		multiLine: NewMultiLine(),
		onEvent:   onEvent,
	}

	editor.init(title)
	return &editor
}

func (mle *MultiLineEditor) Clear() {
	mle.multiLine.Clear()
	mle.textView.SetText(mle.multiLine.ContentWithCursor())
}

func (mle *MultiLineEditor) Reset(content string) {
	mle.multiLine.Reset(content)
	mle.textView.SetText(mle.multiLine.ContentWithCursor())
}

func (mle *MultiLineEditor) init(title string) {
	mle.textView.SetBorder(true).SetTitle(title)
	mle.textView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Key() {
		case tcell.KeyLeft:
			mle.multiLine.MoveCursorLeft()
		case tcell.KeyRight:
			mle.multiLine.MoveCursorRight()
		case tcell.KeyUp:
			mle.multiLine.MoveCursorUp()
		case tcell.KeyDown:
			mle.multiLine.MoveCursorDown()
		case tcell.KeyEnter:
			mle.multiLine.Split()
		case tcell.KeyTab:
			mle.multiLine.Add(' ')
			mle.multiLine.Add(' ')
		case tcell.KeyRune:
			mle.multiLine.Add(event.Rune())
		case tcell.KeyBackspace, tcell.KeyBackspace2:
			mle.multiLine.Remove()
		}

		mle.textView.SetText(mle.multiLine.ContentWithCursor())
		mle.textView.GetText(false)
		mle.handleEvent()

		return event
	})
}

func (mle *MultiLineEditor) handleEvent() {
	content := mle.multiLine.content()
	e, ok := findEvent(content)
	if !ok {
		return
	}

	content = removeEvent(content, e)
	mle.multiLine.Reset(content)
	mle.textView.SetText(mle.multiLine.ContentWithCursor())

	// time.Sleep(time.Second * 2)
	mle.onEvent(e, content)
}

func (mle *MultiLineEditor) View() *tview.TextView {
	return mle.textView
}

func (mle *MultiLineEditor) OnFocus() {
	mle.app.SetFocus(mle.textView)
}

func removeEvent(content string, e Event) string {
	return content[:len(content)-len(string(e))]
}

func findEvent(content string) (Event, bool) {
	if strings.HasSuffix(content, string(SubmitEvent)) {
		return SubmitEvent, true
	}
	if strings.HasSuffix(content, string(NextEvent)) {
		return NextEvent, true
	}
	if strings.HasSuffix(content, string(CopyEvent)) {
		return CopyEvent, true
	}
	return "", false
}
