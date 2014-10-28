package main

import "fmt"

// see http://stackoverflow.com/questions/24894479/golangdifferent-about-call-a-method-with-pointer-type-using-point-type-or-str?rq=1

type User struct {
	Name string
}

func (u *User) Greeting() string {
	return fmt.Sprintf("Greetings %s [%p]!", u.Name, u)
}

func main() {
	p := &User{"cppgohan by pointer"}
	u := User{"cppgohan by value"}
	pu := &u
	fmt.Printf("[%p] %v %v\n", pu, pu.Greeting(), pu)
	fmt.Printf("[%p] %v %v\n", &u, u.Greeting(), u)
	fmt.Printf("[%p] %v %v\n", p, p.Greeting(), p)
}

