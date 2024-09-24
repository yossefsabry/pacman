package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

const (
    ESC = "\033"
    UP = "[A"
    DOWN = "[B"
    LEFT = "[D"
    RIGHT = "[C"
)


// ReadInput : reading the key pressing from user
func ReadInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	} else if cnt >= 3 {
		if buffer[0] == 0x1b && buffer[1] == '[' {
			switch buffer[2] {
			case 'A', 'h':
				return "UP", nil
			case 'B':
				return "DOWN", nil
			case 'C':
				return "RIGHT", nil
			case 'D':
				return "LEFT", nil
			}
		}
	}

	return "", nil
}

// MakeMove: the  main function for moving a player or ghost
func MakeMove(oldRow, oldCol int, dir string) (newRow, newCol int) {
	newRow, newCol = oldRow, oldCol

	switch dir {
	case "UP":
		newRow = newRow - 1
		if newRow < 0 {
			newRow = len(maze) - 1
		}
	case "DOWN":
		newRow = newRow + 1
		if newRow == len(maze) {
			newRow = 0
		}
	case "RIGHT":
		newCol = newCol + 1
		if newCol == len(maze[0]) {
			newCol = 0
		}
	case "LEFT":
		newCol = newCol - 1
		if newCol < 0 {
			newCol = len(maze[0]) - 1
		}
	}

	if maze[newRow][newCol] == '#' || maze[newRow][newCol] == '$' {
		newRow = oldRow
		newCol = oldCol
	}

	return
}

// MovePlayer: moving a player
func MovePlayer(dir string) {
	player.row, player.col = MakeMove(player.row, player.col, dir)
	removeDot := func(row, col int) {
		maze[row] = maze[row][0:col] + " " + maze[row][col+1:]
	}
	switch maze[player.row][player.col] {
	case '.':
		numDots--	
		score++
		// remove the dots from screen
		removeDot(player.row, player.col)
	case 'X':
		score += 10
		removeDot(player.row, player.col)
	}
}

// DrawDirection: random dir
func DrawDirection() string {
	dir := rand.Intn(4)
	move := map[int]string{
		0: "UP",
		1: "DOWN",
		2: "RIGHT",
		3: "LEFT",
	}
	return move[dir]
}

// MoveGhosts: moving the ghost
func MoveGhosts() {
	for _, g := range ghosts {
		dir := DrawDirection()
		g.row, g.col = MakeMove(g.row, g.col, dir)
	}
}

// Initialise: initzliaztion the terminal for the game
func Initialise() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cbreak mode:", err)
	}
}

// Cleanup: for clean up the code after the main end
func Cleanup() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalln("unable to activate cooked mode:", err)
	}
}

// moveCursor: for moving the cursor for the new place
func moveCursor(row, col int) {
    if cfg.UseEmoji {
        simpleansi.MoveCursor(row, col*2)
    } else {
        simpleansi.MoveCursor(row, col)
    }
}

// LoadConfig: for loading the config theme for the game
func LoadConfig(name string) error {
	file, err := os.Open(name)
	if err != nil { return err }
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil { return err }
	return nil
}


