class Universal {
  static getInformacoes() {
    return Informacoes.allItems();
  }
}

class Informacoes {
  static items = [];

  static addItems(i) {
    Informacoes.items.push(i);
  }

  static allItems() {
    return Informacoes.items;
  }
}
const informacoesGoRequest = {
  linguagem: "Go",
  revelaTipo: "reflect.TypeOf( <target> )",
  revelaInstancia:
    `use o metodo Name do typeOf e terao nome do struct em string ex:
    if reflect.TypeOf( TARGET ).Name() != "NomeDostruct" { ..faça algo`,
};
const informacoesJsRequest = {
  linguagem: "Js",
  revelaTipo: `typeof operando === "tipoEmstring"`,
  revelaInstancia: "objeto instanceof Construtor",
};

Informacoes.addItems(informacoesGoRequest);
Informacoes.addItems(informacoesJsRequest);

console.log(Universal.getInformacoes());
