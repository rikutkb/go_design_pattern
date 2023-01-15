package main

import "fmt"

type Trouble struct {
	number int
}

func NewTrouble(number int) Trouble {
	return Trouble{number: number}
}
func (t Trouble) GetNumber() int {
	return t.number
}

type ISupport interface {
	SetNext(next ISupport) ISupport
	Support(t Trouble)
	resolve(t Trouble) bool
	done(t Trouble)
	fail(t Trouble)
}

type NoSupport struct {
	name string
	next ISupport
}

func NewNoSupport(name string) NoSupport {
	return NoSupport{name: name}
}
func (ns *NoSupport) SetNext(next ISupport) ISupport {
	ns.next = next
	return next
}

func (ns *NoSupport) Support(t Trouble) {
	if ns.resolve(t) {
		ns.done(t)
	} else if ns.next != nil {
		ns.next.Support(t)
	} else {
		ns.fail(t)
	}
}

func (ns *NoSupport) resolve(t Trouble) bool {
	return false
}

func (ns *NoSupport) done(t Trouble) {
	fmt.Println(fmt.Sprintln("%s is resolved by %s", t, ns))
}

func (ns *NoSupport) fail(t Trouble) {
	fmt.Println(fmt.Sprintln("%s cannot be resolved", t))
}
func (ns NoSupport) String() string {
	return fmt.Sprintf("[%s]", ns.name)
}

func main() {

}
