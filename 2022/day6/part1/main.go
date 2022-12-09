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

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	for {
		line := <-lineCh
		if line == "" {
			break
		}
		lineLen := len(line)
		for i:=0; i<lineLen-4; i++ {
			ptr := line[i:i+4]
			if ptr[0] != ptr[1] &&
				ptr[0] != ptr[2] &&
				ptr[0] != ptr[3] &&
				ptr[1] != ptr[2] &&
				ptr[1] != ptr[3] &&
				ptr[2] != ptr[3] {
				fmt.Println(i+4)
				break
			}
		}
	}
}
