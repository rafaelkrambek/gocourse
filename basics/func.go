package main


import "fmt"

func main () {

	// sum:= add(1,2)




	greet := func(){
		fmt.Println("Hello!")
	}

	greet()



	// operation := add

	// result := operation(40,30)

	// fmt.Println(result)

	result := applyOperation(5,3,add)

	fmt.Println(result)


	multiply2 := creatrMultiplier(2)

	fmt.Println("6 * 2 = ", multiply2(6))
	

	

	
	
	
}
	
func add(a,b int ) int {
	return a+b

}

func applyOperation(x int, y int , operation func(int,int) int ) int {
	return operation(x,y)

}

func creatrMultiplier(factor int ) func(int) int {

	return func(x int) int{
		return x*factor
	}

}

func divide (a,b int) (int,int) {

		q := a/b
		r := a%b

		return q,r

	}

