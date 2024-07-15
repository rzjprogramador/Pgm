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
				Molde:    `cria-se uma interface com somente metodos que vai querer que os objetos que SEJAM UM DESTA INTERFACE tenham este metodo cada um implementando de um tipo mas com mesmo contratod e tipo de retorno.`,

				Exemplos: []string{
					`Esta AcaoFoo() foi definida na interface que esses objetos SÃO UM - ex: Veiculo e implementada por estes subobjetos vinculando assim como metodo deles.`,
					``,
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
