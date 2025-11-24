# 📦 Estoka WMS: Mini Warehouse Management System (Go & Flutter)

> O Estoka WMS é uma solução de gerenciamento de armazém simplificada e eficiente, projetada para atender às necessidades básicas de controle de estoque de pequenos negócios e operações autônomas. Focado na **Agilidade**, o Produto Mínimo Viável (MVP) foi concebido para ser entregue em apenas **10 dias**.

## 🎯 Status do Projeto

| Status Global | Fase Atual | Foco |
| :--- | :--- | :--- |
| 🟡 **Planejamento / Pré-Desenvolvimento** | 📅 Fase 1: O MVP | Ciclo de vida essencial do produto: CRUD de Produtos + Registro de Entrada/Saída. |

## ✨ Tecnologias Principais

O projeto adota uma arquitetura Cliente-Servidor focada em performance e portabilidade.

| Categoria | Tecnologia | Justificativa de Escolha (MVP) |
| :--- | :--- | :--- |
| **Backend** | **Go (Golang)** | Alta performance, tipagem forte e simplicidade. |
| **Router API** | **Chi** | Leve, idiomático e alinhado com o `net/http` padrão do Go. |
| **Banco de Dados** | **SQLite** | *Serverless*, configuração zero, ideal para MVPs locais e portabilidade. |
| **Frontend** | **Dart (Flutter)** | Desenvolvimento multiplataforma nativo e UI fluida. |
| **Gerência de Estado** | **Provider** | Curva de aprendizado baixa e redução de boilerplate para um MVP rápido. |

## 🛠️ Escopo do MVP (Fase 1)

O foco principal é o ciclo de vida essencial do produto.

### Backend (Go)

  * ✅ Setup Inicial (Go Mod, Estrutura de Pastas, Configuração SQLite).
  * \[ \] Core Domain: Definir Struct `Product`.
  * \[ \] Gestão de Produtos: Implementar Repositório, Serviço e Handler para **CRUD** de produtos.
  * \[ \] Controle de Estoque: Implementar endpoints **POST** `/inventory/entries` (Incrementar Saldo) e **POST** `/inventory/exits` (Decrementar Saldo e validar saldo negativo).

### Frontend (Flutter Mobile)

  * \[ \] Setup Inicial (Projeto Flutter, Estrutura Feature-First).
  * \[ \] Feature Produtos: Criar `ProductModel` e `ProductProvider`.
  * \[ \] Feature Produtos: Implementar Tela de **Listagem** com Card de produtos.
  * \[ \] Feature Produtos: Implementar Tela de **Cadastro/Edição** (Formulário com validação).
  * \[ \] Feature Movimentação: Criar modais/telas para Registro de **Entrada** e **Saída**.

## 💻 Estrutura de Diretórios

O projeto será dividido em dois repositórios, seguindo padrões de organização específicos para cada stack:

### 4.1. Backend (`mini-wms-api/`)

```
mini-wms-api/
├── cmd/
│   └── api/
│       └── main.go           # Entrypoint: Router, DI, Server Start
├── internal/
│   ├── core/                 # Entidades e Interfaces
│   ├── handlers/             # Controllers HTTP (Chi)
│   ├── services/             # Regras de Negócio
│   └── repositories/         # Queries SQL (SQLite)
├── migrations/               # Scripts SQL de criação de tabelas
└── go.mod
```

### 4.2. Frontend (`mini-wms-app/`)

Estrutura **Feature-First** para melhor manutenibilidade.

```
mini-wms-app/  
├── lib/  
│   ├── main.dart             # Setup Global (Providers, Theme)  
│   ├── core/                 # Utilitários e Widgets genéricos
│   ├── features/  
│   │   ├── products/         # Feature: CRUD Produtos  
│   │   └── inventory/        # Feature: Entradas/Saídas  
│   └── navigation/           # Definição de Rotas (AppRouter)  
└── pubspec.yaml
```

## 🔗 Endpoints Principais (API Go)

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| **POST** | `/products` | Cria um novo produto |
| **GET** | `/products` | Lista todos os produtos |
| **PUT** | `/products/{id}` | Atualiza um produto existente |
| **POST** | `/inventory/entries` | Registra uma entrada (Incrementa saldo) |
| **POST** | `/inventory/exits` | Registra uma saída (Decrementa saldo) |
| **GET** | `/inventory/products/{productId}/stock` | Consulta o saldo atual de um produto |

