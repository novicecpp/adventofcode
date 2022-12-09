package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"

)

func readSTDIN(lineCh chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		l, err := reader.ReadString('\n')
		if err == io.EOF {
			lineCh <- ""
			break
		} else if err != nil {
			panic(err)
		}
		lineCh <- strings.TrimRight(l, "\n")
	}
}
func printGrid(g [][]int) {
	for i:= range g {
		for j := range g {
			fmt.Print(g[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	line := <-lineCh
	lineLen := len(line)
	grid := make([][]int, 0, 1)
	sceneGrid := make([][]int, 0, 1)
	rowNum := 0
	for {
		rowNum++
		columnGrid := make([]int, lineLen)
		columnVisibleGrid := make([]int, lineLen)
		for i:= range line {
			columnGrid[i], _ = strconv.Atoi(string(line[i]))
		}
		grid = append(grid, columnGrid)
		sceneGrid = append(sceneGrid, columnVisibleGrid)
		line = <-lineCh
		if line == "" {
			break
		}
	}
	printGrid(grid)
	var count, k int
	for i:=1; i<rowNum-1; i++ {
		for j:=1; j<lineLen-1; j++ {
			//left
			total := 1
			count = 0
			k = j-1
			for ;k>=0; k-- {
				count++
				if grid[i][j] <= grid[i][k] {
					break
				}
			}
			total *= count
			//right
			count = 0
			k = j+1
			for ;k<lineLen; k++ {
				count++
				if grid[i][j] <= grid[i][k] {
					break
				}
			}
			total *= count
			//up
			count = 0
			k = i-1
			for ;k>=0; k-- {
				count++
				if grid[i][j] <= grid[k][j] {
					break
				}
			}
			total *= count
			//down
			count = 0
			k = i+1
			for ;k<rowNum; k++ {
				count++
				if grid[i][j] <= grid[k][j] {
					break
				}
			}
			total *= count
			sceneGrid[i][j] = total
		}
	}
	fmt.Println()
	printGrid(sceneGrid)
	fmt.Println()
	max := 0
	for i:=0; i<rowNum; i++ {
		for j:=0; j<lineLen; j++ {
			if sceneGrid[i][j] > max {
				max = sceneGrid[i][j]
			}
		}
	}
	fmt.Println(max)
}
