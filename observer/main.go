package main

import (
	"fmt"
	"math/rand"
	"time"
)

type INumberGenerator interface {
	getNumber() int
	execute()
	addObserver(observer Observer)
	deleteObserver(observer Observer)
	NotiryObservers()
}
type RandomNumberGenerator struct {
	Observers []Observer
	rand      rand.Rand
	number    int
}

func NewRandomNumberGenerator() RandomNumberGenerator {
	rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return RandomNumberGenerator{rand: *rand}
}

func (rng *RandomNumberGenerator) addObserver(observer Observer) {
	rng.Observers = append(rng.Observers, observer)
}

func (rng *RandomNumberGenerator) deleteObserver(observer Observer) {
}
func (rng *RandomNumberGenerator) NotiryObservers() {
	for _, o := range rng.Observers {
		o.Update(rng)
	}
}

func (rng RandomNumberGenerator) getNumber() int {
	return rng.number
}
func (rng *RandomNumberGenerator) execute() {
	for i := 0; i < 20; i++ {
		rng.number = rand.Int() % 50
		rng.NotiryObservers()
	}
}

type Observer interface {
	Update(ng INumberGenerator)
}
type DigitObserver struct{}

func (db DigitObserver) Update(ng INumberGenerator) {
	fmt.Println("Digit Observer: ", ng.getNumber())
	time.Sleep(100)
}

type GraphObserver struct{}

func (gb GraphObserver) Update(ng INumberGenerator) {
	fmt.Print("Graph Observer: ")
	for i := 0; i < ng.getNumber(); i++ {
		fmt.Print("*")
	}
	fmt.Println()
	time.Sleep(100)

}

func main() {
	generator := NewRandomNumberGenerator()
	observer1 := DigitObserver{}
	observer2 := GraphObserver{}
	generator.addObserver(observer1)
	generator.addObserver(observer2)
	generator.execute()
}
