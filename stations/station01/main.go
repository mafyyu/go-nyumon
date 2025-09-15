package main

import "fmt"

func main() {
	fmt.Println("Start FizzBuzz....")

	FizzBuzz() // この関数を実装する

	fmt.Println("FizzBuzz completed!")
}

func FizzBuzz(){
	for i:=1;i<=100;i++{
		if i%15==0 {
			fmt.Println("FizzBuzz")
		}else if i%5==0{
			fmt.Println("Buzz")
		}else if i%3==0{
			fmt.Println("Fizz")
		}else{
			fmt.Println(i)
		}
	}
}
