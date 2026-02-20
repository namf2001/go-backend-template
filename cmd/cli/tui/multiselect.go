package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// MultiSelectModel is a Bubble Tea model for multi-selection
type MultiSelectModel struct {
	header   string
	options  []string
	descs    map[string]string
	cursor   int
	selected map[string]bool
	result   map[string]bool
	quitting bool
}

// NewMultiSelect creates a new multi-select model
func NewMultiSelect(header string, options []string, descs map[string]string, result map[string]bool) MultiSelectModel {
	selected := make(map[string]bool)
	// Pre-select based on existing result
	for k, v := range result {
		selected[k] = v
	}
	return MultiSelectModel{
		header:   header,
		options:  options,
		descs:    descs,
		selected: selected,
		result:   result,
	}
}

func (m MultiSelectModel) Init() tea.Cmd {
	return nil
}

func (m MultiSelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		case " ":
			opt := m.options[m.cursor]
			m.selected[opt] = !m.selected[opt]
		case "a":
			// Toggle all
			allSelected := true
			for _, opt := range m.options {
				if !m.selected[opt] {
					allSelected = false
					break
				}
			}
			for _, opt := range m.options {
				m.selected[opt] = !allSelected
			}
		case "enter":
			for k, v := range m.selected {
				m.result[k] = v
			}
			m.quitting = true
			return m, tea.Quit
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MultiSelectModel) View() string {
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
		}

		checked := "[ ]"
		if m.selected[option] {
			checked = SuccessStyle.Render("[✓]")
		}

		desc := ""
		if m.descs != nil {
			if d, ok := m.descs[option]; ok {
				desc = DimStyle.Render(fmt.Sprintf(" - %s", d))
			}
		}

		if m.cursor == i {
			b.WriteString(fmt.Sprintf("%s%s %s%s\n", cursor, checked, HighlightStyle.Render(option), desc))
		} else {
			b.WriteString(fmt.Sprintf("%s%s %s%s\n", cursor, checked, option, desc))
		}
	}

	b.WriteString("\n")
	b.WriteString(DimStyle.Render("(↑/↓ move, space toggle, 'a' toggle all, enter confirm, q quit)"))

	return b.String()
}
