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
    length int
}

func newSnake(x int, y int, l int) *Snake {
    t := &Node{typ: Tail, x: x, y: y}
    h := t
    for i := 0; i < l / 2 - 1; i++ {
        h.prev = &Node{typ: Tail, x: h.x + 1, y: y, next: h}
        h = h.prev
    }
    for i := 0; i < l / 2; i++ {
        h.prev = &Node{typ: Head, x: h.x + 1, y: y, next: h}
        h = h.prev
    }
    s := Snake{head: h, tail: t, length: l}

    return &s
}

func matrixInit(matrix [][]BodyType, snake *Snake) {
    cur := snake.head
    for cur != nil {
        matrix[cur.x][cur.y] = cur.typ
        cur = cur.next
    }
}
