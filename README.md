# No code api

O conceito do projeto é ser um site que possibilite desenvolvedores front-end, mobile e profissionais de UX/UI criem APIs com recursos dinâmicos, operações de CRUD de forma fácil, rápida e sem precisar escrever uma linha de código ou ter que se preocupar em conexão com banco de dados, configurações de CORS e outras coisas.

## Pré-requisitos

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Go](https://go.dev/) versão 1.23.0

## Configuração

### Variáveis de Ambiente

O projeto requer um arquivo `.env` com as seguintes variáveis de ambiente. Use o arquivo `.env.template` como referência para criar o seu `.env`.

```
SERVER_PORT=:8080
POSTGRE_HOST=localhost
POSTGRE_PORT=PORTA_DOCKER
POSTGRE_USER_NAME=USERNAME_DOCKER
POSTGRE_PASSWORD=PASSWORD_DOCKER
POSTGRE_DB_NAME=SUA_ESCOLHA
POSTGRE_SSL_MODE=disable
JWT_SECRET=CHAVE_JWT
REDIS_HOST=localhost:PORTA_DOCKER
REDIS_PASSWORD=PASSWORD_DOCKER
REDIS_DB=0
```

Para a chave JWT utilizei o site [JWTSecret.com](https://jwtsecret.com/generate)

Importante: Certifique-se de preencher todos os campos necessários no seu arquivo .env antes de iniciar o projeto.

### Arquivo obrigatório do MongoDB
Crie um arquivo mongo-init.js na raiz do projeto com o seguinte conteúdo:

```javascript
db.createUser(
    {
        user: "root",
        pwd: "SuaSenha",
        roles: [
            {
                role: "readWrite",
                db: "my_db"
            }
        ]
    }
);
db.createCollection("collection");
```
Esse arquivo é utilizado para configurar o usuário e a coleção padrão no MongoDB.

### Docker
Para facilitar a configuração das dependências (PostgreSQL, Redis e MongoDB), use o arquivo docker-compose.yml. Este arquivo cria os serviços necessários para o projeto.

#### Configurações utilizadas
```
version: "3.7"
services:
  postgres:
    container_name: postgres
    image: postgres:15.3-alpine
    shm_size: 128mb
    environment:
        POSTGRES_PASSWORD: SuaSenha
        POSTGRES_USER: postgres
    ports:
      - "5432:5432"
  redis:
      image: redis:7.4.1
      container_name: redis
      ports:
        - '6379:6379'
      command: redis-server --loglevel warning --requirepass SuaSenha
      volumes: 
        - cache:/data
  mongo:
    image: mongo:latest
    container_name: mongo
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: SuaSenha
    ports:
      - "27017:27017"
    volumes:
      - mongodbdata:/data/db
      - ./init-mongo.js:/docker-entrypoint-initdb.d/mongo-init.js:ro
volumes:
  mongodbdata:
    driver: local
  cache:
    driver: local
```

#### Subindo os serviços com Docker
Para rodar todos os serviços necessários, use o comando abaixo:

```bash
docker-compose up -d
```
Este comando irá subir o PostgreSQL, pgAdmin, Redis e MongoDB conforme configurado no docker-compose.yml.

Para verificar se os contêineres estão funcionando corretamente, use:

```bash
docker-compose ps
```

### Executando o Projeto em Go
Instale as dependências do Go, e então rode o projeto:

```bash
go mod tidy
go run main.go
```
O servidor estará disponível na porta definida em SERVER_PORT.

### Dependências Go
Abaixo estão as dependências do projeto, com uma breve descrição de cada uma:

| Lib | Versão | Descrição |
| :--- | :---: | :---|
| [GIN](https://github.com/gin-gonic/gin) | v1.10.0 | Framework web em Go, usado para criar rotas e gerenciar requisições HTTP |
| [JWT](https://github.com/golang-jwt/jwt) | v5.2.1 | Biblioteca para geração e validação de tokens JWT, essencial para autenticação |
| [gotoenv](https://github.com/joho/godotenv) | v1.5.1 | Utilizada para carregar variáveis de ambiente a partir de arquivos .env |
| [crypto](https://golang.org/x/crypto) | v0.26.0 | Pacote com implementações de criptografia e hashing para maior segurança |
| [uuid](https://github.com/google/uuid) | v1.6.0 | Pacote com implementação de geração de IDs unicos padrão uuid |
| [POSTGRE](https://github.com/jackc/pgx) | v5.7.1 | Driver PostgreSQL, usado para comunicação com o banco PostgreSQL|
| [Redis](https://github.com/redis/go-redis) | v9.6.2 | Client Regis, usado para comunicação com o banco Redis|
| [MONGODB](https://go.mongodb.org/mongo-driver/mongo) | v1.17.1 | Driver MongoDB, usado para comunicação com o banco MongoDb|

### Makefile
Para utilizar ele é preciso instalar o pacote [make](https://community.chocolatey.org/packages/make) através do [chocolatey](https://chocolatey.org/install).
Para rodar o projeto possui dois comandos

```bash
make run-dev
```
Esse comando irá apenas rodar o GO utilizando o comando `go run ./main.go`

```bash
make run
```
Esse comando irá realizar o build do executável do projeto para a pasta `bin` e logo em seguida irá rodar o projeto.

## Acesso aos Serviços
### Docker Compose.
PostgreSQL:
Acesso disponível na porta 5432.

Redis:
Acesso disponível na porta 6379.

MongoDB:
Acesso disponível na porta 27017.

## Observações
Certifique-se de substituir "SuaSenha" nos arquivos de configuração pelas senhas de sua escolha.
