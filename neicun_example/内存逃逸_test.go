package escape_example

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func TestEscape(t *testing.T) {
	s := GetUser("tom", 12)
	fmt.Println(s)
}

func GetUser(name string, age int) *User {
	s := new(User) //局部变量逃逸到堆里
	s.Name = name
	s.Age = age
	return s
}
