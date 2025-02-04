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
    box_style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorWhite)
    apple_style := tcell.StyleDefault.Background(tcell.ColorRed).Foreground(tcell.ColorReset)
    banana_style := tcell.StyleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorReset)

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

    width := 23
    height := 21
    border := newWin(s, 0, 0, width * 2 + 2, height + 2)
    game_win := newWin(s, border.x + 1, border.y + 1, width * 2, height)
    game := newGame(width, height)

    // Event loop
    eventq := make(chan tcell.Event)
    quitq := make(chan struct{})
    go s.ChannelEvents(eventq, quitq)

    for !game.dead {
        s.Clear()
        drawBorder(border, box_style)
        drawMatrix(game_win, game.matrix, head_style, tail_style, apple_style, banana_style)
        // drawStat()
        s.Show()

        select {
        case ev := <-eventq:
            switch ev := ev.(type) {
            case *tcell.EventResize:
                s.Sync()
            case *tcell.EventKey:
                if ev.Rune() == 'q' || ev.Key() == tcell.KeyCtrlC {
                    return
                } else if ev.Key() == tcell.KeyCtrlL {
                    s.Sync()
                } else if ev.Rune() == 'w' {
                    turnSnake(game.snake, UP)
                } else if ev.Rune() == 'a' {
                    turnSnake(game.snake, LEFT)
                } else if ev.Rune() == 's' {
                    turnSnake(game.snake, DOWN)
                } else if ev.Rune() == 'd' {
                    turnSnake(game.snake, RIGHT)
                } else if ev.Rune() == ' ' {
                    flipSnake(game.snake)
                }
            }
        default:
        }
        gameTick(game)
        time.Sleep(time.Millisecond * 20)
    }
}
