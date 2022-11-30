package tui

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/cqroot/ternote/pkg/config"
	"github.com/cqroot/ternote/pkg/types"
)

type tableModel struct {
	table table.Model

	WidgetMsg widgetMsg

	err      error
	selected int
}

func (m *tableModel) initModel(width, height int) {
	m.selected = m.table.Cursor()

	columns := []table.Column{
		{Title: "#", Width: 4},
		{Title: "ID", Width: 15},
		{Title: "Category", Width: 10},
		{Title: "Title", Width: width - 38},
	}
	var rows []table.Row

	for index, note := range tn.Notes() {
		err := tn.UpdateNoteMetadata(&note)
		if errors.Is(err, types.ErrorNoteFileNoteFound) {
			continue
		} else if err != nil {
			continue
		}

		rows = append(rows, table.Row{
			fmt.Sprintf("%d", index), note.Id, note.Category, note.Title,
		})
	}

	m.table = table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithWidth(width),
		table.WithHeight(height),
	)
	m.table.SetCursor(m.selected)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("15")).
		Bold(false)
	m.table.SetStyles(s)
}

func (m *tableModel) refreshModel() {
	m.initModel(m.table.Width(), m.table.Height())
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.WidgetMsg = emptyMsg

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			m.WidgetMsg = quitMsg
			return m, tea.Quit

		case "e", "l", "enter":
			m.WidgetMsg = editMsg
			return m, editNote(m.table.SelectedRow()[1])

		case "d":
			tn.RemoveNote(m.table.SelectedRow()[1])
			m.refreshModel()
			return m, cmd

		case "n":
			m.WidgetMsg = newInputMsg
			m.table, cmd = m.table.Update(msg)
			return m, cmd
		}
	}

	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error() + "\n"
	}

	return m.table.View()
}

func editNote(id string) tea.Cmd {
	notePath, err := config.NotePath(id)
	if err != nil {
		return tea.Quit
	}

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	edcmd := exec.Command(editor, notePath)

	return tea.ExecProcess(edcmd, func(err error) tea.Msg {
		return editorFinishedMsg{err}
	})
}
