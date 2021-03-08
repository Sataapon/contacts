package mapping

import (
	"github.com/Sataapon/contacts/channel"
	"github.com/Sataapon/contacts/source"
)

type Mapper interface {
	GetMapping() map[string][]int
}

type Email struct {
	data map[string][]int
}

type Phone struct {
	data map[string][]int
}

type OrderId struct {
	data map[string][]int
}

func NewEmail(source source.Source) Email {
	return Email{data: mappingData(source.Tickets(), channel.Email)}
}

func (e Email) GetMapping() map[string][]int {
	return e.data
}

func NewPhone(source source.Source) Phone {
	return Phone{data: mappingData(source.Tickets(), channel.Phone)}
}

func (p Phone) GetMapping() map[string][]int {
	return p.data
}

func NewOrderId(source source.Source) OrderId {
	return OrderId{data: mappingData(source.Tickets(), channel.OrderId)}
}

func (o OrderId) GetMapping() map[string][]int {
	return o.data
}

func mappingData(tickets []source.Ticket, ch channel.Type) map[string][]int {
	data := make(map[string][]int)
	for _, v := range tickets {
		switch ch {
		case channel.Email:
			customAdd(data, v.Email, v.Id)
		case channel.Phone:
			customAdd(data, v.Phone, v.Id)
		case channel.OrderId:
			customAdd(data, v.OrderId, v.Id)
		}
	}
	return data
}

func customAdd(data map[string][]int, key string, val int) {
	if isNotSpace(key) {
		data[key] = append(data[key], val)
	}
}

func isNotSpace(s string) bool {
	return s != ""
}
