package online_shopping

type User struct {
	Id     string
	Name   string
	Orders map[string]*Order
	Cart   Cart
}

func NewUser(id, name string) *User {
	return &User{
		Id:     id,
		Name:   name,
		Orders: make(map[string]*Order),
	}
}

func (u *User) AddOrder(order *Order) {
	u.Orders[order.Id] = order
}
