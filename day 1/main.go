package main

import "fmt"
import "time"

var delayTime time.Duration = 1

func main() {
	fmt.Println("Tell me some things about yourself, man!")
	delay(delayTime / 2)
	fmt.Println("Well that's a nice name, " + name() + ".")
	delay(delayTime)
	fmt.Println("You're feeling " + feelings() + "? That's cool, man.")
	delay(delayTime)
	fmt.Println("You're " + age() + " years old? Wow, so young!")
	delay(delayTime)
	fmt.Println("Well then, goodbye man!")
}

func name() string {
	var name string
	fmt.Println("What's your name, man?")
	fmt.Scanln(&name)
	delay(delayTime / 2)
	return name
}

func feelings() string {
	var feelings string
	fmt.Println("How are you feeling, man?")
	fmt.Scanln(&feelings)
	delay(delayTime / 2)
	return feelings
}

func age() string {
	var age string
	fmt.Println("How old are you, man?")
	fmt.Scanln(&age)
	delay(delayTime / 2)
	return age
}

func delay(delay time.Duration) {
	time.Sleep(delay * time.Second)
}