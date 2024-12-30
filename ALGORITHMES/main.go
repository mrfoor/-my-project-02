package main

import "fmt"

func main() {
	var x, y int
	var z float64

	fmt.Print("Give the values ​​of X and Y: ")
	fmt.Scanf("%d %d", &x, &y)
	if x == 0 {
		fmt.Println("Error: The value of 'a' cannot be 0.")
		return
	}

	z = float64(-y) / float64(x)

	fmt.Printf("the solution of the value %d + z + %d = 0 est z = %.2f\n", x, y, z)
}
