package memento

import (
	"fmt"
	"math/rand"
	"time"
)

type Memento struct {
	money  int
	fruits []string
}

func (m Memento) GetMoney() int {
	return m.money
}
func (m *Memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}
func (m Memento) GetFruit() []string {
	return m.fruits
}
func NewMemento(money int) Memento {
	return Memento{money: money}
}

var fruitsName = []string{"りんご", "ブドウ"}

type Gamer struct {
	money  int
	fruits []string
	rand   rand.Rand
}

func NewGamer(money int) Gamer {
	_rand := rand.New(rand.NewSource(time.Now().UnixMicro()))
	return Gamer{money: money, rand: *_rand}
}

func (g Gamer) GetMoney() int {
	return g.money
}
func (g Gamer) GetFruit() string {
	r := rand.Int() % (len(fruitsName))
	return fruitsName[r]
}
func (g *Gamer) Bet() {
	dice := rand.Int()%6 + 1
	switch dice {
	case 1:
		g.money *= 2
		fmt.Println("所持金が増えました")
	case 2:
		g.money /= 2
		fmt.Println("所持金が半分になりました")
	case 6:
		f := g.GetFruit()
		fmt.Println(fmt.Sprintf("フルーツ(%s)を取得", f))
		g.fruits = append(g.fruits, f)
	default:
		fmt.Println("何も起こりませんでした")
	}

}
func (g Gamer) CreateMemento() Memento {
	m := NewMemento(g.money)
	for _, f := range g.fruits {
		m.addFruit(f)
	}
	return m
}
func (g *Gamer) RestoreMemento(m Memento) {
	g.fruits = m.GetFruit()
	g.money = m.GetMoney()
}
func (g Gamer) String() string {
	return fmt.Sprintf("所持金: %d, frutits: %s", g.money, g.fruits)
}
