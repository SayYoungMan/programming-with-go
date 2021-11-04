package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c *Cow) Eat() {
	fmt.Println("grass")
	return
}

func (c *Cow) Move() {
	fmt.Println("walk")
	return
}

func (c *Cow) Speak() {
	fmt.Println("moo")
	return
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b *Bird) Eat() {
	fmt.Println("worms")
	return
}

func (b *Bird) Move() {
	fmt.Println("fly")
	return
}

func (b *Bird) Speak() {
	fmt.Println("peep")
	return
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s *Snake) Eat() {
	fmt.Println("mice")
	return
}

func (s *Snake) Move() {
	fmt.Println("slither")
	return
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
	return
}

func main() {
	var (
		comm  string
		name  string
		third string
	)

	aniMap := make(map[string]Animal)

	for {
		fmt.Print("> ")
		fmt.Scanf("%s %s %s", &comm, &name, &third)

		if comm == "newanimal" {
			a := MakeAnimal(third)
			if a != nil {
				aniMap[name] = a
				fmt.Println("Created it!")
			}
		} else if comm == "query" {
			a, prs := aniMap[name]
			if prs {
				printRelevant(a, third)
			} else {
				fmt.Println("Invalid name please try again")
			}
		} else {
			fmt.Println("Invalid command please try again")
		}
	}
}

func MakeAnimal(animal string) Animal {
	switch animal {
	case "cow":
		return new(Cow)
	case "bird":
		return new(Bird)
	case "snake":
		return new(Snake)
	default:
		fmt.Println("Invalid animal please try again")
		return nil
	}
}

func printRelevant(ani Animal, act string) {
	switch act {
	case "eat":
		ani.Eat()
	case "move":
		ani.Move()
	case "speak":
		ani.Speak()
	default:
		fmt.Println("Invalid action please try again")
	}

	return
}
