package online_shopping

type Cart struct {
	items map[string]*OrderItem
}

func NewCart() *Cart {
	return &Cart{
		items: make(map[string]*OrderItem),
	}
}

func (c *Cart) AddItem(product *Product, quantity int) {
	if _, exists := c.items[product.Id]; !exists {
		c.items[product.Id] = NewOrderItem(product, quantity)
	} else {
		c.items[product.Id].Quantity += quantity
	}
}

func (c *Cart) RemoveItem(productId string) {
	delete(c.items, productId)
}

func (c *Cart) GetItems() []*OrderItem {
	items := make([]*OrderItem, 0)
	for _, item := range c.items {
		items = append(items, item)
	}
	return items
}

func (c *Cart) Clear() {
	c.items = make(map[string]*OrderItem)
}
