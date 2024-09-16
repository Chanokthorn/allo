package style

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryColor     lipgloss.Color
	PrimaryFontColor lipgloss.Color
	SecondaryColor   lipgloss.Color
	SpecialFontColor lipgloss.Color
)

func init() {
	PrimaryColor = lipgloss.Color("#FFE0F8")
	PrimaryFontColor = lipgloss.Color("#381616")
	// Style.SecondaryColor = lipgloss.
	SpecialFontColor = lipgloss.Color("#F10065")
}
