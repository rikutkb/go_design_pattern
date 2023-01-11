package main

import (
	"fmt"
)

type IDisplay interface {
	RawOpen()
	RawPrint()
	RawClose()
}

type Display struct {
	IDisplay IDisplay
}

func (d *Display) Open() {
	d.IDisplay.RawOpen()
}

func (d *Display) Close() {
	d.IDisplay.RawClose()
}
func (d *Display) Print() {
	d.IDisplay.RawPrint()
}
func (d *Display) Display() {
	d.Open()
	d.Print()
	d.Close()
}
func NewDisplay(iDisplay IDisplay) Display {
	return Display{IDisplay: iDisplay}
}

type CountDisplay struct {
	Display Display
}

func NewCountDisplay(display Display) CountDisplay {
	return CountDisplay{Display: display}
}

func (cd *CountDisplay) MultiDisplay(times int) {
	cd.Display.Open()
	for i := 0; i < times; i++ {
		cd.Display.Print()
	}
	cd.Display.Close()
}

type StringDisplay struct {
	Str   string
	width int
}

var _ IDisplay = StringDisplay{}

func (sd StringDisplay) RawOpen() {
	sd.printLine()

}
func (sd StringDisplay) RawPrint() {
	fmt.Println("|" + sd.Str + "|")
}
func (sd StringDisplay) RawClose() {
	sd.printLine()
}
func (sd *StringDisplay) printLine() {
	line := "+"
	for i := 0; i < sd.width; i++ {
		line += "-"
	}
	line += "+"
	fmt.Println(line)
}

func NewStringDisplay(str string, width int) StringDisplay {
	return StringDisplay{Str: str, width: width}
}

func main() {
	display := NewDisplay(NewStringDisplay("test", 3))
	countDisplay := NewCountDisplay(NewDisplay(NewStringDisplay("a", 3)))
	display.Display()
	countDisplay.Display.Display()
	countDisplay.MultiDisplay(5)

}
