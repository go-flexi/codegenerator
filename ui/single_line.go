package ui

// SingleLine is a single line of text
type SingleLine struct {
	Content string
	Cursor  int
}

// NewSingleLine creates a new SingleLine
func NewSingleLine() SingleLine {
	return SingleLine{
		Content: "",
		Cursor:  0,
	}
}

// NewSingleLineWithContent creates a new SingleLine with content
func NewSingleLineWithContent(content string) SingleLine {
	return SingleLine{
		Content: content,
		Cursor:  0,
	}
}

// Add adds a rune to the SingleLine at the cursor position
func (s *SingleLine) Add(r rune) {
	s.Content = s.Content[:s.Cursor] + string(r) + s.Content[s.Cursor:]
	s.Cursor++
}

// Remove removes a rune from the SingleLine at the cursor position
func (s *SingleLine) Remove() (end bool) {
	if s.Cursor == 0 {
		return false
	}

	tmContent := s.Content[:s.Cursor-1]
	if s.Cursor < len(s.Content) {
		tmContent += s.Content[s.Cursor:]
	}
	s.Content = tmContent
	s.Cursor--

	return true
}

// MoveCursorToEnd moves the cursor to the end of the SingleLine
func (s *SingleLine) MoveCursorToEnd() {
	s.Cursor = len(s.Content)
}

// MoveCursorToStart moves the cursor to the start of the SingleLine
func (s *SingleLine) MoveCursorToStart() {
	s.Cursor = 0
}

// MoveCursorLeft moves the cursor to the left
func (s *SingleLine) MoveCursorLeft() (end bool) {
	if s.Cursor > 0 {
		s.Cursor--
		return true
	}
	return false
}

// MoveCursorRight moves the cursor to the right
func (s *SingleLine) MoveCursorRight() (end bool) {
	if s.Cursor < len(s.Content) {
		s.Cursor++
		return true
	}
	return false
}

// MoveCursor moves the cursor to a specific position
func (s *SingleLine) MoveCursor(p int) {
	if p >= 0 && p <= len(s.Content) {
		s.Cursor = p
	} else {
		s.Cursor = len(s.Content)
	}
}

// Split splits the SingleLine at the cursor position
func (s *SingleLine) Split() SingleLine {
	if s.Cursor == 0 {
		nextSingleLine := NewSingleLineWithContent(s.Content)
		s.Content = ""
		s.MoveCursorToStart()

		return nextSingleLine
	}

	if s.Cursor == len(s.Content) {
		s.MoveCursorToEnd()
		return NewSingleLineWithContent("")
	}

	nextSingleLine := NewSingleLineWithContent(s.Content[s.Cursor:])
	s.Content = s.Content[:s.Cursor]
	s.MoveCursorToEnd()
	return nextSingleLine
}

// Merge merges a SingleLine into another SingleLine
func (s *SingleLine) Merge(line SingleLine) {
	s.Cursor = len(s.Content)
	s.Content += line.Content
}

// ContentWithCursor returns the content with the cursor
func (s *SingleLine) ContentWithCursor() string {
	return s.Content[:s.Cursor] + "|" + s.Content[s.Cursor:]
}
