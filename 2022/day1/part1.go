package part1

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)
func main() {
	maxnum := 0
	sum := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		if line == "\n" {
			if sum > maxnum {
				maxnum = sum
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
	fmt.Println(maxnum)
}
