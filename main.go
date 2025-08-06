package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tab int

const (
	tabTyping tab = iota
	tabStats
	tabSettings
	tabHelp
	minHeight = 17
)

var (
	borderStyleActive lipgloss.Style = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(1, 2)

	borderStyleDefault lipgloss.Style = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("1")).
				Padding(1, 2)
)

var tabNames = []string{"Typing", "Stats", "Settings", "Help"}

type model struct {
	currentTab  tab
	width       int
	height      int
	textarea    textarea.Model
	typingTab   *typing
	settingsTab *settings
	allStyles   styles
}

type styles struct {
	centreStyle lipgloss.Style
	borderStyle lipgloss.Style
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Change settings here..."
	ta.SetWidth(40)
	ta.SetHeight(8)
	ta.Blur()

	m := model{
		currentTab:  tabTyping,
		textarea:    ta,
		typingTab:   &typing{gameMode: "countdown", gameCount: 30},
		allStyles:   styles{borderStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62")).Padding(1, 2)},
		settingsTab: &settings{mode: "countdown", count: 30, time: 30},
	}

	m.typingTab.initTyping()
	m.settingsTab.initSettings()

	return m
}

func (m model) Init() tea.Cmd {
	tick()
	return tea.EnterAltScreen
}

func (m *model) updateFromSettings() tea.Cmd {
	m.typingTab.gameMode = m.settingsTab.mode
	m.typingTab.gameCount = m.settingsTab.count
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// non tab specific commands
		case "ctrl+c":
			return m, tea.Quit
		case "left":
			if m.currentTab > 0 {
				m.currentTab--
			}
			return m, cmd
		case "right":
			if m.currentTab < tab(len(tabNames)-1) {
				m.currentTab++
			}
			return m, cmd
		case "ctrl+r":
			// make a new typing Tab
			// check gamemode
			var newTypingTab *typing
			if m.settingsTab.mode == "countdown" {
				newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.time}
			} else {
				newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.count}
			}
			m.typingTab = newTypingTab
			m.typingTab.initTyping()
			return m, nil
		case "enter":
			// make a new typing Tab
			// check gamemode
			var newTypingTab *typing
			if m.settingsTab.mode == "countdown" {
				newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.time}
			} else {
				newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.count}
			}
			m.typingTab = newTypingTab
			m.typingTab.initTyping()
			m.currentTab = tabTyping
			return m, nil
		default:
			switch m.currentTab {
			case tabTyping:
				return m, runTypingUpdate(m.typingTab, msg.String())
			case tabSettings:
				m.typingTab = m.updateSettings(msg.String())
				return m, nil
			}

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.allStyles.centreStyle = lipgloss.NewStyle().Width(m.width - 6).Height(m.height - 4).Align(lipgloss.Center)
		return m, nil
	}
	return m, tick()

}

func (m model) View() string {
	if m.height > 0 && m.height < minHeight {
		return "Window too small please resize"
	}
	header := renderTabs(m.currentTab)
	body := renderTabContent(m)
	rows := len(strings.Split(body, "\n"))
	if m.height > 0 {
		padding := (m.height-rows)/2 - 2
		body = strings.Repeat("\n", padding) + body
	}
	content := fmt.Sprintf("%s\n\n%s", header, body)

	if m.height > 0 && m.width > 0 {
		return m.allStyles.borderStyle.Render((m.allStyles.centreStyle.Render(content)))
	}

	return content
}

func renderTabs(current tab) string {
	var out string
	for i, name := range tabNames {
		style := lipgloss.NewStyle().Padding(0, 1)
		if current == tab(i) {
			style = style.Bold(true).Underline(true)
		}
		out += style.Render(name) + " "
	}
	return out
}

func renderTabContent(m model) string {
	switch m.currentTab {
	case tabTyping:
		return m.typingTab.viewTypingTab()
	case tabStats:
		return "Stats will be displayed here."
	case tabSettings:
		return m.settingsTab.viewSettings()
	case tabHelp:
		return "Use ← → to change tabs. Press q to quit. CTRL R to restart test."
	default:
		return "Unknown tab."
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func main() {
	go func() {
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Fatalf("error %v", err)
		}
	}()
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
