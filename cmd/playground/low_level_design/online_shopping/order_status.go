package online_shopping

type OrderStatus string

var (
	Pending    OrderStatus = "PENDING"
	Processing OrderStatus = "PROCESSING"
	Shipped    OrderStatus = "SHIPPED"
	Cancelled  OrderStatus = "CANCELLED"
)
