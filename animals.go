//package main
package animals

import "fmt"
import "os"

type Animal struct {
	food, locomotion, noise string
}

func (obj Animal) Eat() {
	fmt.Println(obj.food)
}

func (obj Animal) Move() {
	fmt.Println(obj.locomotion)
}

func (obj Animal) Speak() {
	fmt.Println(obj.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	snake := Animal{"mice", "slither", "hiss"}
	bird := Animal{"worms", "fly", "peep"}

	var name, info string
	var obj Animal

	for {
		fmt.Print("\nEnter two strings- Animal name and Animal info; Type \"exit\" to exit\n>")
		fmt.Scanln(&name, &info)
		if name == "exit" {
			os.Exit(0)
		}

		switch name {
		case "cow":
			obj = cow
		case "snake":
			obj = snake
		case "bird":
			obj = bird
		default:
			fmt.Println("Please try again.")
			continue
		}

		switch info {
		case "eat":
			obj.Eat()
		case "move":
			obj.Move()
		case "speak":
			obj.Speak()
		default:
			fmt.Println("Please try again.")
			continue
		}
	}
}
