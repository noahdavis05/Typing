package main

import (
	"os"
	"path/filepath"
)

const (
	configDirName  = "typingTester"
	configFilename = "config.json"
)

func (m model) loadConfig() error {
	// get the config dir
	configDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// make the filepath
	appConfigDir := filepath.Join(configDir, configDirName, configFilename)

	// try and open the file
	file, err := os.Open(appConfigDir)
	if err != nil {
		if os.IsNotExist(err) {
			// the newly created config file will be empty - so just return from function
			return m.createConfig(configDir)
		}
		return err
	}
	defer file.Close()
	data, err := os.ReadFile(appConfigDir)
	if err != nil {
		return err
	}
	println(string(data))
	return nil

}

func (m model) createConfig(configDir string) error {
	// create directory if doesn't exist
	appConfigDir := filepath.Join(configDir, configDirName)
	err := os.MkdirAll(appConfigDir, 0755)
	if err != nil {
		return err
	}

	// now create the file
	configFilePath := filepath.Join(appConfigDir, configFilename)
	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	writeInitialConfig(file)
	return nil
}

func writeInitialConfig(f *os.File) {
	f.WriteString("hello world")
}
