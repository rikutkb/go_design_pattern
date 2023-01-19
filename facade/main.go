package main

import (
	"fmt"

	"github.com/rikutkb/go_design_pattern/facade/pagemaker"
)

func main() {
	pageMaker := pagemaker.PageMaker{}
	if err := pageMaker.MakeWelComPage("test@test.com"); err != nil {
		fmt.Println(err)
	}
}
