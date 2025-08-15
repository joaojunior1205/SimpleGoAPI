SimpleGoAPI

Uma API RESTful em Go com autenticação JWT, roteamento modular e middleware personalizável. Projetada para ser leve, escalável e clara, ideal para portfólio.

Tecnologias

Go 1.25.0+

Golang JWT v5

HTTP nativo (net/http)

Middleware e roteamento customizados

Estrutura do projeto
SimpleGoAPI/
├── auth/               # Geração e validação de JWT
├── controllers/        # Controllers HTTP
├── handlers/           # Lógica de negócio separada dos controllers
├── middlewares/        # Middlewares (ex: AuthMiddleware)
├── routes/             # Router customizado com suporte a parâmetros e grupos
├── main.go             # Ponto de entrada da aplicação

Funcionalidades

Autenticação JWT

Login com email e senha.

Tokens expiram em 1 hora.

Middleware verifica token e injeta userID no contexto da request.

Roteamento modular

Rotas separadas por funcionalidade (auth, users, profile).

Suporte a parâmetros via URL (/users/:id) e métodos HTTP (GET, POST, PUT, DELETE).

Middleware aplicável globalmente ou por grupo de rotas.

Suporte a contexto

userID do JWT disponível em qualquer handler protegido.

Parâmetros de rota acessíveis via helper routes.Param(r, "paramName").

Exemplos de uso
Login

Request:

POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "123456"
}


Response:

{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}

Acessando rota protegida

Request:

GET /api/v1/users/123
Authorization: Bearer <token>


Handler:

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
    userID := middlewares.GetUserID(r) // Pega userID do JWT
    fmt.Fprintf(w, "Usuário logado: %s", userID)
}

Como rodar localmente

Clone o repositório:

git clone https://github.com/seu-usuario/SimpleGoAPI.git
cd SimpleGoAPI


Instale dependências:

go mod tidy


Execute a aplicação:

go run main.go


A API estará disponível em http://localhost:8000.

Boas práticas demonstradas

Separação de responsabilidades: controllers, handlers, middlewares, auth.

Contexto de request para dados de usuário (userID) e parâmetros de rota.

Modularidade: fácil adicionar novas rotas e middlewares.

Middleware global e por grupo de rotas.

Tratamento correto de erros HTTP e status codes.

Próximos passos (portfólio)

Integração com banco de dados (PostgreSQL ou MongoDB).

Refresh token JWT.

Registro de usuários, recuperação de senha e roles.

Testes unitários para handlers, middleware e serviços.