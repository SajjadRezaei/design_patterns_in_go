package observer

import "fmt"

type Observer interface {
	notify(string)
	getID() string
}

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type Customer struct {
	id string
}

type Item struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *Item) deregister(o Observer) {
	i.observerList = removeFromSlice(i.observerList, o)
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) notifyAll() {
	for _, observer := range i.observerList {
		observer.notify(i.name)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	length := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[length-1], observerList[i] = observerList[i], observerList[length-1]
			return observerList[:length-1]
		}
	}

	return observerList
}

func (c *Customer) notify(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

func Run() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
