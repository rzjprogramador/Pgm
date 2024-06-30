package tester

import (
	"fmt"

	s "github.com/rzjprogramador/Pgm_Universal/Pgm/literals/seed/seed_universal"
)

func UniversalPGM_Tester() {
	fmt.Println(
		s.NomePastaGuardaOutrasPastas(),
		s.NomePastaDeArquivosSoltos(),
		s.QuebraDeLinhas(),
		s.Implementacoes(),
	)
}
