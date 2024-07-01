package partseed_universal

import u "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"

var SignificadosSEED = u.Significados{
	Computador: u.Computador{
		O_Computador_entende: `somente binarios 0 e 1 onde 0 é falso , circuito apagado e 1 é verdadeiro circuito acesso.`,
	},

	Linguagem: u.Linguagem{
		Linguagem_De_Programacao: `Linguagem é o idioma onde criaremos o codigo fonte das nossas insyrucoes, e para o pc entender precisara de um interpretador ou um compilador para entregar ao pc em binario.`,

		Linguagem_Compilada: `Sao todas que cokpilamos o codigo fonte para binario, obs: as compiladas sao mais rapidas por ja entregarem em binario e o pc processa naturalmentevmais rapido.
		Ex: [ C, Golang, python, java,]`,

		Linguagem_Interpretada: `Sao todas que nao compilam para binario, para serem entendidas pelo pc precisa de um interpretador
		Ex: [ Javascript, typescript,  ]`,
	},

	Algoritmos: u.Algoritmos{
		Pensamento_Declaracoes: `Escopo TalNome : É UMA estrutura { E VAI TER ... }`,
		Metodos_Computaveis:    `São metodos/funcoes ligadas a estrutura que criou o tipo - para ser usado basta desencadear com . ponto a instancia que estará cjhamando o metodo/funcao disponivel, obs: toda instancia herda metodos da sua estrutura.`,
		Valor:                  []string{"representacao de um tipo", "variavel", "objeto", "funcao que retorna algo"},
		Valor_e_seus_Tipos: u.Tipos_Valor{
			Conceito: `o valor singular da origem a inferencia de tipo ( ao preencher uma variavel ela ja se torna o tipo do valor inserido)
			e na maioria das linguagens a variavel a ser preenchida com valor vira um objeto e seus metodos_herdados prototipos vem da estrutura/classe que criou este tipo e quando esta variavel/objeto é preenchida pode-se desencadear com .ponto para acessar os metodos prototipos deste tipo que tem o poder fazer algo com este tipo de dado/valor.`,
			QualquerValor:                  `na maioria das linguagens tipadas é do tipo any (aceita qualquer tipo de valor) ou interface{} vazia.`,
			TextoSingular:                  `è um conjunto de char caracteres, é iteravel por ser um array de caracteres char/unidade de strings.`,
			TextoColecao:                   `#TODO`,
			TextoEmGeral:                   `#TODO`,
			NumeroInteiroSingular:          `#TODO`,
			NumeroInteiroColecao:           `#TODO`,
			NumeroDecimalSingular:          `#TODO`,
			NumerodecimalColecao:           `#TODO`,
			Numero_Inteiro_SomentePositivo: `#TODO`,
			Numero_Complexo:                `#TODO`,
			Logico:                         `#TODO`,
		},
	},

	Dev: u.Dev{
		Lib: []string{"funcoes/modulos externas importaveis", "pacote"},
	},
}