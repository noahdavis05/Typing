package main

import (
	"strconv"

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
		{title: "Time Limit", position: 0, options: []string{"15", "30", "60", "90", "120"}},
		{title: "Word Limit", position: 0, options: []string{"15", "30", "50", "60", "100"}},
	}
}

func (s *settings) viewSettings() string {
	fullContent := ""
	for pos, set := range s.sets {
		tempContent := set.title
		for i, content := range set.options {
			if set.position == i {
				tempContent += "\n" + content + " (X)"
			} else {
				tempContent += "\n" + content + " ( )"
			}
		}
		if s.active == pos {
			fullContent += borderStyleActive.Render(tempContent) + "\n"
		} else {
			fullContent += borderStyleDefault.Render(tempContent) + "\n"
		}

	}
	return fullContent
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
