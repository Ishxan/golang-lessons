package main

import "fmt"

func printint(x int) {
	fmt.Println(x)
}
func printstring(y string) {
	fmt.Println(y)
}
func calc(z int, t int) {
	fmt.Println(z + t)
	fmt.Println(z - t)
	fmt.Println(z / t)
	fmt.Println(z * t)
}
func calc1(z int, t int) (int, int, int, int) {
	return z + t, z - t, z / t, z * t

}
