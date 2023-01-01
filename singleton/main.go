package main

import (
	"fmt"
	"time"
)

type Singleton struct {
}

type GoodSingleton struct {
	Singleton
}

func (gs GoodSingleton) Hello() (message string) {
	return "hello"
}
func GetGoodInstance() *GoodSingleton {
	time.Sleep(3)
	return gInstance
}

type BadSingleton struct {
	Singleton
}

func (bs BadSingleton) Hello() (message string) {
	return "hello"
}
func GetBadInstance() *BadSingleton {

	if bInstance == nil {
		time.Sleep(3)
		fmt.Println("BadSingleton is initialized.")
		bInstance = &BadSingleton{}
	}
	return bInstance
}

var gInstance *GoodSingleton
var bInstance *BadSingleton

func init() {
	// シングルトンのインスタンスをスレッドセーフにするため、init関数内で初期化を行う。
	fmt.Println("GoodSingleton is initialized.")
	gInstance = &GoodSingleton{}

}

func run(ch chan<- interface{}, name string) {
	hello := GetGoodInstance()
	fmt.Println(name, hello.Hello())
	ch <- hello

}
func badRun(ch chan<- interface{}, name string) {
	hello := GetBadInstance()
	fmt.Println(name, hello.Hello())
	ch <- hello

}

func main() {
	ch := make(chan interface{})
	go run(ch, "g1")
	go run(ch, "g2")
	<-ch
	<-ch
	fmt.Println("------bad Pattern-----")
	go badRun(ch, "b1")
	go badRun(ch, "b2")
	go badRun(ch, "b3")
	<-ch
	<-ch
	<-ch
}

// ---実行結果----
// singleton[master] % go run main.go
// GoodSingleton is initialized.
// g2 hello
// g1 hello
// ------bad Pattern-----
// BadSingleton is initialized.
// b3 hello
// b2 hello
// BadSingleton is initialized.
// b1 hello
// singleton[master] %
