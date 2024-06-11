class Universal {
  static getInformacoes() {
    return Informacoes.allItems();
  }

  static getArrays() {
    return Arrays.allItems();
  }zz
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

class Arrays {
  static items = [];

  static addItems(i) {
    Arrays.items.push(i);
  }

  static allItems() {
    return Arrays.items;
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

const arraysGoRequest = {
  linguagem: "Go",
  add_item: `Add( item ) // ou // append(array, item)`,
};

const arraysJsRequest = {
  linguagem: "js",
  add_item: `array.push(item)`,
};

Informacoes.addItems(informacoesGoRequest);
Informacoes.addItems(informacoesJsRequest);

Arrays.addItems(arraysGoRequest);
Arrays.addItems(arraysJsRequest);

console.log(Universal.getInformacoes());
console.log(Universal.getArrays());
