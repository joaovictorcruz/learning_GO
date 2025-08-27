package main

import (
	"fmt"
)

func main() {
	var student string = "João Victor" // variavel tipada

	var student2 = "Jetta" //variaveis inferidas, o compilador decide o tipo com base no valor. Tipo python
	x := 2

	fmt.Println(student)
	fmt.Println(student2)
	fmt.Println(x)

	//declaração de multiplas variaveis
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//variaveis inferidas
	var e, f = 8, "tipou"
	g, h := 10, "pega"
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	var (
		i int    //output = 0
		j int    = 1
		k string = "hello"
	)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	//constantes
	const PI float32 = 3.14
	const constante = 150

	fmt.Println(PI)
	fmt.Println(constante)

	//declaração de multiplas const
	const (
		l int = 20
		m     = 50
		n     = 70
	)

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
}
