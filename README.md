
# Api rest de conversão de moeda

A api tem como objetivo armazenar uma transação de compra e converter o valor da compra em diferentes moedas.

As tecnologias escolhidas para a implementação são:

-Go: Go é uma linguagem de programação moderna, desenvolvida pela Google, que se destaca pela sua simplicidade, desempenho e escalabilidade. Projetada para lidar com sistemas de alto desempenho e grandes volumes de dados.

-Docker: Docker é uma plataforma de contêinerização que permite criar, testar e implantar aplicações de maneira consistente em qualquer ambiente.

-Postgres: O PostgreSQL é um sistema de gerenciamento de banco de dados relacional de código aberto, altamente confiável e robusto


## Funcionalidades Principais

- Criação de uma transação.
- Conversão da transação para outra moeda.


## Como utilizar

## Go e Docker
Primeiramente, certifique-se de ter o Go e o docker instalados na sua máquina.

Clone o projeto:

```bash
  git clone https://github.com/MurilojrMarques/api-transaction
```

Crie um arquivo .env na pasta raiz do projeto com as seguintes informações:

DB_SERVER=

DB_USER=

DB_PASSWORD=

DB_DATABASE=

DB_PORT=

No diretório raiz do projeto, execute o seguinte comando para iniciar o PostgreSQL e a API:

```bash
  docker compose up
```

Será necessário entrar na pasta da migration:

```bash
  cd database/migration
```

e rodar o comando substituindo os campos para subir a migration para o banco de dados:

```bash
  goose postgres "user=<user_db> dbname=<name_db> password=<password_db> host=<host_db> sslmode=disable" up 
```

## Rodando os testes

Para rodar os testes, rode o comando go run main.go e teste as rotas via postman/insomnia ou entre no link do swagger: 

http://localhost:8080/swagger/index.html

