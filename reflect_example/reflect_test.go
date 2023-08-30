package reflect_example

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflectType(t *testing.T) {
	var i int
	i = 10
	var a any = i
	ar := reflect.TypeOf(a)
	fmt.Println(ar)
}

func TestReflectValue(t *testing.T) {
	var i int
	i = 10
	var a any = i
	ar := reflect.ValueOf(a)
	fmt.Println(ar)
}

func TestAnyReflectToType(t *testing.T) {
	var i int64
	i = 10
	var a any = i
	ar := reflect.ValueOf(a)
	var j int64
	j = ar.Int()
	fmt.Println(j)
}

//通过反射修改值

func TestReflectChangeTypeValue(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println(x)
}
