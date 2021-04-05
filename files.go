package main

import "fmt"

func calc(x int, y int) (int, int){
	var out1 = x+y
	var out2 =x-y
	return out1, out2
}

func main(){
	var i int = 1
	var j int = 45
	fmt.Print(j+i)
	fmt.Printf("\n%v,%T\n",i,i)

	for k :=1;k<=5;k++{
		fmt.Println("hai jackson",k)
	}

	num1 := 56
	num2 := 67
	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)
}