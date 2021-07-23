package main

import (
    "fmt"
    "time"

    "github.com/go-vgo/robotgo"
    hook "github.com/robotn/gohook"
)

const (
    impairedColorPixelX1920x1080 = 1297
    impairedColorPixelY1920x1080 = 1016
    impairedColorPixel           = "c02826"
)

var (
    qssToggle     bool
    cleanseToggle bool
)

var (
    qssKey     string
    cleanseKey string
)

func main() {
    fmt.Print("Insert qss key: ")
    fmt.Scan(&qssKey)
    fmt.Print("Insert cleanse key: ")
    fmt.Scan(&cleanseKey)

    go autoClick(33)
    add()
}

func add() {
    fmt.Println("====AUTO QSS/CLEANSE====")
    fmt.Println("After the qss/cleanse has been used, it will auto toggle off")
    fmt.Println("\tPress T to toggle auto qss")
    fmt.Println("\tPress X to toggle auto cleanse")
    fmt.Println("\tPress M to toggle off")

    robotgo.EventHook(hook.KeyDown, []string{"t"}, func(e hook.Event) {
        qssToggle = true
        cleanseToggle = false
    })

    robotgo.EventHook(hook.KeyDown, []string{"x"}, func(e hook.Event) {
        qssToggle = false
        cleanseToggle = true
    })

    robotgo.EventHook(hook.KeyDown, []string{"m"}, func(e hook.Event) {
        qssToggle = false
        cleanseToggle = false
    })

    s := robotgo.EventStart()
    <-robotgo.EventProcess(s)
}

func autoClick(delay time.Duration) {
    for {
        if qssToggle && isPixelColor(impairedColorPixelX1920x1080, impairedColorPixelY1920x1080,
            impairedColorPixel) {
            robotgo.KeyTap(qssKey)
            qssToggle = false
        }
        if cleanseToggle && isPixelColor(impairedColorPixelX1920x1080, impairedColorPixelY1920x1080,
            impairedColorPixel) {
            robotgo.KeyTap(cleanseKey)
            cleanseToggle = false
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
