package main

import "fmt"

func main(){
	m := make(map[string]string) //antes de usar um map é necessario instaciar e alocar a memoria dele na declaração, 
								//por isso utilizamos make para fazer essa alocação, podemos fazer isso com slices tambem.								
	m["Pedro"] = "pessoa"

	m2 := map[string]string{ // tambem podemos instanciar diretamente dessa forma
		"Pedro": "Pessoa",
		"Joao": "GOAT",
	}
	fmt.Println(m)
	fmt.Println(m2)
	
	sliceMap := map[string][3]int{ // aqui estamos criando um map de slices, nesse exemplo ele tem a chave "Numeros" que aponta para um slice de inteiros
		"Numeros": {1,2,3},
		"NumerosPares": {2,4,6},
	}

	pedro, ok := m2["Joao"] //para pegar o valor de slices basta atribuilos em variaves, esse ok que passamos é uma 
							//variavel boll que geralmente declaram para ver o retorno do slice, é convencional utilizar ok nesse quesito.
	numero := sliceMap["Numeros"]
	delete(sliceMap, "NumerosPares") // funcao build-in para deletar chave-valor de maps
	
	fmt.Println(sliceMap)
	fmt.Println(pedro, ok)
	fmt.Println(numero)
	
	clear(m) //caso queriamos limpar nosso map porem manter seu tamanho, podemos utilizar o clear ao inves do delete.
			//o clear tambem foi criado para remover flots de clear, pois o delete nao te essa capacidade
	fmt.Println(m)
}