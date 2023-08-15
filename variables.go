package main

import (
	. "fmt"
)

func variables() {
	var num int = 3
	var example int
	Println(num, example)
	example = 2
	Println(num, example)

	var a, b, c = 1, 2, "hello"
	Println(a, b, c)

	var (
		d int    = 22
		e string = "Hello World!"
		f bool   = true
		g        = "And another string value."
	)

	Println(d, e, f, g)

	greeting := "Hi, how are you?"
	score, phrase := 231, "Winning the lot!"
	Println(greeting, score, phrase)

}
