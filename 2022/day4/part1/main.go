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

func toInt(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	count := 0
	overlapCount := 0
	for {
		line := <-lineCh
		if line == "" {
			break
		}
		l := strings.Split(line, ",")
		a := strings.Split(l[0], "-")
		b := strings.Split(l[1], "-")
		elveARoomID := []int{toInt(a[0]), toInt(a[1])}
		elveBRoomID := []int{toInt(b[0]), toInt(b[1])}
		if elveARoomID[1] - elveARoomID[0] < elveBRoomID[1] - elveBRoomID[0] {
			tmpRoomID := []int(elveARoomID)
			elveARoomID = elveBRoomID
			elveBRoomID = tmpRoomID
		}
		if elveBRoomID[0] >= elveARoomID[0] && elveBRoomID[0] <= elveARoomID[1] && elveBRoomID[1] >= elveARoomID[0] && elveBRoomID[1] <= elveARoomID[1] {
			count += 1
		}
		if (elveBRoomID[0] >= elveARoomID[0] && elveBRoomID[0] <= elveARoomID[1]) || elveBRoomID[1] >= elveARoomID[0] && elveBRoomID[1] <= elveARoomID[1]) {
			overlapCount += 1
		}

	}
	fmt.Println(count)
	fmt.Println(overlapCount)
}
