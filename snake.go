package main

type BodyType int

const (
    None = iota
    Head
    Tail
)

type Movement struct {
    dx int
    dy int
}

type Node struct {
    typ  BodyType
    x    int
    y    int
    next *Node
    prev *Node
}

type Snake struct {
    head   *Node
    tail   *Node
    mid    *Node
    length int
    move   Movement
    end    int
}

func newSnake(x int, y int, l int) *Snake {
    t := &Node{typ: Tail, x: x, y: y}
    h := t
    for i := 0; i < l / 2 - 1; i++ {
        h.prev = &Node{typ: Tail, x: h.x + 1, y: y, next: h}
        h = h.prev
    }
    m := h
    for i := 0; i < l / 2; i++ {
        h.prev = &Node{typ: Head, x: h.x + 1, y: y, next: h}
        h = h.prev
    }
    snake := Snake{head: h, tail: t, length: l, move:Movement{dx: 1}, end: Head, mid: m}

    return &snake
}

func matrixInit(matrix [][]BodyType, snake *Snake) {
    cur := snake.head
    for cur != nil {
        if cur.x >= len(matrix) || cur.y >= len(matrix[0]) {
            continue
        }
        matrix[cur.x][cur.y] = cur.typ
        cur = cur.next
    }
}
