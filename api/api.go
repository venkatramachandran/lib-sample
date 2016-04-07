package api

import "fmt"

func Hello(s string) {
	fmt.Println("Hello " + s + "!")
}

func Hello2(s string) string {
	return "Hello " + s + "!"
}
