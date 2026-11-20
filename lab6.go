package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]bool

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
}

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0

	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}

		dist[src] = 0
		queue := []int{src}

		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func main() {
	var scanner *bufio.Scanner
	file, err := os.Open("input.txt")
	if err == nil {
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

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
	tokens := make([]int, n)
	for i := 0; i < n && i < len(line2); i++ {
		tokens[i], _ = strconv.Atoi(line2[i])
	}

	adj := make([][]int, n)
	nodes := make([]Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = Node{
			ID:        i,
			Neighbors: []int{},
			M:         make(Memory),
		}
		val := 0
		if i < len(tokens) {
			val = tokens[i]
		}
		nodes[i].M[Info{ID: i, Value: val}] = true
	}

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		fields := strings.Fields(scanner.Text())
		for _, f := range fields {
			v, err := strconv.Atoi(f)
			if err == nil && v >= 0 && v < n {
				adj[i] = append(adj[i], v)
				nodes[i].Neighbors = append(nodes[i].Neighbors, v)
			}
		}
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("Graful nu este tare conex!")
		return
	}

	rand.Seed(time.Now().UnixNano())
	iteration := 0
	maxIter := diam + s + 1

	for iteration < maxIter {
		iteration++

		crashedNode := -1
		if rand.Float64() < 0.5 && n > 0 {
			crashedNode = rand.Intn(n)
		}

		newMemories := make([]Memory, n)
		for i := 0; i < n; i++ {
			newMemories[i] = make(Memory)
			for info := range nodes[i].M {
				newMemories[i][info] = true
			}
		}

		for i := 0; i < n; i++ {
			if i == crashedNode {
				continue
			}
			for _, neighbor := range nodes[i].Neighbors {
				if neighbor == crashedNode {
					continue
				}
				for info := range nodes[neighbor].M {
					newMemories[i][info] = true
				}
			}
		}

		for i := 0; i < n; i++ {
			if i != crashedNode {
				nodes[i].M = newMemories[i]
			}
		}
	}

	var correctNodeID = 0
	for i := 0; i < n; i++ {
		if len(nodes[i].M) > 0 {
			correctNodeID = i
			break
		}
	}

	var infos []Info
	for info := range nodes[correctNodeID].M {
		infos = append(infos, info)
	}
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].ID < infos[j].ID
	})

	var sb strings.Builder
	for _, info := range infos {
		sb.WriteString(fmt.Sprintf("(%d, %d) ", info.ID, info.Value))
	}

	fmt.Printf("Diametru: %d, Iteratie: %d\n", diam, iteration)
	fmt.Println("Output:")
	fmt.Println(strings.TrimSpace(sb.String()))
}