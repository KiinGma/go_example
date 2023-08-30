package struct_example

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	name string `tage1:"name" tage2:"名字"`
	age  int    `tage1:"age" tage2:"年龄"`
}

// 获取struct tag
func TestGetStructTag(t *testing.T) {
	u := user{}
	ur := reflect.TypeOf(u)
	fmt.Println(ur.NumField())
	for i := 0; i < ur.NumField(); i++ {
		fmt.Println(ur.Field(i).Tag.Get("tage1"))
		fmt.Println(ur.Field(i).Tag.Get("tage2"))
	}
}
