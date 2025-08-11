package main 

import (
	"github.com/charmbracelet/lipgloss"
)

type colourTheme struct {
	borderStyleDefault lipgloss.Style
	borderStyleActive lipgloss.Style
	tabTextDefault lipgloss.Style
	tabTextActive lipgloss.Style
	typeTextIncorrect lipgloss.Style
	typeTextCorrect lipgloss.Style
	typeTextDefault lipgloss.Style
	normalText lipgloss.Style
	countDownBar lipgloss.Style
}

func NewColourTheme(profile int) (colourTheme, error) {
	ct := colourTheme{}
	if profile == 0 {
		// default theme
		ct.loadDefaultColourTheme()
		return ct, nil
	}
	return ct, nil
}


func (ct *colourTheme) loadDefaultColourTheme() {
	ct.borderStyleActive = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("62")).
		Padding(1, 2)

	ct.borderStyleDefault = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#B0C4DE")). 
		Padding(1, 2)

	ct.tabTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#5F9EA0")).
		Padding(0, 1)

	ct.tabTextActive = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4682B4")). 
		Bold(true).
		Underline(true).
		Padding(0, 1)

	ct.typeTextIncorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6F61")). 
		Bold(true).
		Underline(true)

	ct.typeTextCorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#65F527")). 
		Bold(true)

	ct.typeTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAE5EB")) 

	ct.normalText = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DAE5EB")) 

	ct.countDownBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4682B4")) 
}