package main
import "fmt"
func init (){


}

func main () {



	process()
	process2(-5)
	process3()
	fmt.Println("Returned from process")

}


func process (){

	defer fmt.Println("Deferred")
	fmt.Println("Normal execution")
}

func process2 (input int){
	
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	
	
	
	if input <0 {
		panic("Input must be positive and no 0")
		defer fmt.Println("Deferred 3")
	}
	
	
	fmt.Println("Processing input: ",input)

}

func process3 () {


	defer func() {
		if r := recover() ; r != nil {
			fmt.Println("Recovered", r)
		}
	}()

		fmt.Println("Start")
		panic("Something Wrong")



}