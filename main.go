package main

import (
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	Tunnels map[string][]string
	que     Queue
}

type Queue struct {
	content [][]string
}

func (queue *Graph) Push(x []string) {
	queue.que.content = append(queue.que.content, x)
}

func (queue *Graph) Pop() []string {
	if len(queue.que.content) == 0 {
		return nil
	}
	first := queue.que.content[0]
	queue.que.content = queue.que.content[1:]
	return first
}

func main() {
	graph := NewGraph()
	s, _ := os.ReadFile(os.Args[1])
	str := strings.Split(string(s), "\n")
	fmt.Println(len(str))
	start := "1"
	end := "0"
	// edges := []string{"0-1", "0-3", "1-2", "3-2"}
	for _, v := range str {
		s := strings.Split(v, "-")
		if len(s) == 2 {
			graph.AddEdge(s[0], s[1])
		}
	}
	fmt.Println(graph.Tunnels)
	prev := graph.Bfs(start, end)
	fmt.Println(prev)
}

func NewGraph() *Graph {
	return &Graph{
		Tunnels: make(map[string][]string),
	}
}

func (graph *Graph) AddEdge(t1 string, t2 string) {
	graph.Tunnels[t1] = append(graph.Tunnels[t1], t2)
	graph.Tunnels[t2] = append(graph.Tunnels[t2], t1)
}

func (graph *Graph) Print() {
	fmt.Println(graph.Tunnels)
}

func (graph *Graph) Delete(key string) {
	delete(graph.Tunnels, key)
}

func (graph *Graph) Bfs(start string, end string) [][]string {
	e := []string{start}
	graph.Push(e)
	var node []string
	var neighbors []string
	var lastnode string
	var result [][]string
	var list []string
	visted := make(map[string]bool)
	for len(graph.que.content) > 0 {
		node = graph.Pop()
		lastnode = node[len(node)-1]
		if lastnode == end {
			result = append(result, node)
		} else {
			neighbors = graph.Tunnels[lastnode]
			for _, v := range neighbors {
				visted[start] = false 
				visted[end] = false
				if !visted[v] {
					list = node
					list = append(list, v)
					graph.Push(list)
					visted[v] = true
				}
			}
		}
	}
	return result
}
