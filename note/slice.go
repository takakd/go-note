package main

import (
	"fmt"
)

// Note: Slice

func main() {
	s1 := []string{"a", "b"}

	printSlice(s1)

	// NOTE: s2 does not have new data("a", "b"), only has a value of cap and length.
	s2 := s1
	printSlice(s1)
	printSlice(s2)

	strSlice := []string{"word", "desk", "corn", "door", "sky", "car"}

	// "desk", "corn", "door"
	printSlice(strSlice[1:4])

	// "word", "desk"
	printSlice(strSlice[:2])

	// "desk", "corn", "door", "sky", "car"
	printSlice(strSlice[1:])


	// NOTE: s points strSlice
	s := strSlice[:0]
	// []
	printSlice(s)
	// "desk", "corn", "door", "sky", "car"
	printSlice(strSlice)

	s = s[:4]
	// "word", "desk", "corn", "door"
	printSlice(s)
	// "desk", "corn", "door", "sky", "car"
	printSlice(strSlice)

	s = s[2:]
	// "corn", "door"
	printSlice(s)
	// "desk", "corn", "door", "sky", "car"
	printSlice(strSlice)

	strSlice[2] = "burned-corn"
	// "burned-corn", "door"
	printSlice(s)
	// "desk", "burned-corn", "door", "sky", "car"
	printSlice(strSlice)

	changeSlice(s, 0, "changed1")
	changeSlice(strSlice, 3, "changed2")
	// "changed1", "changed2"
	printSlice(s)
	// "work", "desk", "changed1", "changed2", "sky", "car"
	printSlice(strSlice)

	// Nil
	var ns []string
	println(ns == nil)
	printSlice(ns)

	// Cap
	b := make([]string, 0, 5)
	// len=2, cap=5(not changed)
	c := b[:2]
	printSlice(c)
	// len=3, cap=3:(5 - 2), 2 is 0,1
	d := c[2:5]
	printSlice(d)

	// append
	var ns1 []string
	// append more than one element.
	ns1 = append(ns1, "desk", "pen")
	printSlice(ns1)
}

func changeSlice(s []string, idx int, v string) {
	s[idx] = v
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
