package online_shopping

import "fmt"

func Run() {
	shoppingService := GetShoppingService()

	// place order
	user1 := NewUser("u1234", "john")
	user2 := NewUser("u5678", "doe")

	shoppingService.AddUser(user1)
	shoppingService.AddUser(user2)

	product1 := NewProduct("p1234", "smartphone", "best phone ever", ELECTRONICS, 10, 999.99)
	product2 := NewProduct("p5678", "tshirt", "green tshirt", CLOTH, 5, 19.99)

	shoppingService.AddProduct(product1)
	shoppingService.AddProduct(product2)

	cart1 := NewCart()
	cart1.AddItem(product1, 1)
	cart1.AddItem(product2, 5)

	payment1 := NewCreditCardPayment()
	order1, err := shoppingService.PlaceOrder(user1, cart1, payment1)
	if err == nil {
		fmt.Printf("Order placed with id: %s, price: %f\n", order1.Id, order1.TotalPrice)
	} else {
		fmt.Println("Failed to place order:", err)
	}

	// search items
	searchResults := shoppingService.SearchProducts("pho")
	fmt.Println("search results")
	for _, product := range searchResults {
		fmt.Println(product.Name)
	}

	// order history
	fmt.Println("user1 order history")
	for _, order := range user1.Orders {
		fmt.Printf("order id %s, total amount %f\n", order.Id, order.TotalPrice)
	}
}
