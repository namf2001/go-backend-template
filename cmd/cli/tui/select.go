package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// SelectModel is a Bubble Tea model for single selection
type SelectModel struct {
	header   string
	options  []string
	descs    map[string]string
	cursor   int
	result   *string
	quitting bool
}

// NewSelect creates a new select model
func NewSelect(header string, options []string, descs map[string]string, result *string) SelectModel {
	return SelectModel{
		header:  header,
		options: options,
		descs:   descs,
		result:  result,
	}
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter":
			*m.result = m.options[m.cursor]
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m SelectModel) View() string {
	if m.quitting {
		return ""
	}

	var b strings.Builder
	b.WriteString(TitleStyle.Render(m.header))
	b.WriteString("\n\n")

	for i, option := range m.options {
		cursor := "  "
		if m.cursor == i {
			cursor = HighlightStyle.Render("▸ ")
			desc := ""
			if m.descs != nil {
				if d, ok := m.descs[option]; ok {
					desc = DimStyle.Render(fmt.Sprintf(" - %s", d))
				}
			}
			b.WriteString(fmt.Sprintf("%s%s%s\n", cursor, HighlightStyle.Render(option), desc))
		} else {
			desc := ""
			if m.descs != nil {
				if d, ok := m.descs[option]; ok {
					desc = DimStyle.Render(fmt.Sprintf(" - %s", d))
				}
			}
			b.WriteString(fmt.Sprintf("%s%s%s\n", cursor, option, desc))
		}
	}

	b.WriteString("\n")
	b.WriteString(DimStyle.Render("(↑/↓ to move, enter to select, q to quit)"))

	return b.String()
}
