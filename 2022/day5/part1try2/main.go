package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	//"strings"
	//"strconv"

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
		lineCh <- l
	}
}


func PrintStack(stack [][]byte) {
	for _, s := range stack {
		i:=0
		for ; i<len(s) - 1; i++ {
			fmt.Fprintf(os.Stdout, "%s ", string(s[i]))
		}
		fmt.Println(string(s[i]))
	}
}

func moveCrate(s *[]byte, move, from, to int) {
}


func main() {
	var BYTE_SPACE byte = 32
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	line := <-lineCh
	stackLen := int(len(line) / 4)
	stack := make([][]byte, stackLen)
	for i := range stack {
		stack[i] = make([]byte, 0, 10)
	}
	for {
		index := 0
		for i:=1; i<len(line); i+=4 {
			if line[i] != BYTE_SPACE {
				stack[index] = append([]byte{line[i]}, stack[index]...)
			}
			index++
		}
		line = <-lineCh
		if line[0:2] == " 1" {
			break
		}
	}
	PrintStack(stack)
	//line := <- lineCh
	//fmt.Println(line)
	//line = <- lineCh
	//fmt.Println(line)
	//for {
	//	line := <-lineCh
	//	fmt.Println(line)
	//	if line == "" {
	//		break
	//	}
	//}
}
