package main

import (
	"fmt"
	"time"
)

const (
	block      rune = '\u2588'
	emptyBlock rune = '\u2591'
)

type tickMsg time.Time

// struct for a timer which goes up - used in 'words' gamemode where player is timed
// how long it takes to complete a set of words
type timerUp struct {
	start      time.Time
	started    bool
	finished   bool
	finishTime float64
	wpm        float64
}

// struct for a countdown timer - used in 'countdown' gamemode
// this also includes a countdown bar as well
type timerDown struct {
	seconds  int
	start    time.Time
	started  bool
	finished bool
	wpm      float64
}

// interface for both types of timer
type timer interface {
	displayTimer(designStyles colourTheme) string
	startTimer()
	stopTimer(ty *typing)
	isFinished() bool
	isActive() bool
}

func (t *timerUp) isActive() bool {
	if t.started && !t.finished{
		return true
	}
	return false
}

func (t *timerDown) isActive() bool {
	if t.started && !t.finished{
		return true
	}
	return false
}


func (t *timerUp) startTimer() {
	if !t.started {
		t.started = true
		t.start = time.Now()
		t.wpm = -1
	}
}

func (t *timerDown) startTimer() {
	if !t.started {
		t.started = true
		t.start = time.Now()
		t.wpm = -1
	}
}

// calculates words per minute with formula
// wpm = (characters typed - incorrect characters) / 5 * 60/time taken
func calcWPM(ty *typing, time float64) float64 {
	errorCount := 0
	var i int
	for i = 0; i < len(ty.content); i++ {
		if ty.characterColours[i] == defaultKey {
			break
		}
		if ty.characterColours[i] == incorrectKey {
			errorCount += 1
		}
	}
	return (float64(i) - float64(errorCount)) / 5 * (float64(60) / time)
}

func (t *timerUp) stopTimer(ty *typing) {
	if !t.finished {
		t.finished = true
		t.finishTime = time.Since(t.start).Seconds()
		// calculate words per minute
		// iterate over content for every space we add 1 to word count, also make an error count as we go through
		t.wpm = calcWPM(ty, t.finishTime)
	}
}

func (t *timerDown) stopTimer(ty *typing) {
	if !t.finished {
		t.finished = true
		t.wpm = calcWPM(ty, float64(t.seconds))
	}
}

func (t *timerUp) isFinished() bool {
	return t.finished
}

func (t *timerDown) isFinished() bool {
	return t.finished
}

func (t *timerUp) displayTimer(designStyles colourTheme) string {
	if t.started && !t.finished {
		return designStyles.normalText.Render(fmt.Sprintf("%.2f s", time.Since(t.start).Seconds()))
	} else if t.finished {
		return designStyles.normalText.Render(fmt.Sprintf("%.2f s\nWPM = %.2f", t.finishTime, t.wpm))
	}
	return "0s"
}

func (t *timerDown) displayTimer(designStyles colourTheme) string {
	// add the border
	// create the bar
	if t.started && !t.finished {
		return t.displayBar((float64(t.seconds)-time.Since(t.start).Seconds())/float64(t.seconds)*100, designStyles) + designStyles.normalText.Render(fmt.Sprintf(" %.2f s", float64(t.seconds)-time.Since(t.start).Seconds()))
	} else if t.finished {
		return t.displayBar(0, designStyles) + designStyles.normalText.Render(fmt.Sprintf(" %.2f s\nWPM = %.2f", 0.0, t.wpm))
	}
	return t.displayBar(100, designStyles) + designStyles.normalText.Render(fmt.Sprintf(" %v.00 s", t.seconds))
}

func (t *timerDown) displayBar(percent float64, designStyles colourTheme) string {
	res := ""
	for i := 0; i < int(percent); i++ {
		res += string(block)
	}
	for i := 0; i < 100-int(percent); i++ {
		res += string(emptyBlock)
	}
	return designStyles.countDownBar.Render(res)
}
