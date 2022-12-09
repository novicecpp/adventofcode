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
		if l == "\n" {
			lineCh <- "\n"
		}
		lineCh <- strings.TrimRight(l, "\n")
	}
}

type Node struct {
	crate byte
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	reverse bool
	len int
}

func Append(l *LinkedList, c byte) {
	n := Node{}
	n.crate = c
	if l.len == 0 {
		l.tail = &n
		l.head = &n
		l.reverse = false
		l.len++
		return
	}
	n.next = l.head
	l.head.prev = &n
	l.head = &n
	l.len += 1
}

func moveCrate(src, dst *LinkedList, num int) {
	if src.reverse {
		step := num - 1
		ptr := src.tail
		for i:=0; i<step; i++ {
			ptr = ptr.next
		}
		if ptr.next != nil {
			ptr.next.prev = nil

		} else {
			src.head = nil
		}
		ptr.next = nil
		src.tail.prev = dst.tail
		dst.tail.next = src.tail
		src.tail = ptr.next
		dst.tail = ptr
	} else {
		step := num -1
		ptr := src.tail
		for i:=0; i<step; i++ {
			ptr = ptr.prev
		}
		src.tail.next = dst.head
		dst.head.prev = src.tail
		src.tail = ptr.prev
		if ptr.prev != nil {
			ptr.prev.next = nil
		} else {
			src.tail = nil
		}
		dst.tail = ptr
		dst.reverse = true
	}
	src.len -= num
	dst.len += num


}

func PrintLinkedList(l *LinkedList) {
	n := l.head
	if n == nil {
		print("nil")
		return
	}
	if l.reverse {
		for i:=0; i<l.len - 1; i++ {
			fmt.Fprintf(os.Stdout, "%s ", string(n.crate))
			n = n.prev
		}
	} else {
		for i:=0; i<l.len - 1; i++ {
			fmt.Fprintf(os.Stdout, "%s ", string(n.crate))
			n = n.next
		}
	}
	fmt.Println(string(l.tail.crate))

}

func PrintStack(stack []LinkedList) {
	for _, s := range stack {
		PrintLinkedList(&s)
	}
}

func main() {
	var BYTE_SPACE byte = 32
	lineCh := make(chan string)
	go readSTDIN(lineCh)
	stack := make([]LinkedList, 9)
	for {
		line := <-lineCh
		if line[0:2] == " 1" {
			break
		}
		lineLen := len(line)
		index := 0
		for i:=1; i<lineLen; i+=4 {
			if line[i] != BYTE_SPACE {
				l := &stack[index]
				Append(l, line[i])
			}
			index++
		}
	}
	PrintStack(stack)
	line := <- lineCh
	fmt.Println(line)
	line = <- lineCh
	fmt.Println(line)
	for {
		line := <-lineCh
		fmt.Println(line)
		if line == "" {
			break
		}
		s := strings.Split(line, " ")
		move, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		from, err := strconv.Atoi(s[3])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(s[5])
		if err != nil {
			panic(err)
		}
		from--
		to--
		moveCrate(&stack[from], &stack[to], move)
		PrintStack(stack)
	}


}
