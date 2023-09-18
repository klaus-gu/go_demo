package main

type User struct {
	name string
	age  int
}

func NewUserConstructor() (user *User) {
	var u = new(User)
	u.name = ""
	u.age = 0
	return u
}

func (user *User) getName() string {
	return user.name
}

func (user *User) getAge() int {
	return user.age
}

func (user *User) setName(name string) {
	user.name = name
}

func (user *User) setAge(age int) {
	user.age = age
}
