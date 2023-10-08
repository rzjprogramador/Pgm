import { Args_Dinamico } from "../types/args.ts";
import { Comput_Fixo } from "../types/comput.type.ts";

export class Pgm {

  public comput_fixo: Comput_Fixo = 

  constructor(
    private args_dinamico: Args_Dinamico,
  ) { }

  criarPgm(args_dinamico: Args_Dinamico) {
    return new Pgm(args_dinamico)
  }

}