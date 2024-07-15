package exemplos

import "fmt"

type ObjetoGO struct {
	C1 string
	C2 string
}

func (o ObjetoGO) ComportamentoGO1() string {
	return fmt.Sprintf("%s fez a Acao do ComportamentoGO1", o.C1)
}

func Tester_ObjetoGO() {
	instancia1 := ObjetoGO{
		C1: "CAMPOUM",
		C2: "CAMPODOIS",
	}
	// fmt.Println(instancia1)
	fmt.Println(instancia1.ComportamentoGO1())
}
