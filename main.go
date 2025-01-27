package main

import (
    "fmt"
    "log"

    "github.com/gdamore/tcell/v2"
    "time"
)

func drawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
    row := y1
    col := x1
    for _, r := range []rune(text) {
        s.SetContent(col, row, r, nil, style)
        col++
        if col >= x2 {
            row++
            col = x1
        }
        if row > y2 {
            break
        }
    }
}

func main() {
    defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

    // Initialize screen
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }
    if err := s.Init(); err != nil {
        log.Fatalf("%+v", err)
    }
    s.SetStyle(defStyle)
    s.Clear()

    quit := func() {
        maybePanic := recover()
        s.Fini()
        if maybePanic != nil {
            panic(maybePanic)
        }
    }
    defer quit()

    // Event loop
    eventq := make(chan tcell.Event)
    quitq := make(chan struct{})
    go s.ChannelEvents(eventq, quitq)

    x := 0
    for {
        // Update screen
        drawText(s, 0, 0, 20, 20, tcell.StyleDefault, fmt.Sprintf("%d", x))
        x++
        s.Show()

        select {
        case ev := <- eventq:
            // Process event
            switch ev := ev.(type) {
            case *tcell.EventResize:
                s.Sync()
            case *tcell.EventKey:
                if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
                    return
                } else if ev.Key() == tcell.KeyCtrlL {
                    s.Sync()
                }
            }
        default:
        }
        s.Clear()
        time.Sleep(time.Millisecond * 16)
    }
}
