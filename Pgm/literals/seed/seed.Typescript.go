package seed

import "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"

var Typescript = linguagem.Linguagem{
	Linguagem: "Typescript",
	Tipos: linguagem.Tipos{
		TextoSingular: linguagem.TiposProps{
			Contrato:                 "string",
			ComoUsar:                 " dentro de 2 aspas ou 1 aspas",
			Interpolacao_De_Variavel: "naoTem",
		},
		TextoColecao: linguagem.TiposProps{
			Contrato:                 "[]string",
			ComoUsar:                 "[ \"\", \"\",]",
			Interpolacao_De_Variavel: "naoTem",
		},
		TextoEmGeral: linguagem.TextoEmGeral{
			QuebraDeLinha: "Dentro de Aspas/crazes ex: `texto em cada linha aqui sem precisar usar \n a cada fim de linha desejada`"},

		NumeroInteiroSingular: linguagem.TiposProps{
			Contrato:                 "int",
			ComoUsar:                 "somente numero, ex: 10",
			Interpolacao_De_Variavel: "naoTem",
		},
		NumeroInteiroColecao: linguagem.TiposProps{
			Contrato:                 "[]int",
			ComoUsar:                 "[10, 20, 30]",
			Interpolacao_De_Variavel: "naoTem",
		},
		NumeroDecimalSingular: linguagem.TiposProps{
			Contrato:                 "float",
			ComoUsar:                 "somente numero com ponto flutuante, ex: 10",
			Interpolacao_De_Variavel: "naoTem",
		},
		NumerodecimalColecao: linguagem.TiposProps{
			Contrato:                 "[]float",
			ComoUsar:                 "[10, 20, 30]",
			Interpolacao_De_Variavel: "naoTem",
		},
		QualquerValor: linguagem.TiposProps{
			Contrato:                 "[]any",
			ComoUsar:                 "[]any{ 10, 20, 30}",
			Interpolacao_De_Variavel: "`${ variavel }` // sempre dentro de crazes ",
		},
	},
	ObterInformacao: linguagem.ObterInformacao{
		Debugg: linguagem.Debug{
			Printar_no_Console:                           "console.log( target)",
			Printar_no_Console_com_VariaveisInterpoladas: "console.log( variaveis, variaveis, `${variavelInterpolada dentro de crazes}`)",
		},
		ObterTipo:      "typeof operando === \"tipoEmstring\"",
		ObterInstancia: "objeto instanceof Construtor/NomeClasseOrigem",
	},
	Colecoes: linguagem.Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "const items = []",
		Add_item_na_colecao:                 "colecao.push( item )",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "(identificador) => { #todo }",
	},
	Libs: linguagem.Libs{
		PacotesConfiaveis: []string{"https://www.w3schools.com/js/", "https://developer.mozilla.org/pt-BR/docs/Web/Typescript"},
	},
	InterfaceDefinicaoDeNovosTipos: linguagem.InterfaceDefinicaoDeNovosTipos{
		Conceito: `
		#TODO
		`,
	},
	Regras: linguagem.Regras{
		Virgula_Apos_Fechamento_de_Objetos_Encadeados: "Opcional",
	},

	Recursos: linguagem.Recursos{
		Apelido_Para_Importacoes: `
		em typescript uso a clausula < as > para apelidar ou afirmar ser de um Tipo // OU : dois pontos para apelidar um objeto,
		ex_1_apelidar_tipo : objetoFoo: TipoA  as TipoQueSera
		ex_2_apelidar_objeto : import { objetoA: objetoApelido } from "./algumaLugar/"
		`,
	},

	Codigo: linguagem.Codigo{
		Estruturar_NovoTipo_Para_Gerar_Instancia: `Criar uma classe ex: class NomeDoNovoTipo {}`,
		Gerar_Instancia:                          `const i1 = new NomeDoNovoTipo()`,
		Campos_Fixo_na_Instancia:                 `na parte alta da classe inserir campos com valores pré-determinados`,
		Campos_Dinamicos_na_Instancia:            `Receber via constructor parametros que serao so campos dinamicos da intancia ao ser gerada`,
		Dar_Intelegencia_Ha_Instancia:            `Criar Metodos_Computaveis. ex: `,
	},
}
