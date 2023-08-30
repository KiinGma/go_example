package map_example

import (
	"fmt"
	"testing"
)

func TestMapNil(t *testing.T) {
	var m map[string]string
	v, ok := m["test"]
	fmt.Println(ok)
	//m["test"] = "test"
	if ok {
		fmt.Println(v)
	} else {

	}
	fmt.Println(m)
}
