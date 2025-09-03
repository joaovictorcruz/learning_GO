package main

import "fmt"

func main(){
	arr := [5]int{1,2,3,4,5} // é uma lista de elementos do mesmo tipo, de tamanho fixo e constante.
	slice := arr[1:4]		//slice é um array dinamico, ou seja ele pode crescer de tamanho, diferente do array, por de baixo dos panos eles o slice tem um ponteiro para um array.
						//quando colocando nosso array em um slice, obrigatoriamente temos que passar o limite inferior e superior, caso voce deixe apenas : e nao insira nada
						//o slice irá utilizar os valores padrões que são indice inferior 0 e superior será o tamanho do array
						
	sliceIndependente := []int{1,2,3} // para criar um slice diretamente, é so nao passar o tamanho dele  
	
	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Println(sliceIndependente)

}