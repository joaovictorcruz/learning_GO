package main

import (
	"fmt"
	"math"
)

func main(){
	if x:= math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1{
		fmt.Println("maior que zero")
	} else{
		fmt.Println("caiu no else")
	}
}