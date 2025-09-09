package main

import "fmt"

type MinhaInterface interface{
	Teste(name string) string
}

type MinhaStruct struct{
	Name string
}

func (m *MinhaStruct) Teste(name string) string{
	m.Name = name	

	return name
}

func main(){
	m := &MinhaStruct{}
	
	fmt.Println(m.Teste("João Victor"))
}