package main

import (
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type settings struct {
	mode   string
	time   int
	count  int
	active int
	sets   []*setting
}

type setting struct {
	title    string
	position int
	options  []string
}

func (s *settings) initSettings() {
	s.mode = "countdown"
	s.count = 30
	s.active = 0
	s.sets = []*setting{
		{title: "Game Mode", position: 0, options: []string{"Time Limit", "Word Limit"}},
		{title: "Time Limit", position: 1, options: []string{"15", "30", "60", "90", "120"}},
		{title: "Word Limit", position: 1, options: []string{"15", "30", "50", "60", "100"}},
	}
}

func (s *settings) viewSettings() string {
	height := 6
	fullContent := []string{}
	for pos, set := range s.sets {
		tempContent := set.title
		for i := 0; i < height; i++ {
			if i < len(set.options){
				if set.position == i {
					tempContent += "\n" + set.options[i] + " (X)"
				} else {
					tempContent += "\n" + set.options[i]  + " ( )"
				}
			} else {
				tempContent += "\n"
			}
			
		}
		if s.active == pos {
			tempContent = borderStyleActive.Render(tempContent)
		} else {
			tempContent = borderStyleDefault.Render(tempContent)
		}
		fullContent = append(fullContent, tempContent)
	}
	// make a 2d slice for each block and each row in each block
	twoDimensionContent := [][]string{}
	for _, item := range fullContent {
		tempContent := strings.Split(item, "\n")
		twoDimensionContent = append(twoDimensionContent, tempContent)
	}

	// iterate over each line of each settings tab and add them to a string line by line
	finalString := ""
	for i := 0; i < len(twoDimensionContent[0]); i++ {
		for _, block := range twoDimensionContent {
			finalString += block[i] 
			finalString += " "
		}
		finalString += "\n"
	}
	return finalString
}

func runSettingsUpdate(s *settings, char string) tea.Cmd {
	return func() tea.Msg {
		s.updateSettings(char)
		return nil
	}
}

func (s *settings) updateSettings(key string) {
	switch key {
	case "tab":
		s.active += 1
		if s.active == len(s.sets) {
			s.active = 0
		}
	case "down":
		// get the setting tab
		setting := s.sets[s.active]
		setting.position += 1
		if setting.position == len(setting.options) {
			setting.position = 0
		}
	}
}

func (m *model) updateSettings(key string) *typing {
	s := m.settingsTab
	switch key {
	case "tab":
		s.active += 1
		if s.active == len(s.sets) {
			s.active = 0
		}
		// now change the actual typing settings
		return m.typingTab
	case "down":
		// get the setting tab
		setting := s.sets[s.active]
		setting.position += 1
		if setting.position == len(setting.options) {
			setting.position = 0
		}

		m.updateSettingsValues()
		var gc int
		if m.typingTab.gameMode == "countdown" {
			gc = m.settingsTab.time
		} else {
			gc = m.settingsTab.count
		}

		newTypingTab := &typing{gameMode: m.settingsTab.mode, gameCount: gc}
		newTypingTab.initTyping()
		return newTypingTab
	}
	return m.typingTab
}

func (m *model) updateSettingsValues() {
	// get the current setting
	set := m.settingsTab.sets[m.settingsTab.active]
	switch set.title {
	case "Game Mode":
		if m.typingTab.gameMode == "countdown" {
			m.settingsTab.mode = "words"
		} else {
			m.settingsTab.mode = "countdown"
		}
	case "Time Limit":
		m.settingsTab.time, _ = strconv.Atoi(set.options[set.position])
	case "Word Limit":
		m.settingsTab.count, _ = strconv.Atoi(set.options[set.position])
	}
}
