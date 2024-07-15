package partseed_go

import l "github.com/rzjprogramador/Pgm/Pgm/linguagem"

var Algoritmo_GO = l.Algortimo_LING{
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

	Composicao: l.Composicao{
		Conceito: "Campo de um objeto que Delega a outro objeto fazer a acao por ele.",
		Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes: l.Contexto{
			Titulo:   "Campo_Que_Pode_Variar_Tendo_Mesmas_Acoes",
			Objetivo: `ter um objPai. carro ou moto . vao_Ter_Acao_Foo()`,
			Molde:    `cria-se uma interface com somente metodos que vai querer que os objetos que SEJAM UM DESTA INTERFACE tenham este metodo cada um implementando de um tipo mas com mesmo contratod e tipo de retorno.`,

			Exemplos: []string{
				`Esta AcaoFoo() foi definida na interface que esses objetos SÃO UM - ex: Veiculo e implementada por estes subobjetos vinculando assim como metodo deles.`,
				``,
			},
		},
	},
}
