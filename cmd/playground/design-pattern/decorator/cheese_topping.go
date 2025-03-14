package decorator

type CheeseTopping struct {
	ingredient PizzaIngredient
}

func (ct *CheeseTopping) GetMakingSteps() []string {
	step := "cheese topping"
	if ct.ingredient != nil {
		return append(ct.ingredient.GetMakingSteps(), step)
	}
	return []string{step}
}

func (ct *CheeseTopping) GetCost() int {
	oldCost := 0
	if ct.ingredient != nil {
		oldCost += ct.ingredient.GetCost()
	}
	return oldCost + 10
}
