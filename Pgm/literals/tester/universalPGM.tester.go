package tester

import (
	"fmt"

	"github.com/rzjprogramador/Pgm_Universal/Pgm/literals/seed"
)

func UniversalPGM_Tester() {
	fmt.Println(
		seed.NomePastaGuardaOutrasPastas(),
		seed.NomePastaDeArquivosSoltos(),
		seed.QuebraDeLinhas(),
		seed.Implementacoes(),
	)
}
