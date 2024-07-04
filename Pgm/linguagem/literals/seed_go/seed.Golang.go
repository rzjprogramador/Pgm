package seed_go

import (
	l "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"
	"github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem/literals/seed_go/partseed_go"
)

var GolangSEED = l.Linguagem{
	Linguagem: "Golang",

	Valor: l.Valor{
		Declaracao_Completa: `var Nome = "valor"`,
		Declaracao_Inferida: `nome := "valor"`,

		Tipos_de_Valor: partseed_go.Tipos_Golang_SEED,
	},

	ObterInformacao: l.ObterInformacao{
		Debugg: l.Debug{
			Printar_no_Console:                           "fmt.Println( target) // obs: fmt é uma lib internal importavel",
			Printar_no_Console_com_VariaveisInterpoladas: "fmt.Printf ( \"texto com %demarcacaoVariveis\", variaveis por justa posicao conforme demarcacoes no texto) // obs: fmt é uma lib internal importavel",
		},
		ObterTipo:      "reflect.TypeOf( <target> ) // ou via console : fmt.Printf(\"Tipo: %T\", VARIAVEL_ALVO)",
		ObterInstancia: "use o metodo Name do typeOf e terao nome do struct em string ex: if reflect.TypeOf( TARGET ).Name() != \"NomeDostruct\" { ..faça algo",
	},
	Colecoes: l.Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "var Items_Tipo = []Tipo{}",
		Add_item_na_colecao:                 "append( item ) ||ou|| Add(item)",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "#todo",
	},
	Libs: l.Libs{
		PacotesConfiaveis: []string{"https://pkg.go.dev/"},
	},
	InterfaceDefinicaoDeNovosTipos: l.InterfaceDefinicaoDeNovosTipos{
		Conceito: `
		contrato_interface: em golang para formar um obj onde o campo pode ser de um tipo ou de outro tipo de obj
		usamos a interface.
		o valor destes campos vem de funcoes_de_interface
		- esta interface tera funcoes que devolveram o valor desejado esta interface recebera o parametro Generico e pra ele o tipo que ele aceitara.
		- para definir este tipo se vai ser UmTipo | OutroTipo : criamos um type tipo struct mas ao inves da clausula struct usamos interface e dentro definimos que vai ser UMTipo ou OutroTipo desejado.
		`,
	},
	Regras: l.Regras{
		Virgula_Apos_Fechamento_de_Objetos_Encadeados: "Sim",
	},

	Recursos: l.Recursos{
		Apelido_Para_Importacoes: `
		em golang posso dar apelido para encurtar importacoes funciona como um obj e uso ele nas linahs abaixo ex: antes da importacao crio o apelido e uso ans linahs babaixo ex:
		import u "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"
		// uso o < u > como um objeto para acessar suas props importadas.
		`,
	},

	Codigo: l.Algortimo_LING{
		Pensamento_Declaracoes:                               `Escopo TalNome - É UMa estrutura { E VAI TER ... } ex: type TalNome struct { tera... } `,
		Estrutura_Modeladora_de_Objeto_Instancia_de_NovoTipo: `Criar um struct/estrutura. ex: type NomeDoNovoTipo struct { ... campos doTipo ... }`,
		Gerar_Instancia:                                      `func funcaoNomeDoNovoTipo(t Tipo) { return t}`,
		Campos_Fixo_na_Instancia:                             `#TODO`,
		Campos_Dinamicos_na_Instancia:                        `#TODO`,
		Dar_Intelegencia_Ha_Instancia:                        `Criar Metodos_Computaveis. ex: `,

		Campos_Com_Valores_Default: l.Algoritmo_Type_LING{
			Algoritmo: `no struct inserir campo default: valor`,
			Exemplo:   `#todo`,
		},
	},
}
