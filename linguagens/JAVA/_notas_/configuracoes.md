
extensoes_java_spring_para_vscode:
para_o_java: Extension Pack for Java , da microsoft
para_spring_boot: Spring Boot Extension Pack da VMmwre.com: ja vem 3 extensoes nele o spring Initializr, Spring tools, Spring Dashboard


para_funcionar_lombok: extensao no vscode : Lombok Annotations Support for VS Code obs: #melhor usar o record do java ao inves do lombok,

criar_projeto:
spring_via_vscode_extensao_pacotes_spring : nao precisa fazer no site do initialzr , pela extensao entao:[ paleta de comandos, spring initialzr maven project, especificar a versao do spring, nomeInvertido com.empresa, nome artefato pacote principal, tipo arquivo build final : jar, versao do java, escolher dependencias plugins que vai usar, escolher a pasta workspace que ficar o projeto]
[tutorial: ](https://www.youtube.com/watch?v=mhLkn84qp6k&list=PLk4L0Yd2ljy4vVl1JsEEgA-zCtKjV6zxY&index=12)

adicionar_dependencPlugin_n_projeto: [caixa de comando, maven, add dependencia, nome da dependencia]


no pom.xml
adicionar_plugins:
plugin do maven para evitar conflito de versao do maven
`
// obs: na tag release insira a versao do java em uso no caso aqui esta para java 17.

<plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-compiler-plugin</artifactId>
                <version>3.8.1</version>
                <configuration>
                    <release>17</release>
                    <compilerArgs>
                        <arg>--enable-preview</arg>
                    </compilerArgs>
                </configuration>
            </plugin>

            <plugin>
                <groupId>org.apache.maven.plugins</groupId>
                <artifactId>maven-surefire-plugin</artifactId>
                <version>3.0.0-M4</version>
                <configuration>
                    <argLine>--enable-preview</argLine>
                </configuration>
            </plugin>
`

adicionar_dependencias:
`para funcionar dependencias lombok
<dependency>
			<groupId>org.projectlombok</groupId>
			<artifactId>lombok</artifactId>
			<optional>true</optional>
		</dependency>
`