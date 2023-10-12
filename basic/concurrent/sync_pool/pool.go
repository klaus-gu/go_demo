package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var userPool = sync.Pool{
	New: func() any {
		fmt.Println("重新生成对象")
		user := new(User)
		user.Name = "zs"
		user.Age = 11
		return user
	},
}

func main() {
	user1 := userPool.Get().(*User)
	user1.Age = 12
	ub, _ := json.Marshal(user1)
	fmt.Println(string(ub))
	userPool.Put(user1)
	user2 := userPool.Get().(*User)
	fmt.Println(user2)
	userPool.Put(user2)
	user3 := userPool.Get().(*User)
	fmt.Println(user3)
}

type User struct {
	Name string
	Age  int
}
