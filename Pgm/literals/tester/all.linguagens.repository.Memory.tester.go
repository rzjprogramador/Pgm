package tester

import (
	"fmt"

	"github.com/rzjprogramador/Pgm_Universal/Pgm/linguagem"
	seed "github.com/rzjprogramador/Pgm_Universal/Pgm/literals/seed"
)

func AllLinguagemRepositoryMemory_Tester() {
	// usar : instanciar o obj estrutura
	repo := linguagem.LinguagemRepositoryMemory()

	// usar : seus metodos

	fmt.Println(repo.Createlinguagem(seed.Typescript))

	fmt.Println(repo.GetAll())
	// fmt.Println(reflect.TypeOf(repo.GetAll()))
	fmt.Printf("Tipo: %T", repo.GetAll())

}
