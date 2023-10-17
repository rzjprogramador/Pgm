// // OBJETOS REQUEST ENVIADAS PELO USER

// const request_newDebuggar_Deno = {
//   no_Console: "deno run ENDERECO_ARQUIVO.ts",
// };

// //

// class NewDebuggar {
//   constructor(no_Console) {
//     this.no_Console = no_Console;
//   }
// }

// // function newBaseDebuggar(n) {
// //   return new NewDebuggar(n);
// // }

// // const debuggarDeno = newBaseDebuggar(req_debuggar_no_console_deno);

// class BaseSintaxe {
//   obrigatorio_escopo_principal;
//   qual_escopo_principal;
//   obrigatorio_aspas_duplas_em_string;
// }

// // 2- Modelo do que sera inserido dinamicamente na Main
// // class New_Linguagem {
// //   linguagem;
// //   extensao;
// //   sintaxe = new BaseSintaxe();
// //   debuggar = new BaseDebuggar();
// // }

// // function newDebuggar(d) {
// //   return new NewDebuggar(d);
// // }

// class New_Linguagem {
//   // linguagem;
//   // extensao;
//   // sintaxe = new BaseSintaxe();
//   // debuggar = new NewDebuggar();

//   constructor(
//     newDebuggar,
//   ) {
//     this.newDebuggar = new NewDebuggar(newDebuggar);
//   }
// }

// // 1- Main
// class MainLinguagem {
//   constructor(new_linguagem) {
//     this.linguagem = new New_Linguagem(new_linguagem);
//   }
// }

// // Request para 1 - Main
// const request_linguagem_Typescript = {
//   // linguagem: "typescript",
//   // extensao: ".ts",
//   // sintaxe: {
//   //   obrigatorio_ponto_e_virgula_setencas: false,
//   //   obrigatorio_escopo_principal: false,
//   //   qual_escopo_principal: "funcao main()",
//   //   obrigatorio_aspas_duplas_em_string: false,
//   // },
//   // debuggar: debuggarDeno.no_Console,
// };

// // RECEBENDO REQUESTS - PREENCHIMENTOS DO USER --
// // tenho que receber via funcao e ter uma funcao agrupadora que monta meu objeto atraves do que recebi via estas funcoes.

// function montaMeuObjeto(d) {
//   return new MainLinguagem(d);
// }

// // const linguagem_Typescript = new MainLinguagem(request_linguagem_Typescript);
// // console.log(linguagem_Typescript);

// console.log(montaMeuObjeto(request_newDebuggar_Deno.no_Console));
