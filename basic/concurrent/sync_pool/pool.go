package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var userPool = sync.Pool{
	New: func() any {
		return new(User)
	},
}

func main() {
	user1 := userPool.Get().(*User)
	ub, _ := json.Marshal(user1)
	fmt.Println(string(ub))

	userPool.Put(user1)
}

type User struct {
	Name string
	Age  int
}
