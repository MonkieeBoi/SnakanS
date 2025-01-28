package main

import (
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
    def_style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
    head_style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorReset)
    tail_style := tcell.StyleDefault.Background(tcell.ColorGreen).Foreground(tcell.ColorReset)
    box_style := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorReset)

    // Initialize screen
    s, err := tcell.NewScreen()
    if err != nil {
        log.Fatalf("%+v", err)
    }
    if err := s.Init(); err != nil {
        log.Fatalf("%+v", err)
    }
    s.SetStyle(def_style)
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

    border := newWin(s, 0, 0, 50, 23)
    game_win := newWin(s, 2, 1, 46, 21)
    matrix := make([][]BodyType, 23)
    for i := range matrix {
        matrix[i] = make([]BodyType, 21)
    }
    snake := newSnake(1, 11, 4)
    matrixInit(matrix, snake)

    for {
        drawBorder(border, box_style)
        drawMatrix(game_win, matrix, head_style, tail_style)
        // drawStat()
        s.Show()

        select {
        case ev := <-eventq:
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
