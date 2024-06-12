package pgm

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

type RepoPgm struct {
	Items *arraylist.List
	// Items []Pgm
	// Create string

}

func (r RepoPgm) CreatePgm2(p Pgm) string {
	r.Items.Add(p)
	return "Criado com sucesso!"
}

func (r RepoPgm) GetAll() *arraylist.List {
	return r.Items
}

func RepoInMemoryPgm() RepoPgm {
	repo := RepoPgm{
		Items: arraylist.New(),
	}
	return repo
}
