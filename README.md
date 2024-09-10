## Desafio Código Morse

### Descrição

A presente aplicação tem o intuito de implementar um decodificador de código Morse. A abordagem escolhida foi criar um pacote exportável que tem como produto final um comando Cobra, para ser importado em qualquer aplicação. A única dependência externa instalada foi o pacote cobra command. O caminho escolhido para implementar o desafio foi utilizar uma estratégia de buffer, onde, independentemente do tamanho do input, ele será sempre lido em partes menores, evitando que todo o conteúdo do código seja carregado na memória. Dessa forma, o decodificador converte o código Morse palavra por palavra, pois quem controla o fluxo de leitura é um buffer externo, que pode ler tanto de arquivos quanto do stdin
### Stack

- Golang 1.22
- Docker
- Docker compose
- Github actions
- Pacotes externos  
    - Cobra command 1.8.1

### Utilização

Para subir a app
```
docker compose run morse
```

**Rodar dentro do container** \
Com stdin e stdout padrão

```
./morse decode ".... . .-.. .-.. ---   .-- --- .-. .-.. -.."
```

Com arquivo como entrada e stdout como saída

```
./morse decode --file arquivo.txt
```

Com arquivo como entrada e arquivo como saída

```
./morse decode --file arquivo.txt --output saida.txt
```

Com stdin como entrada e arquivo como saída

```
./morse decode --output lol.txt ".... . .-.. .-.. ---   .-- --- .-. .-.. -.."
```

### Rodar testes

Para rodar os testes

```
docker compose run test
```

### Simular pipeline 

```
docker compose run ci
```