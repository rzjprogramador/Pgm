import {
  Base_Linguagem
} from "../master_pgm_universal.ts";
import {
  PropsBaseLinguagem, Propriedade
} from "../master_pgm_universal_types.ts";

const propriedade_via_construtor_TS: Propriedade = {
  tem_que_ter_a_propriedade: true,
  a_propriedade_pode_ser_criada_no: 'construtor',
  popularAPropriedade: {
    receber: 'recebo a propriedade de fora a ser criada via parametro',
    atribuir_a_propriedade_existente: 'dou para esta this.propriedade da classe a propriedade recebida de fora'
  },
  devolver: 'no construtor nao precisa devolver a propria atribuicao ja devolve.'
}

const propriedade_via_metodo_TS: Propriedade = {
  tem_que_ter_a_propriedade: true,
  a_propriedade_pode_ser_criada_no: 'no construtor digo que vai ter a propriedade privada mas nao configuro seu valor no construtor. essa atribuicao faço no metodo.',
  popularAPropriedade: {
    receber: 'no metodo o metodo tem que ser public e static: recebo a propriedade de fora a ser criada via parametro',
    atribuir_a_propriedade_existente: 'dou para esta this.propriedade da classe a propriedade recebida de fora'
  },
  devolver: 'retorno a instanciacao da classe com o que recebeu'
}

const typescript: PropsBaseLinguagem = {
  linguagem: "typescript",
  precisa_funcao_main: false,
  gerador_novo_objeto: {
    viaConstrutor: {
      propriedade: propriedade_via_construtor_TS
    },
    viaMetodo: {
      propriedade: propriedade_via_metodo_TS
    }
  }
};

const typescript_base = new Base_Linguagem(typescript);

console.log(typescript_base);
