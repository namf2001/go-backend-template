package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TextInputModel is a Bubble Tea model for text input
type TextInputModel struct {
	textInput textinput.Model
	header    string
	result    *string
	quitting  bool
}

// NewTextInput creates a new text input model
func NewTextInput(header, placeholder string, result *string) TextInputModel {
	ti := textinput.New()
	ti.Placeholder = placeholder
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return TextInputModel{
		textInput: ti,
		header:    header,
		result:    result,
	}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			*m.result = m.textInput.Value()
			m.quitting = true
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	if m.quitting {
		return ""
	}
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		TitleStyle.Render(m.header),
		m.textInput.View(),
		DimStyle.Render("(press enter to confirm, esc to quit)"),
	)
}
