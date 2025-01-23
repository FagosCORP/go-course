package main

func main(){
	slice := []int{1,2,3}
	fmt.Println(slice)
	// slice define um pedaço de um array
	// estrutura que acessa os tamanhos mas não se cria o um novo array
	s3 := slice[2:3]
	// Do inicio até o indice 3 mas sem incluir
	s3 := slice[:3]
}