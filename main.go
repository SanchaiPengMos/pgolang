package main

import (
	"fmt"	
	"math"
)


func main() {

	var u float64
	var s float64
	var time float64

	// u = 15
	fmt.Println("กรอกความเร่งที่กำหนด")
	fmt.Scan(&u)
	fmt.Println("กรอกเวลา")
	fmt.Scan(&time)


	s = u * time
	fmt.Println("โพรเจ็คไทม์แกน X ", s)

	var v float64
	var g float64
	var v2 float64

	fmt.Println("กรอกการตกอย่างอิสระ")
	fmt.Scan(&g)
	// g = 9.8

	v = u + g*time

	s = v+u/2*time

	s = u *time+0.5*g*math.Pow(time,2)
	v2 = math.Pow(u,2)+2*g*s

	fmt.Println("โพรเจ็คไทม์แกน Y", v2)

}
