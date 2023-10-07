
# TECNICAS -> CODIGO LIMPO : BOM

1. bom:
    1. bom: ter coesão no artefato para fazer somente uma coisa especifica e relacionadas a si ex: podeSeFazEspecificamenteSomente: [ falarsomenteDeSi_NaoOutrasEntidades, representar_OMundoReal_NoCodigo_CriarPropriedades, mostrar,  ].
    1. bom: proteger entrada de dados com private ou protected assim evita entrar dados sujos invalidos.
    2. validadores: fazer metodos {set para entrar o dado} e {metodo get para ver o dado} porque estão com proteção. obs: estes metodos getter e setter servem como validadores.
    3. acessos propriedades : acesse as maioria das props por metodos get e set assim fica possivel criar regras para setar ou para visualizar/get conforme necessidade de negocio, obs: só usar os get e set fora da classe dentro da classe nao precisa quando pquer somente mostrar a propridade.
    4. #GetterSetter : Só fazer SET em prop. que deseja que alterem o valor fora, só fazer get em prop. que deseja que possam ver visualizar foram, então NÃO precisa criar get e set para todos, Propriedade que não deseja ALTERAR o valor NÃO FAZER SET , NÃO DESEJA MOSTRAR FORA NÃO FAZER GET , tutorial: [https://www.youtube.com/watch?v=kuvg8JixRp4&list=PL62G310vn6nFIsOCC0H-C2infYgwm8SWW&index=57]

    6. Uma classe/em_arquivo cria entidade propriedades, outra classe/em_arquivo test invoca e executa #importante essa execução ocorre sempre dentro do metodo principal main().
    7. chegagens:
        1. comparacao_entre_valores:
            1. mesmoValor: ""não compare um valor consigo mesmo. Se você precisar de um comportamento constante, use os literais booleanos true e false, em vez de codificá-los obscuramente como ou similares. ex: 1 == 1 prefira : true == false",


2. refactor_refatore:
    1. metodos que usam um valor fixoMarretado no codigo, pode #refatorar e fazer deste valorFixo uma propriedade protegida com valor default para uso na classe.


1. ruim
    1. ruim: [artefato não ter coesao, fazer mais que uma tarefa, deixar campos publicos vulneraveis a entrada de dados sujos, ]



