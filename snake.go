package main

type FieldType int

const (
    None = iota
    Head
    Tail
    Apple
    Banana
)

type Movement struct {
    dx int
    dy int
}

type Node struct {
    typ  FieldType
    x    int
    y    int
    next *Node
    prev *Node
}

type Snake struct {
    head   *Node
    tail   *Node
    mid    *Node
    move   *Movement
    length int
    end    int
    ms     int
}

var UP = &Movement{dy: -1}
var DOWN = &Movement{dy: 1}
var LEFT = &Movement{dx: -1}
var RIGHT = &Movement{dx: 1}

func newSnake(x int, y int, l int, ms int) *Snake {
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
    snake := Snake{head: h, tail: t, length: l, move:RIGHT, end: Head, mid: m, ms: ms}

    return &snake
}

func turnSnake(snake *Snake, move *Movement) {
    invalid := &Movement{dx: -snake.move.dx, dy: -snake.move.dy}
    if *move != *invalid {
        snake.move = move
    }
}

func flipSnake(snake *Snake) {
    snake.end = (snake.end % 2) + 1
    cur := snake.head
    last := cur.next
    if snake.end == Tail {
        cur = snake.tail
        last = cur.prev
    }
    snake.move = &Movement{dx: cur.x - last.x, dy: cur.y - last.y}
}

func matrixInit(matrix [][]FieldType, snake *Snake) {
    cur := snake.head
    for cur != nil {
        if cur.x >= len(matrix) || cur.y >= len(matrix[0]) {
            continue
        }
        matrix[cur.x][cur.y] = cur.typ
        cur = cur.next
    }
}
