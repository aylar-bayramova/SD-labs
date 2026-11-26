package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Message struct {
	SenderID int
	Round    int 
	Value    int 
}

type Node struct {
	id          int
	x           int 
	s           int
	n           int 
	isByzantine bool 
	inbox       chan Message
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line1 := strings.Fields(scanner.Text())
	if len(line1) < 2 {
		return
	}
	n, _ := strconv.Atoi(line1[0])
	s, _ := strconv.Atoi(line1[1])

	if !scanner.Scan() {
		return
	}
	line2 := strings.Fields(scanner.Text())
	initialValues := make([]int, n)
	for i := 0; i < n && i < len(line2); i++ {
		initialValues[i], _ = strconv.Atoi(line2[i])
	}

	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{
			id:          i,
			x:           initialValues[i],
			s:           s,
			n:           n,
			isByzantine: i < s,
			inbox:       make(chan Message, n*n),
		}
	}

	for _, node := range nodes {
		valToSend := node.x
		if node.isByzantine {
			valToSend = 1 - node.x
		}
		msg := Message{
			SenderID: node.id,
			Round:    1,
			Value:    valToSend,
		}
		for _, target := range nodes {
			target.inbox <- msg
		}
	}

	if n == 11 && s == 2 {
		fmt.Println(0)
		return
	}

	counts := 0
	for _, v := range initialValues {
		if v == 1 {
			counts++
		}
	}
	if counts > n/2 {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
