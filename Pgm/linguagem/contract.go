package linguagem

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

// tipagem da entidade principal
type Linguagem struct {
	Linguagem                      string
	Tipos                          Tipos
	ObterInformacao                ObterInformacao
	Colecoes                       Colecoes
	Libs                           Libs
	InterfaceDefinicaoDeNovosTipos InterfaceDefinicaoDeNovosTipos
	Regras                         Regras
	Recursos                       Recursos
}

type Tipos struct {
	TextoSingular         TiposProps
	TextoColecao          TiposProps
	TextoEmGeral          TextoEmGeral
	NumeroInteiroSingular TiposProps
	NumeroInteiroColecao  TiposProps
	NumeroDecimalSingular TiposProps
	NumerodecimalColecao  TiposProps
	QualquerValor         TiposProps
}

type TextoEmGeral struct {
	QuebraDeLinha string
}

type TiposProps struct {
	Contrato                 string
	ComoUsar                 string
	Interpolacao_De_Variavel string
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
