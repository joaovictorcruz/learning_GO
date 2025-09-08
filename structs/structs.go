package main

import "fmt"

// NAO EXISTEM CLASSES EM GO!! Mas existem structs, que sao semelhantes ás classes.

type User struct{ //uma struct é basicamente uma coleção de campos, nao declaramos metodos dentro
	Name string 
	Id uint64
}

//em metodos passamos uma variavel no parametro que referencia a struct para poder acessar seus campos, geralmente é utilizado a inicial da struct referenciada
func (u User) foo(){ //para declarar um metodo temos que passar um parametro que referencia a struct que estamos criando o metodo
				   //um ponto importante é que em go, podemos criar um metodo para qualquer type e não só para structs, que é o convencional em ling como java e c# metodos serem especificos de classes.
	fmt.Println(u.Name)
}

func (u *User) UpdateName(){ //ao referenciar a struct no parametro estamos COPIANDO o valor dela que está em memoria, se quisermos criar um metodo como atualizar precisamos apontar para ela com o *.
	u.Name = "João Novo"
}

func main(){
	user := User{Name: "João Victor", Id: 1} // inserindo dados nos campos de uma struct
	fmt.Println(user.Name) //acessando o dado de uma struct, muito semelhante ao POO classico
	user.UpdateName() 
	fmt.Println(user.Name)
}