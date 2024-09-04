/* addition.go
   Author   : Mr. Jake Rodriguez Pomperada,MAED-IT, MIT
   Date     :  March 20, 2021 Saturday  10:12 PM
   Location : Bacolod City, Negros Occidental
   Websites :www.jakerpomperada.comandwww.jakerpomperada.blogspot.com
   Emails   : jakerpomperada@gmail.com and jake_pomperada@tup.edu.ph
*/
//credit to the author as seen via: https://www.youtube.com/watch?v=ktNUIYaFd8M

package main

import "fmt"

func main() {
	var num1 int32
	var num2 int32
	var sum int32
	var reply1 bool = true
	var reply string

	// fmt.Print("\n")
	fmt.Print("Adding Two Numbers in Go Language Interactively\n")
	// fmt.Print("\n\n")
	fmt.Print("Enter any two numbers: ")
	fmt.Scanf("%d%d", &num1, &num2)
	sum = (num1 + num2)
	// fmt.Print("\n")
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
	// fmt.Print("\n")
	// fmt.Print("\tEnd of Program")
	// fmt.Print("\n")
	fmt.Print("Still interested in another task? (yes/no)")
	fmt.Scanln(&reply)

	fmt.Print("You have entered your response as a\t" + reply)

	if reply == "yes" {
		reply1 = true
	} else {
		reply1 = false
	}
	fmt.Print("\nYour entered response of:"+reply+"  "+"\tis equivalent to\t", reply1)

	if reply1 == true {
		fmt.Print("Adding Two Numbers in Go Language Interactively\n")
		// fmt.Print("\n\n")
		fmt.Print("Enter any two numbers: ")
		fmt.Scanf("%d%d", &num1, &num2)
		sum = (num1 + num2)
		// fmt.Print("\n")
		fmt.Println("The sum of", num1, "and", num2, "is", sum)

		// fmt.Print("Still interested in another task? (yes/no)")
		// fmt.Scanln(&reply)

	}

	// else {

	// 		fmt.Print("\nEnd of Program")
	// }

	// //While DO
	// for reply = "yes" {

	// }

	// fmt.Print("You have entered the value:" + reply)
	// //   if 7%2 == 0 {
	// // 	//         fmt.Println("7 is even")

}
