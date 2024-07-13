interface EntidadeFuncionalidade {
  c1: number
  c2: number
}

type EntidadeFuncionalidadeFNReturn = number
type EntidadeFuncionalidadeRequest = EntidadeFuncionalidade | EntidadeFuncionalidadeStruct
type Test_EntidadeFuncionalidadeType = Error | void

class EntidadeFuncionalidadeStruct {
  public readonly c1: number = 0
  public readonly c2: number = 0

  private constructor(private x: EntidadeFuncionalidade) {
    this.c1 = x.c1
    this.c2 = x.c2
  }

  static New(x: EntidadeFuncionalidade) {
    return new EntidadeFuncionalidadeStruct(x)
  }
}

const EntidadeFuncionalidadeRequest_OK = EntidadeFuncionalidadeStruct.New({ c1: 10, c2: 10 })
const EntidadeFuncionalidadeRequest_FAIL = EntidadeFuncionalidadeStruct.New({ c1: 20, c2: 20 })
const EntidadeFuncionalidadeRequest_HARD = { c1: 10, c2: 10 }

function EntidadeFuncionalidadeFN(x: EntidadeFuncionalidadeRequest): number {
  const formula = x.c1 + x.c2
  return formula
}

function Test_EntidadeFuncionalidadeFN(op: EntidadeFuncionalidadeFNReturn): Test_EntidadeFuncionalidadeType {
  if (op != 20) {
    throw new Error(`OPS..O TEST_FALHOU : O resultado = ${op}`)
  }
  console.log(`Ok..O TEST_PASSOU : O resultado = ${op}`)
}

const resultTest_EntidadeFuncionalidade_OK = () => Test_EntidadeFuncionalidadeFN(EntidadeFuncionalidadeFN(EntidadeFuncionalidadeRequest_OK))
const resultTest_EntidadeFuncionalidade_FAIL = () => Test_EntidadeFuncionalidadeFN(EntidadeFuncionalidadeFN(EntidadeFuncionalidadeRequest_FAIL))

function RunTests() {
  resultTest_EntidadeFuncionalidade_OK()
  // resultTest_EntidadeFuncionalidade_FAIL()
}

RunTests()

/*
ALGORITMOS PARA ENTIDADE OU FUNCIONALIDADES

*/