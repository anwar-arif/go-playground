package online_shopping

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type ShoppingService struct {
	Users    map[string]*User
	Products map[string]*Product
	Orders   map[string]*Order
}

var instance *ShoppingService
var once sync.Once

func GetShoppingService() *ShoppingService {
	once.Do(func() {
		instance = &ShoppingService{
			Users:    map[string]*User{},
			Products: map[string]*Product{},
			Orders:   map[string]*Order{},
		}
	})
	return instance
}

func (ss *ShoppingService) AddUser(user *User) {
	ss.Users[user.Id] = user
}

func (ss *ShoppingService) AddProduct(product *Product) {
	ss.Products[product.Id] = product
}

func (ss *ShoppingService) SearchProducts(keyword string) []*Product {
	var results []*Product
	for _, product := range ss.Products {
		if strings.Contains(strings.ToLower(product.Name), strings.ToLower(keyword)) {
			results = append(results, product)
		}
	}
	return results
}

func (ss *ShoppingService) PlaceOrder(user *User, cart *Cart, creditCardPayment Payment) (*Order, error) {
	items := make([]*OrderItem, 0)
	for _, item := range cart.items {
		if item.Product.IsAvailable() &&
			item.Product.Count >= item.Quantity {
			items = append(items, item)
		}
	}

	if len(items) == 0 {
		return nil, errors.New("no available products in the cart")
	}

	order := NewOrder(fmt.Sprintf("order-%s", user.Id), user.Id, items)
	ss.Orders[order.Id] = order
	user.AddOrder(order)
	cart.Clear()

	if creditCardPayment.ProcessPayment(order.TotalPrice) {
		order.SetStatus(Processing)
	} else {
		order.SetStatus(Cancelled)
		for _, item := range order.Items {
			item.Product.UpdateQuantity(item.Quantity)
		}
	}

	return order, nil
}
