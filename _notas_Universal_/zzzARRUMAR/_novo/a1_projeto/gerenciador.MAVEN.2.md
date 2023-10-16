# GERENCIADOR DE PROJETOS JAVA :  MAVEN

## Instalacoes
obrigatorio: "para instalar o maven precisa ter o java jdk instalado primeiro."
site_oficial_maven: "https://maven.apache.org/install.html"

desinstalar_maven_instalado_com_sudo:  ```sudo apt-get purge --auto-remove maven```
// tutorial desinstalar maven : https://www.thelinuxfaq.com/ubuntu/ubuntu-17-04-zesty-zapus/maven?type=uninstall

instalacao: maven , tutorial: "https://www.youtube.com/watch?v=w0sANDwLUcA"
site_oficial_donwload: pegar o Binary tar.gz archive : "https://maven.apache.org/download.cgi"

fluxo:
  1. baixar e extrir a pasta maven - renomear apra maven e colar na home do meu linux
  2. na pasta bin/ posso experimenatr ver a versao do maven encima do arquivo mvn ```./mvn --version```
  3. no bashrc : # MAVEN - dizer que Path recebe o Path : diretorio onde esta o binario mvn do maven
export PATH=$PATH:/home/rzj/maven/bin/

comandosMaven:
```mvn -version```
```mvn``` // vai criar um projeto maven onde eu estiver

criar_projeto:
tutorial: maven no vscode: "https://www.youtube.com/watch?v=FNmZIxAbkT8"
  1.barra de comandos do vscode shift+f1 / escolhe novo projeto java /
  projeto maven / architeture quickstart / usar versao mais recente 1.4
  1. configurar_projeto: com.nomeDoDominio / nomeProjetoArtefato tudojunto / escolher local do projeto.
  2. no terminal do editor: vai perguntar versao do projeto confirma ou da nova/ y para concordar/ enter abrir projeto.
  3. arrumar a importacao do pacote com endereco deste o nome do projeto o do test muda de pasta main. para test. : ex: projectmaveum.src.main.java.com.siterzj
  4. na sidebar Painel MAVEN : utilizar os comandos prontos do maven / primeiramente lifeCycle install
  5. no sidebar MENU Testting habilitar o test java : que vai resolver as dependencias do junit

# Projetos com gerenciador de Projetos Maven
pastas fundamentais projeto java são somente o [src/ do codigo fonte , pow.xml, pasta target/ com o build para produção criada depois de compilar] para dependencias maven., o restante são arquivos das IDE que vc abrir o projeto.

# Tags para projeto maven pow
> saber versao atual do java no oficial da oracle: [https://www.oracle.com/br/java/technologies/downloads/]
> acompanhar realeses java [https://www.java.com/releases/] em 2023 a LTS é a versao 21
> ver versao no pc: ```java --version```
> definir versao do java no projeto add ao pow:
```
<properties>
    <java.version>21</java.version>
</properties>
```

# NomesProjetos
conceito: seguir o modelo de pacotes do java.io.
pacotes_modelo_top: <meusite> <extensaoDominio> <Objetivo>
exemplos: [ java.io.ClasseObjetivoDaUtilidade, meusite.com.Objetivo1 ]

