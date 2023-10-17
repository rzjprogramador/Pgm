// 1 - Funcao que vai receber as request do User e montar com a estrutura do Principal e devolver montado

import { New_Linguagem } from "./B.main.mjs";
import { NewDebuggar } from "./C.partesCompostos.mjs";
import { request_newDebuggar_Deno } from "./Z.requests_user.mjs";

function montaMeuObjeto(l, e, d) {
  return new New_Linguagem({
    newLinguagem: l,
    extensao: e,
    debuggar: new NewDebuggar(d).no_Console,
  });
}

const deno = montaMeuObjeto(
  "Deno",
  "ts",
  request_newDebuggar_Deno,
);

console.log(deno);
