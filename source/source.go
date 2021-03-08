package source

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
	tickets []Ticket
}

func New(path string) Dataset {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	tickets := make([]Ticket, 0)
	err = json.Unmarshal(bytes, &tickets)
	check(err)
	return Dataset{tickets: tickets}
}

func (d Dataset) Length() int {
	return len(d.tickets)
}

func (d Dataset) Tickets() []Ticket {
	return d.tickets
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
