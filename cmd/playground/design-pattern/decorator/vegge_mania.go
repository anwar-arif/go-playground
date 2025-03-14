package decorator

type VeggeMania struct {
	ingredient PizzaIngredient
}

func (vm *VeggeMania) GetMakingSteps() []string {
	step := "vegge mania"
	if vm.ingredient != nil {
		return append(vm.ingredient.GetMakingSteps(), step)
	}
	return []string{step}
}

func (vm *VeggeMania) GetCost() int {
	oldCost := 0
	if vm.ingredient != nil {
		oldCost += vm.ingredient.GetCost()
	}
	return oldCost + 15
}
