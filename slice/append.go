package main

func main(){
	slice := []int
	// não posso copiar em um array
	append(slice, 4,5,6)
	
	// ele pega quantos elementos dá por causa do seu tamanho de elementos
	slice2:= make([]int, 2)
	copy(slice2, slice1)
}