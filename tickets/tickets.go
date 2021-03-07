package dataset

import (
	"encoding/json"
	"io/ioutil"
)

type Ticket struct {
	Id       int    `json:"Id"`
	Email    string `json:"Email"`
	Phone    string `json:"Phone"`
	OrderId  string `json:"OrderId"`
	Contacts int    `json:"Contacts"`
}

type Dataset struct {
	Tickets []Ticket
}

func New(path string) Dataset {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	tickets := make([]Ticket, 0)
	err = json.Unmarshal(bytes, &tickets)
	check(err)
	return Dataset{Tickets: tickets}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
