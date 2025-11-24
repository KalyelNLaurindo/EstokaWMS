# **Documentação do Projeto Estoka WMS (Mini WMS Simplificado)**

## **Resumo**

O Projeto Estoka WMS é uma solução de gerenciamento de armazém simplificada e eficiente, desenvolvida com backend em Go e frontend em Flutter. O sistema foca na entrega de um Produto Mínimo Viável (MVP) em 10 dias, oferecendo funcionalidades essenciais como controle de inventário, registro de entradas/saídas e consulta de estoque para pequenos negócios e operações autônomas.

## **Abstract**

The Estoka WMS Project is a simplified and efficient warehouse management solution, developed with a Go backend and Flutter frontend. The system focuses on delivering a Minimum Viable Product (MVP) within 10 days, offering essential functionalities such as inventory control, entry/exit registration, and stock queries for small businesses and autonomous operations.

## **1\. Introdução**

### **1.1. Contexto e Problema**

Pequenos comerciantes e autônomos frequentemente carecem de ferramentas digitais acessíveis para gerenciar seus estoques, recorrendo a processos manuais propensos a erros. Sistemas WMS tradicionais são excessivamente complexos e caros para operações que necessitam apenas controlar o fluxo básico de mercadorias.

### **1.2. Objetivos e Escopo**

O objetivo do Estoka WMS é fornecer uma ferramenta intuitiva e leve para tarefas críticas de estoque.

1. **Agilidade:** Desenvolvimento de um MVP funcional em 10 dias.  
2. **Simplicidade:** Foco exclusivo no ciclo de vida do produto (Entrada \-\> Estoque \-\> Saída).  
3. **Performance:** Utilização de tecnologias performáticas (Go) e portáteis (SQLite).

O escopo do MVP abrange desde o CRUD de produtos até a atualização de saldos via movimentações, excluindo autenticação complexa ou relatórios avançados nesta fase.

### **1.3. Público-Alvo**

* Pequenos comerciantes (lojas de bairro, mini-mercados).  
* Empreendedores de e-commerce iniciantes.  
* Uso pessoal para organização de inventários domésticos ou de coleções.

## **2\. Requisitos**

### **2.1. Requisitos Funcionais (RF)**

* **RF01 \- Gerenciamento de Produtos:** O sistema deve permitir criar, ler, atualizar e excluir (CRUD) produtos com atributos como SKU, nome e unidade.  
* **RF02 \- Registro de Entradas:** O sistema deve permitir registrar a entrada de itens, incrementando automaticamente o saldo em estoque.  
* **RF03 \- Registro de Saídas:** O sistema deve permitir registrar a saída de itens, decrementando o saldo em estoque.  
* **RF04 \- Consulta de Estoque:** O sistema deve exibir a quantidade atual disponível de cada produto em tempo real.  
* **RF05 \- Validação de Dados:** O sistema deve validar entradas (ex: quantidades não negativas, campos obrigatórios) antes da persistência.

### **2.2. Requisitos Não Funcionais (RNF)**

* **RNF01 \- Portabilidade:** O sistema deve utilizar SQLite embarcado para facilitar o deploy e configuração local.  
* **RNF02 \- Usabilidade:** A interface móvel deve seguir os princípios do Material Design para familiaridade imediata.  
* **RNF03 \- Desempenho:** A API deve ser *stateless* e responder rapidamente, alavancando a concorrência do Go.  
* **RNF04 \- Manutenibilidade:** O código deve seguir estruturas limpas (Feature-First no Flutter, Camadas no Go) para facilitar evoluções futuras.

## **3\. Arquitetura e Modelagem**

### **3.1. Arquitetura e Tecnologias**

| Categoria | Tecnologia | Justificativa |
| :---- | :---- | :---- |
| **Linguagem Backend** | Go (Golang) | Alta performance, tipagem forte e simplicidade. |
| **Router API** | Chi | Leve, idiomático e compatível com net/http padrão. |
| **Banco de Dados** | SQLite | Serverless, configuração zero, ideal para MVPs locais. |
| **Linguagem Frontend** | Dart (Flutter) | Desenvolvimento multiplataforma nativo e UI fluida. |
| **Gerência de Estado** | Provider | Curva de aprendizado baixa e redução de boilerplate. |
| **Comunicação** | REST / JSON | Padrão de indústria para interoperabilidade. |

### **3.2. Estrutura da Camada de Dados**

O MVP utiliza um modelo simplificado focado na atualização direta de saldos, sem tabelas complexas de histórico transacional nesta fase.

#### **📦 Schema do Banco de Dados (SQLite)**

| Tabela | Coluna | Tipo | Descrição |
| :---- | :---- | :---- | :---- |
| products | id | INTEGER (PK) | Identificador único. |
|  | code | TEXT (Unique) | SKU do produto. |
|  | name | TEXT | Nome descritivo. |
|  | unit | TEXT | Unidade (UN, KG, L). |
|  | quantity | INTEGER | Saldo atual (Default 0). |
|  | updated\_at | DATETIME | Timestamp de modificação. |

### **3.3. Fluxograma (Data Flow)**

graph TD  
    subgraph "Frontend (Mobile)"  
        A\[App Flutter\] \--\>|JSON Request| B(API Client)  
    end

    subgraph "Backend (Server)"  
        B \--\>|HTTP/REST| C\[Chi Router\]  
        C \--\>|Handler| D(Product Service)  
        C \--\>|Handler| E(Inventory Service)  
        D \--\>|CRUD Ops| F{SQLite DB}  
        E \--\>|Update Stock| F  
    end

    subgraph "Persistência"  
        F \--\>|Read/Write| G\[(miniwms.db)\]  
    end

