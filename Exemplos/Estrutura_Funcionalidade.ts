interface Funcionalidade {
  c1: number
  c2: number
}

type FuncionalidadeFNReturn = number
type FuncionalidadeRequest = Funcionalidade | FuncionalidadeStruct
type Test_FuncionalidadeType = Error | void

class FuncionalidadeStruct {
  public readonly c1: number = 0
  public readonly c2: number = 0

  private constructor(private x: Funcionalidade) {
    this.c1 = x.c1
    this.c2 = x.c2
  }

  static New(x: Funcionalidade) {
    return new FuncionalidadeStruct(x)
  }
}

const funcionalidadeRequest_OK = FuncionalidadeStruct.New({ c1: 10, c2: 10 })
const funcionalidadeRequest_FAIL = FuncionalidadeStruct.New({ c1: 20, c2: 20 })
const funcionalidadeRequest_HARD = { c1: 10, c2: 10 }

function FuncionalidadeFN(x: FuncionalidadeRequest): number {
  const formula = x.c1 + x.c2
  return formula
}

function Test_FuncionalidadeFN(op: FuncionalidadeFNReturn): Test_FuncionalidadeType {
  if (op != 20) {
    throw new Error(`OPS..O TEST_FALHOU : O resultado = ${op}`)
  }
  console.log(`Ok..O TEST_PASSOU : O resultado = ${op}`)
}

const resultTest_Funcionalidade_OK = () => Test_FuncionalidadeFN(FuncionalidadeFN(funcionalidadeRequest_OK))
const resultTest_Funcionalidade_FAIL = () => Test_FuncionalidadeFN(FuncionalidadeFN(funcionalidadeRequest_FAIL))

function RunTests() {
  resultTest_Funcionalidade_OK()
  // resultTest_Funcionalidade_FAIL()
}
RunTests()