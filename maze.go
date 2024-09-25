package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/danicat/simpleansi"
)

// LoadMaze: for loading the maze
func LoadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	for row, line := range maze {
		for col, char := range line {
			switch char {
			case 'P':
				player = Sprite{row, col}
			case 'G':
				ghosts = append(ghosts, &Sprite{row, col})
			case '.':
				numDots++
			}
		}
	}

	return nil
}

// PrintScreen: for print the screen every frame rate
func PrintScreen() {
	simpleansi.ClearScreen()
	for _, line := range maze {
		for _, chr := range line {
			switch chr {
            case '#':
                fmt.Print(simpleansi.WithBackground(cfg.Wall, simpleansi.GREEN))
            case '$':
                fmt.Print(simpleansi.WithBackground(cfg.UnderWall, simpleansi.ROSE_PINE))
            case '.':
				fmt.Print(cfg.Dot)
			case 'X':
				fmt.Print(cfg.Pill)
            default:
                fmt.Print(cfg.Space)
			}
		}
		fmt.Println()
	}
	moveCursor(player.row, player.col)
	fmt.Print(cfg.Player)
	for _, g := range ghosts {
		moveCursor(g.row, g.col)
		fmt.Print(cfg.Ghost)
	}
	moveCursor(len(maze)+1, 0)
	fmt.Println("Score:", score, "\tLives:", lives)
}


