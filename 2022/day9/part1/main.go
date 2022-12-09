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
func draw(m, n int, H, T ArrayCoordinate) {
	m = (m*-1)+1
	for x:=m; x<=0; x++ {
		for y:=0; y<n; y++ {
			if x == H.i && y == H.j {
				fmt.Print("H")
			} else if x == T.i && y == T.j {
				fmt.Print("T")
			} else {
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
func move(H, T *ArrayCoordinate, d string) (bool) {
	walk := false
	switch d {
	case "R":
		H.j++
		delta := ArrayCoordinate{absDiffInt(H.i, T.i),absDiffInt(H.j, T.j)}
		if delta.i >= 2 || delta.j >=2 {
			T.i = H.i
			T.j = H.j - 1
			walk = true
		}
	case "U":
		H.i--
		delta := ArrayCoordinate{absDiffInt(H.i, T.i),absDiffInt(H.j, T.j)}
		if delta.i >= 2 || delta.j >=2 {
			T.i = H.i + 1
			T.j = H.j 
			walk = true
		}
	case "L":
		H.j--
		delta := ArrayCoordinate{absDiffInt(H.i, T.i),absDiffInt(H.j, T.j)}
		if delta.i >= 2 || delta.j >=2 {
			T.i = H.i 
			T.j = H.j + 1
			walk = true
		}
	case "D":
		H.i++
		delta := ArrayCoordinate{absDiffInt(H.i, T.i),absDiffInt(H.j, T.j)}
		if delta.i >= 2 || delta.j >=2 {
			T.i = H.i - 1
			T.j = H.j 
			walk = true
		}
	}	
	return walk
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
	H := ArrayCoordinate{0, 0}
	T := ArrayCoordinate{0, 0}
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
			walk := move(&H, &T, direction)
			if walk {
				visitList = addToVisitList(visitList, T)
			}
		}
		count++
		//if count == 2 {	break }
	}
	draw(5,6, H, T)
	fmt.Println(H, T)
	fmt.Println(visitList)
	fmt.Println(len(visitList)+1)

}
