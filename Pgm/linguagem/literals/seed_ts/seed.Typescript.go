package seed_ts

import (
	l "github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"
	"github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem/literals/seed_ts/partseed_ts"
)

var TypescriptSEED = l.Linguagem{
	Linguagem: "Typescript",

	Valor: l.Valor{
		Declaracao_Completa: `const nome: string = "valor"`,
		Declaracao_Inferida: `nome = "valor"`,

		Tipos_de_Valor: partseed_ts.Tipos_Typescript_SEED,
	},

	ObterInformacao: l.ObterInformacao{
		Debugg: l.Debug{
			Printar_no_Console:                           "console.log( target)",
			Printar_no_Console_com_VariaveisInterpoladas: "console.log( variaveis, variaveis, `${variavelInterpolada dentro de crazes}`)",
		},
		ObterTipo:      "typeof operando === \"tipoEmstring\"",
		ObterInstancia: "objeto instanceof Construtor/NomeClasseOrigem",
	},
	Colecoes: l.Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "const items = []",
		Add_item_na_colecao:                 "colecao.push( item )",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "(identificador) => { #todo }",
	},
	Libs: l.Libs{
		PacotesConfiaveis: []string{"https://www.w3schools.com/js/", "https://developer.mozilla.org/pt-BR/docs/Web/Typescript"},
	},
	InterfaceDefinicaoDeNovosTipos: l.InterfaceDefinicaoDeNovosTipos{
		Conceito: `
		#TODO
		`,
	},
	Regras: l.Regras{
		Virgula_Apos_Fechamento_de_Objetos_Encadeados: "Opcional",
	},

	Recursos: l.Recursos{
		Apelido_Para_Importacoes: `
		em typescript uso a clausula < as > para apelidar ou afirmar ser de um Tipo // OU : dois pontos para apelidar um objeto,
		ex_1_apelidar_tipo : objetoFoo: TipoA  as TipoQueSera
		ex_2_apelidar_objeto : import { objetoA: objetoApelido } from "./algumaLugar/"
		`,
	},

	Codigo: l.Algortimo_LING{
		Pensamento_Declaracoes:                               `Escopo TalNome : É UMA estrutura { E VAI TER ... } ex:  const talNome: Object = { tera: 10, teratambem: true}`,
		Estrutura_Modeladora_de_Objeto_Instancia_de_NovoTipo: `Criar uma classe ex: class NomeDoNovoTipo {}`,
		Gerar_Instancia:                                      `const i1 = new NomeDoNovoTipo()`,
		Campos_Fixo_na_Instancia:                             `na parte alta da classe inserir campos com valores pré-determinados`,
		Campos_Dinamicos_na_Instancia:                        `Receber via constructor parametros que serao so campos dinamicos da intancia ao ser gerada`,
		Dar_Intelegencia_Ha_Instancia:                        `Criar Metodos_Computaveis. ex: `,

		Campos_Com_Valores_Default: l.Algoritmo_Type_LING{
			Algoritmo: `no alto da classe inserir campo default: valor`,
			Exemplo:   `#todo`,
		},
	},
}
