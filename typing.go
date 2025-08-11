package main

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultKey        = "default"
	correctKey        = "correct"
	incorrectKey      = "incorrect"
	gameModeCountdown = "countdown"
	gameModeWords     = "words"
)

var (
	red         = lipgloss.NewStyle().Foreground(lipgloss.Color("1"))
	green       = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	blue        = lipgloss.NewStyle().Foreground(lipgloss.Color("4"))
	commonWords = []string{
		"the", "be", "and", "of", "a",
		"in", "to", "have", "it", "that",
		"for", "you", "he", "with", "on",
		"do", "at", "by", "not", "this",
		"but", "from", "or", "which", "one",
		"would", "all", "there", "say", "who",
		"like", "when", "make", "them", "know",
		"if", "time", "no", "take", "people",
		"out", "into", "just", "see", "him",
		"your", "come", "could", "now", "than",
		"other", "how", "then", "its", "our",
		"two", "more", "these", "no", "way",
		"well", "only", "my", "other", "could",
		"some", "than", "first", "last", "over",
		"such", "my", "made", "make", "after",
		"may", "much", "where", "need", "should",
		"why", "long", "know", "using", "good",
		"think", "see", "look", "want", "better",
		"here", "need", "much", "want", "more",
		"he's", "she's", "they're", "we're", "you're",
		"I", "we", "you", "he", "she",
	}
)

type typing struct {
	content          string
	position         int
	characterColours []string
	extraKeys        int    // this is how many extra key presses the user did after the end of a word - it resets every time they press space after finishing a word
	gameMode         string // either words or countdown
	gameCount        int    // this is either how many words to complete or how long you have to type as many words as possible depending on game
	time             timer
}

func runTypingUpdate(t *typing, char string) tea.Cmd {
	return func() tea.Msg {
		if !t.time.isFinished() {
			t.updateTypingTab(char)
		}
		return nil
	}
}

func (t *typing) initTyping() {
	// options are words (how long to do n words) or countdown (how many words in n time)
	switch t.gameMode {
	case gameModeWords:
		rand.Seed(time.Now().UnixNano())
		words := ""
		for i := 1; i < t.gameCount+1; i++ {
			words += commonWords[rand.Intn(len(commonWords))]
			if i%15 == 0 {
				words += "\n"
			} else {
				words += " "
			}

		}
		t.content = words

		t.time = &timerUp{started: false, finished: false}
	case gameModeCountdown:
		rand.Seed(time.Now().UnixNano())
		words := ""
		for i := 0; i < 300; i++ {
			words += commonWords[rand.Intn(len(commonWords))]
			words += " "
		}
		t.content = words
		// generate content
		t.time = &timerDown{started: false, finished: false, seconds: t.gameCount}
	}

	for i := 0; i < len(t.content); i++ {
		t.characterColours = append(t.characterColours, defaultKey)
	}
}

func (t *typing) updateTypingTab(key string) {
	switch key {
	case "backspace":
		if t.position > 0 {
			if t.extraKeys == 1 {
				// In this case the user finally removed all extra chars
				t.extraKeys = 0
				t.characterColours[t.position] = correctKey
				t.characterColours[t.position-1] = correctKey
			} else if t.extraKeys > 1 {
				t.extraKeys -= 1
				// mkae sure the previosu character is red
				t.characterColours[t.position] = incorrectKey
			} else {
				t.characterColours[t.position-1] = defaultKey
				t.position -= 1
			}
		}
	default:
		if t.position < len(t.content) {
			switch t.content[t.position] {
			case ' ':
				if key == " " {
					t.characterColours[t.position] = correctKey
					t.position += 1
					t.extraKeys = 0
				} else {
					// incorrect characters after word
					t.extraKeys += 1
					t.characterColours[t.position-1] = incorrectKey
				}
			default:
				if key == string(t.content[t.position]) {
					t.characterColours[t.position] = correctKey
				} else {
					t.characterColours[t.position] = incorrectKey
				}
				t.position += 1
				// make timer set to started as user must have pressed a key now
				t.time.startTimer()
			}
		}
	}
}

func (t typing) viewTypingTab(designStyles colourTheme) string {
	output := ""
	count := 0
	lineCount := 0
	for pos, val := range t.characterColours {
		if lineCount < 3 {
			switch val {
			case defaultKey:
				output += designStyles.typeTextDefault.Render(string(t.content[pos]))
			case correctKey:
				output += designStyles.typeTextCorrect.Render(string(t.content[pos]))
			case incorrectKey:
				output += designStyles.typeTextIncorrect.Render(string(t.content[pos]))
			}
			if t.content[pos] == ' ' && t.gameMode == gameModeCountdown {
				count += 1
				// check if the colour is default or not - if not default we don't want this line and remove it
				if count%15 == 0 {
					output = output[:len(output)-1]
					output += "\n"
					lineCount += 1
					if t.characterColours[pos] != defaultKey {
						output = ""
						lineCount = 0
					}
				}

			}
		}

	}

	output = output + "\n\n\n" + t.time.displayTimer(designStyles)
	return output
}

func (t *typing) roundFinished() bool {
	// check what type of timer and return true if the round has been finished
	switch v := t.time.(type) {
	case *timerDown:
		if v.started && (float64(v.seconds)-time.Since(v.start).Seconds()) <= 0 {
			return true
		}
	case *timerUp:
		if t.position == len(t.content)-1 {
			return true
		}
	}
	return false
}
