package partseed_ts

import l "github.com/rzjprogramador/Pgm/Pgm/linguagem"

var Algoritmo_TS = l.Algortimo_LING{
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

	Composicao: l.Composicao{
		Conceito: "Campo de um objeto que Delega a outro objeto fazer a acao por ele.",
		Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes: l.Contexto{
			Titulo: "Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes",
			Objetivo: `ter um objPai. carro ou moto . vao_Ter_Acao_Foo()`,
			Molde:  `#todo`,

			Exemplos: []string{
				`#todo`,
				``,
			},
		},
	},
}
