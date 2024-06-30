package seed

import u "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"

var VariavelTipoValorSEED = u.Variavel{
	Conceito: `é uma CELULA unidade na memoria já preenchida com o valor zero dependendo do seu tipo de dado esperado, alias sempre será de um TIPO de dado prometido/demarcado pelo dev e sempre será REPRESENTADO por um VALORcorrespondente ao tipo prometido`,

	Tipos_para_Variaveis: []string{
		"Caractere_Unico",
		"Texto_Conjunto_de_Caracteres",
		"Numero_Inteiro",
		"Numero_Real_Decimal",
		"Objeto",
		"Erro",
		"Funcao",
		"Indefinido",
		"Nulo",
		"Livre_QualquerValor_ANY",
	},

	Uso_de_variaveis: `O pc é como uma calculadora guarda 1 valor na celula/memoria, guarda a operacao pedida, guarda outros valores pedidos  , guarda o pedido de resultado ex: = igual, Manda pro processador que devolve outra celula com o resultado.`,

	Consequencia_ao_usar_Variavel: `ao guardar um valor na celula somente apos uma execucao e por padrao guarda na memoria volatil, para esse valor persistir tem que fazer funcionalidade pra guardar e gravar na memoria naoVolatil
	Obs: Cada passo estara em uma celula variavel volatil por padrao ao ser executado/chamado ou naoVolatil se for gravado os dados.
	Entao só vai ocupar a memoria se for chamado e desocupará quando for reiniciado ou fechado o local de memoria usada.`,
}
