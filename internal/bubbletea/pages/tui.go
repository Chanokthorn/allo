package pages

import (
	"fmt"
	"os"

	"allo/internal/bubbletea/components"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	// Define your application state here
	pipeline        []string
	banner          components.BannerModel
	pipelineBuilder components.PipelineBuilderModel
	alloSelector    components.AlloSelectorModel
}

func RunTUI() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	mainModel := model{}

	mainModel.alloSelector = components.InitialAlloSelectorModel(
		[]string{"YYYY-MM-DD", "MM-DD", "DD", "file type"},
	)

	p := tea.NewProgram(mainModel)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting program:", err)
	}
}

func (m model) Init() tea.Cmd {
	// Perform any initial setup here
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle key message
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			var cmd tea.Cmd
			m.alloSelector, cmd = m.alloSelector.Update(msg)
			return m, cmd
		}
	case components.AlloSelectedMsg:
		m.pipeline = append(m.pipeline, msg.Allo)

		var cmd tea.Cmd
		m.pipelineBuilder, cmd = m.pipelineBuilder.Update(components.PipelineBuilderMsg{Pipeline: m.pipeline})
		return m, cmd
	}

	return m, nil
}

func (m model) View() string {
	// Render your application view here
	s := m.banner.View()
	s += m.pipelineBuilder.View()
	s += m.alloSelector.View()
	return s
}
