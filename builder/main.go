package main

import (
	"bytes"
	"fmt"
)

type Builder interface {
	MakeTitle(title string)
	MakeString(str string)
	MakeItems(items []string)
	Close()
}

type Director struct {
	Builder Builder
}

func NewDirector(builder Builder) Director {
	return Director{Builder: builder}
}
func (d *Director) Construct() {
	d.Builder.MakeTitle("Greeting")
	d.Builder.MakeString("朝から昼にかけて")
	d.Builder.MakeItems([]string{"おはようございます。", "こんにちは。"})
	d.Builder.MakeString("夜に")
	d.Builder.MakeItems([]string{"こんばんは", "おやすみなさい", "さようなら"})
	d.Builder.Close()
}

var _ Builder = &TextBuilder{}

type TextBuilder struct {
	Buffer bytes.Buffer
}

func NewTextBuilder() TextBuilder {
	return TextBuilder{}
}

func (t *TextBuilder) MakeTitle(title string) {
	t.Buffer.WriteString("========\n")
	t.Buffer.WriteString("[" + title + "]")
	t.Buffer.WriteString("\n")
}
func (t *TextBuilder) MakeString(str string) {
	t.Buffer.WriteString("*" + str + "\n")
	t.Buffer.WriteString("\n")
}
func (t *TextBuilder) MakeItems(items []string) {
	for _, item := range items {
		t.Buffer.WriteString(" - " + item + "\n")
	}
	t.Buffer.WriteString("\n")
}
func (t *TextBuilder) Close() {
	t.Buffer.WriteString("=========\n")
}
func (t *TextBuilder) GetResult() string {
	return t.Buffer.String()
}
func main() {
	builder := NewTextBuilder()
	director := NewDirector(&builder)
	director.Construct()
	fmt.Println(builder.GetResult())
}

// 実行結果
// builder[master] % go run main.go
// ========
// [Greeting]
// *朝から昼にかけて

//  - おはようございます。
//  - こんにちは。

// *夜に

//  - こんばんは
//  - おやすみなさい
//  - さようなら

// =========
