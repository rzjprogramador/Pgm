package partseed_ts

import l "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"

var Tipos_Typescript_SEED = l.Tipos{
	Conceito: ``,

	Texto_Caractere_Unico: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Texto_Conjunto_de_Caracteres: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Inteiro: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Real_Decimal: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_SomentePositivo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Numero_Complexo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Logico: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Objeto: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Erro: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Funcao: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Indefinido: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},

	Nulo: l.TiposProps{
		Tipo_Contrato:            ``,
		Representacao_ComoUsar:   ``,
		ValorZeroDefault:         ``,
		Interpolacao_De_Variavel: ``,
		ConjuntoDesteTipoDeValor: ``,
		Doc:                      ``,
		Metodos_Prototipos:       []string{""},
	},
}

/*
linguagem.Tipos{
			Conceito: `em typescript por default a variavel ao ser preenchida ou tipada se transformará em um objeto desencadeavel para acessar seus metodos prototipos herdados. ex: variavel.METODO_HERDADO()`,
			TextoSingular: linguagem.TiposProps{
				Tipo_Contrato:            "string",
				Representacao_ComoUsar:   " dentro de 2 aspas ou 1 aspas",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			TextoColecao: linguagem.TiposProps{
				Tipo_Contrato:            "[]string",
				Representacao_ComoUsar:   "[ \"\", \"\",]",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			TextoEmGeral: linguagem.TextoEmGeral{
				QuebraDeLinha: "Dentro de Aspas/crazes ex: `texto em cada linha aqui sem precisar usar \n a cada fim de linha desejada`"},

			NumeroInteiroSingular: linguagem.TiposProps{
				Tipo_Contrato:            "int",
				Representacao_ComoUsar:   "somente numero, ex: 10",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			NumeroInteiroColecao: linguagem.TiposProps{
				Tipo_Contrato:            "[]int",
				Representacao_ComoUsar:   "[10, 20, 30]",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			NumeroDecimalSingular: linguagem.TiposProps{
				Tipo_Contrato:            "float",
				Representacao_ComoUsar:   "somente numero com ponto flutuante, ex: 10",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			NumerodecimalColecao: linguagem.TiposProps{
				Tipo_Contrato:            "[]float",
				Representacao_ComoUsar:   "[10, 20, 30]",
				Interpolacao_De_Variavel: "naoTem",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},
			Livre_QualquerValor_ANY: linguagem.TiposProps{
				Tipo_Contrato:            "[]any",
				Representacao_ComoUsar:   "[]any{ 10, 20, 30}",
				Interpolacao_De_Variavel: "`${ variavel }` // sempre dentro de crazes ",
				Doc:                      `#TODO`,
				Metodos_Prototipos:       []string{""},
			},


*/
