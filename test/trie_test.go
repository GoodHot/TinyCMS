package test

import (
	"fmt"
	"github.com/armon/go-radix"
	"testing"
)

func TestTrie1(t *testing.T) {

	r := radix.New()
	r.Insert("foo", 1)
	r.Insert("bar", 2)
	r.Insert("foobar", 2)
	r.Insert("fuck", 2)

	var list []string

	r.WalkPrefix("f", func(s string, v interface{}) bool {
		fmt.Println(s, v)
		list = append(list, s)
		return false
	})
	fmt.Println(list)
	//str, _ := json.Marshal(r.ToMap())
	//fmt.Println(string(str))
}