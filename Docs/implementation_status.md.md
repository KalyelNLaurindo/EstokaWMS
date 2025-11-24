# **Implementation Status \- Estoka WMS**

**Projeto:** Estoka WMS (Mini Warehouse Management System)

**Stack Principal:** Go (Chi), SQLite, Flutter (Provider)

**Status Global:** 🟡 Planejamento / Pré-Desenvolvimento

## **📅 Fase 1: O MVP**

*Foco: Ciclo de vida essencial do produto (CRUD \+ Entrada/Saída) e funcionamento local.*

### **🛠️ Backend (Go \+ SQLite)**

1. **Setup Inicial**  
   1.1. ✅ Inicializar módulo Go (`go mod init`).  
   1.2. ✅ Configurar estrutura de pastas (`cmd`, `internal`, `configs`).  
   1.3. ✅ Configurar driver SQLite (`go-sqlite3`) e conexão com banco local.  

2. **Core Domain**  
   2.1. \[ \] Definir Struct Product (ID, Code, Name, Unit, Quantity, Timestamps).  
   2.2. \[ \] Criar script de migração SQL `001_create_products_table.sql`.  

3. **Funcionalidade: Gestão de Produtos**  
   3.1. \[ \] Implementar ProductRepository (Insert, GetByID, GetAll, Update, Delete).  
   3.2. \[ \] Implementar ProductService (Validações de negócio).  
   3.3. \[ \] Implementar ProductHandler e rotas HTTP (Chi Router).  

4. **Funcionalidade: Controle de Estoque**  
   4.1. \[ \] Implementar endpoint POST `/inventory/entries` (Incrementar saldo).  
   4.2. \[ \] Implementar endpoint POST `/inventory/exits` (Decrementar saldo e validar saldo negativo).  
   4.3. \[ \] Implementar endpoint GET `/inventory/stock/{id}`.

### **📱 Frontend (Flutter Mobile)**

1. **Setup Inicial**  
   1.1. \[ \] Inicializar projeto Flutter (`flutter create`).  
   1.2. \[ \] Configurar estrutura *Feature-First* (`features/products`, `features/inventory`).  
   1.3. \[ \] Instalar dependências (`provider`, `http`, `intl`).  

2. **Feature: Produtos**  
   2.1. \[ \] Criar ProductModel e ProductProvider (State Management).  
   2.2. \[ \] Implementar Tela de Listagem (ListView com Cards de produtos).  
   2.3. \[ \] Implementar Tela de Cadastro/Edição (Formulário com validação).  
   2.4. \[ \] Integrar http client com API de Produtos.  

3. **Feature: Movimentação**  
   3.1. \[ \] Criar modais ou telas para Registro de Entrada e Saída.  
   3.2. \[ \] Integrar atualização de saldo em tempo real na listagem.

## **🔒 Fase 2: Segurança & Auditoria (v1.1)**

*Foco: Transformar o MVP em um produto seguro e rastreável.*

### **🛠️ Backend**

1. **Autenticação & Autorização**  
   1.1. \[ \] Criar tabela `users` (ID, Username, PasswordHash, Role).  
   1.2. \[ \] Implementar endpoint POST `/login` gerando JWT Token.  
   1.3. \[ \] Criar Middleware de Auth para proteger rotas críticas.  

2. **Histórico Transacional (Logs)**  
   2.1. \[ \] Criar tabela `inventory_logs` (ID, ProductID, Type, Quantity, Date, UserID).  
   2.2. \[ \] Refatorar serviços de Entrada/Saída para gravar log na mesma transação DB.  
   2.3. \[ \] Implementar endpoint GET `/reports/history/{product_id}`.

### **📱 Frontend**

1. **Controle de Acesso**  
   1.1. \[ \] Implementar Tela de Login.  
   1.2. \[ \] Implementar persistência segura do Token (usando `flutter_secure_storage`).  
   1.3. \[ \] Adicionar Interceptor HTTP para injetar Token no header Authorization.  

2. **Visualização de Histórico**  
   2.1. \[ \] Criar tela de "Extrato do Produto" (Timeline de entradas e saídas).

## **🚀 Fase 3: Eficiência Operacional (v1.2)**

*Foco: Agilidade no chão de loja e inteligência de dados.*

### **🛠️ Backend**

1. **Inteligência de Estoque**  
   1.1. \[ \] Adicionar campos `min_stock` e `max_stock` na tabela de produtos.  
   1.2. \[ \] Criar endpoint GET `/dashboard/alerts` (Produtos abaixo do mínimo).  
   1.3. \[ \] Implementar lógica de Curva ABC (SQL Query analítica).  

2. **Performance**  
   2.1. \[ \] Adicionar índices nas colunas `product_code` e `inventory_logs.date`.

### **📱 Frontend**

1. **Leitura de Código de Barras**  
   1.1. \[ \] Integrar pacote `mobile_scanner` ou similar.  
   1.2. \[ \] Implementar botão de "Scan" na busca de produtos.  
   1.3. \[ \] Implementar "Scan-to-Action" (Ler código abre modal de entrada/saída).  

2. **Dashboard Mobile**  
   2.1. \[ \] Implementar gráficos simples (ex: `fl_chart`) mostrando volume de movimentação.  
   2.2. \[ \] Criar widgets de alerta visual (Ex: Card vermelho para estoque crítico).

## **🌐 Fase 4: Ecossistema & Escala (v2.0)**

*Foco: Gestão remota e multi-plataforma.*

### **☁️ Infraestrutura & Backend**

1. **Containerização**  
   1.1. \[ \] Criar Dockerfile multistage para a API Go (Build leve).  
   1.2. \[ \] Criar `docker-compose.yml` para orquestração local.  

2. **Migração de Banco de Dados**  
   2.1. \[ \] Adicionar suporte a PostgreSQL (via variáveis de ambiente) para deploy em nuvem.  
   2.2. \[ \] Criar scripts de migração SQLite \-\> Postgres.  

3. **API Pública**  
   3.1. \[ \] Implementar Swagger/OpenAPI (`swaggo`) para documentação automática.

### **🖥️ Frontend Web (Novo Client)**

1. **Painel Administrativo (React/Next.js)**  
   1.1. \[ \] Inicializar projeto Web para gestão desktop.  
   1.2. \[ \] Implementar tabelas de dados avançadas (Data Grids com filtros complexos).  
   1.3. \[ \] Implementar relatórios exportáveis (CSV/PDF).  
   1.4. \[ \] Criar gestão de múltiplos usuários e permissões.