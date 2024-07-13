interface EntidadeFuncionalidade {
  c1: number
  c2: number
}

type EntidadeFuncionalidadeFNReturn = number
type EntidadeFuncionalidadeInput = EntidadeFuncionalidade | EntidadeFuncionalidadeStruct
type Test_EntidadeFuncionalidadeType = Error | void

// Contrato_Estruturador
class EntidadeFuncionalidadeStruct {
  public readonly c1: number = 0
  public readonly c2: number = 0

  private constructor(private x: EntidadeFuncionalidade) {
    this.c1 = x.c1
    this.c2 = x.c2
  }

  // acao_das_intancias_filhas
  static New(x: EntidadeFuncionalidade) {
    return new EntidadeFuncionalidadeStruct(x)
  }

  //
  static ActionSoma(x: EntidadeFuncionalidadeInput): number {
    const formula = x.c1 + x.c2
    return formula
  }
}


// input de dados sejam eles externos, hardcorde, doCodigo
const EntidadeFuncionalidadeInput_OK = EntidadeFuncionalidadeStruct.ActionSoma({ c1: 10, c2: 10 })
const EntidadeFuncionalidadeInput_FAIL = EntidadeFuncionalidadeStruct.ActionSoma({ c1: 20, c2: 20 })
const EntidadeFuncionalidadeInput_HARD = { c1: 10, c2: 10 }


function Test_EntidadeFuncionalidadeFN(op: EntidadeFuncionalidadeFNReturn): Test_EntidadeFuncionalidadeType {
  if (op != 20) {
    throw new Error(`OPS..O TEST_FALHOU : O resultado = ${op}`)
  }
  console.log(`Ok..O TEST_PASSOU : O resultado = ${op}`)
}

const resultTest_EntidadeFuncionalidade_OK = () => Test_EntidadeFuncionalidadeFN(EntidadeFuncionalidadeInput_OK)
const resultTest_EntidadeFuncionalidade_FAIL = () => Test_EntidadeFuncionalidadeFN(EntidadeFuncionalidadeInput_FAIL)

function RunTests() {
  resultTest_EntidadeFuncionalidade_OK()
  // resultTest_EntidadeFuncionalidade_FAIL()
}

RunTests()

/*
ALGORITMOS PARA ENTIDADE OU FUNCIONALIDADES

conceito a entidade seja ela: [ pessoa, funcionalidade, produto], sera no final a representacao de um objeto.

*/