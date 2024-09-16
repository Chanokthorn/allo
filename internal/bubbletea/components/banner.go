package components

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type BannerModel struct {
	// Define your application state here
	showHelp bool
}

type BannerMsg struct {
	showHelp bool
}

func (m BannerModel) Update(msg BannerMsg) (BannerModel, tea.Cmd) {
	// Perform any initial setup here
	return m, nil
}

func (m BannerModel) View() string {
	s := fmt.Sprint(`
====================================
 _______  ___      ___      _______ 
|   _   ||   |    |   |    |       |
|  |_|  ||   |    |   |    |   _   |
|       ||   |    |   |    |  | |  |
|       ||   |___ |   |___ |  |_|  |
|   _   ||       ||       ||       |
|__| |__||_______||_______||_______|
---- an image allocator service ----
====================================
	`)

	if m.showHelp {
		s += `type "allo -h" for help`
	}
	return s
}
