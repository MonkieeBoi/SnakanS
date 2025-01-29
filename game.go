package main

type Game struct {
    snake  *Snake
    matrix [][]BodyType
    dead   bool
}

func newGame(mw int, mh int) *Game {
    snake := newSnake(1, mh / 2 + mh & 1, 4)
    matrix := make([][]BodyType, mw)
    for i := range matrix {
        matrix[i] = make([]BodyType, mh)
    }
    matrixInit(matrix, snake)
    return &Game{snake: snake, matrix: matrix}
}

func moveSnake(game *Game) bool {
    s := game.snake
    if s.end == Head {
        game.matrix[s.tail.x][s.tail.y] = None

        s.tail = s.tail.prev
        s.tail.next = nil

        s.head.prev = &Node{x: s.head.x + s.move.dx, y: s.head.y + s.move.dy, typ: Head, next: s.head}
        s.head = s.head.prev
        if s.head.x < 0 || s.head.y < 0 || s.head.x >= len(game.matrix) || s.head.y >= len(game.matrix[0]) {
            return false
        }
        game.matrix[s.head.x][s.head.y] = Head

        s.mid = s.mid.prev
        s.mid.typ = Tail
        game.matrix[s.mid.x][s.mid.y] = Tail
    } else {
        game.matrix[s.head.x][s.head.y] = None

        s.head = s.head.next
        s.head.prev = nil

        s.tail.next = &Node{x: s.tail.x + s.move.dx, y: s.tail.y + s.move.dy, typ: Tail, prev: s.tail}
        s.tail = s.tail.next
        if s.tail.x < 0 || s.tail.y < 0 || s.tail.x >= len(game.matrix) || s.tail.y >= len(game.matrix[0]) {
            return false
        }
        game.matrix[s.tail.x][s.tail.y] = Tail

        game.matrix[s.mid.x][s.mid.y] = Head
        s.mid.typ = Head
        s.mid = s.mid.next

    }
    return true
}

func gameTick(game *Game) {
    // bouncing for now
    s := game.snake
    if s.head.x == len(game.matrix) - 1 {
        s.move.dx = -1
        s.end = Tail
    } else if s.tail.x == 0  {
        s.move.dx = 1
        s.end = Head
    }
    if !moveSnake(game) {
        game.dead = true
    }
}
