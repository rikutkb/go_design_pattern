package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	iterator() Iterator
}

type Book struct {
	name string
}

func (b Book) GetName() string {
	return b.name
}

type BookShelf struct {
	books []Book
	len   int
}

func (bs BookShelf) GetBookAt(index int) Book {
	return bs.books[index]
}
func (bs *BookShelf) AppendBook(book Book) {
	bs.books = append(bs.books, book)
	bs.len += 1
}
func (bs BookShelf) GetLength() (length int) {
	return bs.len
}
func (bs *BookShelf) iterator() Iterator {
	return &BookShelfIterator{bookShelf: bs}
}

type BookShelfIterator struct {
	bookShelf *BookShelf
	index     int
}

func (bsi BookShelfIterator) HasNext() bool {
	if bsi.index < bsi.bookShelf.GetLength() {
		return true
	} else {
		return false
	}
}
func (bsi *BookShelfIterator) Next() interface{} {
	book := bsi.bookShelf.GetBookAt(bsi.index)
	bsi.index += 1
	return book

}

func main() {
	bookShelf := BookShelf{}
	bookShelf.AppendBook(Book{name: "test1"})
	bookShelf.AppendBook(Book{name: "test2"})
	bookShelf.AppendBook(Book{name: "test3"})
	it := bookShelf.iterator()
	for it.HasNext() {
		book, ok := it.Next().(Book)
		if ok {
			fmt.Println(book.GetName())
		}
	}

}
