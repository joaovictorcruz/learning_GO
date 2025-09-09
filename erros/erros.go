package main

import (
	"errors"
	"fmt"
)

// Erros em Go são basicamente interfaces que guardam valores dentro dela, por isso em go dizem que erros são valores.

func main(){
	a := 12
	b := 10
	res, err := dividir(a, b)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(a, "/", b,res)
}

func dividir(a, b int) (int, error){
	if b == 0{
		return  0, errors.New("Não pode dividir por zero")
	}
	return a / b, nil
}