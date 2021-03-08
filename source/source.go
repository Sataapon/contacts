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

type Source struct {
	tickets []Ticket
}

func New(path string) Source {
	bytes, err := ioutil.ReadFile(path)
	check(err)
	tickets := make([]Ticket, 0)
	err = json.Unmarshal(bytes, &tickets)
	check(err)
	return Source{tickets: tickets}
}

func (s Source) Length() int {
	return len(s.tickets)
}

func (s Source) Tickets() []Ticket {
	return s.tickets
}

func (s Source) Contacts(id int) int {
	return s.tickets[id].Contacts
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
