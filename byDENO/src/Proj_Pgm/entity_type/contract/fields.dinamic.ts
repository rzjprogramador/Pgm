import { Conceito } from "../../../helpers/core/deps.ts";

export interface Fields_Dinamics {
  dado_Referencia: Dado_Referencia
}

// partes nivel3

export interface Dado_Referencia {
  linguagem: string
  gerado_viaClasse: Conceito
  gerado_viaCStruct: Conceito
  gerado_viaFuncao: Conceito
  comportamentos_Metodos_que_PODE_FAZER: Conceito
}