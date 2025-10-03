# 📸 Photo Album API

API para gerenciamento de um álbum de fotos de viagens.  
Permite organizar, cadastrar e consultar fotos de forma simples e prática.

---

## 🚀 Como rodar o projeto

### Pré-requisitos
- [Go](https://go.dev/dl/) instalado (versão 1.18+ recomendada)
- [Docker](https://www.docker.com/) (opcional, caso queira subir banco/serviços)

### Iniciando a aplicação
No diretório raiz do projeto, execute:

```bash
go run .\cmd\
```
### 📖 Documentação da API

A documentação interativa (Swagger UI) pode ser acessada em:

👉 http://localhost:8080/swagger/index.html

### ✨ Funcionalidades

📸 Cadastro e gerenciamento de álbuns de fotos de viagens

🔑 Autenticação com JWT (JSON Web Token)

♻️ Sistema de Refresh Token para renovação de sessões

⬆️ Upload de imagens

📂 Organização das fotos em álbuns

### 📂 Estrutura do Projeto
```bash
├── cmd/               # Ponto de entrada da aplicação
├── internal/          # Regras de negócio e lógica interna
├── pkg/               # Pacotes reutilizáveis
├── docs/              # Arquivos de documentação (Swagger/OpenAPI)
├── db/                # Migrations do banco de dados
├── go.mod             # Dependências do Go
├── go.sum             # Checksum das dependências
└── README.md          # Documentação do projeto
```
### 🛠 Tecnologias

- Go (Golang)

- Swagger/OpenAPI para documentação

- Docker para containerização

- PostgreSQL como banco de dados

- JWT (JSON Web Token) para autenticação

- Refresh Token para controle de sessões


