package pgm

import "fmt"

var Items_Pgm = []Pgm{}

func CreatePgm(p Pgm) string {
	_ = append(Items_Pgm, p)
	return fmt.Sprintf("Criado com sucesso! %s")
}

func GetAllPGM() *[]Pgm {
	return &Items_Pgm

	// res := []Pgm{}
	// for _, x := range Items_Pgm {
	// 	res = []Pgm{x}
	// }
	// return res
}
