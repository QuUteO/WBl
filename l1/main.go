package main

import (
	"fmt"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Action struct {
	Human
	action string
}

func (h Human) SayHi() string {
	resultStr := fmt.Sprintf("Human: Hi, I am %s! I am %v! My nubmer %s!", h.name, h.age, h.phone)
	return resultStr
}

func (a Action) Do() string {
	resultStr := fmt.Sprintf("action: %s", a.action)
	return resultStr
}

func main() {
	act := Action{
		Human: Human{
			name:  "Михаил",
			age:   19,
			phone: "13800138000",
		},
		action: "Привет!",
	}

	fmt.Println(act.SayHi()) // метод Human
	fmt.Println(act.Do())    // метод Action
}
