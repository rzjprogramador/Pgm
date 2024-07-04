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

	Algoritmos: PartSeed_Algoritmos,

	Dev: u.Dev{
		Lib: []string{"funcoes/modulos externas importaveis", "pacote"},
	},
}