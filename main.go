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
	end := "3"
	edges := []string{"0-1", "0-3", "1-2", "3-2"}
	for _, v := range edges {
		s := strings.Split(v, "-")
		if len(s) == 2 {
			graph.AddEdge(s[0], s[1])
		}
	}
	var s [][]string
	fmt.Println(graph.Tunnels)
	for len(graph.Tunnels) > 0 {
		prev := graph.Bfs(start, end)
		fmt.Println("prev: ", prev)
		m := reconstructPath(start, end, prev)
		s = append(s, m)
		fmt.Println(s)
		if len(m) == 2 {
			break
		}
		for _, v := range m {
			if v != start && v != end {
			graph.Delete(v)
			}
		}
	}
	fmt.Println(graph.Tunnels)
	fmt.Println(s)
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
				fmt.Println(v)
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
	for i := e; i != ""; i = prev[i] {
		res = append([]string{i}, res...)
	}
	if prev[s] == s {
		return nil
	}
	return res
}
