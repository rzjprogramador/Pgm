package partseed_go

import l "github.com/rzjprogramador/Pgm/Pgm/linguagem"

var Objeto_GO = l.Objeto{
	Objeto_Contexto: l.Objeto_Contexto{

		Molde_Gerador_da_Informacao: "struct",

		Gerar_Instancia:               `func funcaoNomeDoNovoTipo(t Tipo) { return t}`,
		Campos_Fixo_na_Instancia:      `#TODO`,
		Campos_Dinamicos_na_Instancia: `#TODO`,
		Dar_Intelegencia_Ha_Instancia: `Criar Metodos_Computaveis. ex: `,

		Campos_Com_Valores_Default: l.Algoritmo_Type_LING{
			Algoritmo: `no struct inserir campo default: valor`,
			Exemplo:   `#todo`,
		},

		Composicao: l.Composicao_Contexto{
			Conceito: "Campo de um objeto que Delega a outro objeto fazer a acao por ele.",

			Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes: l.Contexto{
				Titulo:   "Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes",
				Objetivo: `ter um objPai. carro ou moto . vao_Ter_Acao_Foo()`,
				Molde: `cria-se uma interface com somente metodos que vai querer que os objetos que SEJAM UM DESTA INTERFACE tenham .

				Obs: cada SubTipo tem que replicar um metodo com o mesmmo nome da funcao definida na interface , vai ser polimorfico cada tipo pode implementar do seu jeito só tem que devolve ro emsmo tipo prometido no contrato da interface.

				Quem replicar o comportamento da interface inserindo o seu This SERA_UM desta interface ex: os 2 Tipos replicaram o metodo da interface - quer dizer que eles SAO UM AGORA .. assim cada instancia podera fazer composicao do objeto instanciado que desejar basta ele ser do tipo prometido.
				`,

				Exemplos: []string{
					`Esta AcaoFoo() foi definida na interface que esses objetos SÃO UM - ex: Veiculo e implementada por estes subobjetos vinculando assim como metodo deles.`,
					`
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
					`,
				},
			},
		},

		Criar_Metodo_Para_o_Objeto: l.Implementacao{
			Contexto: "#todo",
			Macete:   "faz a funcao normal dps da clausula func entre parenteses -( defina a variavel que sera o this para acessar os campos do Objeto e qual é o tipo deste objeto a sera acessado)",
			Ditado:   "só de usar a var this entre parenteses referenciando o Tipo ja define como um metodo deste tipo para todas suas intancias usarem.",
			Exemplos: []string{
				`func (o ObjetoGO) ComportamentoGO1() string {
				return fmt.Sprintf("%s fez a Acao do ComportamentoGO1", o.C1)
			}`,
				``,
			},
		},
	},
}
