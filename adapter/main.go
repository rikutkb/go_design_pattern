package main

import "fmt"

type IPrint interface {
	PrintWeak()
	PrintStrong()
}

// 提供されている元のstruct
type Banner struct {
	message string
}

func (b Banner) showWithParen() {
	fmt.Print("(" + b.message + ")")
}
func (b Banner) showWithAster() {
	fmt.Print("*" + b.message + "*")
}

type PrintBanner struct {
	Banner
}

func (pb PrintBanner) PrintWeak() {
	pb.Banner.showWithParen()
}
func (pb PrintBanner) PrintStrong() {
	pb.Banner.showWithAster()

}
func main() {
	var printer IPrint
	printer = PrintBanner{Banner: Banner{message: "test"}}
	printer.PrintWeak()
	fmt.Println()
	printer.PrintStrong()

}
