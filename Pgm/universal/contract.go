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
	Significados                     Significados
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
	Computador Computador
	Linguagem  Linguagem
	Algoritmos Algoritmos
	Dev Dev
}

type Computador struct {
	O_Computador_entende string
}

type Linguagem struct {
	Linguagem_De_Programacao string
	Linguagem_Compilada      string
	Linguagem_Interpretada   string
}
type Algoritmos struct {
	Pensamento_Declaracoes string
	Metodos_Computaveis string
	Valor []string
	Valor_e_seus_Tipos Tipos_Valor
}

type Tipos_Valor struct {
	Conceito string
	QualquerValor         string
	TextoSingular         string
	TextoColecao          string
	TextoEmGeral          string
	NumeroInteiroSingular string
	NumeroInteiroColecao  string
	NumeroDecimalSingular string
	NumerodecimalColecao  string
	Numero_Inteiro_SomentePositivo string
	Numero_Complexo string
	Logico string

}

type Dev struct {
	Lib []string
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
