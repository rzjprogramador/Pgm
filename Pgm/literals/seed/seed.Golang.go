package seed

import "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"

var Golang = linguagem.Linguagem{
	Linguagem: "Golang",
	Tipos: linguagem.Tipos{
		TextoSingular: linguagem.TiposProps{
			Contrato:                 "string",
			ComoUsar:                 " dentro de 2 aspas",
			Interpolacao_De_Variavel: "%s",
		},
		TextoColecao: linguagem.TiposProps{
			Contrato:                 "[]string",
			ComoUsar:                 "[]string{ \"\", \"\",}",
			Interpolacao_De_Variavel: "todo",
		},
		TextoEmGeral: linguagem.TextoEmGeral{
			QuebraDeLinha: "Dentro de Aspas/crazes ex: `texto em cada linha aqui sem precisar usar \n a cada fim de linha desejada`"},
		NumeroInteiroSingular: linguagem.TiposProps{
			Contrato:                 "int64",
			ComoUsar:                 "somente numero, ex: 10",
			Interpolacao_De_Variavel: "%d",
		},
		NumeroInteiroColecao: linguagem.TiposProps{
			Contrato:                 "[]int64",
			ComoUsar:                 "[]int64{ \"\", \"\",}",
			Interpolacao_De_Variavel: "todo",
		},
		NumeroDecimalSingular: linguagem.TiposProps{
			Contrato:                 "float64",
			ComoUsar:                 "somente numerocom ponto flutuante, ex: 10",
			Interpolacao_De_Variavel: "%.2f - explicacao: por padrao o %f exibe 6 casas decimais - vc pode definir quantas casas quer mostrar após o % com ponto numeroDecasas e o f - exemplo: %2.f vai mostrar 2 numeros depois do ponto.",
		},
		NumerodecimalColecao: linguagem.TiposProps{
			Contrato:                 "[]float64",
			ComoUsar:                 "[]float64{ 10, 20, 30}",
			Interpolacao_De_Variavel: "naoTem",
		},
		QualquerValor: linguagem.TiposProps{
			Contrato:                 "[]any",
			ComoUsar:                 "[]any{ 10, 20, 30}",
			Interpolacao_De_Variavel: "%v",
		},
	},
	ObterInformacao: linguagem.ObterInformacao{
		Debugg: linguagem.Debug{
			Printar_no_Console:                           "fmt.Println( target) // obs: fmt é uma lib internal importavel",
			Printar_no_Console_com_VariaveisInterpoladas: "fmt.Printf ( \"texto com %demarcacaoVariveis\", variaveis por justa posicao conforme demarcacoes no texto) // obs: fmt é uma lib internal importavel",
		},
		ObterTipo:      "reflect.TypeOf( <target> ) // ou via console : fmt.Printf(\"Tipo: %T\", VARIAVEL_ALVO)",
		ObterInstancia: "use o metodo Name do typeOf e terao nome do struct em string ex: if reflect.TypeOf( TARGET ).Name() != \"NomeDostruct\" { ..faça algo",
	},
	Colecoes: linguagem.Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "var Items_Tipo = []Tipo{}",
		Add_item_na_colecao:                 "append( item ) ||ou|| Add(item)",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "#todo",
	},
	Libs: linguagem.Libs{
		PacotesConfiaveis: []string{"https://pkg.go.dev/"},
	},
	InterfaceDefinicaoDeNovosTipos: linguagem.InterfaceDefinicaoDeNovosTipos{
		Conceito: `
		contrato_interface: em golang para formar um obj onde o campo pode ser de um tipo ou de outro tipo de obj
		usamos a interface.
		o valor destes campos vem de funcoes_de_interface
		- esta interface tera funcoes que devolveram o valor desejado esta interface recebera o parametro Generico e pra ele o tipo que ele aceitara.
		- para definir este tipo se vai ser UmTipo | OutroTipo : criamos um type tipo struct mas ao inves da clausula struct usamos interface e dentro definimos que vai ser UMTipo ou OutroTipo desejado.
		`,
	},
	Regras: linguagem.Regras{
		Virgula_Apos_Fechamento_de_Objetos_Encadeados: "Sim",
	},

	Recursos: linguagem.Recursos{
		Apelido_Para_Importacoes: `
		em golang posso dar apelido para encurtar importacoes funciona como um obj e uso ele nas linahs abaixo ex: antes da importacao crio o apelido e uso ans linahs babaixo ex:
		import u "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"
		// uso o < u > como um objeto para acessar suas props importadas.
		`,
	},
}
