package main

// não consigo mudar o tipo do array por causa das suas posições posso referenciar mas não posso modificar.
// não cresce de tamanho também.
// homogeneização dos mesmo tipo e valores do mesmo tipo.
// estrutura indexadas

func main(){
	var notas [3] float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.0, 8.0, 9.0

	for i:= 0; i < len(notas); i++ {
		total += notas[i]
	}

	media := total / float64(len(notas))
	fmt.Printf("Média %.2f\n", media)
}