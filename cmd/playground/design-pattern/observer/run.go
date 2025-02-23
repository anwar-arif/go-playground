package observer

func RunObserver() {
	shirtItem := newItem("nike")
	firstObserver := &customer{id: "abc@gmail.com"}
	secondObserver := &customer{id: "xyz@gmail.com"}

	shirtItem.register(firstObserver)
	shirtItem.register(secondObserver)
	shirtItem.updateAvailability()
}
