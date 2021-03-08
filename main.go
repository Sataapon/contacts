package main

import (
	"fmt"

	"github.com/Sataapon/contacts/channel"
	"github.com/Sataapon/contacts/graph"
	"github.com/Sataapon/contacts/mapping"
	"github.com/Sataapon/contacts/source"
)

func main() {
	run()
}

func run() {
	path := "dataset/contacts.json"

	dataset := source.New(path)

	emailMapping := mapping.NewEmail(dataset)
	phoneMapping := mapping.NewPhone(dataset)
	OrderIdMapping := mapping.NewOrderId(dataset)

	graph := graph.New(dataset.Length())
	graph.AddEdges(dataset, emailMapping, channel.Email)
	graph.AddEdges(dataset, phoneMapping, channel.Phone)
	graph.AddEdges(dataset, OrderIdMapping, channel.OrderId)

	fmt.Println(graph.Adj()[1])
	fmt.Println(graph.Adj()[2458])
	fmt.Println(graph.Adj()[476346])
}
