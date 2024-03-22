
extensoes_java_spring_para_vscode:
Spring Boot Extension Pack da VMmwre.com: ja vem 3 extensoes nele o spring Initializr, Spring tools, Spring Dashboard

criar_projeto:
spring_via_vscode_extensao_pacotes_spring : nao precisa fazer no site do initialzr , pela extensao entao:[ paleta de comandos, spring initialzr maven project, especificar a versao do spring, nomeInvertido com.empresa, nome artefato pacote principal, tipo arquivo build final : jar, versao do java, escolher dependencias plugins que vai usar, escolher a pasta workspace que ficar o projeto]
[tutorial: ](https://www.youtube.com/watch?v=mhLkn84qp6k&list=PLk4L0Yd2ljy4vVl1JsEEgA-zCtKjV6zxY&index=12)

adicionar_dependencPlugin_n_projeto: [caixa de comando, maven, add dependencia, nome da dependencia]


no pom.xml adicionar
plugin do maven para evitar conflito de versao do maven
`
<plugin>
  <groupId>org.apache.maven.plugins</groupId>
  <artifactId>maven-resources-plugin</artifactId>
  <version>3.1.0</version>
</plugin>
`