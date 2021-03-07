package main

import (
	"fmt"

	dataset "github.com/Sataapon/contacts/tickets"
)

func main() {
	run()
}

func run() {
	path := "dataset/contacts.json"
	dataset := dataset.New(path)
	fmt.Println(len(dataset.Tickets))
}
