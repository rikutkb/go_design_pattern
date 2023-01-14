package main

import (
	"fmt"

	"github.com/rikutkb/go_design_pattern/memento/memento"
)

func main() {
	gamer := memento.NewGamer(100)
	mem := gamer.CreateMemento()
	for i := 0; i < 100; i++ {
		if i%5 == 0 {
			if mem.GetMoney() > gamer.GetMoney() {
				fmt.Println("復元をしました")
				gamer.RestoreMemento(mem)
			} else {
				fmt.Println("現状を保存")
				mem = gamer.CreateMemento()
			}
		}
		fmt.Println("====", i+1)
		gamer.Bet()
		fmt.Println("現状:", gamer)
	}
}
