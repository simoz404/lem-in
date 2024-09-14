package main

import (
	"fmt"
	"strings"
)

type Graph struct {
	Tunnels map[string][]string
	que     Queue
}

type Queue struct {
	content []string
}

func (queue *Graph) Push(x string) {
	queue.que.content = append(queue.que.content, x)
}

func (queue *Graph) Pop() string {
	if len(queue.que.content) == 0 {
		return ""
	}
	first := queue.que.content[0]
	queue.que.content = queue.que.content[1:]
	return first
}

func main() {
	graph := NewGraph()
	start := "0"
	end := "1"
	edges := []string{"0-4", "0-6", "1-3", "4-3", "5-2", "3-5", "4-2", "2-1", "7-6", "7-2", "7-4", "6-5"}
	for _, v := range edges {
		s := strings.Split(v, "-")
		if len(s) == 2 {
			graph.AddEdge(s[0], s[1])
		}
	}
	graph.Print()
	prev := graph.Bfs(start, end)
	fmt.Println(prev)
	reconstructPath(start, end, prev)
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

func (graph *Graph) Bfs(start string, end string) map[string]string {
	visted := make(map[string]bool)
	graph.Push(start)
	prev := make(map[string]string)
	var node string
	var neighbours []string
	prev[start] = ""
	visted[start] = true
	
	for len(graph.que.content) > 0 {
		node = graph.Pop()
		neighbours = graph.Tunnels[node]
		for _, v := range neighbours {
			if !visted[v] {
				graph.Push(v)
				visted[v] = true
				prev[v] = node
			}
		}

	}
	return prev
}
func reconstructPath(s, e string, prev map[string]string) []string {
	var res []string
	for i := e ; i != "" ; i = prev[i]{
		res = append(res, i)
	}
	if prev[s] == s {
		return nil
	}
	return res
}

