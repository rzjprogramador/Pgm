package linguagem

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

// tipagem da entidade principal
type Linguagem struct {
	Linguagem                      string
	Valor                          Valor
	ObterInformacao                ObterInformacao
	Colecoes                       Colecoes
	Libs                           Libs
	InterfaceDefinicaoDeNovosTipos InterfaceDefinicaoDeNovosTipos
	Regras                         Regras
	Recursos                       Recursos
	Codigo                         Codigo
}

type Valor struct {
	Declaracao_Completa string
	Declaracao_Inferida string
	Tipos_de_Valor      Tipos
}

type Tipos struct {
	Conceito                     string
	Texto_Caractere_Unico        TiposProps
	Texto_Conjunto_de_Caracteres TiposProps
	Numero_Inteiro               TiposProps
	Numero_Real_Decimal          TiposProps
	Numero_SomentePositivo       TiposProps
	Numero_Complexo              TiposProps
	Logico                       TiposProps
	Objeto                       TiposProps
	Erro                         TiposProps
	Funcao                       TiposProps
	Indefinido                   TiposProps
	Nulo                         TiposProps
}

type TiposProps struct {
	Tipo_Contrato            string
	Representacao_ComoUsar   string
	ValorZeroDefault   string
	Interpolacao_De_Variavel string
	ConjuntoDesteTipoDeValor string
	Doc                      string
	Metodos_Prototipos       []string

	// representacao: , tipo: , doc: , metodos: ,
}

type TextoEmGeral struct {
	QuebraDeLinha string
}

type ObterInformacao struct {
	Debugg         Debug
	ObterTipo      string
	ObterInstancia string
}

type Debug struct {
	Printar_no_Console                           string
	Printar_no_Console_com_VariaveisInterpoladas string
}

type Colecoes struct {
	Criar_prop_de_colecao_inicial_vazia string
	Add_item_na_colecao                 string
	Remover_item_na_colecao             string
	Mostrar_items_da_colecao            string
	Mostrar_item_Especifico_da_colecao  string
}

type Libs struct {
	PacotesConfiaveis []string
}

// Tipagem Repositorio -- para salvar dados da entidade principal
type linguagemRepository struct {
	Items *arraylist.List
	// Obs: Items nao tem problema ser um tipo externo de lib - porque é usado somente em memoria - na real o bd nao tem items
}

// metodos para ser um :Repositorio
type IlinguagemRepository[C ICollectionlinguagem] interface {
	Createlinguagem(p Linguagem) string
	GetAll() C
}

// opcoes dinamicas de tipos
type ICollectionlinguagem interface {
	[]Linguagem | *arraylist.List
}

type InterfaceDefinicaoDeNovosTipos struct {
	Conceito string
}

type Regras struct {
	Virgula_Apos_Fechamento_de_Objetos_Encadeados string
}

type Recursos struct {
	Apelido_Para_Importacoes string
}

type Codigo struct {
	Pensamento_Declaracoes                               string
	Estrutura_Modeladora_de_Objeto_Instancia_de_NovoTipo string
	Gerar_Instancia                                      string
	Campos_Fixo_na_Instancia                             string
	Campos_Dinamicos_na_Instancia                        string
	Dar_Intelegencia_Ha_Instancia                        string
}
