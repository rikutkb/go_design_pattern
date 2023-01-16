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
func (t Trouble) String() string {
	return fmt.Sprintf("%d", t.number)
}

type ISupport interface {
	SetNext(next ISupport) ISupport
	resolve(t Trouble) bool
	handle(t Trouble)
}
type Support struct {
	name string
	own  ISupport
	next ISupport
}

func (s *Support) SetNext(next ISupport) ISupport {
	s.next = next
	return next
}

func (s *Support) handle(t Trouble) {
	if s.own.resolve(t) {
		s.done(t)
	} else if s.next != nil {
		s.next.handle(t)
	} else {
		s.fail(t)
	}
}

func (s *Support) done(t Trouble) {
	fmt.Println(fmt.Sprintf("%s is resolved by %s", t, s))
}

func (s *Support) fail(t Trouble) {
	fmt.Println(fmt.Sprintf("%s cannot be resolved", t))
}
func (s *Support) String() string {
	return fmt.Sprintf("%s", s.name)
}

func (s *Support) resolve(t Trouble) bool {
	return false
}

func NewSupport(name string) *Support {
	return &Support{name: name}
}

type NoSupport struct {
	*Support
}

func NewNoSupport(name string) *NoSupport {
	noSupport := &NoSupport{Support: NewSupport(name)}
	noSupport.Support.own = noSupport
	return noSupport
}
func (ns *NoSupport) resolve(t Trouble) bool {
	return false
}

type OddSupport struct {
	*Support
}

func NewOddSupport(name string) *OddSupport {
	oddSupport := &OddSupport{Support: NewSupport(name)}
	oddSupport.Support.own = oddSupport
	return oddSupport
}
func (os *OddSupport) resolve(t Trouble) bool {
	return t.GetNumber()%2 == 0
}

type SpecialSupport struct {
	*Support
	number int
}

func (ss *SpecialSupport) resolve(t Trouble) bool {
	return ss.number == t.GetNumber()
}
func NewSpecialSupport(name string, number int) *SpecialSupport {
	specialSupport := &SpecialSupport{Support: NewSupport(name)}
	specialSupport.number = number
	specialSupport.Support.own = specialSupport
	return specialSupport
}

func main() {
	alice := NewNoSupport("Alice")
	bob := NewOddSupport("Bob")
	charlie := NewSpecialSupport("Charlie", 17)
	elmo := NewSpecialSupport("Elmo", 19)
	alice.SetNext(bob).SetNext(charlie).SetNext(elmo)
	for i := 0; i < 20; i++ {
		alice.handle(NewTrouble(i))
	}

}