## 🚀 Próximos Passos (Pós-MVP)

Após a entrega do MVP funcional, as seguintes fases de evolução estão planejadas:

| Fase | Foco | Atividades Chave |
| :--- | :--- | :--- |
| **Fase 2: Segurança & Auditoria (v1.1)** | Transformar o MVP em produto seguro e rastreável. | Implementar **Autenticação JWT** e criar a tabela `inventory_logs` para **Histórico Transacional**. |
| **Fase 3: Eficiência Operacional (v1.2)** | Agilidade no chão de loja e inteligência de dados. | Integrar **Leitor de Código de Barras** (via `mobile_scanner`) e gerar **Alertas** para produtos abaixo do estoque mínimo. |
| **Fase 4: Ecossistema & Escala (v2.0)** | Gestão remota e multi-plataforma. | Criar **Dockerfile** e adicionar suporte a **PostgreSQL** para deploy em nuvem. Desenvolver **Painel Administrativo Web (React)**.Com base nos documentos do projeto Estoka WMS (Plano de Projeto, Status de Implementação e Documentação) e seguindo a estrutura de um bom `README.md` conforme as boas práticas de engenharia de software, abaixo está o conteúdo sugerido para o `README.md` do repositório principal do projeto.

-----

# 📦 Estoka WMS: Mini Warehouse Management System (Go & Flutter)

> O Estoka WMS é uma solução de gerenciamento de armazém simplificada e eficiente, projetada para atender às necessidades básicas de controle de estoque de pequenos negócios e operações autônomas. Focado na **Agilidade**, o Produto Mínimo Viável (MVP) foi concebido para ser entregue em apenas **10 dias**.

## 🎯 Status do Projeto

| Status Global | Fase Atual | Foco |
| :--- | :--- | :--- |
| 🟡 **Planejamento / Pré-Desenvolvimento** | 📅 Fase 1: O MVP | Ciclo de vida essencial do produto: CRUD de Produtos + Registro de Entrada/Saída. |

## ✨ Tecnologias Principais

O projeto adota uma arquitetura Cliente-Servidor focada em performance e portabilidade.

| Categoria | Tecnologia | Justificativa de Escolha (MVP) |
| :--- | :--- | :--- |
| **Backend** | **Go (Golang)** | Alta performance, tipagem forte e simplicidade. |
| **Router API** | **Chi** | Leve, idiomático e alinhado com o `net/http` padrão do Go. |
| **Banco de Dados** | **SQLite** | *Serverless*, configuração zero, ideal para MVPs locais e portabilidade. |
| **Frontend** | **Dart (Flutter)** | Desenvolvimento multiplataforma nativo e UI fluida. |
| **Gerência de Estado** | **Provider** | Curva de aprendizado baixa e redução de boilerplate para um MVP rápido. |

## 🛠️ Escopo do MVP (Fase 1)

O foco principal é o ciclo de vida essencial do produto.

### Backend (Go)

  * ✅ Setup Inicial (Go Mod, Estrutura de Pastas, Configuração SQLite).
  * \[ \] Core Domain: Definir Struct `Product`.
  * \[ \] Gestão de Produtos: Implementar Repositório, Serviço e Handler para **CRUD** de produtos.
  * \[ \] Controle de Estoque: Implementar endpoints **POST** `/inventory/entries` (Incrementar Saldo) e **POST** `/inventory/exits` (Decrementar Saldo e validar saldo negativo).

### Frontend (Flutter Mobile)

  * \[ \] Setup Inicial (Projeto Flutter, Estrutura Feature-First).
  * \[ \] Feature Produtos: Criar `ProductModel` e `ProductProvider`.
  * \[ \] Feature Produtos: Implementar Tela de **Listagem** com Card de produtos.
  * \[ \] Feature Produtos: Implementar Tela de **Cadastro/Edição** (Formulário com validação).
  * \[ \] Feature Movimentação: Criar modais/telas para Registro de **Entrada** e **Saída**.

## 💻 Estrutura de Diretórios