### **3.4. Diagramas de Modelagem**

#### **Diagrama de Sequência (Fluxo de Entrada de Estoque)**

sequenceDiagram  
    participant User  
    participant App as App Flutter  
    participant API as API Go  
    participant DB as SQLite

    User-\>\>App: Seleciona Produto e Qtd (Entrada)  
    App-\>\>API: POST /inventory/entries {id, qty}  
    API-\>\>API: Valida Requisição  
    API-\>\>DB: UPDATE products SET quantity \+= qty  
    DB--\>\>API: Success (RowsAffected)  
    API--\>\>App: 200 OK {new\_quantity}  
    App--\>\>User: Exibe Saldo Atualizado

#### **Diagrama de Classes (Estrutura Backend)**

classDiagram  
    class ProductHandler {  
        \+CreateProduct(w, r)  
        \+ListProducts(w, r)  
    }  
    class ProductService {  
        \+Create(ctx, product)  
        \+List(ctx)  
    }  
    class ProductRepository {  
        \+Insert(product)  
        \+FindAll()  
    }  
      
    ProductHandler \--\> ProductService  
    ProductService \--\> ProductRepository

## **4\. Estrutura do Projeto**

### **4.1. Backend (Go)**

Estrutura organizada em camadas (cmd, internal), isolando domínio, serviços e repositórios.

mini-wms-api/  
├── cmd/  
│   └── api/  
│       └── main.go           \# Entrypoint: Router, DI, Server Start  
├── internal/  
│   ├── core/                 \# Entidades (Structs) e Interfaces  
│   ├── handlers/             \# Controllers HTTP (Chi)  
│   ├── services/             \# Regras de Negócio  
│   └── repositories/         \# Queries SQL (SQLite)  
├── migrations/               \# Scripts SQL de criação de tabelas  
├── go.mod                    \# Gerenciador de dependências  
└── README.md

### **4.2. Frontend (Flutter)**

Estrutura **Feature-First**, agrupando arquivos por funcionalidade e não por tipo técnico, facilitando a escalabilidade.

mini-wms-app/  
├── lib/  
│   ├── main.dart             \# Setup Global (Providers, Theme)  
│   ├── core/                 \# Utilitários compartilhados (API Client, Widgets genéricos)  
│   ├── features/  
│   │   ├── products/         \# Feature: CRUD Produtos  
│   │   │   ├── presentation/ \# Screens, Widgets, Providers  
│   │   │   └── data/         \# Models específicos  
│   │   └── inventory/        \# Feature: Entradas/Saídas  
│   │       └── presentation/  
│   └── navigation/           \# Definição de Rotas (AppRouter)  
└── pubspec.yaml

## **5\. Design de Interface e Experiência (UI/UX)**

### **Guia de Estilo (Style Guide)**

* **Filosofia:** Utilitarismo e Clareza. Interface limpa focada em produtividade rápida.  
* **Componentes:** Material Design (Google) padrão.  
* **Navegação:** Stack Navigation simples com rotas nomeadas.

#### **Telas Principais**

1. **Dashboard/Lista de Produtos:**  
   * Visualização rápida de lista com cards.  
   * Exibição clara do SKU, Nome e **Saldo Atual**.  
   * FAB (Floating Action Button) para adicionar produto.  
2. **Detalhe/Edição:**  
   * Formulário limpo com validação em tempo real.  
3. **Movimentação (Modal ou Tela):**  
   * Seletor de operação (Entrada/Saída).  
   * Input numérico grande para quantidade.  
   * Feedback visual de sucesso (SnackBar).

## **6\. Plano de Desenvolvimento (Roadmap MVP 10 Dias)**

Uma abordagem ágil e intensiva para garantir a entrega dentro do prazo curto.

### **Cronograma de Sprints (Diário)**

| Dia | Foco | Atividades Chave |
| :---- | :---- | :---- |
| **Dia 1** | **Backend Setup** | Init Go Module, Configurar SQLite, Definir Structs e Migrations. |
| **Dia 2** | **Backend Core** | Implementar CRUD de Produtos (Handlers, Services, Repos). Testes unitários básicos. |
| **Dia 3** | **Frontend Setup** | Init Flutter, Estrutura de Pastas, Tela de Listagem (Mockada). |
| **Dia 4** | **Integração I** | Conectar Listagem do App com API Real. Tela de Cadastro de Produtos. |
| **Dia 5** | **Backend Logic** | Implementar lógica de Movimentação (Entrada/Saída) e atualização de saldo. |
| **Dia 6** | **Frontend Logic** | Implementar Telas de Movimentação. Integração com API de Inventário. |
| **Dia 7** | **Refinamento** | Validações de formulário, Feedback de erros (Toasts/Snackbars). |
| **Dia 8** | **Testes** | Testes manuais de fluxo completo (E2E). Correção de bugs críticos. |
| **Dia 9** | **Polimento** | Ajustes visuais, ícones, melhorias na navegação. |
| **Dia 10** | **Entrega** | Documentação final (README), Build de APK, Limpeza de código. |

## **7\. Próximos Passos e Evolução (Pós-MVP)**

Após a validação do MVP de 10 dias, o roadmap de evolução inclui:

* **Histórico Transacional:** Criar tabela de logs para auditoria de cada movimentação (Data, Qtd, Tipo).  
* **Autenticação:** Implementar JWT para suporte a múltiplos usuários/dispositivos.  
* **Dashboard Web:** Interface React para gestão administrativa em desktop.  
* **Leitor de Código de Barras:** Utilizar a câmera do celular para buscar produtos rapidamente.  
* **Dockerização:** Criar Dockerfile para deploy fácil do backend em servidores Linux.