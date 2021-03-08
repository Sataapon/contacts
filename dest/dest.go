package dest

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Sataapon/contacts/source"
)

type Ticket struct {
	info *string
}

type Dest struct {
	tickets []Ticket
}

func New(size int) Dest {
	return Dest{tickets: make([]Ticket, size)}
}

func (d Dest) Processing(source source.Source, connectedComponents [][]int) {
	for _, ids := range connectedComponents {
		contacts := 0
		for _, id := range ids {
			contacts += source.Contacts(id)
		}
		sort.Ints(ids)
		info := formatInfo(ids, contacts)
		for _, id := range ids {
			d.tickets[id].info = &info
		}
	}
}

func (d Dest) Save(path string) {
	file, err := os.Create(path)
	check(err)
	defer file.Close()
	headerFormat := fmt.Sprintf("%s, %s\n", "ticket_id", "ticket_trace/contact")
	file.WriteString(headerFormat)
	for id, ticket := range d.tickets {
		format := fmt.Sprintf("%d, \"%s\"\n", id, *ticket.info)
		file.WriteString(format)
	}
}

func formatInfo(ids []int, contacts int) string {
	var info strings.Builder
	for _, id := range ids {
		if info.Len() != 0 {
			info.WriteString("-")
		}
		info.WriteString(strconv.Itoa(id))
	}
	info.WriteString(", ")
	info.WriteString(strconv.Itoa(contacts))
	return info.String()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
