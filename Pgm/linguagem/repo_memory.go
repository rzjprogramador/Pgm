package linguagem

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

// Instanciar Obj da estrutura
func LinguagemRepositoryMemory() linguagemRepository {
	repo := linguagemRepository{
		Items: arraylist.New(),
	}
	return repo
}

// Implementar metodos
func (r linguagemRepository) Createlinguagem(p Linguagem) string {
	r.Items.Add(p)
	return "Criado com sucesso!"
}

func (r linguagemRepository) GetAll() *arraylist.List {
	return r.Items
}
