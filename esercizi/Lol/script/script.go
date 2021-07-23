package main

import (
    "time"
    _"fmt"
    _"log"
    "github.com/go-vgo/robotgo"
)

const HeraldHP = 7125.0
const SmiteDamage6 = 510.0

// 7125 : 100 = 510 : x     x = 51000 / 7125
func main() {
    time.Sleep(4 * time.Second)

    perc := SmiteDamage6 * 100 / HeraldHP
    for {
        if getHPPercentageFromScreen() <= perc {
            smite()
            break
        }
        time.Sleep(33 * time.Millisecond)
    }
}

func smite() {
    robotgo.KeyTap("d")
}

func getHPPercentageFromScreen() float64 {
    var health float64 = 10.0
    hpPerPixel := 100.0 / (342 - 203)

    for x := 203 + 14 - 1; x >= 203; x-- {

        color := robotgo.GetPixelColor(x, 40)

        if string(color[0]) != "6" && string(color[0]) != "7" && string(color[0]) != "8" && string(color[0]) != "9" &&
            string(color[0]) != "a" && string(color[0]) != "b" && string(color[0]) != "c" {
            health -= hpPerPixel
        } else {
            break
        }

    }

    return health
}
