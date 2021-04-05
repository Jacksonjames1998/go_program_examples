package main

import "fmt"

func add(x int, y int) int{
	var sum = x+y
	return sum
}

func main(){
	var i int = 1
	var j int = 45
	fmt.Print(j+i)
	fmt.Printf("\n%v,%T\n",i,i)

	for k :=1;k<=5;k++{
		fmt.Println("hai jackson",k)
	}

	num1 := 1
	num2 := 2
	result := add(num1, num2)
	fmt.Println(result)
}