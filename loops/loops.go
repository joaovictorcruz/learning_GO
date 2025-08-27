package main

import "fmt"

func main() {
	
	for i:= 0; i<=10; i++{ //for padrao
		fmt.Println(i)
	}

	var (
		res int
		i int
	)
	for ; i<10;{ // meio inicializado, um parametro funciona como um while()
		res++
		i++
		fmt.Println(res)
	}

	for {   //for sem parametros, é o while(true) go ja que no go nao existe while      
		break	
	}
	
	arr := [5]int {1,2,3,4,5}

	for i := range arr {     //forma para iterar um array com go, dessa forma pegamos o indicie
    fmt.Println("Índice:", i)
	}

	for range 10{			//itera diretamente
		fmt.Println("seila")
	}

	for i, v:= range arr{   //mesma coisa porem agora estamos pegando o valor do array com v
		fmt.Println(i, v)
	}

		for _, v:= range arr{   //apenas exibindo o valor, utilizamos o _ para ignorar o indice. basicamente o _ diz para o compilador descatar a variavel que ele utilizaria naquele caso
		fmt.Println(v)
	}
}