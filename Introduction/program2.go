/* program2.go */
package main

import ( "fmt" ; "math" )

func curve(x, A, B float64) float64 {
	y := A * x * math.Exp(-1.1*x)
	if x > B {
		y = y + B*(x-B)*math.Exp(-1.1*(x-B))
	}
	return y
}

func main() {
   var A, B float64
	fmt.Scan(&A, &B)
	for i := 0; i < 100; i++ {
		x := float64(i) / 10.0
		fmt.Println(x, curve(x, A, B))
	}
}
