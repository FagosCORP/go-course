package main

func main() {
	sub := func(a int, b int) int {
		return a - b 
	}

	subtrair(sub)
	
}

func subtrair(funcao func(int, int) int)  {
	return funcao(2-3)
}