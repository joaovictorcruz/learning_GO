package main

import "fmt"

// Um ponteiro é endereço de memória
// ponteiros em stack, são ponteiros que apontaram para a memoria da stack da propria funcao

func main(){
	
	fmt.Println(exemple())

	x := create()
	fmt.Println(*x)
}

func exemple() (*int, int){
	
	x := 10
	p :=  &x //utilizar & para utilizar o ponteiro e pegar o endereço de memoria da variavel
	return p, *p//utiliza * para fazer uma dereferencia do ponteiro, ou seja estamos pegando o valor daquele endereço de memoria
}

func create() *int{ //quando colocarmos um asterisco no tipo, estamos dizendo que o tipo da variavel ou retorno é do tipo ponteiro, quando é em uma varivavel é uma dereferencia
	x := 10
	return &x
}