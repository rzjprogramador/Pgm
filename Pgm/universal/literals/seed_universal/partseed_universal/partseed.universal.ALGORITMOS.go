package partseed_universal

import "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"

var PartSeed_Algoritmos = universal.Algoritmos{

	Pensamento_Declaracoes: `Escopo TalNome : É UMA estrutura { E VAI TER ... }`,

	Metodos_Computaveis: `São metodos/funcoes ligadas a estrutura que criou o tipo - para ser usado basta desencadear com . ponto a instancia que estará cjhamando o metodo/funcao disponivel, obs: toda instancia herda metodos da sua estrutura.`,

	Valor: []string{"representacao de um tipo", "variavel", "objeto", "funcao que retorna algo"},

	Valor_e_seus_Tipos: universal.Tipos_Valor{
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

	Campos_Com_Valores_Default_UNIV: universal.Algoritmo_Type_UNIV{
		Algoritmo: `Fazer uma funcao que avalia no caso de texto : se o tamanho do texto recebido for menor ou igual a 0 da um valor fixo ao campo - caso contrario vai deixar o valor recebido.`,

		Exemplo: `// Exemplo:

		const objRequest = {
		 C1: "a",
		 C2: "",
	 }

	 function configurando_Obj(o) {

		 if (o.C1.length <= 0) {
			 o.C1 = "valorFIXO_1"
		 }
		 if (o.C2.length <= 0) {
			 o.C2 = "valorFIXO_2"
		 }

		 const res = {
			 C1: o.C1,
			 C2: o.C2,
		 }
		 return res
	 }

	 console.log( configurando_Obj(objRequest) )
 `,
	},
}
