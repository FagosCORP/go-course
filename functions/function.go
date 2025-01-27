package main

func fNotReturnNotParam() {
	fmt.Println("definir bons identificadores")
}

func fParam(p1 string, p2 float64) {
	fmt.Println("definir bons identificadores")
}
// em todos ramos você necessita retornar

func fReturn() string {
	return "a"
}

// em todos ramos você necessita retornar
func fReturn(p1 string, p2 float64) string {
	return "a"
}

//funções pura que não gera efeitos colaterais
// não alterar nada que está externo
// preferência receber por parâmetro 

func fReturn(p1 string, p2 float64) (string, float64) {
	return "teste", 50.3
}

function main() {

}