O projeto será dividido em dois repositórios, seguindo padrões de organização específicos para cada stack:

### 4.1. Backend (`mini-wms-api/`)

```
mini-wms-api/
├── cmd/
│   └── api/
│       └── main.go           # Entrypoint: Router, DI, Server Start
├── internal/
│   ├── core/                 # Entidades e Interfaces
│   ├── handlers/             # Controllers HTTP (Chi)
│   ├── services/             # Regras de Negócio
│   └── repositories/         # Queries SQL (SQLite)
├── migrations/               # Scripts SQL de criação de tabelas
└── go.mod
```

### 4.2. Frontend (`mini-wms-app/`)

Estrutura **Feature-First** para melhor manutenibilidade.

```
mini-wms-app/  
├── lib/  
│   ├── main.dart             # Setup Global (Providers, Theme)  
│   ├── core/                 # Utilitários e Widgets genéricos
│   ├── features/  
│   │   ├── products/         # Feature: CRUD Produtos  
│   │   └── inventory/        # Feature: Entradas/Saídas  
│   └── navigation/           # Definição de Rotas (AppRouter)  
└── pubspec.yaml
```

## 🔗 Endpoints Principais (API Go)

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| **POST** | `/products` | Cria um novo produto |
| **GET** | `/products` | Lista todos os produtos |
| **PUT** | `/products/{id}` | Atualiza um produto existente |
| **POST** | `/inventory/entries` | Registra uma entrada (Incrementa saldo) |
| **POST** | `/inventory/exits` | Registra uma saída (Decrementa saldo) |
| **GET** | `/inventory/products/{productId}/stock` | Consulta o saldo atual de um produto |

## 🚀 Próximos Passos (Pós-MVP)

Após a entrega do MVP funcional, as seguintes fases de evolução estão planejadas:

| Fase | Foco | Atividades Chave |
| :--- | :--- | :--- |
| **Fase 2: Segurança & Auditoria (v1.1)** | Transformar o MVP em produto seguro e rastreável. | Implementar **Autenticação JWT** e criar a tabela `inventory_logs` para **Histórico Transacional**. |
| **Fase 3: Eficiência Operacional (v1.2)** | Agilidade no chão de loja e inteligência de dados. | Integrar **Leitor de Código de Barras** (via `mobile_scanner`) e gerar **Alertas** para produtos abaixo do estoque mínimo. |
| **Fase 4: Ecossistema & Escala (v2.0)** | Gestão remota e multi-plataforma. | Criar **Dockerfile** e adicionar suporte a **PostgreSQL** para deploy em nuvem. Desenvolver **Painel Administrativo Web (React)**. |

Com base nos documentos do projeto Estoka WMS (Plano de Projeto, Status de Implementação e Documentação) e seguindo a estrutura de um bom `README.md` conforme as boas práticas de engenharia de software, abaixo está o conteúdo sugerido para o `README.md` do repositório principal do projeto.

-----

# 📦 Estoka WMS: Mini Warehouse Management System (Go & Flutter)

> O Estoka WMS é uma solução de gerenciamento de armazém simplificada e eficiente, projetada para atender às necessidades básicas de controle de estoque de pequenos negócios e operações autônomas. Focado na **Agilidade**, o Produto Mínimo Viável (MVP) foi concebido para ser entregue em apenas **10 dias**.

## 🎯 Status do Projeto

| Status Global | Fase Atual | Foco |
| :--- | :--- | :--- |
| 🟡 **Planejamento / Pré-Desenvolvimento** | 📅 Fase 1: O MVP | Ciclo de vida essencial do produto: CRUD de Produtos + Registro de Entrada/Saída. |

## ✨ Tecnologias Principais

O projeto adota uma arquitetura Cliente-Servidor focada em performance e portabilidade.

| Categoria | Tecnologia | Justificativa de Escolha (MVP) |
| :--- | :--- | :--- |
| **Backend** | **Go (Golang)** | Alta performance, tipagem forte e simplicidade. |
| **Router API** | **Chi** | Leve, idiomático e alinhado com o `net/http` padrão do Go. |
| **Banco de Dados** | **SQLite** | *Serverless*, configuração zero, ideal para MVPs locais e portabilidade. |
| **Frontend** | **Dart (Flutter)** | Desenvolvimento multiplataforma nativo e UI fluida. |
| **Gerência de Estado** | **Provider** | Curva de aprendizado baixa e redução de boilerplate para um MVP rápido. |

