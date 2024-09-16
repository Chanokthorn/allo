package components

import (
	"fmt"

	"allo/internal/style"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AlloSelectorModel struct {
	choices []string
	cursor  int
}

type AlloSelectedMsg struct {
	Allo string
}

func InitialAlloSelectorModel(choices []string) AlloSelectorModel {
	return AlloSelectorModel{
		choices: choices,
		cursor:  0,
	}
}

func (m AlloSelectorModel) Update(msg tea.Msg) (AlloSelectorModel, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			msg := AlloSelectedMsg{Allo: m.choices[m.cursor]}
			m.cursor = 0
			return m, func() tea.Msg { return msg }
		}
		// default:
		// 	var cmd tea.Cmd
		// 	m.spinner, cmd = m.spinner.Update(msg)
		// 	return m, cmd
	}
	return m, nil
}

// View implements tea.Model.
func (m AlloSelectorModel) View() string {
	// The header
	s := fmt.Sprint(`
select allocation steps:
`)

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = "->" // cursor!
		}

		// Render the row
		selectedStyle := lipgloss.NewStyle().
			Background(style.PrimaryColor).
			Foreground(style.SecondaryColor).
			PaddingLeft(1).
			Width(50)

		unselectedStyle := lipgloss.NewStyle().
			PaddingLeft(1).
			Width(50)

		var rowStyle lipgloss.Style
		if m.cursor == i {
			rowStyle = selectedStyle
		} else {
			rowStyle = unselectedStyle
		}

		s += fmt.Sprintf("%s %s\n",
			lipgloss.NewStyle().Bold(true).Foreground(style.SpecialFontColor).Render(cursor),
			rowStyle.Render(choice),
		)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
