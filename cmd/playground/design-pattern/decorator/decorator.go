package decorator

import "fmt"

func printPizzaDetails(pizza PizzaIngredient) {
	for idx, step := range pizza.GetMakingSteps() {
		fmt.Printf("step %d: %s\n", idx+1, step)
	}
	fmt.Printf("total cost: %d\n", pizza.GetCost())
}

func Run() {
	veggePizzaWithCheeseAndTomato := &VeggeMania{
		ingredient: &CheeseTopping{
			ingredient: &TomatoTopping{},
		},
	}
	fmt.Println("details of veggePizzaWithCheeseAndTomato...")
	printPizzaDetails(veggePizzaWithCheeseAndTomato)

	veggePizzaWithCheese := &VeggeMania{
		ingredient: &CheeseTopping{},
	}
	fmt.Println("details of veggePizzaWithCheese...")
	printPizzaDetails(veggePizzaWithCheese)
}
