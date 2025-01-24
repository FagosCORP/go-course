package main

func main() {
	//estrutura chave e valor
	//mapas devem ser inicializados

	//var aprovados map[int]string
	aprovados := make(map[int]string)

	aprovados[1] = "Maria"
	fmt.Println(aprovados)

	for cpf, nome := range(aprovados) {
		fmt.Println(cpf)
		fmt.Println(nome)
	}

	//excluir elementos
	fmt.Println(aprovados[1])
	delete(aprovados, 1)
	delete(aprovados, 1)

}