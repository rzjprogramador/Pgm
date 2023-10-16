
# tutoriais
by_loiane : https://www.youtube.com/watch?v=dkmlOi_MNb4&list=PLGxZ4Rq3BOBqLCGkhtvD7-NVtx2J4VoUj&index=2

# Framework Spring Boot
conceito:
arquivo_gerador_de_saida_app: Jar mesmo para apps web api,
extensoes_vscode
  1. Spring Boot Extension Pack
  2. Spring Boot Dashboard by Microsoft : identifica projeto spring no workspace, abre painel no sidebar com o mesmo nome para rodar ferraenats da app tipo run, tests, este botao aparecerá nos botoes da sidebar.
  3. lombok Annotations by Microsoft : para anotacoes que ja implementam getters, setters e outros.

# projeto springboot
iniciando_via_web_opcao2: https://start.spring.io/ o emsmo pode ser feito no vscode, veja:
inicializador: paletaComandos/ spring: initializer
gerenciadorProjeto: maven project
> informacoes_do_projeto


especificar_versao_do_gerenciador: escolher a estavel,MENOS a SNAPSHOT que não esta estavel por isto pe snapshot
linguagem: escolha a linguagem : [ java, kotlin, Goove ]


tipoEmpacotamentoSaida: Jar que é bytecode de java
versao_do_java: em 2023 escolher >> 20
dependencias_iniciais:
  1. SpringBoot DevTools : ferramentas Spring , hotreload,
  2. lombok : reduz codigo d ebolepaint getters,setters, outros
  3. Spring Web : deixa disponivel utilizar o Rest

dependencias_decorrer_implementacoes:
  1. Rest Repositories : anotacoes de repositorio
  2. H2 DataBase SQL : bando de dados em memoria


finalizacao: vai selecioando as extensoes com mouse, ao acabar da ENTER para confirmar, abra o vscode na pasta do projeto gerado com <OPEN> da sugestão.

# futuras_dependencias_add_no_pom
1. clica com direito no arquivo pom / escolher <add starters : adicionar aos Iniciantes>
2. finalizando: clica em <proceed : continuar> para processar confirma com <always : sempre>.

# outras_dependencias_projeto
1. javax para validacoes mudou para jakarta

# instalacoes_projeto
1. instalar_dependencias: no terminal ir até o nivel do arquivo pom , byMaven: ```mvn install```

# run_rodar
1. rodar_aplicacao: ```mvn spring-boot:run``` // ou no botão ```Run``` acima da classe.
2. porta_default_server: http://localhost:8080
3. mudar_porta_default_server: no arquivo application.properties // adicionar : server,port = 8081 ou outra.

# estrutur
# inicial
src_main_resorces_aplication_properties: onde inserimos informacoes do bancoDeDados,
classe_main_principal: src/main/java/com/<DEV>/<NOME_DO_PROJETO>/<NOME_DO_PROJETO>Application.java ANOTADA COM @SpringBootApplication
classe_de_test: anotacao >> @SpringBootTest de org.springframework.boot.test.context.SpringBootTest;
cada_suite_test: anotacao: @Test de import org.junit.jupiter.api.Test;

# subir_servidor_api_no_ar:
viaDashboard: no botao do dashboard clicar em play e subir a app, obs tem que ter o toncat instalado para subir o spring app.

pom_config_banco:
  1. senao for criar um banco externo antes comente a dependencia spring-data-jpa porque ao subir o server ele espera um ocnexao com um BD. só o indique no pom se ja definiu o BD.