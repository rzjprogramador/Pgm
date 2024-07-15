package exemplos

import "fmt"

type ObjetoGO struct {
	C1                string
	C2                string
	ComposicaoVariada Variado
}

type OpcaoVariada1 struct {
	C1 string
}
type OpcaoVariada2 struct {
	C1 string
}

// definindo interface - metodos polimorficos que podem ser implemenatdos diferenets mas retornando o tipo prometido.
type Variado interface {
	ComportamentoSejaQualForOObjeto() string
}

// criando metodo
func (o ObjetoGO) ComportamentoGO1() string {
	return fmt.Sprintf("%s fez a Acao do ComportamentoGO1", o.C1)
}

// .

func (o OpcaoVariada1) ComportamentoSejaQualForOObjeto() string {
	return fmt.Sprintf("Meu Jeito1:: var Interpolada >> %s ", o.C1)
}

func (o OpcaoVariada2) ComportamentoSejaQualForOObjeto() string {
	return fmt.Sprintf("Meu Jeito2:: var Interpolada >> %s", o.C1)
}

func Tester_ObjetoGO() {
	opcaovariada1 := OpcaoVariada1{C1: "FooVariada1"}
	opcaovariada2 := OpcaoVariada2{C1: "FooVariada1"}

	instancia1 := ObjetoGO{
		C1:                "CAMPOUM",
		C2:                "CAMPODOIS",
		ComposicaoVariada: opcaovariada1,
	}

	instancia2 := ObjetoGO{
		C1:                "CAMPOUM",
		C2:                "CAMPODOIS",
		ComposicaoVariada: opcaovariada2,
	}

	// fmt.Println(instancia1)
	// fmt.Println(instancia1.ComportamentoGO1())
	fmt.Println(instancia1.ComposicaoVariada.ComportamentoSejaQualForOObjeto())
	fmt.Println(instancia2.ComposicaoVariada.ComportamentoSejaQualForOObjeto())
}
