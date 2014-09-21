package main

type User struct {
	Name string
}

func main() {
	u1 := &User{Name: "Leto"}
	println(u1.Name) //prints "Leto"
	Modify(u1)
	println(u1.Name) //prints "Leto"

	u2 := User{Name: "Leto"}
	Modify2(u2)
	println(u2.Name) //prints "Leto"

	u3 := &User{Name: "Leto"}
	Modify3(&u3)
	println(u3.Name) //prints "Bob"
}

func Modify(u *User) {
	u = &User{Name: "Paul"}
}

func Modify2(u User) {
	u.Name = "Duncan"
}

func Modify3(u **User) {
	*u = &User{Name: "Bob"}
}
