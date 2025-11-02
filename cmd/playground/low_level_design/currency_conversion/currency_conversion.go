package currency_conversion

import "errors"

var (
	NoConversionFoundErr = errors.New("no conversion found")
)

type Rate struct {
	From         string
	To           string
	ExchangeRate float64
}

type Engine struct {
	Adj map[string][]Rate
}

func NewEngine(rates []Rate) *Engine {
	engine := Engine{
		Adj: make(map[string][]Rate),
	}
	for _, rate := range rates {
		engine.Adj[rate.From] = append(engine.Adj[rate.From], rate)
	}

	return &engine
}

type Edge struct {
	Currency string
	Amount   float64
}

type Queue struct {
	items []Edge
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]Edge, 0),
	}
}

func (q *Queue) Push(val Edge) {
	q.items = append(q.items, val)
}

func (q *Queue) Pop() Edge {
	old := q.items
	item := old[0]
	q.items = old[1:]
	return item
}

func (q *Queue) Front() Edge {
	return q.items[0]
}

func (q *Queue) Len() int {
	return len(q.items)
}

func (e *Engine) Convert(from, to string, amount float64) (float64, error) {
	visited := make(map[string]bool)
	visited[from] = true
	qu := NewQueue()
	qu.Push(Edge{from, amount})

	for qu.Len() > 0 {
		top := qu.Front()
		qu.Pop()
		if top.Currency == to {
			return top.Amount, nil
		}

		for _, nei := range e.Adj[top.Currency] {
			if _, ok := visited[nei.To]; !ok {
				visited[nei.To] = true
				qu.Push(Edge{nei.To, top.Amount * nei.ExchangeRate})
			}
		}
	}

	return 0, NoConversionFoundErr
}

func Convert(from, to string, amount float64, rates []Rate) (float64, error) {
	engine := NewEngine(rates)
	return engine.Convert(from, to, amount)
}
