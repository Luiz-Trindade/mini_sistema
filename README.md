# 🚀 Mini Sistema (Go + Vue 3 Quasar PWA)

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D)
![Quasar](https://img.shields.io/badge/Quasar-1976D2?style=for-the-badge&logo=quasar&logoColor=white)
![PWA](https://img.shields.io/badge/PWA-5A0FC8?style=for-the-badge&logo=pwa&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![License AGPLv3](https://img.shields.io/badge/License-AGPLv3-blue.svg?style=for-the-badge)

Uma aplicação web & mobile moderna, ultrarrápida e resiliente que combina um backend em **Go (Gin + GORM)** com um frontend **Quasar (Vue 3)** compilado como **PWA (Progressive Web App)**, totalmente embutido no próprio binário executável via `//go:embed`.

---

## 💡 Por que esta Arquitetura?

### 📲 Progressive Web App (PWA) Nativo

- **Instalável:** Pode ser instalado diretamente no smartphone (Android/iOS) ou Desktop como se fosse um aplicativo nativo.
- **Recursos Offline:** Utiliza Service Workers (`sw.js`) e armazenamento em cache inteligente para manter a interface funcional e responsiva mesmo com oscilações de rede.
- **Manifest Personalizado:** Experiência de tela cheia (_standalone_) sem barras de navegação do browser.

### ⚡ Performance Bruta & Baixo Consumo

- **Zero I/O de Disco:** O build PWA (`frontend/dist/pwa`) é compilado diretamente dentro do binário Go. A interface é servida diretamente da memória RAM com respostas na casa dos **microsegundos**.
- **Zero Overhead de Servidores Web:** Não necessita de Nginx ou Node.js em produção. O executável Go lida com os arquivos estáticos e os endpoints de API.
- **Concorrência do Go:** Lida com milhares de requisições simultâneas utilizando _goroutines_ leves.

### 🛢️ Banco de Dados Escalável

- **PostgreSQL + PgBouncer:** Utiliza pooling de conexões em modo _transaction_ via PgBouncer na porta `6432` para suportar alto volume de conexões simultâneas sem sobrecarregar o banco.

---

## 🛠️ Tecnologias Utilizadas

- **Backend:** Go 1.22+, [Gin Web Framework](https://gin-gonic.com/), [GORM](https://gorm.io/)
- **Frontend:** Vue 3 (Composition API), [Quasar Framework](https://quasar.dev/), Bun
- **PWA:** Service Worker Workbox, App Manifest, Offline Caching
- **Autenticação:** JWT com controle de versão de tokens (`token_version`) e Revogação
- **Banco de Dados:** PostgreSQL 16+, PgBouncer
- **Containerização:** Docker Compose

---

## 📁 Estrutura do Projeto

```text
mini_sistema/
├── auth/                 # Módulo de Autenticação (Models, Services, API, Selectors)
├── database/             # Configuração GORM e PgBouncer
├── frontend/             # Código-fonte Vue 3 / Quasar PWA
│   ├── src/              # Páginas, Componentes e Boot Files
│   ├── src-pwa/          # Configurações do Service Worker e Manifest
│   └── dist/pwa/         # Build PWA gerado pelo Quasar
├── main.go               # Entrypoint Go com //go:embed e servidor Gin
├── docker-compose.yml    # Serviços Postgres e PgBouncer
└── README.md
```

---

## 🚀 Como Rodar o Projeto

### Pró-requisitos

- **Go** (v1.22+)
- **Bun** ou **Node.js**
- **Docker & Docker Compose**

### 1. Iniciar a Infraestrutura de Banco de Dados

```bash
docker compose up -d

```

### 2. Desenvolvimento Local (com Live-Reload / HMR)

Rode o backend e o frontend em terminais separados:

- **Terminal 1 (Backend Go):**

```bash
go run main.go

```

_Servidor rodando em `http://localhost:8080_`

- **Terminal 2 (Frontend Quasar PWA):**

```bash
cd frontend
bun quasar dev -m pwa

```

_Interface rodando em `http://localhost:9000` com proxy automático de API para o Go._

---

## 📦 Compilação para Produção (Binário Único)

1. **Compilar o PWA do Frontend:**

```bash
cd frontend
bun quasar build -m pwa
cd ..

```

2. **Compilar o Executável Go:**

```bash
go build -o mini_sistema main.go

```

3. **Executar:**

```bash
./mini_sistema

```

Acesse `http://localhost:8080` para visualizar a PWA completa rodando a partir de um único arquivo binário!

---

## 📄 Licença

Este projeto está licenciado sob a licença **GNU Affero General Public License v3.0 (AGPL-3.0)**.
