package main

func main(){

}

func media(numeros ...float64) {
	total := 0.0
	for _, num := numeros {
		total += num
	}
	return total / float64(len(numeros))
}