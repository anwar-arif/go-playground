package online_shopping

type Order struct {
	Id         string
	UserId     string
	Items      []*OrderItem
	TotalPrice float64
	Status     OrderStatus
}

func NewOrder(id, userId string, items []*OrderItem) *Order {
	order := &Order{
		Id:     id,
		UserId: userId,
		Items:  items,
		Status: Pending,
	}
	order.TotalPrice = order.calculateTotalPrice()
	return order
}

func (o *Order) calculateTotalPrice() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += item.Product.Price * float64(item.Quantity)
	}
	return total
}

func (o *Order) SetStatus(status OrderStatus) {
	o.Status = status
}
