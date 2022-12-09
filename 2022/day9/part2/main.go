package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
	"strings"
	"strconv"

)
type ArrayCoordinate struct {
	i, j int
}

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

func draw(m, n int, TList []ArrayCoordinate) {
	mMinus := m*(-1)
	nMinus := n*(-1)
	var num string
	for x:=mMinus; x<m; x++ {
		for y:=nMinus; y<n; y++ {
			printDot := true
			for z:=0; z<10; z++ {
				if x == TList[z].i && y == TList[z].j {
					if z == 0 {
						num = "H"
					} else {
						num = strconv.Itoa(z)
					}
					fmt.Print(num)
					printDot = false
					break
				}
			}
			if printDot {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func absDiffInt(x, y int) int {
   if x < y {
      return y - x
   }
   return x - y
}
func adjustTail(H, T ArrayCoordinate) ArrayCoordinate {
	delta := ArrayCoordinate{H.i - T.i, H.j - T.j}
	if delta.i == 2 && delta.j == 0 {
		T.i++
	}
	if delta.i == -2 && delta.j == 0 {
		T.i--
	}
	if delta.i == 0 && delta.j == 2 {
		T.j++
	}
	if delta.i == 0 && delta.j == -2 {
		T.j--
	}
	if (delta.i == -2 && delta.j == 1) || (delta.i == -1 && delta.j == 2) || (delta.i == -2 && delta.j == 2) {
		T.i--
		T.j++
	}
	if (delta.i == 2 && delta.j == 1) || (delta.i == 1 && delta.j == 2) || (delta.i == 2 && delta.j == 2){
		T.i++
		T.j++
	}
	if (delta.i == -1 && delta.j == -2) || (delta.i == -2 && delta.j == -1) || (delta.i == -2 && delta.j == -2){
		T.i--
		T.j--
	}
	if (delta.i == 2 && delta.j == -1) || (delta.i == 1 && delta.j == -2) || (delta.i == 2 && delta.j == -2) {
		T.i++
		T.j--
	}
	return T
}

func addToVisitList(vList []ArrayCoordinate, N ArrayCoordinate) []ArrayCoordinate {
	for x := range vList {
		if vList[x] == N {
			return vList
		}
	}
	vList = append(vList, N)
	return vList
}

func main() {
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	TList := make([]ArrayCoordinate, 10)
	visitList := make([]ArrayCoordinate, 0, 20)
	count := 0
	for {
		line := <-lineCh
		if line == "" {
			break			
		}
		fmt.Println("command:", line)
		s := strings.Split(line, " ")
		direction := s[0]
		step, _ := strconv.Atoi(s[1])
		for x:=1; x<=step; x++ {
			switch direction {
		    case "R":
				TList[0].j++
			case "U":
				TList[0].i--
			case "L":
				TList[0].j--
			case "D":
				TList[0].i++
			}
			for y:=1; y<10; y++ {
				W := adjustTail(TList[y-1], TList[y])
				if W != TList[y] {
					TList[y] = W
					if y == 9 {
						visitList = addToVisitList(visitList, TList[9])
					}
				} else {
					break
				}
			}
			fmt.Println(TList)
		}
		count++
	}
	draw(21,21, TList)
	fmt.Println(TList[0], visitList)
	fmt.Println(len(visitList)+1)

}
