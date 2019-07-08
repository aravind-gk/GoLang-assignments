//package main
package displacement

import "fmt"

func GenDisplaceFn(acc, vel, s0 float64) func(time float64) float64 {
	fn := func(time float64) float64 {
		displacement := (0.5 * acc * time * time) + (vel * time) + s0
		return displacement
	}
	return fn
}

func main() {
	var acc, vel, s0, t float64

	fmt.Print("Enter acceleration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter initial velocity: ")
	fmt.Scan(&vel)
	fmt.Print("Enter initial displacement: ")
	fmt.Scan(&s0)

	ComputeDisplacement := GenDisplaceFn(acc, vel, s0)

	fmt.Print("Enter time: ")
	fmt.Scan(&t)

	fmt.Printf("Displacement at time=%.2f is %f units", t, ComputeDisplacement(t))
}
