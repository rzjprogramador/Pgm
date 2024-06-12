package pgm

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

// var Items_Pgm = arraylist.New()

func ItemsPgm() *arraylist.List{
	items := arraylist.New()
	return items
}

func (t Pgm) CreatePgm(p Pgm) string {
	ItemsPgm().Add(p)
	return "Criado com sucesso!"
}


// func GetAllPGM() *arraylist.List {
// 	return Items_Pgm
// }
