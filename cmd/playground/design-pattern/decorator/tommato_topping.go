package decorator

type TomatoTopping struct {
	ingredient PizzaIngredient
}

func (tt *TomatoTopping) GetMakingSteps() []string {
	step := "tomato topping"
	if tt.ingredient != nil {
		return append(tt.ingredient.GetMakingSteps(), step)
	}
	return []string{step}
}

func (tt *TomatoTopping) GetCost() int {
	oldCost := 0
	if tt.ingredient != nil {
		oldCost += tt.ingredient.GetCost()
	}
	return oldCost + 7
}