## 🛠️ Escopo do MVP (Fase 1)

O foco principal é o ciclo de vida essencial do produto.

### Backend (Go)

  * ✅ Setup Inicial (Go Mod, Estrutura de Pastas, Configuração SQLite).
  * \[ \] Core Domain: Definir Struct `Product`.
  * \[ \] Gestão de Produtos: Implementar Repositório, Serviço e Handler para **CRUD** de produtos.
  * \[ \] Controle de Estoque: Implementar endpoints **POST** `/inventory/entries` (Incrementar Saldo) e **POST** `/inventory/exits` (Decrementar Saldo e validar saldo negativo).

### Frontend (Flutter Mobile)

  * \[ \] Setup Inicial (Projeto Flutter, Estrutura Feature-First).
  * \[ \] Feature Produtos: Criar `ProductModel` e `ProductProvider`.
  * \[ \] Feature Produtos: Implementar Tela de **Listagem** com Card de produtos.
  * \[ \] Feature Produtos: Implementar Tela de **Cadastro/Edição** (Formulário com validação).
  * \[ \] Feature Movimentação: Criar modais/telas para Registro de **Entrada** e **Saída**.

## 💻 Estrutura de Diretórios

O projeto será dividido em dois repositórios, seguindo padrões de organização específicos para cada stack:

### 4.1. Backend (`mini-wms-api/`)

```
mini-wms-api/
├── cmd/
│   └── api/
│       └── main.go           # Entrypoint: Router, DI, Server Start
├── internal/
│   ├── core/                 # Entidades e Interfaces
│   ├── handlers/             # Controllers HTTP (Chi)
│   ├── services/             # Regras de Negócio
│   └── repositories/         # Queries SQL (SQLite)
├── migrations/               # Scripts SQL de criação de tabelas
└── go.mod
```

### 4.2. Frontend (`mini-wms-app/`)

Estrutura **Feature-First** para melhor manutenibilidade.

```
mini-wms-app/  
├── lib/  
│   ├── main.dart             # Setup Global (Providers, Theme)  
│   ├── core/                 # Utilitários e Widgets genéricos
│   ├── features/  
│   │   ├── products/         # Feature: CRUD Produtos  
│   │   └── inventory/        # Feature: Entradas/Saídas  
│   └── navigation/           # Definição de Rotas (AppRouter)  
└── pubspec.yaml
```

## 🔗 Endpoints Principais (API Go)

| Método | Rota | Descrição |
| :--- | :--- | :--- |
| **POST** | `/products` | Cria um novo produto |
| **GET** | `/products` | Lista todos os produtos |
| **PUT** | `/products/{id}` | Atualiza um produto existente |
| **POST** | `/inventory/entries` | Registra uma entrada (Incrementa saldo) |
| **POST** | `/inventory/exits` | Registra uma saída (Decrementa saldo) |
| **GET** | `/inventory/products/{productId}/stock` | Consulta o saldo atual de um produto |

## 🚀 Próximos Passos (Pós-MVP)

Após a entrega do MVP funcional, as seguintes fases de evolução estão planejadas:

| Fase | Foco | Atividades Chave |
| :--- | :--- | :--- |
| **Fase 2: Segurança & Auditoria (v1.1)** | Transformar o MVP em produto seguro e rastreável. | Implementar **Autenticação JWT** e criar a tabela `inventory_logs` para **Histórico Transacional**. |
| **Fase 3: Eficiência Operacional (v1.2)** | Agilidade no chão de loja e inteligência de dados. | Integrar **Leitor de Código de Barras** (via `mobile_scanner`) e gerar **Alertas** para produtos abaixo do estoque mínimo. |
| **Fase 4: Ecossistema & Escala (v2.0)** | Gestão remota e multi-plataforma. | Criar **Dockerfile** e adicionar suporte a **PostgreSQL** para deploy em nuvem. Desenvolver **Painel Administrativo Web (React)**.

**Built with ❤️ by KOYOS STUDIOS**