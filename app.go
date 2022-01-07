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

var (
	click  bool
	clicks = 0
	delay  = 500
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
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
	hook.End()
}

func (b *App) hooks() {
	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "t"}, func(e hook.Event) {
		click = !click
		if click {
			clicks = 0
			go b.click()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func (b *App) Clicks() string {
	return fmt.Sprint(clicks)
}

func (b *App) ClicksPerSecond() string {
	return fmt.Sprintf("%d clicks per second", (1000 / delay))
}

func (b *App) SetDelay(d string) {
	n, err := strconv.ParseInt(d, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	delay = int(n)
}

func (b *App) click() {
	for {
		if !click {
			break
		}

		robotgo.Click()

		clicks++

		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
