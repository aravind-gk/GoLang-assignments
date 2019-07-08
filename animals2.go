//package main
package animals2

import "fmt"
import "os"

//Declaring Interface
//Every animal can eat, speak and move
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Declaring Animal types
//Every animal has a name
type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

//Declaring three methods for each animal type
func (Cow) Eat() {
	fmt.Println("grass")
}

func (Bird) Eat() {
	fmt.Println("worms")
}

func (Snake) Eat() {
	fmt.Println("mice")
}

func (Cow) Move() {
	fmt.Println("walk")
}

func (Bird) Move() {
	fmt.Println("fly")
}

func (Snake) Move() {
	fmt.Println("slither")
}

func (Cow) Speak() {
	fmt.Println("moo")
}

func (Bird) Speak() {
	fmt.Println("peep")
}

func (Snake) Speak() {
	fmt.Println("hsss")
}

//Main function
func main() {
	var animalMap map[string]Animal
	animalMap = make(map[string]Animal)

	for {
		var commandType, animalType, animalName, animalInfo string
		var newAnimal Animal
		fmt.Print("\nEnter your command; Enter 'quit' to quit\n>")
		fmt.Scan(&commandType)
		switch commandType {
		case "newanimal":
			fmt.Scan(&animalName)
			fmt.Scan(&animalType)

			switch animalType {
			case "cow":
				newAnimal = Cow{animalName}
			case "bird":
				newAnimal = Bird{animalName}
			case "snake":
				newAnimal = Snake{animalName}
			default:
				fmt.Println("Wrong animal type, please try again with cow/bird/snake")
				continue
			}

			if _, ok := animalMap[animalName]; ok {
				fmt.Println("Sorry, an animal with the same name already exists!\nPlease try again with a different name.")
				continue
			} else {
				animalMap[animalName] = newAnimal
				fmt.Println("Animal creation successful!")
			}

		case "query":
			fmt.Scan(&animalName)
			fmt.Scan(&animalInfo)

			if tempAnimal, ok := animalMap[animalName]; ok {
				switch animalInfo {
				case "eat":
					tempAnimal.Eat()
				case "move":
					tempAnimal.Move()
				case "speak":
					tempAnimal.Speak()
				default:
					fmt.Println("Wrong animal info queried, please try again with eat/move/speak.")
					continue
				}
			} else {
				fmt.Println("Animal with given name does not exist. Please try again.")
				continue
			}

		case "quit":
			os.Exit(0)
		default:
			fmt.Println("Wrong command, Please try again with newanimal/query/quit.")
		}
	}

}
