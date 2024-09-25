
/* __TODO__
	1- adding new phase in the game(more maze)
	2- adding final message (ascii message)
	3- handle new and better colorscheme
	4- slove problem in rander the game
	5- KILL MY SELF 
*/


package main

import (
	"fmt"
	"log"
	"time"
)

// PUBLIC VAR
var player Sprite
var ghosts []*Sprite
var maze []string
var score int
var numDots int
var lives = 1
var cfg Config

// Init: function for initailzatoin the game
func Init() {
	// initialise game
	Initialise()
	defer Cleanup()

	// load resources
	err := LoadMaze("maze01.txt")
	if err != nil { log.Println("failed to load maze:", err) }
	err = LoadConfig("config.json")
	if err != nil { log.Println("failed to load maze:", err) }

	// process input (async) because input stop
	input := make(chan string)
	go func(ch chan<- string) {
		for {
			input, err := ReadInput()
			if err != nil {
				log.Println("error reading input:", err)
				ch <- "ESC"
			}
			ch <- input
		}
	}(input)

	// game loop
	for {
		var prevDir string
		// process movement
		select {
		case inp := <-input:
			if inp == "ESC" {
				lives = 0
			}
			// process movement
			MovePlayer(inp)
		default:
			// TODO : DO IT
			if prevDir != "" { 
				MovePlayer(prevDir)
			}
		}

		// update screen
		PrintScreen()

		// process movement
		MoveGhosts()

		// process collisions
		for _, g := range ghosts {
			if player.row == g.row  && player.col == g.col {
				lives = 0
			}
		}


		// check game over
		if numDots <= 0 || lives <= 0{
			if lives == 0 {
				moveCursor(player.row, player.col)
				fmt.Print(cfg.Death)
				moveCursor(len(maze)+1, 0)
				fmt.Println("Score:", score, "\tLives:", lives)
				fmt.Println("GAME OVER LOOSSER......")
				moveCursor(len(maze)+2, 0)// move cursor out

				break
			}
			fmt.Println("GOOD GAME FOO......")
			break
		}

		// repeat
		time.Sleep(130 * time.Millisecond)
	}
}


