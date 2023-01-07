package main

// クラスからインスタンスを作るのではなく、インスタンスから別のインスタンスを作成
import (
	"fmt"

	"github.com/rikutkb/go_design_pattern/prototype/framework"
)

var _ framework.Product = UnderLinePen{}

type UnderLinePen struct {
	ulChar string
}

func (ul UnderLinePen) Use(str string) {
	fmt.Println(ul.ulChar)
	fmt.Println(str)
	fmt.Println(ul.ulChar)

}
func (ul UnderLinePen) CreateClone() framework.Product {
	return ul
}

func main() {
	manager := framework.NewManager()
	strongPen := UnderLinePen{ulChar: "*"}
	manager.Register("strong", strongPen)
	warningPen := UnderLinePen{ulChar: "~"}
	manager.Register("warning", warningPen)

	p1 := manager.Create("warning")
	p2 := manager.Create("strong")
	fmt.Println("-------warning-----")

	p1.Use("hello")
	fmt.Println("-------strong-----")
	p2.Use("hello")
}
