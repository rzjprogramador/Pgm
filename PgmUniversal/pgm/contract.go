package pgm

type Pgm struct {
	Items           []any
	Linguagem       string
	ObterInformacao ObterInformacao
	Colecoes        Colecoes
}

type ObterInformacao struct {
	ObterTipo      string
	ObterInstancia string
}

type Colecoes struct {
	Criar_prop_de_colecao_inicial_vazia string
	Add_item_na_colecao                 string
	Remover_item_na_colecao             string
	Mostrar_items_da_colecao            string
	Mostrar_item_Especifico_da_colecao  string
}
