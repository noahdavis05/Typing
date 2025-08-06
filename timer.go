package main

import (
	"fmt"
	"time"
)

const (
	block      rune = '\u2588'
	emptyBlock      = '\u2591'
)

type tickMsg time.Time

type timerUp struct {
	start      time.Time
	started    bool
	finished   bool
	finishTime float64
	wpm float64
}

type timerDown struct {
	seconds  int
	start    time.Time
	started  bool
	finished bool
	wpm float64
}

type timer interface {
	displayTimer(ty *typing) string
	startTimer()
	stopTimer(ty *typing)
	isFinished() bool
}

func (t *timerDown) progressPercent() float64 {
	if t.started && !t.finished {
		return (float64(t.seconds) - time.Since(t.start).Seconds()) / float64(t.seconds) * 100
	}
	return 0
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

func calcWPM(ty *typing, time float64) float64{
	wordCount := 0
	errorCount := 0
	for i := 0; i < len(ty.content); i++{
		if ty.characterColours[i] == "default"{
			break
		}
		if ty.content[i] == ' '{
			wordCount += 1
		}
		if ty.characterColours[i] == "incorrect"{
			errorCount += 1
		}
	}
	return float64(wordCount) * (float64(60) / time) 
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
	if !t.finished{
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

func (t *timerUp) displayTimer(ty *typing) string {
	if t.started && !t.finished {
		return fmt.Sprintf("%.2f s", time.Since(t.start).Seconds())
	} else if t.finished {
		return fmt.Sprintf("%.2f s, WPM = %v", t.finishTime, t.wpm)
	}
	return "0s"
}

func (t *timerDown) displayTimer(ty *typing) string {
	// add the border
	// create the bar
	if t.started && !t.finished {
		if (float64(t.seconds) - time.Since(t.start).Seconds()) < 0 {
			t.stopTimer(ty)
		}
		return (fmt.Sprintf("%s %.2f s", t.displayBar((float64(t.seconds)-time.Since(t.start).Seconds())/float64(t.seconds)*100), float64(t.seconds)-time.Since(t.start).Seconds()))
	} else if t.finished {
		return (fmt.Sprintf("%s %.2f s, WPM = %v", t.displayBar(0), 0.0, t.wpm))
	}
	return (fmt.Sprintf("%s %v.00 s", t.displayBar(100), t.seconds))
}

func (t *timerDown) displayBar(percent float64) string {
	res := ""
	for i := 0; i < int(percent); i++ {
		res += blue.Render(string(block))
	}
	for i := 0; i < 100-int(percent); i++ {
		res += string(emptyBlock)
	}
	return res
}
