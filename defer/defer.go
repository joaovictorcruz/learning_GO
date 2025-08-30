package main

import "fmt"

func main(){
	tempoDefer()
	doDefer()
}

func tempoDefer(){  // o defer statement adia a chamda de uma função até que a função principal retorne. 
					// um ponto importante é que o defer está atrelado diretamente ao escopo da função//
	defer fmt.Println("world")
	fmt.Println("hello")
}

func doDefer(){
	defer fmt.Println(3)
	defer fmt.Println(2)
	fmt.Println(1)
}