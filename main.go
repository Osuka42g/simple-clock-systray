package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/getlantern/systray"
)

var (
	timezone string
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	timezone = "Local"

	systray.SetIcon(getIcon("assets/clock.ico"))

	localTime := systray.AddMenuItem("Local time", "Local time")
	hcmcTime := systray.AddMenuItem("Ho Chi Minh time", "Asia/Ho_Chi_Minh")
	sydTime := systray.AddMenuItem("Sydney time", "Australia/Sydney")
	gdlTime := systray.AddMenuItem("Guadalajara time", "America/Mexico_City")
	sfTime := systray.AddMenuItem("San Fransisco time", "America/Los_Angeles")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	go func() {
		for {
			systray.SetTitle(getClockTime(timezone))
			systray.SetTooltip(timezone + " timezone")
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			select {
			case <-localTime.ClickedCh:
				timezone = "Local"
			case <-hcmcTime.ClickedCh:
				timezone = "Asia/Ho_Chi_Minh"
			case <-sydTime.ClickedCh:
				timezone = "Australia/Sydney"
			case <-gdlTime.ClickedCh:
				timezone = "America/Mexico_City"
			case <-sfTime.ClickedCh:
				timezone = "America/Los_Angeles"
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}

func getClockTime(tz string) string {
	t := time.Now()
	utc, _ := time.LoadLocation(tz)

	hour, min, sec := t.In(utc).Clock()
	return ItoaTwoDigits(hour) + ":" + ItoaTwoDigits(min) + ":" + ItoaTwoDigits(sec)
}

// ItoaTwoDigits time.Clock returns one digit on values, so we make sure to convert to two digits
func ItoaTwoDigits(i int) string {
	b := "0" + strconv.Itoa(i)
	return b[len(b)-2:]
}

func getIcon(s string) []byte {
	b, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	return b
}
