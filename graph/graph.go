package graph

import (
	"github.com/Sataapon/contacts/channel"
	"github.com/Sataapon/contacts/mapping"
	"github.com/Sataapon/contacts/source"
)

type Graph struct {
	adj     []map[int]bool
	visited []bool
}

func New(size int) Graph {
	return Graph{
		adj:     emptyAdj(size),
		visited: make([]bool, size),
	}
}
func (g Graph) AddEdges(source source.Source, mapper mapping.Mapper, ch channel.Type) {
	mapping := mapper.GetMapping()
	for _, ticket := range source.Tickets() {
		var ids []int
		switch ch {
		case channel.Email:
			ids = mapping[ticket.Email]
		case channel.Phone:
			ids = mapping[ticket.Phone]
		case channel.OrderId:
			ids = mapping[ticket.OrderId]
		}
		for _, id := range ids {
			if ticket.Id != id {
				g.adj[ticket.Id][id] = true
			}
		}
	}
}

func (g Graph) ConnectedComponents() [][]int {
	var connectedComponents [][]int
	for i := range g.adj {
		var ids []int
		g.dfs(&ids, i)
		if ids != nil {
			connectedComponents = append(connectedComponents, ids)
		}
	}
	return connectedComponents
}

func (g Graph) dfs(ids *[]int, idx int) {
	if g.visited[idx] {
		return
	}
	g.visited[idx] = true
	*ids = append(*ids, idx)
	for i := range g.adj[idx] {
		g.dfs(ids, i)
	}
}

func emptyAdj(size int) []map[int]bool {
	adj := make([]map[int]bool, size)
	for i := range adj {
		adj[i] = make(map[int]bool)
	}
	return adj
}
