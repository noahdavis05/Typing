package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

const (
	configDirName  = "typingTester"
	configFilename = "config.json"
)

type themeConfig struct {
	Name               string `json:"name"`
	BorderActiveColor  string `json:"border_active_color"`
	BorderDefaultColor string `json:"border_default_color"`
	TabActiveColor     string `json:"tab_active_color"`
	TabDefaultColor    string `json:"tab_default_color"`
	TextIncorrectColor string `json:"text_incorrect_color"`
	TextCorrectColor   string `json:"text_correct_color"`
	TextDefaultColor   string `json:"text_default_color"`
	NormalTextColor    string `json:"normal_text_color"`
	CountDownBarColor  string `json:"countdown_bar_color"`
}

func (m model) loadConfig() ([]colourTheme, error) {
	// get the config dir
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	// make the filepath
	appConfigDir := filepath.Join(configDir, configDirName, configFilename)

	// try and open the file
	file, err := os.Open(appConfigDir)
	if err != nil {
		if os.IsNotExist(err) {
			cTs, err := m.createConfig(configDir)
			if err != nil {
				return nil, err
			}
			return cTs, nil
		}
		return nil, err
	}
	defer file.Close()
	data, err := os.ReadFile(appConfigDir)
	if err != nil {
		return nil, err
	}

	var themeConfigs []themeConfig
	if err := json.Unmarshal(data, &themeConfigs); err != nil {
		return nil, err
	}
	var themes []colourTheme
	for _, tc := range themeConfigs {
		themes = append(themes, convertStructs(tc))
	}
	return themes, nil

}

func (m model) createConfig(configDir string) ([]colourTheme, error) {
	// create directory if doesn't exist
	appConfigDir := filepath.Join(configDir, configDirName)
	err := os.MkdirAll(appConfigDir, 0755)
	if err != nil {
		return nil, err
	}

	// now create the file
	configFilePath := filepath.Join(appConfigDir, configFilename)
	file, err := os.Create(configFilePath)
	if err != nil {
		return nil, err
	}
	ct := writeInitialConfig(file)
	return ct, nil
}

func writeInitialConfig(f *os.File) []colourTheme {
	// Create a slice of themeConfig with 5 default configs (customize as needed)
	defaultConfigs := []themeConfig{
		{
			Name:               "Theme1",
			BorderActiveColor:  "#7F5AF0",
			BorderDefaultColor: "#c0c8d2ff",
			TabDefaultColor:    "#c0c8d2ff",
			TabActiveColor:     "#7F5AF0",
			TextIncorrectColor: "#FF6F61",
			TextCorrectColor:   "#2CB67D",
			TextDefaultColor:   "#c0c8d2ff",
			NormalTextColor:    "#c0c8d2ff",
			CountDownBarColor:  "#7F5AF0",
		},
		{
			Name:               "Theme2",
			BorderActiveColor:  "#7F5AF0",
			BorderDefaultColor: "#c0c8d2ff",
			TabDefaultColor:    "#c0c8d2ff",
			TabActiveColor:     "#7F5AF0",
			TextIncorrectColor: "#FF6F61",
			TextCorrectColor:   "#2CB67D",
			TextDefaultColor:   "#c0c8d2ff",
			NormalTextColor:    "#c0c8d2ff",
			CountDownBarColor:  "#7F5AF0",
		},
		{
			Name:               "Theme3",
			BorderActiveColor:  "#7F5AF0",
			BorderDefaultColor: "#c0c8d2ff",
			TabDefaultColor:    "#c0c8d2ff",
			TabActiveColor:     "#7F5AF0",
			TextIncorrectColor: "#FF6F61",
			TextCorrectColor:   "#2CB67D",
			TextDefaultColor:   "#c0c8d2ff",
			NormalTextColor:    "#c0c8d2ff",
			CountDownBarColor:  "#7F5AF0",
		},
		{
			Name:               "Theme4",
			BorderActiveColor:  "#7F5AF0",
			BorderDefaultColor: "#c0c8d2ff",
			TabDefaultColor:    "#c0c8d2ff",
			TabActiveColor:     "#7F5AF0",
			TextIncorrectColor: "#FF6F61",
			TextCorrectColor:   "#2CB67D",
			TextDefaultColor:   "#c0c8d2ff",
			NormalTextColor:    "#c0c8d2ff",
			CountDownBarColor:  "#7F5AF0",
		},
		{
			Name:               "Theme5",
			BorderActiveColor:  "#7F5AF0",
			BorderDefaultColor: "#c0c8d2ff",
			TabDefaultColor:    "#c0c8d2ff",
			TabActiveColor:     "#7F5AF0",
			TextIncorrectColor: "#FF6F61",
			TextCorrectColor:   "#2CB67D",
			TextDefaultColor:   "#c0c8d2ff",
			NormalTextColor:    "#c0c8d2ff",
			CountDownBarColor:  "#7F5AF0",
		},
	}
	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	encoder.Encode(defaultConfigs)

	res := []colourTheme{}
	for _, tc := range defaultConfigs {
		res = append(res, convertStructs(tc))
	}
	return res
}

func convertStructs(tc themeConfig) colourTheme {
	colT := colourTheme{}
	colT.borderStyleActive = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(tc.BorderActiveColor)).
		Padding(1, 2)

	colT.borderStyleDefault = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(tc.BorderDefaultColor)).
		Padding(1, 2)

	colT.tabTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.TabDefaultColor)).
		Padding(0, 2)

	colT.tabTextActive = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.TabActiveColor)).
		Bold(true).
		Underline(true).
		Padding(0, 2)

	colT.typeTextIncorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.TextIncorrectColor)).
		Bold(true).
		Underline(true)

	colT.typeTextCorrect = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.TextCorrectColor)).
		Bold(true)

	colT.typeTextDefault = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.TextDefaultColor))

	colT.normalText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.NormalTextColor))

	colT.countDownBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color(tc.CountDownBarColor))

	return colT
}
