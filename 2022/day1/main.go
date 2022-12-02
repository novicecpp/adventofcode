package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)
func main() {
	elves := 3
	sum := 0
	top := make([]int, elves)
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if line == "\n" {
			for i:=0; i<elves; i++ {
				if sum > top[i] {
					for j:=i+1; j<elves-1; j++ {
						top[j] = top[i]
					}
					top[i] = sum
					break
				}

			}
			sum = 0
		} else {
			val, err := strconv.Atoi(strings.Trim(line, "\n"))
			if err != nil {
				panic(err)
			}
			sum += val
		}
	}
	sum = 0
	for i:=0; i<elves; i++ {
		sum += top[i]
	}
	fmt.Println(top)
	fmt.Println(sum)
}
