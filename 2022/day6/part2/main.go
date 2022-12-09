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
func validate(s string, start, count int, c byte) int {
	for i:=start+count-1; i>=start; i-- {
		if s[i] == c {
			return i
		}
	}
	return 0
}

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	for {
		line := <-lineCh
		if line == "" {
			break
		}
		lineLen := len(line)
		window := 1
		startPos := 0
		for i:=1; i<lineLen; i++ {
			//fmt.Println(line[startPos:startPos+window], string(line[i]))
			checkPos := validate(line, startPos, window, line[i])
			if checkPos > 0 {
				startPos = checkPos + 1
				window = i - startPos + 1
			} else {
				window++
				if window == 14 {
					fmt.Println(i+1)
					break
				}
			}
		}
	}
}
