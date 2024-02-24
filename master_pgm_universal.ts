import type {
  PropsBaseLinguagem
} from './master_pgm_universal_types.ts'

class Base_Linguagem {

  constructor(
    private readonly base_linguagem: PropsBaseLinguagem,
  ) { }

  public static criar(
    base_linguagem: PropsBaseLinguagem
  ) {
    return new Base_Linguagem(base_linguagem)
  }


}

export {
  Base_Linguagem
};


