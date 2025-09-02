package main

import "fmt"

func main(){
	padrao(1)
	noCondition("foo")
}

func padrao (x int){
	switch x{
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("teste")
	}
}

func noCondition(x any){
	switch {
	case x == 1:
		fmt.Println(1)
	case x == "foo":
		fmt.Println(2)
	default:
		fmt.Println("outra coisa")
	}
}