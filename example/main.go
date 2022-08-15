package main

import (
	"fmt"

	"github.com/YusukeKishino/go-linq"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	list := linq.From(s)
	fmt.Println(list.MustFirst())
	fmt.Println(list.MustLast())
	fmt.Println(list.MustAt(2))

	s2 := []struct {
		n int
	}{
		{1}, {2}, {3}, {4}, {5},
	}
	list2 := linq.From(s2)
	fmt.Println(list2.MustFirst())
	fmt.Println(list2.MustLast())
	fmt.Println(list2.MustAt(2))
}
