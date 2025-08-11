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
	tabSettings
	tabHelp
	minHeight = 17
)

// generic styles
var (
	borderStyleActive lipgloss.Style = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("62")).
				Padding(1, 2)

	borderStyleDefault lipgloss.Style = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#616060")).
				Padding(1, 2)
)

var tabNames = []string{"Typing", "Settings", "Help"}

// bubbletea model struct - contains the sub structs for given tabs
type model struct {
	currentTab  tab
	width       int
	height      int
	typingTab   *typing
	settingsTab *settings
	centreStyle lipgloss.Style
	designStyles colourTheme
}



// initialise the initial model and its sub structs
func initialModel() model {
	ta := textarea.New()
	ta.SetWidth(40)
	ta.SetHeight(8)
	ta.Blur()

	m := model{
		currentTab:  tabHelp,
		typingTab:   &typing{gameMode: "countdown", gameCount: 30},
		settingsTab: &settings{mode: "countdown", count: 30, time: 30},
	}

	m.typingTab.initTyping()
	m.settingsTab.initSettings()
	m.designStyles, _ = NewColourTheme(0)

	return m
}

// Init the app and set it to full screen
func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

// main update function - updates model and calls functions on key presses
// within case tea.Msg all general keys (ones used everywhere) defined here
// all specific keys to certain tabs are checked in their own update functions
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	if m.typingTab.time.isActive() {
		cmd = tick()
	}

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
			m = m.startRound()
			return m, cmd
		case "enter":
			m = m.startRound()
			m.currentTab = tabTyping
			return m, cmd
		default:
			switch m.currentTab {
			case tabTyping:
				return m, tea.Batch(runTypingUpdate(m.typingTab, msg.String()), tick())
			case tabSettings:
				m.typingTab = m.updateSettings(msg.String())
				return m, cmd
			}

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.centreStyle = lipgloss.NewStyle().Width(m.width - 6).Height(m.height - 4).Align(lipgloss.Center)
		return m, cmd

	case tickMsg:
		if m.typingTab.roundFinished() {
			m.typingTab.time.stopTimer(m.typingTab)
		}
		return m, cmd
	}

	return m, cmd
}

// initialises new typing tab struct within model and returns it
func (m model) startRound() model {
	var newTypingTab *typing
	if m.settingsTab.mode == "countdown" {
		newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.time}
	} else {
		newTypingTab = &typing{gameMode: m.settingsTab.mode, gameCount: m.settingsTab.count}
	}
	m.typingTab = newTypingTab
	m.typingTab.initTyping()
	return m
}

// display the content
func (m model) View() string {
	if m.height > 0 && m.height < minHeight {
		return "Window too small please resize"
	}
	header := m.renderTabs()
	body := renderTabContent(m)
	rows := len(strings.Split(body, "\n"))
	if m.height > 0 {
		padding := (m.height-rows)/2 - 2
		body = strings.Repeat("\n", padding) + body
	}
	content := fmt.Sprintf("%s\n\n%s", header, body)

	if m.height > 0 && m.width > 0 {
		return m.designStyles.borderStyleDefault.Render((m.centreStyle.Render(content)))
	}

	return content
}

// function which returns the string to display the tabs at top of screen
func (m model) renderTabs() string {
	var out string
	for i, name := range tabNames {
		if m.currentTab == tab(i) {
			out += m.designStyles.tabTextActive.Render(name )
		} else {
			out += m.designStyles.tabTextDefault.Render(name )
		}
		
	}
	return out
}

// function which renders content from inside a tab
func renderTabContent(m model) string {
	switch m.currentTab {
	case tabTyping:
		return m.typingTab.viewTypingTab(m.designStyles)
	case tabSettings:
		return m.settingsTab.viewSettings(m.designStyles)
	case tabHelp:
		return m.displayHelp()
	default:
		return "Unknown tab."
	}
}

func (m model) displayHelp() string {
	res := ""
	res += m.designStyles.normalText.Render("← → to change tabs") + "\n\n"
	res += m.designStyles.normalText.Render("CTRL C to quit") + "\n\n"
	res += m.designStyles.normalText.Render("CTRL R restart test") + "\n\n"
	res += m.designStyles.normalText.Render("TAB toggle new setting") + "\n\n"
	res += m.designStyles.normalText.Render("↑ ↓ change current setting")
	return res
}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*10, func(t time.Time) tea.Msg {
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
