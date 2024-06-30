package partseed_go

import l "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"

var Tipos_Golang_SEED = l.Tipos{
	Conceito: ``,

	Texto_Caractere_Unico: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Texto_Conjunto_de_Caracteres: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Inteiro: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Real_Decimal: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_SomentePositivo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Complexo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Logico: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Objeto: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Erro: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Funcao: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Indefinido: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Nulo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault: ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},
}

/*
linguagem.Tipos{
			Conceito: `em golang por default a variavel com valor nNÃO vira um objeto desencadeavel para acessar seus metodos herdados
			os metodos de cada tipo vem do objeto importavel correspondente ao tipo do valor ex: para usar
			metodos para :`` importamos [ o objeto :``, para numero number, para matematica math, ] e usamos o objeto e deste objeto desencadeamos seus metodos :``.METODO_HERDADO()

			METODOS_PROTOTIPOS : algumas funcoes para todos os tipos estao soltas na linguagem ex: para ver tamanho o length é a funcao solat len( alvo )

			`,
			TextoSingular: linguagem.l.TiposProps{
				Tipo_Contrato:            ":``",
				Representacao_ComoUsar:   " dentro de 2 aspas",
				Interpolacao_De_Variavel: "%s",
				Doc:                      `#TODO`,
				Metodos_Prototipos:      : []:``{"pa"",cote prototipos : :``s https://pkg.go.dev/:``s"},
			},
			TextoColecao: linguagem.l.TiposProps{
				Tipo_Contrato:            "[]:``",
				Representacao_ComoUsar:   "[]:``{ \"\", \"\",}",
				Interpolacao_De_Variavel: "todo",
				Doc:                      `#TODO`,
				Metodos_Prototipos:      : []:``{`#T"",ODO`},
			},
			TextoEmGeral: linguagem.TextoEmGeral{
				QuebraDeLinha: "Dentro de Aspas/crazes ex: `texto em cada linha aqui sem precisar usar \n a cada fim de linha desejada`"},
			NumeroInteiroSingular: linguagem.l.TiposProps{
				Tipo_Contrato:            "int64",
				Representacao_ComoUsar:   "somente numero, ex: 10",
				Interpolacao_De_Variavel: "%d",
				Doc:                      `https://go.dev/tour/basics/11`,
				Metodos_Prototipos:      : []:``{"pa"",cote prototipos number: https://pkg.go.dev/golang.org/x/text/number", "matematica https://pkg.go.dev/math"},
			},
			NumeroInteiroColecao: linguagem.l.TiposProps{
				Tipo_Contrato:            "[]int64",
				Representacao_ComoUsar:   "[]int64{ \"\", \"\",}",
				Interpolacao_De_Variavel: "todo",
				Doc:                      `#TODO`,
				Metodos_Prototipos:      : []:``{`#T"",ODO`},
			},
			NumeroDecimalSingular: linguagem.l.TiposProps{
				Tipo_Contrato:            "float64",
				Representacao_ComoUsar:   "somente numerocom ponto flutuante, ex: 10",
				Interpolacao_De_Variavel: "%.2f - explicacao: por padrao o %f exibe 6 casas decimais - vc pode definir quantas casas quer mostrar após o % com ponto numeroDecasas e o f - exemplo: %2.f vai mostrar 2 numeros depois do ponto.",
				Doc:                      `https://go.dev/tour/basics/11`,
				Metodos_Prototipos:      : []:``{`#T"",ODO`},
			},
			NumerodecimalColecao: linguagem.l.TiposProps{
				Tipo_Contrato:            "[]float64",
				Representacao_ComoUsar:   "[]float64{ 10, 20, 30}",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:      : []:``{`#T"",ODO`},
			},
			Livre_QualquerValor_ANY: linguagem.l.TiposProps{
				Tipo_Contrato:            "[]any",
				Representacao_ComoUsar:   "[]any{ 10, 20, 30}",
				Interpolacao_De_Variavel: "%v",
				Doc:                      `#TODO`,
				Metodos_Prototipos:      : []:``{`#T"",ODO`},
			},
		},
*/
