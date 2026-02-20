package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// DoneMsg signals the spinner to stop
type DoneMsg struct{ Err error }

// SpinnerModel is a Bubble Tea model for a loading spinner
type SpinnerModel struct {
	spinner  spinner.Model
	message  string
	quitting bool
	err      error
}

// NewSpinner creates a new spinner model
func NewSpinner(message string) SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6"))

	return SpinnerModel{
		spinner: s,
		message: message,
	}
}

func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
	case DoneMsg:
		m.err = msg.Err
		m.quitting = true
		return m, tea.Quit
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m SpinnerModel) View() string {
	if m.quitting {
		if m.err != nil {
			return ErrorStyle.Render(fmt.Sprintf("✗ %s\n", m.err.Error()))
		}
		return ""
	}
	return fmt.Sprintf("\n %s %s\n", m.spinner.View(), m.message)
}
