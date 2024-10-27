# User API

Esta aplicação é uma API para gerenciar usuários, construída em Go com uma arquitetura organizada para escalabilidade e modularidade. A API foi desenvolvida seguindo padrões de mercado e utilizando princípios de boas práticas de programação.

### Endpoints Disponíveis
- GET /users: Retorna todos os usuários.
- GET /users/{id}: Retorna um usuário pelo ID.
- POST /users: Cria um novo usuário.
- PUT /users/{id}: Atualiza as informações de um usuário pelo ID.
- DELETE /users/{id}: Deleta um usuário pelo ID.




### Estrutura do Projeto

A estrutura de pastas é organizada da seguinte forma:

```plaintext
go-api
├── cmd
│   └── start
│       └── main.go          # Ponto de entrada da aplicação
├── internal
│   └── app
│       └── api
│           ├── controller
│           │   └── user_controller
│           │       ├── interface.go       # Interface do controlador
│           │       └── user_controller.go # Implementação do controlador de usuário
│           ├── model
│           │   └── user_model
│           │       ├── set.go             # Métodos de definição de campos do modelo de usuário
│           │       └── user.go            # Estrutura do modelo de usuário
│           └── router
│               └── userRouter
│                   ├── deleteUserById.go   # Endpoint para deletar usuário por ID
│                   ├── getAllUser.go       # Endpoint para obter todos os usuários
│                   ├── getUserById.go      # Endpoint para obter usuário por ID
│                   ├── router.go           # Configuração das rotas de usuário
│                   ├── setUser.go          # Endpoint para criar usuário
│                   └── updateUserById.go   # Endpoint para atualizar usuário por ID
├── config
│   ├── configEnv
│   │   ├── config.go        # Configurações do ambiente
│   │   └── module.go        # Inicialização do módulo de configuração
│   ├── db
│   │   ├── migrations       # Scripts de migração do banco de dados
│   │   │   ├── asdasd.sql
│   │   │   └── migrationInit.sql
│   │   ├── db.go            # Configuração do banco de dados
│   │   └── module.go        # Módulo de inicialização do banco de dados
│   ├── httpApi
│   │   ├── module.go        # Configuração do módulo HTTP
│   │   └── server.go        # Servidor HTTP
│   └── repository
│       └── user_repo
│           ├── user_repo.go        # Interface do repositório de usuário
│           └── user_repo_impl.go   # Implementação do repositório de usuário
├── pkg
│   ├── env
│   │   └── env.go            # Configurações de variáveis de ambiente
│   └── projectError
│       └── error.go          # Gerenciamento de erros
```
### Bibliotecas Utilizadas

- **go.uber.org/fx v1.23.0**: Utilizada para gerenciar o ciclo de vida da aplicação e a injeção de dependências, facilitando a modularização.
- **net/http**: Biblioteca padrão do Go para lidar com requisições HTTP.
- **database/sql**: Biblioteca padrão do Go para manipulação de bancos de dados SQL, garantindo uma integração direta e performática.

### Configuração do Ambiente

A configuração do ambiente é realizada através do pacote `configEnv`, onde todas as variáveis de ambiente necessárias para a aplicação são gerenciadas. Isso inclui configurações do banco de dados, porta do servidor, entre outras.

### Configuração do Banco de Dados

O banco de dados é configurado no módulo `db`. A aplicação utiliza scripts de migração SQL armazenados na pasta `migrations` para definir a estrutura inicial do banco. Estes scripts são aplicados durante a inicialização, garantindo que o banco de dados esteja sempre na versão correta.

### Estrutura da API

- **Controller**: Contém a lógica dos controladores, que gerenciam a comunicação entre o modelo, o repositório e as rotas.
- **Model**: Define as estruturas de dados que representam os usuários e seus métodos de manipulação.
- **Router**: Define as rotas da API, separadas por funcionalidades, e mapeia as requisições para os controladores corretos.
- **Repository**: Contém a lógica de acesso ao banco de dados, permitindo a manipulação de dados de forma independente dos controladores.

### Como Executar

1. Clone o repositório:
   ```bash
   git clone git@github.com:lkgiovani/go-api.git

2. Instale as dependências:
   ```bash
   go mod tidy
   ```
   

3. Configure as variáveis de ambiente conforme descrito em configEnv/config.go.

4. Execute a aplicação:
   ```bash
   go run cmd/start/main.go
