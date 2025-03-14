package decorator

type PeepyPaneer struct {
	ingredient PizzaIngredient
}

func (pp *PeepyPaneer) GetMakingSteps() []string {
	step := "peepy paneer"
	if pp.ingredient != nil {
		return append(pp.ingredient.GetMakingSteps(), step)
	}
	return []string{step}
}

func (pp *PeepyPaneer) GetCost() int {
	oldCost := 0
	if pp.ingredient != nil {
		oldCost += pp.ingredient.GetCost()
	}
	return oldCost + 20
}
