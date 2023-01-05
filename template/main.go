package main

import "fmt"

type IDisplay interface {
	open()
	print()
	close()
}
type AbstractDisplay struct {
	iDisplay IDisplay
}

func (ad *AbstractDisplay) Display() {
	ad.iDisplay.open()
	for i := 0; i < 5; i++ {
		ad.iDisplay.print()
	}
	ad.iDisplay.close()
}

type CharDisplay struct {
	char string
}

func (cd CharDisplay) open() {
	fmt.Print("<<")
}
func (cd CharDisplay) print() {
	fmt.Print(cd.char)
}
func (cd CharDisplay) close() {
	fmt.Print(">>")
}

type StringDisplay struct {
	str string
}

func (sd StringDisplay) open() {
	fmt.Println("+-----+")
}
func (sd StringDisplay) print() {
	fmt.Println(sd.str)
}

func (sd StringDisplay) close() {
	fmt.Println("+-----+")
}

// AbstractDisplayで大筋を設定し、具象のDisplayによって詳細な処理を任せる。
func main() {
	cd := AbstractDisplay{iDisplay: CharDisplay{char: "+"}}
	cd.Display()
	fmt.Println()
	fmt.Println("-------")
	sd := AbstractDisplay{iDisplay: StringDisplay{str: "this is test message"}}
	sd.Display()

}
