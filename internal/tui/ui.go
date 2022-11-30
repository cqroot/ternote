package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("8"))
var focusStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("7"))

type model struct {
	tableModel    tableModel
	newInputModel newInputModel

	err      error
	mode     displayMode
	quitting bool
}

func (m *model) resize() {
	screenWidth, screenHeight, _ := term.GetSize(int(os.Stdout.Fd()))
	tableWidth := screenWidth - 1
	tableHeight := screenHeight - 5

	switch m.mode {
	case normalMode:
		m.tableModel.initModel(tableWidth, tableHeight)
	case newInputMode:
		m.tableModel.initModel(tableWidth, tableHeight-3)
		m.newInputModel.initModel(screenWidth-5, 0)
	}
}

func (m model) Init() tea.Cmd {
	m.mode = normalMode
	m.quitting = false

	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.mode {
	case normalMode:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.resize()

		case editorFinishedMsg:
			m.quitting = false
			if msg.err != nil {
				m.err = msg.err
				return m, tea.Quit
			}

			m.resize()
			return m, nil

		default:
			var mm tea.Model
			mm, cmd = m.tableModel.Update(msg)
			m.tableModel, _ = mm.(tableModel)

			switch m.tableModel.WidgetMsg {
			case quitMsg:
				m.quitting = true

			case editMsg:
				m.quitting = true

			case newInputMsg:
				m.mode = newInputMode
				m.resize()
			}
		}

	case newInputMode:
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			m.resize()

		default:
			var mm tea.Model
			mm, cmd = m.newInputModel.Update(msg)
			m.newInputModel, _ = mm.(newInputModel)

			switch m.newInputModel.WidgetMsg {
			case quitMsg:
				m.mode = normalMode
				m.resize()
			}
		}
	}

	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error() + "\n"
	}

	if !m.quitting {
		switch m.mode {
		case normalMode:
			return focusStyle.Render(m.tableModel.View())
		case newInputMode:
			return lipgloss.JoinVertical(
				lipgloss.Top,
				baseStyle.Render(m.tableModel.View()),
				focusStyle.Render(m.newInputModel.View()),
			)
		}
	}

	return ""
}

func Run() error {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	_, err := p.Run()
	if err != nil {
		return err
	}

	return nil
}
