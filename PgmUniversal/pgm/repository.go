package pgm

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

type PgmRepositoy[P PgmArray] interface {
	CreatePgm(r RepoPgm) string
	GetAll() P
}

type PgmArray interface {
	*arraylist.List | []Pgm
}
