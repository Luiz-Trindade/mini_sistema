# Mini Sistema - Monorepo Fullstack (Go + Quasar/Vue 3)

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL_v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.x-4FC08D?style=flat&logo=vuedotjs)](https://vuejs.org/)
[![Quasar Framework](https://img.shields.io/badge/Quasar-v2.x-1976D2?style=flat&logo=quasar)](https://quasar.dev/)
[![Bun](https://img.shields.io/badge/Bun-Latest-000000?style=flat&logo=bun)](https://bun.sh/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql)](https://www.postgresql.org/)
[![PgBouncer](https://img.shields.io/badge/PgBouncer-Enabled-4169E1?style=flat&logo=postgresql)](https://www.pgbouncer.org/)

Uma aplicação web fullstack de alto desempenho que combina a velocidade e a concorrência do **Go** com o rico ecossistema de componentes de interface do **Vue 3** e do **Quasar Framework**. Os arquivos do frontend são compilados em pacotes estáticos e incorporados diretamente ao executável Go via `//go:embed`, proporcionando uma distribuição autocontida em um único binário.

---

## 🚀 Principais vantagens desta arquitetura

### ⚡ Desempenho extremo e baixa latência

- **Zero E/S de disco:** Os arquivos estáticos (`index.html`, JavaScript, CSS e imagens) são incorporados diretamente ao binário Go compilado (`//go:embed`) e servidos diretamente da RAM, com tempos de resposta em microssegundos.
- **Baixo consumo de recursos:** A aplicação fullstack completa (frontend + API backend) consome apenas **15 MB–30 MB de RAM** durante a execução.
- **Concorrência nativa:** O `net/http` e o `Gin` do Go lidam facilmente com dezenas de milhares de conexões simultâneas usando goroutines leves.

### 🛡️ Implantação com mesma origem (sem sobrecarga de CORS)

- Em produção, o frontend SPA e a API REST do backend compartilham exatamente o mesmo protocolo, domínio e porta (`http://localhost:8080`).
- Elimina requisições preliminares `OPTIONS` do navegador, reduz viagens de ida e volta HTTP e dispensa o gerenciamento complexo de cabeçalhos CORS.

### 📦 Implantação simples em um único binário

- **Node.js ou Nginx não são necessários no servidor:** você não precisa do runtime Node, de servidores de desenvolvimento Webpack/Vite ou de configurações complexas do Nginx em produção.
- Um único arquivo executável de aproximadamente 40 MB contém toda a interface do frontend, a camada de banco de dados GORM, a autenticação JWT e os endpoints da API.

### 🔌 Pooling de banco de dados escalável com PgBouncer

- Gerenciamento de conexões PostgreSQL de alta concorrência por meio do **PgBouncer** executado em modo de transação.
- Minimiza a sobrecarga das conexões e otimiza o uso de memória durante picos de tráfego.

---

## 🛠️ Stack tecnológica

### Backend

- **Language:** Go (1.22+)
- **Framework web:** [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/) (driver PostgreSQL com prepared statements desabilitadas para compatibilidade com o PgBouncer)
- **Autenticação:** JWT (JSON Web Tokens) + Bcrypt
- **Banco de dados:** PostgreSQL 16 + PgBouncer

### Frontend

- **Framework:** [Vue 3](https://vuejs.org/) (Composition API / `<script setup>`)
- **Ecossistema de UI:** [Quasar Framework v2](https://quasar.dev/)
- **Runtime e bundler:** [Bun](https://bun.sh/) + Vite
- **Cliente HTTP:** Axios (com interceptors dinâmicos de requisição/resposta)
- **Visualização de dados:** [ECharts](https://echarts.apache.org/) + `vue-echarts`

---

## 📁 Estrutura do projeto

```text
mini_sistema/
├── auth/                 # Authentication domain (models, services, selectors, handlers)
├── database/             # GORM & PostgreSQL / PgBouncer connection module
├── frontend/             # Vue 3 + Quasar frontend (managed via Bun)
│   ├── src/
│   │   ├── boot/         # Quasar boot files (axios, etc.)
│   │   ├── components/   # Reusable Vue components
│   │   ├── pages/        # Application views (Index/Dashboard, Login)
│   │   └── router/       # Vue Router configuration & Navigation Guards
│   ├── quasar.config.js  # Quasar & Vite configuration (Proxy enabled for dev)
│   └── dist/
│       └── spa/          # Compiled static build target
├── docker-compose.yml    # PostgreSQL and PgBouncer services
├── main.go               # Go entrypoint (Gin routes + //go:embed static file server)
├── go.mod
└── go.sum
```

---

## ⚙️ Variáveis de ambiente

O backend lê dinamicamente a configuração das variáveis de ambiente do sistema e usa valores locais padrão quando elas não estão definidas:

| Variável      | Descrição                       | Padrão          |
| :------------ | :------------------------------ | :-------------- |
| `DB_HOST`     | Endereço do host do banco       | `localhost`     |
| `DB_PORT`     | Porta do PgBouncer / PostgreSQL | `6432`          |
| `DB_USER`     | Usuário do banco                | `app_user`      |
| `DB_PASSWORD` | Senha do banco                  | `app_password`  |
| `DB_NAME`     | Nome do banco                   | `mini_sistema`  |
| `DB_SSLMODE`  | Modo de conexão SSL             | `disable`       |
| `DB_TIMEZONE` | Fuso horário do banco           | `America/Belem` |

---

## 🚦 Primeiros passos

### Pré-requisitos

- **Go** 1.22 or higher
- **Bun** (or Node.js 18+)
- **Docker & Docker Compose** (for PostgreSQL + PgBouncer)

### 1. Clonar e configurar o banco de dados

```bash
git clone https://github.com/Luiz-Trindade/mini_sistema.git
cd mini_sistema

# Iniciar os serviços PostgreSQL e PgBouncer
docker compose up -d
```

### 2. Modo de desenvolvimento (Hot Reload / Live Reload)

Para desenvolvimento diário com atualizações instantâneas da interface (HMR):

**Terminal 1 (Backend Go):**

```bash
go run main.go
```

**Terminal 2 (Frontend Quasar com Bun):**

```bash
cd frontend
bun install
bun dev
```

> O servidor de desenvolvimento Quasar é executado em `http://localhost:9000` e encaminha automaticamente as requisições `/auth` e `/api` para o backend Go em `http://localhost:8080`.

---

## 📦 Build de produção e geração do binário único

Para compilar a aplicação inteira em um único executável:

```bash
# 1. Compilar os arquivos estáticos da SPA no diretório frontend
cd frontend
bun run build
cd ..

# 2. Compilar o binário Go unificado com os arquivos do frontend incorporados
go build -o mini_sistema main.go

# 3. Executar o binário independente
./mini_sistema
```

Acesse a aplicação de produção em `http://localhost:8080`.

---

## 📄 Licença

Este projeto está licenciado sob a **GNU Affero General Public License v3.0 (AGPL-3.0)**.

> Você é livre para copiar, modificar e distribuir este software, desde que quaisquer trabalhos derivados ou serviços hospedados em rede baseados neste código também sejam licenciados sob a AGPL-3.0 e disponibilizados aos usuários pela rede.
