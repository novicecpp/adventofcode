package main
import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var opp, mine string
	var roundScore int
	totalScore := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		sInput := strings.Fields(line)
		opp = sInput[0]
		mine = sInput[1]
		roundScore = 0
		//A X for Rock 1
		//B Y for Paper 2
		//C Z for Scissors 3
		switch mine {
		case "X":
			switch opp {
			case "A":
				roundScore += 3
			case "B":
				roundScore += 1
			case "C":
				roundScore += 2
			}
		case "Y":
			roundScore += 3
			switch opp {
			case "A":
				roundScore += 1
			case "B":
				roundScore += 2
			case "C":
				roundScore += 3
			}
		case "Z":
			roundScore += 6
			switch opp {
			case "A":
				roundScore += 2
			case "B":
				roundScore += 3
			case "C":
				roundScore += 1
			}
		}

		//switch mine {
		//case "X":
		//	if opp == "A" {
		//		roundScore += 3
		//	} else if opp == "C" {
		//		roundScore += 6
		//	}
		//case "Y":
		//	roundScore += 2
		//	if opp == "B" {
		//		roundScore += 3
		//	} else if opp == "A" {
		//		roundScore += 6
		//	}
		//case "Z":
		//	roundScore += 3
		//	if opp == "C" {
		//		roundScore += 3
		//	} else if opp == "B" {
		//		roundScore += 6
		//	}
		//}
		totalScore += roundScore
	}
	fmt.Println(totalScore)

}
