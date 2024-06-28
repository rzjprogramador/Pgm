package seed

import u "github.com/rzjprogramador/Pgm_Universal/Pgm/universal"

func NomePastaGuardaOutrasPastas() u.Singular {
	return u.Singular{Singular: "#todo"}
}
func NomePastaDeArquivosSoltos() u.Singular {
	return u.Singular{Singular: "#todo"}
}
func QuebraDeLinhas() string {
	return "#todo"
}

func Implementacoes() u.Implementacoes {
	return u.Implementacoes{
		Significados: u.Significados{
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
<<<<<<< HEAD
				Pensamento_Declaracoes: `Escopo TalNome : É UMA estrutura { E VAI TER ... }`,
				Metodos_Computaveis:    `São metodos/funcoes ligadas a estrutura que criou o tipo - para ser usado basta desencadear com . ponto a instancia que estará cjhamando o metodo/funcao disponivel, obs: toda instancia herda metodos da sua estrutura.`,
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
=======
				Metodos_Computaveis: `São metodos/funcoes ligadas a estrutura que criou o tipo - para ser usado basta desencadear com . ponto a instancia que estará chamando o metodo/funcao disponivel, obs: toda instancia herda metodos da sua estrutura, use para poder dar poderes inteligencia as instancias, entre elas poder COMPUTAR Campos, se modificar, Fazer coisas que só quem é deste tipo pode sem precisar de funcoes externas ou soltas, etc.`,
>>>>>>> 3c6734da78d3b35ae6164ffebb8008011d5656dc
			},
		},

		Obrigatorio_atualizar: `Faça um Wrapper em toda app para que só a possa usar ATUALIZADA - amarre o funcionamento na versao da app.`,
		Log:                   `crie logs para todas funcoes - para controlar quem esta usando-as`,
		Origem_Campo:          ` origem_campo : Se o campo ou 1 de seus campos no caso de obj  for somente um primitivo sem inteligencia nao precisa ser gerado por uma estrutura classificatoria, caso contrario se precisar de inteligencia (operacoes, crud, logicas)  faça sua geracao vir de uma estrutura ou classe.`,
		Debuggar_Console: `Sempre rode com depuracao debugg
		`,

		Gerador_De_Objeto_Via_Construtor: u.Gerador_De_Objeto_Via_Construtor{
			O_que_vai_no_construtor: `No construtor vão os campos com valor : que ganharam , ou viram com valor , ou executam algo, quando a classe é instanciada.`,
		},

		Operacoes: u.Operacoes{
			Ajuntar_Com_Valor_que_Ja_tem: `Ajuntar com o valor que ja tem : += ex: operandoValor operacao operandoDeMudanca <se for string vai concatenar se for numero vai somar.>`,
		},

		Loop_For: `No loop crio a variavel key que a cada volta do loop vai set a chave do objAlvo, com target[key] tenho o valor contido nesta chave capturada e com ele faco algo {},
		*/

		for ( const key in target ) {
			target[key]
		}

		// logica :for in modificando valor de obj

		for( const key in objAlvo) {
				objAlvo[key] += mudanca
				// aqui tenho o valor a cada loop faco alguma operacao nele
				}
			return objAlvo
			// aqui fora do for reutilizando a variavel modificada tenho o alvo modificado ou retorno ou continuo usando-o ja afetado`,

		Formulas: u.Formulas{
			Porcentagem_entre_2_valores: `https://youtu.be/Jd_wrMJKdQg?si=fg8MWYJrFOQ3koJb`,
		},

		CodigoUniversal: u.CodigoUniversal{
			u.Objetos{
				Dar_Inteligencia_Ha_Instancias: `Criar Metodos_Computaveis que daram acoes as instancias com seus proprios campos - ou seja as informacoes que elas tiverem poderam ser COMPUTADAS entre eles mesmos e gerar resultados. obs: use somente para gerar Computados.`,
			},
		},
	}
}

// #TODO Formar um objeto com essas funcoes

// func ConceitosUniversalPGM_FN () u.Universal {

// 	return {
// 		NomeDePastas:         NomePastaGuardaOutrasPastas(),
// 		NomeDeArquivosSoltos: NomePastaDeArquivosSoltos(),
// 		QuebraDeLinhas:       QuebraDeLinhas(),
// 	}
// }
