package main

import (
	"fmt"	
	"math"
)


func main() {

		var number int
		var x float64
		var y float64
		var z float64

		fmt.Scan(&number)

		y = math.Sin(float64(number))
		x = math.Cos(float64(number))
		x = math.Tan(float64(number))


		fmt.Println( "Value Sin",y)
		fmt.Println( "Value Cos",x)
		fmt.Println( "Value tan",z)

		fmt.Println("Value Sin", math.Sin(float64(number)*math.Pi/180))

	

}

