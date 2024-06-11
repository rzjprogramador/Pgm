package pgm

var PgmGolang = Pgm{
	Linguagem: "Golang",
	ObterInformacao: ObterInformacao{
		ObterTipo:      "reflect.TypeOf( <target> )",
		ObterInstancia: "use o metodo Name do typeOf e terao nome do struct em string ex: if reflect.TypeOf( TARGET ).Name() != \"NomeDostruct\" { ..faça algo",
	},
	Colecoes: Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "var Items_Tipo = []Tipo{}",
		Add_item_na_colecao:                 "append( item ) ||ou|| Add(item)",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "#todo",
	},
}

var PgmJavascript = Pgm{
	Linguagem: "javascript",
	ObterInformacao: ObterInformacao{
		ObterTipo:      "typeof operando === \"tipoEmstring\"",
		ObterInstancia: "objeto instanceof Construtor/NomeClasseOrigem",
	},
	Colecoes: Colecoes{
		Criar_prop_de_colecao_inicial_vazia: "const items = []",
		Add_item_na_colecao:                 "colecao.push( item )",
		Remover_item_na_colecao:             "#todo",
		Mostrar_items_da_colecao:            "return items",
		Mostrar_item_Especifico_da_colecao:  "(identificador) => { #todo }",
	},
}
