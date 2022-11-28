package tui

type displayMode int

const (
	normalMode displayMode = iota
	newInputMode
)

type widgetMsg int

const (
	emptyMsg widgetMsg = iota
	quitMsg

	editMsg

	newInputMsg
)

type editorFinishedMsg struct{ err error }
