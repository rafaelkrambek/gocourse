// package main


// import "fmt"

// func main () {

// 	// Simple Iteration
// 	for i:= 1; i<= 5 ; i++ {
// 		fmt.Println(i)
// 	}

// 	// iteration over collection
	
// 	numbers := [] int {1,2,3,4,5,6}

// 	for index, value  := range numbers {
// 		fmt.Printf("Index: %d,  Value:%d\n", index , value)
// 	}

// 	// continue and break
// 	 for i:=1 ; i<= 10 ; i++ {

// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println("Odd number:",i)
// 		if i == 5{
// 			break // break out of the loop
// 		}
// 	 }

// 	//STAR PRINT
// 	rows := 3
	
// 	for i:=1 ; i<= rows ; i++ {
// 		for j := 1 ; j<= rows-i ; j++ {
// 			fmt.Print(" ")
// 		}
// 		for k:=1; k <= 2*i-1;k++{
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// 	//GO NEW FUNC
// 	for i:= range 10{
		
// 		fmt.Println(10-i)
// 	}

	//For as WHILE


	
	 
//}


package main

import (
	"math/rand"
	"time"
	"fmt"

)

func main () {


	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	target := random.Intn(100) + 1





	fmt.Println("Guessing Game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can u guess which number it is?")

	var guess int

	for {
			fmt.Println("Enter you guess")
			fmt.Scanln(&guess)
			if guess == target {
				fmt.Println("You guessed it correct!")
				break
			} else if guess < target{
				fmt.Println("Too low. Guess again")

			}else{
				fmt.Println("Too high. Guess again")
			}



	}



}