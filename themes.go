package main

import (
	"github.com/charmbracelet/lipgloss"
)

type colourTheme struct {
	borderStyleDefault lipgloss.Style
	borderStyleActive  lipgloss.Style
	tabTextDefault     lipgloss.Style
	tabTextActive      lipgloss.Style
	typeTextIncorrect  lipgloss.Style
	typeTextCorrect    lipgloss.Style
	typeTextDefault    lipgloss.Style
	normalText         lipgloss.Style
	countDownBar       lipgloss.Style
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
		BorderForeground(lipgloss.Color("#7F5AF0")).
		Padding(1, 2)

	ct.borderStyleDefault = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#94A1B2")).
		Padding(1, 2)

	ct.tabTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#72757E")).
		Padding(0, 2)

	ct.tabTextActive = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7F5AF0")).
		Bold(true).
		Underline(true).
		Padding(0, 2)

	ct.typeTextIncorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF6F61")).
		Bold(true).
		Underline(true)

	ct.typeTextCorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#2CB67D")).
		Bold(true)

	ct.typeTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#94A1B2"))

	ct.normalText = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#94A1B2"))

	ct.countDownBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7F5AF0"))
}
