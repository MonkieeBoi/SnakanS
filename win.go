package main

import (
    "github.com/gdamore/tcell/v2"
)

type Win struct {
    x int
    y int
    w int
    h int
    s tcell.Screen
}

func newWin(s tcell.Screen, x int, y int, w int, h int) *Win {
    win := Win{x: x, y: y, w: w, h: h, s: s}
    return &win
}

func drawBorder(w *Win, style tcell.Style) {
    // Top
    for i := 0; i < w.w; i++ {
        w.s.SetContent(w.x+i, w.y, ' ', nil, style)
    }
    // Bot
    for i := 0; i < w.w; i++ {
        w.s.SetContent(w.x+i, w.y+w.h, ' ', nil, style)
    }
    // Left
    for j := 0; j < w.h; j++ {
        w.s.SetContent(w.x, w.y+j, ' ', nil, style)
        w.s.SetContent(w.x+1, w.y+j, ' ', nil, style)
    }
    // Right
    for j := 0; j < w.h; j++ {
        w.s.SetContent(w.x+w.w-1, w.y+j, ' ', nil, style)
        w.s.SetContent(w.x+w.w-2, w.y+j, ' ', nil, style)
    }
}

func drawCell(w *Win, x int, y int, c rune, style tcell.Style) {
    if x >= w.w || y >= w.h {
        return
    }
    w.s.SetContent(w.x + x, w.y + y, c, nil, style)
}

func drawMatrix(w *Win, matrix [][]BodyType, h_style tcell.Style, t_style tcell.Style) {
    for i, col := range matrix {
        if i >= w.w {
            break
        }
        for j, typ := range col {
            if j >= w.h {
                break
            }
            switch typ {
            case Head:
                drawCell(w, i * 2, j, ' ', h_style)
                drawCell(w, i * 2 + 1, j, ' ', h_style)
            case Tail:
                drawCell(w, i * 2, j, ' ', t_style)
                drawCell(w, i * 2 + 1, j, ' ', t_style)
            }
        }
    }
}
