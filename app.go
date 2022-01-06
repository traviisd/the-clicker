package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const (
	leftCtrlCode  = 59
	leftShiftCode = 56
	sKey          = 1
	qKey          = 12
)

var (
	click    *bool
	shutdown = false
	clicks   = 0
	delay    *int
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	// default delay
	d := 500
	c := false
	delay = &d
	click = &c
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	go b.hooks()
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	c := false
	click = &c
	shutdown = true
	delay = nil
	clicks = 0
}

func (b *App) hooks() {
	evChan := hook.Start()
	defer hook.End()

	var ctrlHold, ctrlUp, shiftHold, shiftUp, startHold, startUp, stopHold, stopUp bool

	for ev := range evChan {
		switch ev.Kind {
		case hook.KeyHold:
			// ctrl/control
			if ev.Rawcode == leftCtrlCode {
				ctrlHold = true
				ctrlUp = false
			}
			// shift key
			if ev.Rawcode == leftShiftCode {
				shiftHold = true
				shiftUp = false
			}

			if ev.Rawcode == sKey {
				startHold = true
				startUp = false
			}

			if ev.Rawcode == qKey {
				stopHold = true
				stopUp = false
			}

		case hook.KeyUp:
			// ctrl/control
			if ev.Rawcode == leftCtrlCode {
				ctrlUp = true
			}

			// shift key
			if ev.Rawcode == leftShiftCode {
				shiftUp = true
			}

			if ev.Rawcode == sKey {
				startUp = true
			}

			if ev.Rawcode == qKey {
				stopUp = true
			}
		}

		ctrlEnabled := (ctrlHold && !ctrlUp)
		shiftEnabled := (shiftHold && !shiftUp)
		startEnabled := (startHold && !startUp)
		stopEnabled := (stopHold && !stopUp)

		// start action
		if !*click && ctrlEnabled && shiftEnabled && startEnabled {
			ctrlHold = false
			shiftHold = false
			startHold = false
			c := true
			click = &c
			clicks = 0
			go b.click()
		}

		// stop action
		if ctrlEnabled && shiftEnabled && stopEnabled {
			ctrlHold = false
			shiftHold = false
			stopHold = false
			c := false
			click = &c
		}

		if shutdown {
			break
		}
	}
}

func (b *App) Clicks() string {
	return fmt.Sprint(clicks)
}

func (b *App) ClicksPerSecond() string {
	d := *delay
	return fmt.Sprintf("%d clicks per second", (1000 / d))
}

func (b *App) SetDelay(d string) {
	n, err := strconv.ParseInt(d, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	i := int(n)
	delay = &i
}

func (b *App) click() {
	for {
		if !*click {
			break
		}

		robotgo.Click()

		clicks++

		time.Sleep(time.Duration(*delay) * time.Millisecond)
	}
}
