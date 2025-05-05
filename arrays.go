package main

import (
	"fmt"
)

func main() {
	//declaração de arrays
	var arr1 = [6]int{1, 2, 3, 4, 5, 6}
	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr1)
	fmt.Println(arr2)

	var arr3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arr4 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(arr3)
	fmt.Println(arr4)

	//indices e mudanças de valor em arrays
	var cars = [5]string{"BMW", "Ferrari", "Ford", "Mercedes", "Volkswagen"}
	fmt.Println(cars)
	fmt.Println(cars[0])

	cars[0] = "McLaren"
	fmt.Println(cars[0])

	//inicialização de arrays
	arr5 := [5]int{}              //não inicializado
	arr6 := [5]int{1, 2}          //parcialmente inicializado
	arr7 := [5]int{1, 2, 3, 4, 5} //totalmente inicializado

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	//inicialização de elementos específicos
	arr8 := [5]int{1: 10, 2: 40}

	fmt.Println(arr8)
	fmt.Println(len(arr8)) //comprimento do array
}
