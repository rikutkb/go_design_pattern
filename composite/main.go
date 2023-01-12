package main

import "fmt"

type IEntry interface {
	GetName() string
	GetSize() int
	PrintList(prefix string)
	Add(entry *IEntry) error
	String() string
}

type File struct {
	name string
	size int
}

func (f File) GetName() string {
	return f.name
}

func (f File) GetSize() int {
	return f.size
}
func (f File) String() string {
	return fmt.Sprintf("%s (%d)", f.GetName(), f.GetSize())
}
func (f File) Add(entry *IEntry) error {
	return fmt.Errorf("ファイルにエントリを追加することはできません。")
}
func (f File) PrintList(prefix string) {
	fmt.Println(fmt.Sprintf("%s/%s", prefix, f))
}
func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

type Directory struct {
	name  string
	entry [](*IEntry)
}

func (d Directory) GetName() string {
	return d.name
}

func (d Directory) GetSize() int {
	size := 0
	for _, e := range d.entry {
		size += (*e).GetSize()
	}
	return size
}
func (d Directory) String() string {
	return d.name
}
func (d Directory) PrintList(prefix string) {
	fmt.Println(fmt.Sprintf("%s/%s", prefix, d))
	for _, e := range d.entry {
		(*e).PrintList(fmt.Sprintf("%s/%s", prefix, d.name))
	}
}
func (d *Directory) Add(entry *IEntry) error {
	d.entry = append(d.entry, entry)
	return nil
}
func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

var _ IEntry = &Directory{}
var _ IEntry = &File{}

func main() {
	rootDir := NewDirectory("root")
	var bindDir IEntry = NewDirectory("bind")
	var tmpDir IEntry = NewDirectory("tmp")
	var usrDir IEntry = NewDirectory("usr")
	rootDir.Add(&bindDir)
	rootDir.Add(&tmpDir)
	rootDir.Add(&usrDir)
	var viFile IEntry = NewFile("vim", 32)
	var latexFile IEntry = NewFile("latex", 1024)
	bindDir.Add(&viFile)
	bindDir.Add(&latexFile)
	rootDir.PrintList("")

}
