# PGM UNIVERSAL

Arquiteturas

ESRC [ Entidade Servico Repositorio Controller ], aplicacao em camadas em java ::

pacote_entities / todas classes modelo entidades, tera os campos do obj, sera a entidade no BD,

pacote_services / todos casos de uso das entidades , o casodeUso vai receber via metodo unico execute( recebera o obj da entidade para poder mandar o repo salvar o caso de uso solicitado pra esta entidade,  delegando a um contrato de repositorio via composicao para salvar a acao pra ele.

pacote_repositories / todos repos das entidades vai ter contrato com acoes de caso de uso e classe que vai implementa-lo para servir ao service.

pacote_controllers / todos controllers e tests das entidades , obs: de uma entidade teste somente os controllers porque ele ja vem carregando todas acoes da entidade.

---

Significados
- Retorno : a clausula return é um COMPARTILHAR, voce compartilha o que vem depois do return.

Tipo_do_Valor
nulo: indica que nao esta preenchido ainda, [ null, nil, ]

valor_exportavel_e_dinamico_mudavel: qualquer valor literal pode ser usado representado dentro de uma variavel para qualquer situacao, pense literal ele não é exportavel só vai usar aqui, dentro de variavel ele pode viajar ser exportavel, dentro de uma funcao que o retorna vindo de um parametro ele se torna dinamico pode ter outros valores e ainda ser exportavel.

Valores_em_estrutura_Classe: precisa encapsular proteger as variáveis,  deixe as privada, e só deixe acesso a elas via metodos get para pesquisa e set para receber algo e modificar setar, assim podemos fazer acoes de regras antes de pesquisa ou modificacao do pedido recebido

- Funcao : é um RESULTADO, quando esta sendo criada vai configurar um resultado com o que for entrar ou nao e retornar ela, quando esta sendo chamada usada em algum lugar recebera ou nao o que foi configurado e devolvera um resultado.
pense: _ QUANDO <NomeDaFuncao> VOU DEVOLVER <o que vier depois do return ou senao tiver return vou executar a ultima linha aqui configurada que faz algo ja tem return  encapsulada no que estou aqui declarando.>
-
- Parametro : "O QUE VOU PRECISAR, informacoes variaveis editaveis que vou usar na funcao.
-
- Promessa QUANDO TERMINAR < .then( "_quando terminar esta promessa encadeada FAZ ISTO com esta funcao anonima que vou passar neste param." ))

Algumas linguagens como o [java, ] permite ter mais que 1 construtor se permitir entao crie um vazio sem params, isso é util para instancia-la sem argumentar, ou senao crue um metodo statico wue instancia a classe.

Associacao: TEM UM, na estrutura faco uma prop. receber outra classe, isso resultara em objeto aninhado, onde um campo do obj sera um objeto ao inves de tipo primitivo final,
Exemplo: obj pessoa tera um campo perfil com os dados do perfil aninhado.

auxilio_composicao: faça o rasculho de como vai usar o objeto instanciado e veja onde sera substituivel alguns dos compartimentos para assim depois implementar.

- Composicao : é o uso de um objeto operario sob contrato com metodo QUE FAÇA A ACAO para o objeto que o compos, sendo assim [ substituivel, cumpridor de acao, polo unico, metodo mascara de libs externas ]
exemplos : [ carro. Motor. ligar() > aqui carro depende do motor que sera injetado - qualquer intancia de motor vai servir para fazer a acao de ligar() que o carro precisa, ]

- Composicao_Dupla : para gerar opcionais, ex: pagamento  .via <este com os atributos que serao opcoes > e cada um destes camposCompostos serao Compostos por objetos com o metodo de contrato pagar (). Cada objeto podem ter seus campos diferentes o que é obrigatorio é ter o metodo em comum que a classe de autonivel vai usar que é o pagar() exemplo : pagamento.  formaPag < viaCartao | viaPix | viaDinheiro > .pagar() , exemplo tbm serve para ClienteFisico e Juridico.

- Para ter objeto entidade, antes criamos metodos CRUD create, ler um e todos, update , deletar , cada metodo deste pode ser divido em servico e controle, o controle é a ponte de entrada para informacao que vem de fora via metodos de api (ex: graphql , rest)

- Objeto entidade de Configuracao : alimentado pelo ADM
nao fazer singleton objeto unico, porque pode-se querer pesquisar estados anteriores desta informacao, entao esta informacao_atual tem que vir da ultima posicao de uma colecao de informações deste contexto inserido.

- valores de objADM na criacao voce pode cria-los com valores default, e modificar somente via api

- Ordem entidade Prototype com acoes computadas entre seus campos, onde VALIDAMOS OS ARGUMENTOS, Defimos Acoes UseCase
- _Crud (Dividido entre Service e Controller), este controller alimentara a api.

- Repositorios por entidade serao administrado por objeto_AdmDEV nele posso definir onde em qual base de dados sera guardada a informacao da entidade especifica.

- Composicao
  - objeto administravel por partes ou todo : Se quisesse gerar 1 obj só com todos campos instanciava só pelo construtor, mas como quero gerar este objeto por partes independentes faço os metodos para isso.

- autoInstanciao : instancie o obj internamente via metodo factoryMethod, evita outras formas de instanciacao por todo codigo.
Faca um metodo semelhante ao construtor, retornando uma new PropriaClasse como se estivesse instanciando fora.

para validadores criar classes Text, Number, Rule , que recebem objeto com campos :: length, ou value

PLATAFORMAS RUN TIME
Conceito: as lingusgens antigas tem suas ferramemtas somente para aplicacoes internas ou web, nao tem por padrao ferramentas para servidor onde se constroem microservicos internos e externosque precisam se comunicar como apis, comunicacao banco de dados, por isto elas usam suas plataformas ex: js usa o node, deno, java usa o spring boot, quarkus, etc...

versoes: utlize sempre a versao LTS que é a mais garantida.

Gerenciador de Dependencia
Conceito: toda plataforma tem seu gerenciador de dependencia para buscar e baixar na internet as ďependencias usadas no projeto na bersao correta.
No caso do spring utilizo em 2024 o maven

LIBS
funcoes Sync sao assincronas tem que esperar com await na frente, funcoes sem o sync tem que passar callback uma funcao anonima pra fazer algo depois que ela acabar de fazer a tarefa antes.




