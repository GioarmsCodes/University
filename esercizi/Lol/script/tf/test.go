package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var (
	DelayBlueCard   time.Duration = 33
	DelayRedCard    time.Duration = 33
	DelayYellowCard time.Duration = 33
)

var (
	blueCardToggle   bool
	redCardToggle    bool
	yellowCardToggle bool
)

const (
	YellowCardColor = "e7e700"
	BlueCardColor   = "c6cbef"
	RedCardColor    = "f41f1e"
)

const (
	CardColorPixelX1920x1080 = 824
	CardColorPixelY1920x1080 = 975
)

func main() {
	go checkCard(10)
	add()
}

func add() {
	fmt.Println("Press T for toggle yellow card")
	fmt.Println("Press X for toggle blue card")
	fmt.Println("Press N for toggle red card")
	fmt.Println("Press M to untoggle")

	robotgo.EventHook(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		blueCardToggle = false
		redCardToggle = false
		yellowCardToggle = true
	})

	robotgo.EventHook(hook.KeyDown, []string{"x"}, func(e hook.Event) {
		blueCardToggle = true
		redCardToggle = false
		yellowCardToggle = false
	})

	robotgo.EventHook(hook.KeyDown, []string{"n"}, func(e hook.Event) {
		blueCardToggle = false
		redCardToggle = true
		yellowCardToggle = false
	})

	robotgo.EventHook(hook.KeyDown, []string{"m"}, func(e hook.Event) {
		blueCardToggle = false
		redCardToggle = false
		yellowCardToggle = false
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}

func checkCard(delay time.Duration) {
	for {
		if yellowCardToggle && isPixelColor(CardColorPixelX1920x1080, CardColorPixelY1920x1080, YellowCardColor) {
			robotgo.KeyTap("w")
		}
		if blueCardToggle && isPixelColor(CardColorPixelX1920x1080, CardColorPixelY1920x1080, BlueCardColor) {
			robotgo.KeyTap("w")
		}
		if redCardToggle && isPixelColor(CardColorPixelX1920x1080, CardColorPixelY1920x1080, RedCardColor) {
			robotgo.KeyTap("w")
		}

		time.Sleep(delay * time.Millisecond)
	}
}

func isPixelColor(x, y int, color string) bool {
	if robotgo.GetPixelColor(x, y) == color {
		return true
	}
	return false
}
