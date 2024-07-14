package universal

type ConceitosUniversalPGM[N NomePastas, C Conceito, I Implementacoes] interface {
	NomePastaGuardaOutrasPastas() N
	NomePastaDeArquivosSoltos() N
	QuebraDeLinhas() C
	Implementacoes() I
}

type UniversalPGM struct {
	Significados                                         Significados
	Variavel_Guarda_TipoDeDado_Representado_por_um_Valor Variavel
}

type Convencao struct {
	Nomeacao Nomeacao
}

type Nomeacao struct {
	Contexto NovoConceito
}

type NovoConceito struct {
	Titulo    string
	Conceito  []string
	Exemplos  []string
	Tutoriais []Tutoriais
}

type Variavel struct {
	Conceito                      string
	Tipos_para_Variaveis          []string
	Uso_de_variaveis              string
	Consequencia_ao_usar_Variavel string
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
	Formulas                         Formulas
	CodigoUniversal                  CodigoUniversal
}

type Significados struct {
	Contexto NovoConceito
}

type Tutoriais struct {
	By   string
	Link string
}

type Tipos_Valor struct {
	Conceito                       string
	QualquerValor                  string
	TextoSingular                  string
	TextoColecao                   string
	TextoEmGeral                   string
	NumeroInteiroSingular          string
	NumeroInteiroColecao           string
	NumeroDecimalSingular          string
	NumerodecimalColecao           string
	Numero_Inteiro_SomentePositivo string
	Numero_Complexo                string
	Logico                         string
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

type CodigoUniversal struct {
	Objetos Objetos
}

type Objetos struct {
	Dar_Inteligencia_Ha_Instancias string
}
