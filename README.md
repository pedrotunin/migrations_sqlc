# Golang Migrations + SQLC + PostgreSQL

Este repositório demonstra como usar [golang-migrate](https://github.com/golang-migrate/migrate) para gerenciar migrações de banco de dados e [SQLC](https://sqlc.dev/) para gerar código Go a partir de consultas SQL, utilizando um banco de dados PostgreSQL.

## Estrutura do Repositório

```
db/
    migrations/    # Arquivos de migração do banco de dados
    queries/       # Arquivos de consulta SQL para SQLC
internal/
    repository/    # Código gerado pelo SQLC para interação com o banco
main.go           # Arquivo principal da aplicação
sqlc.yaml        # Configuração do SQLC
docker-compose.yml # Configuração do PostgreSQL com Docker
```

## Configuração

### 1. Subindo o Banco de Dados com Docker

Certifique-se de ter o Docker e o Docker Compose instalados. Para iniciar o banco de dados PostgreSQL, execute:

```sh
docker-compose up -d
```

Isso iniciará um contêiner PostgreSQL com as credenciais:
- **Usuário**: test
- **Senha**: test
- **Banco**: test

### 2. Criando e Aplicando Migrações

Para criar uma nova migração:

```sh
migrate create -ext sql -dir db/migrations -seq nome_da_migracao
```

Para aplicar as migrações ao banco:

```sh
migrate -path db/migrations -database "postgres://test:test@localhost:5432/test?sslmode=disable" up
```

### 3. Gerando Código com SQLC

Depois de definir suas consultas SQL em `db/queries/`, gere o código com:

```sh
sqlc generate
```

Isso criará os arquivos Go em `internal/repository/` para interagir com o banco de dados.

Para mais informações sobre como escrever as queries para o sqlc, consulte a [documentação oficial](https://docs.sqlc.dev/en/latest/index.html).

### 4. Executando a Aplicação

Para rodar a aplicação:

```sh
go run main.go
```

## Funcionamento do main.go

O arquivo `main.go` é responsável por:
- Estabelecer uma conexão com o banco de dados PostgreSQL.
- Verificar se a conexão foi bem-sucedida.
- Criar uma instância do driver do `golang-migrate`.
- Aplicar automaticamente as migrações pendentes ao iniciar a aplicação.

Se não houver novas migrações, o programa exibirá a mensagem "No new migrations to apply".

## Tecnologias Utilizadas

- **Golang** - Linguagem principal do projeto
- **PostgreSQL** - Banco de dados
- **golang-migrate** - Gerenciamento de migrações
- **SQLC** - Geração de código a partir de SQL
- **Docker** - Contêiner para o banco de dados

## Contribuição

Sinta-se à vontade para abrir issues e PRs para melhorar este repositório!

## Licença

Este projeto está sob a licença MIT.

