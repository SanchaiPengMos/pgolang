package main

import "fmt"

func main() {

	fmt.Println("You want know about??")

	var About string
	fmt.Scanf("%s", &About)

	if About == "yes" {
		var Speed int
		var time int

		fmt.Println("Your speed (Km/hr)")
		fmt.Scan(&Speed)
		fmt.Println("Speed", Speed)

		fmt.Println("Your time (Km/hr)")
		fmt.Scan(&time)

		fmt.Println("time", time)

		fmt.Println("ความเร็ว", Speed*time)
	}
}
