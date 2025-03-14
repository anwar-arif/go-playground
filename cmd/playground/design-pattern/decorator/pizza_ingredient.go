package decorator

type PizzaIngredient interface {
	GetMakingSteps() []string
	GetCost() int
}
