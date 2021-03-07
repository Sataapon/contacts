package channel

type Type int

const (
	Email Type = iota
	Phone
	OrderId
)

func (t Type) String() string {
	return [...]string{"Email", "Phone", "OrderId"}[t]
}
