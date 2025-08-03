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
}

type timerDown struct {
	seconds  int
	start    time.Time
	started  bool
	finished bool
}

type timer interface {
	displayTimer() string
	startTimer()
	stopTimer()
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
	}
}

func (t *timerDown) startTimer() {
	if !t.started {
		t.started = true
		t.start = time.Now()
	}
}

func (t *timerUp) stopTimer() {
	t.finished = true
	t.finishTime = time.Since(t.start).Seconds()
}

func (t *timerDown) stopTimer() {
	t.finished = true
}

func (t *timerUp) isFinished() bool {
	return t.finished
}

func (t *timerDown) isFinished() bool {
	return t.finished
}

func (t *timerUp) displayTimer() string {
	if t.started && !t.finished {
		return fmt.Sprintf("%.2f s", time.Since(t.start).Seconds())
	} else if t.finished {
		return fmt.Sprintf("%.2f s", t.finishTime)
	}
	return "0s"
}

func (t *timerDown) displayTimer() string {
	// add the border
	// create the bar
	if t.started && !t.finished {
		if (float64(t.seconds) - time.Since(t.start).Seconds()) < 0 {
			t.finished = true
		}
		return (fmt.Sprintf("%s %.2f s", t.displayBar((float64(t.seconds)-time.Since(t.start).Seconds())/float64(t.seconds)*100), float64(t.seconds)-time.Since(t.start).Seconds()))
	} else if t.finished {
		return (fmt.Sprintf("%s %.2f s", t.displayBar(0), 0.0))
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
