package main

import "fmt"

func square(n int) int { return n * n }
func main() {
	var f func(int) int
	f = func(n int) int { return -n }
	fmt.Println(f(3)) // -3
	f = square
	fmt.Println(f(3)) // 9
	g := func(m, n int) int { return m * n }
	fmt.Println(g(2, 3)) // 6
	// f = g :erreur de compil car  'func(int,int) int' et 'func(int) int' incompatibles
	fmt.Println(fcompf(f, 3))                                // 81
	fmt.Println(fcompf(func(n int) int { return 2 * n }, 3)) // 12
}
func fcompf(f func(int) int, n int) int {
	return f(f(n))
}

// end OMIT
