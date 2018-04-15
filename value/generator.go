package value

import (
	"time"
)

// Generator interface is responsible for constantly generating values.
type Generator interface {
	// TODO.
	Value(int) int
	// Starts generating values with intervals given as parameter.
	// Interval is provided in seconds.
	// Returns chanel with generated values.
	Start(int) chan int
}

type gen struct {
	value int
}

func New(initialValue int) Generator {
	return &gen{initialValue}
}

func (g *gen) Value(v int) int {
	oldValue := g.value
	g.value = v
	return oldValue
}

func (g *gen) Start(interval int) chan int {
	ch := make(chan int, 10)

	t := time.NewTicker(time.Second * time.Duration(interval))
	go func() {
		for {
			<-t.C
			ch <- g.value
		}
	}()

	return ch
}
