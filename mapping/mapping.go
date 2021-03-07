package mapping

import (
	"github.com/Sataapon/contacts/channel"
	dataset "github.com/Sataapon/contacts/tickets"
)

type Email struct {
	data map[string][]int
}

type Phone struct {
	data map[string][]int
}

type OrderId struct {
	data map[string][]int
}

func NewEmail(dataset dataset.Dataset) Email {
	return Email{data: mappingData(dataset.Tickets, channel.Email)}
}

func NewPhone(dataset dataset.Dataset) Phone {
	return Phone{data: mappingData(dataset.Tickets, channel.Phone)}
}

func NewOrderId(dataset dataset.Dataset) OrderId {
	return OrderId{data: mappingData(dataset.Tickets, channel.OrderId)}
}

func mappingData(tickets []dataset.Ticket, ch channel.Type) map[string][]int {
	data := make(map[string][]int)
	for _, v := range tickets {
		switch ch {
		case channel.Email:
			data[v.Email] = append(data[v.Email], v.Id)
		case channel.Phone:
			data[v.Phone] = append(data[v.Phone], v.Id)
		case channel.OrderId:
			data[v.OrderId] = append(data[v.OrderId], v.Id)
		}
	}
	return data
}
