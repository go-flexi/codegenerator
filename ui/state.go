package ui

// const (
// 	CoreView    View = "core"
// 	HandlerView View = "handler"
// )

// const (
// 	Model   SubView = "model"
// 	Command SubView = "command"
// )

// type View string

// type SubView string

// type State struct {
// 	View    View
// 	SubView SubView
// 	Model   string
// 	Command string
// 	Content string
// }

// func NewState() State {
// 	return State{
// 		View:    CoreView,
// 		SubView: Model,
// 	}
// }

// func (s *State) NextSubView() {
// 	switch s.SubView {
// 	case Model:
// 		s.SubView = Command
// 		s.Content = ""
// 	case Command:
// 		s.SubView = Model
// 		s.Content = ""
// 	}
// }
