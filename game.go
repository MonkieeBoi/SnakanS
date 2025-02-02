package main

type Game struct {
    snake  *Snake
    matrix [][]BodyType
    dead   bool
    tickc  int
}

func newGame(mw int, mh int) *Game {
    snake := newSnake(1, mh / 2 + mh & 1, 4, 10)
    matrix := make([][]BodyType, mw)
    for i := range matrix {
        matrix[i] = make([]BodyType, mh)
    }
    matrixInit(matrix, snake)
    return &Game{snake: snake, matrix: matrix}
}

func validMove(game *Game) bool {
    s := game.snake
    var x int
    var y int
    var endx int
    var endy int

    if s.end == Head {
        x = s.head.x + s.move.dx
        y = s.head.y + s.move.dy
        endx = s.tail.x
        endy = s.tail.y
    } else {
        x = s.tail.x + s.move.dx
        y = s.tail.y + s.move.dy
        endx = s.head.x
        endy = s.head.y
    }

    if x < 0 ||
        y < 0 ||
        x >= len(game.matrix) ||
        y >= len(game.matrix[0]) ||
        (!(endx == x && endy == y) && game.matrix[x][y] != None) {

        return false
    }
    return true
}

func moveSnake(game *Game) bool {
    if !validMove(game) {
        return false
    }
    s := game.snake
    if s.end == Head {
        game.matrix[s.tail.x][s.tail.y] = None

        s.tail = s.tail.prev
        s.tail.next = nil

        s.head.prev = &Node{x: s.head.x + s.move.dx, y: s.head.y + s.move.dy, typ: Head, next: s.head}
        s.head = s.head.prev
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
        game.matrix[s.tail.x][s.tail.y] = Tail

        game.matrix[s.mid.x][s.mid.y] = Head
        s.mid.typ = Head
        s.mid = s.mid.next
    }
    return true
}

func gameTick(game *Game) {
    game.tickc++
    if game.tickc == game.snake.ms {
        game.tickc = 0
        game.dead = !moveSnake(game)
    }
}
