package decorator

type makhsos struct {
	pizza pizza
}

func (c *makhsos) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
