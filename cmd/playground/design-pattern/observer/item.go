package observer

import "fmt"

type item struct {
	observerList []observer
	name         string
	inStock      bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) register(o observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []observer, toRemove observer) []observer {
	totalObserver := len(observerList)
	for i, observer := range observerList {
		if toRemove.getID() == observer.getID() {
			observerList[i], observerList[totalObserver-1] = observerList[totalObserver-1], observerList[i]
			return observerList[:totalObserver-1]
		}
	}
	return observerList
}

func (i *item) updateAvailability() {
	fmt.Printf("item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}
