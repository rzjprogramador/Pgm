import {
  Base_Linguagem
} from "../struct.Base_Linguagem.ts";
import {
  BaseLinguagemTypes, Propriedade
} from "../types.Base_Linguagem.ts";

const propriedade_via_construtor_kotlin: Propriedade = {
  tem_que_ter_a_propriedade: true,
  a_propriedade_pode_ser_criada_no: 'construtor',
  popularAPropriedade: {
    receber: 'recebo a propriedade de fora a ser criada via parametro',
    atribuir_a_propriedade_existente: 'TODO'
  },
  devolver: 'TODO'
}

const propriedade_via_metodo_kotlin: Propriedade = {
  tem_que_ter_a_propriedade: true,
  a_propriedade_pode_ser_criada_no: 'TODO',
  popularAPropriedade: {
    receber: 'TODO',
    atribuir_a_propriedade_existente: 'TODO'
  },
  devolver: 'TODO'
}

const kotlin: BaseLinguagemTypes = {
  linguagem: "kotlin",
  precisa_funcao_main: true,
  gerador_novo_objeto: {
    viaConstrutor: {
      propriedade: propriedade_via_construtor_kotlin
    },
    viaMetodo: {
      propriedade: propriedade_via_metodo_kotlin
    }
  }
};

const kotlin_base = Base_Linguagem.criar(kotlin);

console.log(kotlin_base);
