package main

type User struct {
	Name string
}

func main() {
	//u := &User{Name: "Leto"}
	//println(u.Name) //prints "Leto"
	//Modify(u)
	//println(u.Name) //prints "Leto"

	//u := User{Name: "Leto"}
	//Modify2(u)
	//println(u.Name) //prints "Leto"

	//u := &User{Name: "Leto"}
	//Modify3(&u)
	//println(u.Name) //prints "Bob"
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
