package main

import (
	"fmt"
	"testing"
)

//for range 的 v值 存储地址不会变 , 所以取值都是最后变化的值

func TestForRangePointer(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	var a []*int
	for _, v := range s {
		a = append(a, &v)
	}
	for _, v := range a {
		fmt.Println(*v)
	}
	fmt.Println(a)
}

func TestForPointer(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	var a []*int
	for i := 0; i < len(s); i++ {
		a = append(a, &s[i])
	}
	for _, v := range a {
		fmt.Println(*v)
	}
	fmt.Println(a)
}
