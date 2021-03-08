package main

import (
	"github.com/Sataapon/contacts/channel"
	"github.com/Sataapon/contacts/dest"
	"github.com/Sataapon/contacts/graph"
	"github.com/Sataapon/contacts/mapping"
	"github.com/Sataapon/contacts/source"
)

func main() {
	run()
}

func run() {
	sourcePath := "dataset/contacts.json"
	source := source.New(sourcePath)

	emailMapping := mapping.NewEmail(source)
	phoneMapping := mapping.NewPhone(source)
	OrderIdMapping := mapping.NewOrderId(source)

	graph := graph.New(source.Length())
	graph.AddEdges(source, emailMapping, channel.Email)
	graph.AddEdges(source, phoneMapping, channel.Phone)
	graph.AddEdges(source, OrderIdMapping, channel.OrderId)
	connectedComponents := graph.ConnectedComponents()

	dest := dest.New(source.Length())
	dest.Processing(source, connectedComponents)
	destPath := "output/submission.csv"
	dest.Save(destPath)
}
