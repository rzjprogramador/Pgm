package types_pgm

type Args_Dinamico struct {
	temNaLinguagem         Habilitado
	linguagem              string
	dadoViaFuncao          Dado_Gerado_Por_Funcao_Classificatorias
	referencia             Referencia
	metodo_Para_NovosTipos Metodo_Para_Novos_Tipos
	tipo_personal_Novo     Tipo_Personal_Novo
	ferramentas_Operacoes  Ferramentas_Operacoes
}

type Dado_Gerado_Por_Funcao_Classificatorias struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	via_Classe     Classe
	via_Struct     Struct
	exemplo        Exemplo
}

type Classe struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	construtor     string
}

type Struct struct {
	temNaLinguagem Habilitado
	conceito       Conceito
	modelagem      string
	exemplo        Exemplo
}

type Referencia struct {
	palavrasChaveReferencia PalavrasChaveReferencia
	declaracaoReferencia    Declaracao_Referencia
	exemplo                 Exemplo
}

type Tipo_Personal_Novo struct {
	temNaLinguagem       Habilitado
	conceito             Conceito
	criar_estrutura      string
	criar_acoes_metodos  string
	uso_em_escopo_de_uso string
	diferenciais         Diferencias_da_Linguagem
	exemplo              Exemplo
}

type Metodo_Para_Novos_Tipos struct {
	acionar_via string
}

type Ferramentas_Operacoes struct {
	temNaLinguagem Habilitado
	devolve_texto  Devolve_Texto
}

type Devolve_Texto struct {
	return_texto                   string
	texto_interpolado_com_variavel string
}
