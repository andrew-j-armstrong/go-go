package reversi

import "fmt"

type Piece int

const (
	EmptyPiece Piece = iota
	Player1Piece
	Player2Piece
)

const BoardHeight int = 8
const BoardWidth int = 8

type Board [BoardHeight][BoardWidth]Piece

func (board *Board) IsEqual(otherBoard *Board) bool {
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			if board[y][x] != otherBoard[y][x] {
				return false
			}
		}
	}

	return true
}

func (board *Board) String() string {
	var output string
	output += "+---+---+---+---+---+---+---+---+---+\n"
	output += "|   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |\n"
	for y := 0; y < BoardHeight; y++ {
		output += "+---+---+---+---+---+---+---+---+---+\n"
		output += fmt.Sprintf("| %c ", 'A'+y)
		for x := 0; x < BoardWidth; x++ {
			output += "| "
			switch board[y][x] {
			case EmptyPiece:
				output += "  "
			case Player1Piece:
				output += "B "
			case Player2Piece:
				output += "W "
			default:
				output += "? "
			}
		}
		output += "|\n"
	}
	output += "+---+---+---+---+---+---+---+---+---+\n"

	return output
}

func (board *Board) Print() {
	print(board.String())
}

func (board *Board) Clone() *Board {
	newBoard := &Board{}

	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			newBoard[y][x] = board[y][x]
		}
	}

	return newBoard
}

func (board *Board) CloneGeneric() interface{} {
	return board.Clone()
}
