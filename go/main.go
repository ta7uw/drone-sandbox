package main

import "fmt"

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func main() {
	fmt.Println(gcd(13, 39))
	fmt.Println(lcm(3, 8))
}
