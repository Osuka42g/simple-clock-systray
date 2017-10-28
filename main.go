package main

import (
	"fmt"
	"io/ioutil"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("assets/clock.ico"))
	systray.SetTitle("I'm alive!")
	systray.SetTooltip("Look at me, I'm a tooltip!")
}

func onExit() {
	// Cleaning stuff here.
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
