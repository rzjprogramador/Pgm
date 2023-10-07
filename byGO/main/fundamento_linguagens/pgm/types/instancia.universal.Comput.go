package types_pgm

import "fmt"

func Comput() {
	comput := Comput_Fixo{
		metodologias: Metodologias{
			metodologia_desenho_de_objeto: Metodologia_Desenho_de_Objeto{
				metodologia_Args_e_Fixos: "campo_que_sera_valor_fixo: marretar o valor campo_que_sera_dinamico_sofrer_variacoes: receber via parametro.",
			},
		},

		linguagem: "universal",

		dado: Dado{
			temNaLinguagem: Habilitado{habilitado: true},
			conceito: Conceito{
				conceito: "",
			},
			valor:     "",
			parteDado: "",
		},
	}
	fmt.Println(comput)
}

func main() {}
