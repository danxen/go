package main

import "fmt"

func main() {
	fmt.Println("print from main ")

	mod1()
	mod2()
	mod3()
}

func mod1() {
	fmt.Println("print from mod a")
}

func mod2() {
	fmt.Println("print from mod b")
}

func mod3() {
	fmt.Println("print from mode c")
}
