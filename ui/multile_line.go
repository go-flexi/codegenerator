package ui

import (
	"bytes"
	"strings"
)

// MultiLine is a multi line of text
type MultiLine struct {
	SingleLines []SingleLine
	Cursor      int
}

// NewMultiLine creates a new MultiLine
func NewMultiLine() MultiLine {
	return MultiLine{
		SingleLines: []SingleLine{NewSingleLine()},
		Cursor:      0,
	}
}

func (m *MultiLine) Clear() {
	m.SingleLines = []SingleLine{NewSingleLine()}
	m.Cursor = 0
}

func (m *MultiLine) Reset(content string) {
	splitedContents := strings.Split(content, "\n")
	m.SingleLines = []SingleLine{}
	for i := range splitedContents {
		m.Cursor = i
		singleLine := NewSingleLineWithContent(splitedContents[i])
		singleLine.MoveCursorToEnd()
		m.SingleLines = append(m.SingleLines, singleLine)
	}

	m.MoveCursorToEnd()
}

// Add adds a rune to the MultiLine at the cursor position
func (m *MultiLine) Add(r rune) {
	m.SingleLines[m.Cursor].Add(r)
}

// Remove removes a rune from the MultiLine at the cursor position
func (m *MultiLine) Remove() {
	if m.SingleLines[m.Cursor].Remove() {
		return
	}

	if m.Cursor == 0 {
		return
	}

	m.SingleLines[m.Cursor-1].Merge(m.SingleLines[m.Cursor])

	tmpSingleLines := []SingleLine{}
	for i := 0; i < len(m.SingleLines); i++ {
		if i == m.Cursor {
			continue
		}
		tmpSingleLines = append(tmpSingleLines, m.SingleLines[i])
	}

	m.SingleLines = tmpSingleLines
	m.Cursor--
}

// MoveCursorLeft moves the cursor to the left
func (m *MultiLine) MoveCursorLeft() {
	if m.SingleLines[m.Cursor].MoveCursorLeft() {
		return
	}

	if m.Cursor == 0 {
		return
	}

	m.Cursor--
	m.SingleLines[m.Cursor].MoveCursorToEnd()
}

// MoveCursorRight moves the cursor to the right
func (m *MultiLine) MoveCursorRight() {
	if m.SingleLines[m.Cursor].MoveCursorRight() {
		return
	}

	if m.Cursor == len(m.SingleLines)-1 {
		return
	}

	m.Cursor++
	m.SingleLines[m.Cursor].MoveCursorToStart()
}

// MoveCursorUp moves the cursor up
func (m *MultiLine) MoveCursorUp() {
	curCursor := m.Cursor
	nextCursor := m.Cursor - 1

	if nextCursor < 0 {
		return
	}

	m.SingleLines[nextCursor].MoveCursor(m.SingleLines[curCursor].Cursor)
	m.Cursor = nextCursor
}

// MoveCursorDown moves the cursor down
func (m *MultiLine) MoveCursorDown() {
	curCursor := m.Cursor
	nextCursor := m.Cursor + 1

	if nextCursor >= len(m.SingleLines) {
		return
	}

	m.SingleLines[nextCursor].MoveCursor(m.SingleLines[curCursor].Cursor)
	m.Cursor = nextCursor
}

// MoveCursorToEnd moves the cursor to the end of the MultiLine
func (m *MultiLine) MoveCursorToEnd() {
	m.Cursor = len(m.SingleLines) - 1
	m.SingleLines[m.Cursor].MoveCursorToEnd()
}

// ContentWithCursor returns the content of the MultiLine with the cursor
func (m *MultiLine) ContentWithCursor() string {
	buf := bytes.Buffer{}

	for i := range m.SingleLines {
		if i == m.Cursor {
			buf.WriteString(m.SingleLines[i].ContentWithCursor())
		} else {
			buf.WriteString(m.SingleLines[i].Content)
		}
		buf.WriteString("\n")
	}

	return buf.String()
}

func (m *MultiLine) content() string {
	buf := bytes.Buffer{}
	for i := range m.SingleLines {
		if i > 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(m.SingleLines[i].Content)

	}
	return buf.String()
}

// Split splits the MultiLine at the cursor position
func (m *MultiLine) Split() {
	nextSingleLine := m.SingleLines[m.Cursor].Split()

	tmpSingleLines := []SingleLine{}

	for i := 0; i <= m.Cursor; i++ {
		tmpSingleLines = append(tmpSingleLines, m.SingleLines[i])
	}
	tmpSingleLines = append(tmpSingleLines, nextSingleLine)

	for i := m.Cursor + 1; i < len(m.SingleLines); i++ {
		tmpSingleLines = append(tmpSingleLines, m.SingleLines[i])
	}

	m.SingleLines = tmpSingleLines
	m.Cursor++
}
