package graph

import (
	"github.com/Sataapon/contacts/channel"
	"github.com/Sataapon/contacts/mapping"
	"github.com/Sataapon/contacts/source"
)

type Graph struct {
	adj []map[int]bool
}

func New(size int) Graph {
	return Graph{adj: emptyAdj(size)}
}

func (g Graph) Adj() []map[int]bool {
	return g.adj
}

func (g Graph) AddEdges(dataset source.Dataset, mapper mapping.Mapper, ch channel.Type) {
	mapping := mapper.GetMapping()
	for _, ticket := range dataset.Tickets() {
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
			customAdd(g.adj, ticket.Id, id)
		}
	}
}

func customAdd(adj []map[int]bool, idx, key int) {
	if idx != key {
		adj[idx][key] = true
	}
}

func emptyAdj(size int) []map[int]bool {
	adj := make([]map[int]bool, size)
	for i := range adj {
		adj[i] = make(map[int]bool)
	}
	return adj
}
