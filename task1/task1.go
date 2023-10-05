package task1

import (
	"fmt"
	"time"
)

// Human Изначальная структура
type Human struct {
	Name string
	Age  int
}

// Talk 1 метод структуры
func (h *Human) Talk() {
	fmt.Println("Hello")
}

// Think 2 метод структуры
func (h *Human) Think() {
	time.Sleep(3 * time.Second)
	fmt.Println("AAaaAA")
}

// Action структура наследующая Human
type Action struct {
	human Human
}

// Do метод структуры Action который обращается к методам структуры Human
func (a *Action) Do() {
	a.human.Think()
	a.human.Talk()
}

// Task1 демонстрация
func Task1() {
	max := Human{
		Name: "Maxim",
		Age:  22,
	}
	act := Action{human: max}

	max.Think()
	max.Talk()
	act.human.Think()
	act.human.Talk()
	act.Do()
}
