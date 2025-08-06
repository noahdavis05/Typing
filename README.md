# Typing Speed Test

A terminal-based typing speed test application built with Go and Bubble Tea. Test your typing speed and accuracy with customizable settings and real-time feedback.

![Go Version](https://img.shields.io/badge/Go-1.19+-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)

## Features

- **Two Game Modes**: 
  - **Time Limit**: Type as many words as possible within a set time
  - **Word Limit**: Complete a specific number of words as fast as possible
- **Real-time Feedback**: See your typing accuracy with color-coded text
- **Customizable Settings**: Adjust time limits, word counts, and more
- **WPM Calculation**: Track your words per minute and accuracy
- **Clean Terminal UI**: Clean interface using Bubble Tea framework

## Installation

### Prerequisites
- Go 1.19 or later
- A terminal that supports ANSI colors

### Option 1: Download Pre-built Binary
1. Download the latest release from the release page (NO RELEASES YET)

### Option 2: Build from Source
```bash
# Clone the repository
git clone repo
cd typing-app/Typing

# Install dependencies
go mod tidy

# Build the application
go build -o typing-test .

# Run the application
./typing-test
```

### Option 3: Install with Go
```bash
go install repo
```

## Usage

### Basic Controls
- **← →** - Navigate between tabs
- **Enter** - Start a new typing test
- **Ctrl+R** - Restart current test
- **Ctrl+C** - Quit application

### Settings Tab
- **Tab** - Switch between different settings (Game Mode, Time, Words)
- **↑ ↓** - Change the current setting's value

### Game Modes

#### Time Limit Mode
Type as many words as possible within the selected time limit:
- Available times: 15s, 30s, 60s, 90s, 120s
- Progress bar shows remaining time
- WPM calculated based on time elapsed

#### Word Limit Mode
Complete a specific number of words as quickly as possible:
- Available word counts: 15, 30, 50, 60, 100
- Timer shows elapsed time
- WPM calculated when all words are completed



## Technical Details

### Dependencies
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components



### Future Configuration Options
Future versions will support user-configurable:
- Custom color schemes
- Custom word lists
- Extended statistics tracking
- Export/import settings

## Development

### Project Structure
```
Typing/
├── main.go        # Main application and UI logic
├── typing.go      # Core typing test logic
├── timer.go       # Timer implementations
├── settings.go    # Settings management
├── go.mod         # Go module dependencies
└── README.md      # This file
```

