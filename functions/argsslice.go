package main

func main() {
	a :=  []float64 {2.3, 3.2} 
	media(a...)
}

func media(numeros ...float64) {
	total := 0.0
	for _, num := numeros {
		total += num
	}
	return total / float64(len(numeros))
}
