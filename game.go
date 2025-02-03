package main

import "math/rand"

type Game struct {
    snake  *Snake
    matrix [][]FieldType
    dead   bool
    tickc  int
    fruits int
}

func newGame(mw int, mh int) *Game {
    snake := newSnake(1, mh / 2 + mh & 1, 4, 10)
    matrix := make([][]FieldType, mw)
    for i := range matrix {
        matrix[i] = make([]FieldType, mh)
    }
    matrixInit(matrix, snake)
    game := &Game{snake: snake, matrix: matrix}
    genFruit(game)
    return game
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
            (!(endx == x && endy == y) &&
            (game.matrix[x][y] == Tail || game.matrix[x][y] == Head)) {

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
        s.head.prev = &Node{x: s.head.x + s.move.dx, y: s.head.y + s.move.dy, typ: Head, next: s.head}
        s.head = s.head.prev

        if game.matrix[s.head.x][s.head.y] == None {
            game.matrix[s.tail.x][s.tail.y] = None
            s.tail = s.tail.prev
            s.tail.next = nil

            s.mid = s.mid.prev
            s.mid.typ = Tail
            game.matrix[s.mid.x][s.mid.y] = Tail
        } else {
            game.snake.length++
            game.fruits--
        }

        game.matrix[s.head.x][s.head.y] = Head
    } else {
        s.tail.next = &Node{x: s.tail.x + s.move.dx, y: s.tail.y + s.move.dy, typ: Tail, prev: s.tail}
        s.tail = s.tail.next

        if game.matrix[s.tail.x][s.tail.y] == None {
            game.matrix[s.head.x][s.head.y] = None
            s.head = s.head.next
            s.head.prev = nil

            game.matrix[s.mid.x][s.mid.y] = Head
            s.mid.typ = Head
            s.mid = s.mid.next
        } else {
            game.snake.length++
            game.fruits--
        }

        game.matrix[s.tail.x][s.tail.y] = Tail
    }
    return true
}

func genFruit(game *Game) {
    if game.fruits > 0 {
        return
    }
    w, h := len(game.matrix), len(game.matrix[0])
    x, y := rand.Intn(w), rand.Intn(h)
    for game.matrix[x][y] != None {
        x, y = rand.Intn(w), rand.Intn(h)
    }
    game.matrix[x][y] = Apple
    if rand.Intn(2) == 1 {
        game.matrix[x][y] = Banana
    }
    game.fruits++
}

func gameTick(game *Game) {
    game.tickc++
    if game.tickc == game.snake.ms {
        game.tickc = 0
        game.dead = !moveSnake(game)
        genFruit(game)
    }
}
