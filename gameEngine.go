package main

import (
    "fmt"

    b "gitlab.com/Arkan501/arkanchesslib/board"
    p "gitlab.com/Arkan501/arkanchesslib/pieces"
)

func GameEngine() {
	playerWhite := p.White
	playerBlack := p.Black
	currentPlayer := playerWhite
	chessBoard := b.NewBoard()

	for chessBoard.GameState(currentPlayer) == 0 {
		clearScreen()
		fmt.Println("Arkan's Chess CLI")
		fmt.Printf("%s to play\n", p.ColourString[currentPlayer])
		chessBoard.DrawBoard()
		chooseMove(currentPlayer, &chessBoard)
		switch currentPlayer {
		case playerWhite:
			currentPlayer = playerBlack
		default:
			currentPlayer = playerWhite
		}
	}

	clearScreen()
	fmt.Println("Arkan's Chess CLI")
	chessBoard.DrawBoard()

    switch chessBoard.GameState(currentPlayer) {
    case 1:
        fmt.Printf(
            "%s is in Checkmate. %s wins!",
            p.ColourString[playerBlack],
            p.ColourString[playerWhite],
            )
    case 2:
        fmt.Printf(
            "%s is in Checkmate. %s wins!",
            p.ColourString[playerWhite],
            p.ColourString[playerBlack],
            )
    case 3:
        fmt.Printf("draw!")
    }
}

func chooseMove(colour p.Colour, chessBoard *b.Board) {
	var fromSquare string
	var toSquare string
	var err error

	for {
		fmt.Println("Choose a piece to move: ")
		fmt.Scan(&fromSquare)
		fmt.Println("Choose a square to move to: ")
		fmt.Scan(&toSquare)

		convertedFromSquare := b.NotationToIndex[fromSquare]
		convertedToSquare := b.NotationToIndex[toSquare]

		err = chessBoard.MakeMove(colour, convertedFromSquare, convertedToSquare)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
}
