package main

func man() {
	// se tirar os ... vira um slice
	numeros := [...]int{1,2,3,4,5,6}

	for i, numero := range numeros {
		fmt.Println(i)
		fmt.Println(numero)
	}
}