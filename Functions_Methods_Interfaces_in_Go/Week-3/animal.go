package main

import "fmt"

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (ani *Animal) Eat() {
	fmt.Println(ani.food)
	return
}

func (ani *Animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

func (ani *Animal) Speak() {
	fmt.Println(ani.noise)
	return
}

func main() {
	var (
		ani string
		act string
		tmp Animal
	)

	cow := Animal{
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	}

	bird := Animal{
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	}

	snake := Animal{
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	}

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s", &ani, &act)

		if ani == "cow" {
			tmp = cow
		} else if ani == "bird" {
			tmp = bird
		} else if ani == "snake" {
			tmp = snake
		} else {
			fmt.Println("Invalid animal please try again")
			continue
		}

		if act == "eat" {
			tmp.Eat()
		} else if act == "move" {
			tmp.Move()
		} else if act == "speak" {
			tmp.Speak()
		} else {
			fmt.Println("Invalid action please try again")
		}
	}
}
