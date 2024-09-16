package components

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type PipelineBuilderModel struct {
	// Define your application state here
	pipelines []string
}

type PipelineBuilderMsg struct {
	Pipeline []string
}

func (m PipelineBuilderModel) Update(msg tea.Msg) (PipelineBuilderModel, tea.Cmd) {
	// Perform any initial setup here
	switch msg := msg.(type) {
	case PipelineBuilderMsg:
		m.pipelines = msg.Pipeline
		log.Println("pipeline updated", len(m.pipelines), strings.Join(m.pipelines, ","))
	}

	return m, nil
}

func (m PipelineBuilderModel) View() string {
	s := fmt.Sprintf(`
current pipeline:
%s
	`, strings.Join(m.pipelines, " -> "))
	return s
}
