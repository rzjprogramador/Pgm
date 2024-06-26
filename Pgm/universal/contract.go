package universal

type ConceitosUniversalPGM[N NomePastas, C Conceito, I Implementacoes] interface {
	NomePastaGuardaOutrasPastas() N
	NomePastaDeArquivosSoltos() N
	QuebraDeLinhas() C
	Implementacoes() I
}

type NomePastas interface {
	Singular | Plural
}

type Singular struct {
	Singular string
}

type Plural struct {
	Plural string
}

type Exemplo struct {
	Exemplo string
}

type Conceito struct {
	Conceito string
}

type Implementacoes struct {
	Obrigatorio_atualizar            string
	Log                              string
	Origem_Campo                     string
	Debuggar_Console                 string
	Gerador_De_Objeto_Via_Construtor Gerador_De_Objeto_Via_Construtor
	Operacoes                        Operacoes
	Loop_For                         string
	Formulas Formulas
}

type Gerador_De_Objeto_Via_Construtor struct {
	O_que_vai_no_construtor string
}

type Operacoes struct {
	Ajuntar_Com_Valor_que_Ja_tem string
}

type Formulas struct {
	Porcentagem_entre_2_valores string
}
