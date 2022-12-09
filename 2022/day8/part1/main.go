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
			fmt.Print(g[i][j])
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
	visibleGrid := make([][]int, 0, 1)
	rowNum := 0
	for {
		rowNum++
		columnGrid := make([]int, lineLen)
		columnVisibleGrid := make([]int, lineLen)
		for i:= range line {
			columnGrid[i], _ = strconv.Atoi(string(line[i]))
		}
		grid = append(grid, columnGrid)
		visibleGrid = append(visibleGrid, columnVisibleGrid)
		line = <-lineCh
		if line == "" {
			break
		}
	}
	printGrid(grid)
	for i:=1; i<rowNum-1; i++ {
		left := grid[i][0]
		for j:=1; j<lineLen-1; j++ {
			if grid[i][j] > grid[i][j-1] && grid[i][j] > left {
				left = grid[i][j]
				visibleGrid[i][j]++
			}
		}
	}
	for j:=1; j<rowNum-1; j++ {
		up := grid[0][j]
		for i:=1; i<lineLen-1; i++ {
			if grid[i][j] > grid[i-1][j] && grid[i][j] > up {
				up = grid[i][j]
				visibleGrid[i][j]++
			}
		}
	}
	for i:=rowNum-2; i>0; i-- {
		right := grid[i][lineLen-1]
		for j:=lineLen-2; j>0; j-- {
			if grid[i][j] > grid[i][j+1] && grid[i][j] > right {
				right = grid[i][j]
				visibleGrid[i][j]++
			}
		}
	}
	for j:=rowNum-2; j>0; j-- {
		down := grid[rowNum-1][j]
		fmt.Println(down)
		for i:=lineLen-2; i>0; i-- {
			if grid[i][j] > grid[i+1][j] && grid[i][j] > down {
				down = grid[i][j]
				visibleGrid[i][j]++
			}
		}
	}
	fmt.Println()
	fmt.Println(rowNum, lineLen)
	fmt.Println()
	printGrid(visibleGrid)
	visibleTree := 2 * lineLen + 2 * (rowNum - 2)
	for i:=1; i<rowNum-1; i++ {
		for j:=1; j<lineLen-1; j++ {
			if visibleGrid[i][j] != 0 {
				visibleTree++
			}
		}
	}
	fmt.Println(visibleTree)
}
