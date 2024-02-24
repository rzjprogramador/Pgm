import {
  Base_Linguagem
} from "../master_pgm_universal.ts";
import {
  PropsBaseLinguagem, Propriedade
} from "../master_pgm_universal_types.ts";

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

const kotlin: PropsBaseLinguagem = {
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
