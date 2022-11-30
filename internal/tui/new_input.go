package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/cqroot/ternote/pkg/ternote"
)

type newInputModel struct {
	textinput textinput.Model

	WidgetMsg widgetMsg
}

func (m *newInputModel) initModel(width, height int) {
	m.textinput = textinput.New()
	m.textinput.Focus()
	m.textinput.CharLimit = 156
	m.textinput.Width = width
	m.textinput.Prompt = " Category: "
	m.textinput.PromptStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("10"))
}

func (m newInputModel) Init() tea.Cmd { return nil }

func (m newInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.WidgetMsg = emptyMsg

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			m.WidgetMsg = quitMsg
			return m, cmd
		case "enter":
			m.WidgetMsg = quitMsg
			ternote.NewNote(m.textinput.Value())
			return m, cmd
		}
	}

	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func (m newInputModel) View() string {
	return m.textinput.View()
}

func (m newInputModel) Focus() tea.Cmd {
	return m.textinput.Focus()
}
