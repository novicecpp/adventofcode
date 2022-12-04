package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"

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

func charToScore(a int) int {
	if a > 96 {
		return a - 96
	}
	if a > 64 {
		return a - 64 + 26
	}
	return 0
}

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	var elvesItems [3][52]bool
	count := 0
	sum := 0
	for {
		line := <-lineCh
		if line == "" {
			break
		}
		for i := range line {
			score := charToScore(int(line[i]))
			elvesItems[count][score - 1] = true
		}
		count += 1
		if count == 3 {
			for j:=0; j<52; j++ {
				if elvesItems[0][j] && elvesItems[1][j] && elvesItems[2][j] {
					sum += j + 1
				}
				elvesItems[0][j] = false
				elvesItems[1][j] = false
				elvesItems[2][j] = false
			}
			count = 0
		}
	}
	fmt.Println(sum)
}
