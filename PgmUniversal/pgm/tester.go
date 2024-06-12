package pgm

import "fmt"

func TesterPgm() {
	var r = RepoInMemoryPgm()

	fmt.Println(r.CreatePgm2(PgmGolang))
	fmt.Println(r.CreatePgm2(PgmJavascript))

	fmt.Println(r.GetAll())

	// fmt.Println(r.CreatePgm(PgmGolang))
	// fmt.Println(r.CreatePgm(PgmJavascript))
	// fmt.Println(Items_Pgm)
}
