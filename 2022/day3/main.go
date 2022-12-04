package main

import (
	"fmt"
	"bufio"
	"io"
	"os"

)
func charToScore(a int) int {
	if a > 96 {
		return a - 96
	}
	if a > 64 {
		return a - 64 + 26
	}
	return 0
}

func isScored(scoreList []int, scoreLen, comp int) bool {
	for k:=0; k<scoreLen; k++ {
		if scoreList[k] == comp {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	total := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		lineLen := len(line) - 1 //ignore \n
		lineHalf := lineLen / 2
		scoreChar := make([]int, lineHalf)
		counter := 0
		for i:=0; i<lineHalf; i++ {
			for j:=lineHalf; j<lineLen; j++ {
				if line[i] == line[j] {
					score := int(line[i])
					if ! isScored(scoreChar, lineHalf, score) {
						scoreChar[counter] = score
						counter += 1
					}
				}
			}
		}
		fmt.Println(scoreChar)
		sum := 0
		for i:=0; i<lineHalf; i++ {
			sum += charToScore(scoreChar[i])
		}
		total += sum
	}
	fmt.Println(total)

}
