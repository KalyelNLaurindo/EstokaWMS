**PLANO DE PROJETO DE SOFTWARE: MINI WMS SIMPLIFICADO (GO \+ FLUTTER)**

**Parte Externa**

* **Capa**  
  * Nome da Instituição:  
  * Título: Plano de Projeto de Software: Mini WMS Simplificado (Go \+ Flutter)  
  * Autor: \[Nome do Autor/Equipe de Planejamento\]  
  * Local: \[Cidade\]  
  * Data:

**Parte Interna**

* **Elementos Pré-textuais**  
  * **Folha de Rosto**  
    * Nome do Autor: \[Nome do Autor/Equipe de Planejamento\]  
    * Título: Plano de Projeto de Software: Mini WMS Simplificado (Go \+ Flutter)  
    * Natureza do Trabalho: Plano de Projeto de Software  
    * Objetivo: Apresentar o planejamento detalhado para o desenvolvimento de um Sistema de Gerenciamento de Armazém Simplificado (Mini WMS) utilizando Go para o backend e Flutter para o frontend, com foco em um Produto Mínimo Viável (MVP) a ser entregue em 10 dias.  
    * Nome do Orientador (se aplicável): \[Nome do Orientador\]  
    * Local: \[Cidade\]  
    * Data:  
  * **RESUMO**  
    Este plano de projeto detalha a concepção e o planejamento para o desenvolvimento de um "Mini WMS Simplificado", uma aplicação de software com backend em Go e frontend em Flutter. O objetivo principal do sistema é oferecer funcionalidades básicas de gerenciamento de armazém, como controle de produtos, registro de entradas e saídas, e consulta de estoque, visando atender às necessidades de pequenos negócios ou operações com demandas simplificadas. O presente documento estabelece a visão geral do sistema, o escopo do Produto Mínimo Viável (MVP) a ser desenvolvido em um prazo de 10 dias, as escolhas arquiteturais e de engenharia de software, as tecnologias a serem empregadas, e um cronograma macro para a execução do MVP. A metodologia de planejamento adota as melhores práticas de engenharia de software, com ênfase na viabilidade e pragmatismo para o curto prazo estipulado, sem negligenciar a qualidade e a estrutura profissional. A documentação segue as diretrizes da Associação Brasileira de Normas Técnicas (ABNT), especificamente a NBR 10719 para relatórios técnico-científicos.1 As escolhas tecnológicas incluem Go com o router Chi para a API RESTful do backend, SQLite para persistência de dados, e Flutter com o gerenciador de estado Provider para a aplicação móvel frontend. O plano visa servir como um guia fundamental para a equipe de desenvolvimento, assegurando o alinhamento e a entrega de um MVP funcional e de valor.  
    **Palavras-chave:** Mini WMS; Go; Flutter; Plano de Projeto; ABNT; MVP; SQLite; Chi; Provider.  
  * **LISTA DE TABELAS**  
    Tabela 1 – Estrutura de Diretórios Sugerida para o Backend Go.......................................... \[Número da Página\]  
    Tabela 2 – Principais Endpoints da API RESTful................................................................... \[Número da Página\]  
    Tabela 3 – Modelo de Dados Simplificado (SQLite)............................................................... \[Número da Página\]  
    Tabela 4 – Estrutura de Diretórios Sugerida para o Frontend Flutter (Feature-First)....... \[Número da Página\]  
    Tabela 5 – Cronograma Detalhado do MVP (10 Dias).......................................................... \[Número da Página\]  
  * **LISTA DE ABREVIATURAS E SIGLAS**  
    ABNT – Associação Brasileira de Normas Técnicas  
    API – Application Programming Interface (Interface de Programação de Aplicações)  
    CRUD – Create, Read, Update, Delete (Criar, Ler, Atualizar, Excluir)  
    DTO – Data Transfer Object (Objeto de Transferência de Dados)  
    HTTP – Hypertext Transfer Protocol (Protocolo de Transferência de Hipertexto)  
    JSON – JavaScript Object Notation (Notação de Objetos JavaScript)  
    JWT – JSON Web Token  
    MVP – Minimum Viable Product (Produto Mínimo Viável)  
    NBR – Norma Brasileira Registrada  
    REST – Representational State Transfer (Transferência de Estado Representacional)  
    SQL – Structured Query Language (Linguagem de Consulta Estruturada)  
    UI – User Interface (Interface do Usuário)  
    URL – Uniform Resource Locator (Localizador Uniforme de Recursos)  
    UX – User Experience (Experiência do Usuário)  
    WMS – Warehouse Management System (Sistema de Gerenciamento de Armazém)  
  * **SUMÁRIO**  
    1 INTRODUÇÃO................................................................................................................. \[Número da Página\]  
    1.1 Apresentação do Projeto "Mini WMS Simplificado".................................................... \[Número da Página\]  
    1.2 Justificativa da Necessidade e Relevância de um Plano de Projeto Detalhado......... \[Número da Página\]  
    1.3 Objetivos do Presente Plano de Projeto....................................................................... \[Número da Página\]  
    1.4 Estrutura do Documento............................................................................................... \[Número da Página\]  
    2 DESENVOLVIMENTO (Plano de Projeto do Software "Mini WMS Simplificado")..... \[Número da Página\]  
    2.1 Visão Geral do Sistema................................................................................................... \[Número da Página\]  
    2.1.1 Propósito e Funcionalidades Principais do Mini WMS.............................................. \[Número da Página\]  
    2.1.2 Escopo do MVP para o Prazo de 10 Dias................................................................ \[Número da Página\]  
    2.1.3 Público-Alvo e Casos de Uso Simplificados............................................................ \[Número da Página\]  
    2.2 Arquitetura do Software................................................................................................ \[Número da Página\]  
    2.2.1 Arquitetura Geral...................................................................................................... \[Número da Página\]  
    2.2.2 Escolhas de Design e Justificativas (Foco na Viabilidade em 10 Dias)................. \[Número da Página\]  
    2.3 Backend (Go)................................................................................................................ \[Número da Página\]  
    2.3.1 Estrutura do Projeto Go............................................................................................ \[Número da Página\]  
    2.3.2 Design da API RESTful............................................................................................. \[Número da Página\]  
    2.3.3 Escolha do Router Go............................................................................................... \[Número da Página\]  
    2.3.4 Camada de Lógica de Negócios (Serviços)............................................................ \[Número da Página\]  
    2.3.5 Persistência de Dados............................................................................................... \[Número da Página\]  
    2.3.6 Validação de Requisições........................................................................................ \[Número da Página\]  
    2.3.7 Testes....................................................................................................................... \[Número da Página\]  
    2.4 Frontend (Flutter)......................................................................................................... \[Número da Página\]  
    2.4.1 Estrutura do Projeto Flutter..................................................................................... \[Número da Página\]  
    2.4.2 Gerenciamento de Estado......................................................................................... \[Número da Página\]  
    2.4.3 Comunicação com API (Backend Go).................................................................... \[Número da Página\]  
    2.4.4 Design de UI/UX (Interface do Usuário e Experiência do Usuário)...................... \[Número da Página\]  
    2.4.5 Navegação............................................................................................................... \[Número da Página\]  
    2.4.6 Testes....................................................................................................................... \[Número da Página\]  
    2.5 Cronograma de Desenvolvimento (MVP \- 10 dias)..................................................... \[Número da Página\]  
    2.6 Versionamento e Boas Práticas de Código............................................................... \[Número da Página\]  
    3 CONSIDERAÇÕES FINAIS.............................................................................................. \[Número da Página\]  
    3.1 Reafirmação da Viabilidade do Plano para o MVP de 10 Dias................................. \[Número da Página\]  
    3.2 Próximos Passos Sugeridos Após a Aprovação do Plano........................................ \[Número da Página\]  
    3.3 Potenciais Evoluções Futuras do Mini WMS (Pós-MVP)........................................ \[Número da Página\]  
    REFERÊNCIAS................................................................................................................. \[Número da Página\]  
* **Elementos Textuais**  
  **1 INTRODUÇÃO**  
  1.1 Apresentação do Projeto "Mini WMS Simplificado"  
  O projeto "Mini WMS Simplificado" consiste no desenvolvimento de um sistema de software destinado a gerenciar operações elementares de um armazém de pequeno porte. Suas funcionalidades centrais incluem o controle de entrada e saída de produtos, bem como o monitoramento do estoque existente. Este sistema é concebido como uma solução leve e acessível, particularmente adequada para pequenos negócios, empreendedores autônomos que lidam com inventário físico, ou mesmo para uso pessoal na organização de itens, os quais não necessitam da complexidade e do custo associados a um Sistema de Gerenciamento de Armazém (WMS) completo e robusto. A proposta é oferecer uma ferramenta intuitiva e eficiente para as tarefas mais críticas do dia a dia de um pequeno estoque.  
  1.2 Justificativa da Necessidade e Relevância de um Plano de Projeto Detalhado  
  A elaboração de um plano de projeto detalhado é fundamental, mesmo para uma iniciativa descrita como "mini" e "simplificada" e com um cronograma de desenvolvimento para o Produto Mínimo Viável (MVP) estipulado em 10 dias. A existência de tal plano é crucial para alinhar expectativas entre os envolvidos, definir com clareza o escopo do que será entregue, orientar o processo de desenvolvimento de forma eficiente, otimizar a alocação de recursos (especialmente o tempo) e, fundamentalmente, garantir que o produto final entregue possua valor real para o usuário.  
  Em contextos de prazos exíguos, como o proposto, o planejamento formal assume um papel ainda mais crítico. Ele atua como um mecanismo de mitigação de riscos, prevenindo desvios de escopo – uma ocorrência comum que pode comprometer prazos curtos – e reduzindo a probabilidade de retrabalho. A ausência de um plano claro pode facilmente levar a um MVP que não é minimamente viável, que excede o prazo ou que entrega funcionalidades desalinhadas com as necessidades primordiais. Este documento, portanto, serve como um roteiro e um ponto de referência, impondo a disciplina necessária para o sucesso de um sprint de desenvolvimento tão conciso. A adesão a padrões de documentação, como os da ABNT, reforça essa disciplina, conferindo estrutura e clareza ao próprio processo de planejamento, o que, por sua vez, beneficia a qualidade do desenvolvimento subsequente.1  
  1.3 Objetivos do Presente Plano de Projeto  
  Este plano de projeto possui os seguintes objetivos específicos:  
  * Detalhar as escolhas de arquitetura de software e engenharia para os componentes de backend, a ser desenvolvido em Go, e de frontend, a ser desenvolvido em Flutter.  
  * Definir de forma precisa o escopo do Produto Mínimo Viável (MVP) que será desenvolvido no prazo de 10 dias.  
  * Estabelecer um cronograma macro realista para o desenvolvimento do referido MVP.  
  * Servir como documento de referência e consulta para a equipe de desenvolvimento, mesmo que esta seja composta por um único desenvolvedor.  
  * Assegurar a conformidade do planejamento e da documentação com as melhores práticas de engenharia de software e com os padrões da ABNT solicitados.

  1.4 Estrutura do DocumentoO presente documento está estruturado em conformidade com as recomendações da norma ABNT NBR 10719 para a apresentação de relatórios técnico-científicos.1 Após esta introdução, a seção de **Desenvolvimento** detalha o plano de projeto do software "Mini WMS Simplificado", abrangendo a visão geral do sistema, arquitetura, especificações do backend e frontend, cronograma e práticas de versionamento. Seguem-se as **Considerações Finais**, que sumarizam a viabilidade do plano e propõem próximos passos. Por fim, são apresentadas as **Referências** bibliográficas utilizadas.**2 DESENVOLVIMENTO (Plano de Projeto do Software "Mini WMS Simplificado")**Esta seção detalha o planejamento para a criação do software "Mini WMS Simplificado", abordando desde a concepção funcional até as escolhas técnicas e o cronograma de execução para o Produto Mínimo Viável (MVP).2.1 Visão Geral do SistemaA compreensão clara do propósito, funcionalidades e limitações do sistema é o primeiro passo para um planejamento eficaz.2.1.1 Propósito e Funcionalidades Principais do Mini WMSO propósito central do Mini WMS Simplificado é fornecer uma solução de software intuitiva e leve para o gerenciamento básico de inventário. Destina-se a simplificar o controle de estoque para operações que não justificam a adoção de sistemas WMS complexos.As funcionalidades essenciais do sistema incluem:

  * **Gerenciamento de Produtos:** Permitir o Cadastro (Create), Leitura (Read), Atualização (Update) e Exclusão (Delete) – operações CRUD – dos itens que compõem o inventário. Cada produto será caracterizado por atributos como Nome, Descrição (opcional), um Código Único para identificação (SKU \- *Stock Keeping Unit*), e sua Unidade de Medida (ex: unidade, caixa, kg).  
  * **Registro de Entradas:** Facilitar o registro da chegada de novos produtos ou de mais unidades de produtos existentes ao estoque. Este registro deverá conter, no mínimo, a identificação do Produto, a Quantidade que está entrando e a Data da entrada. Um campo para Fornecedor poderá ser considerado, mas como opcional para o MVP.  
  * **Registro de Saídas:** Permitir o registro da retirada de produtos do estoque, seja por venda, uso interno ou outra finalidade. Este registro deverá conter, no mínimo, a identificação do Produto, a Quantidade que está saindo e a Data da saída. Um campo para Cliente ou Destino poderá ser considerado, mas como opcional para o MVP.  
  * **Consulta de Estoque:** Oferecer uma forma simples e rápida de visualizar a quantidade atual de cada produto disponível em estoque.

  Uma funcionalidade de autenticação de usuário, se considerada viável dentro do prazo de 10 dias, seria implementada de forma extremamente simplificada, possivelmente com um único tipo de usuário e credenciais fixas ou um mecanismo de login muito básico, sem gerenciamento complexo de papéis ou permissões.2.1.2 Escopo do MVP para o Prazo de 10 DiasDado o prazo rigoroso de 10 dias para a entrega do Produto Mínimo Viável, a definição do que é "mínimo" torna-se crucial. Qualquer funcionalidade que não seja absolutamente central para o propósito fundamental de um WMS – gerenciar o que entra, o que sai e o que há em estoque – deve ser postergada. O escopo do MVP será:

  * **Núcleo Indispensável:** As operações CRUD para Gerenciamento de Produtos e a funcionalidade de Consulta de Estoque são prioritárias e formam a base do MVP.  
  * **Registro de Entradas e Saídas Simplificado:** As funcionalidades de Registro de Entradas e Saídas serão implementadas de maneira a atualizar diretamente o saldo do produto em estoque. Para este MVP, não será mantido um log detalhado e separado de cada transação de movimentação. Esta simplificação reduz significativamente a complexidade do modelo de dados e da lógica de negócios, tornando a implementação mais rápida. Em vez de uma tabela de "movimentações", existirá um campo quantidade\_em\_estoque na tabela de produtos, que será incrementado ou decrementado diretamente.  
  * **Autenticação (Se Viável):** Se a autenticação for incluída, será implementada de forma rudimentar (ex: um único usuário com credenciais fixas no código ou um sistema de login muito simples). A decisão final sobre sua inclusão dependerá do progresso nas funcionalidades centrais. Caso contrário, será adiada para uma fase pós-MVP.  
  * **Interface do Usuário (UI):** A UI da aplicação Flutter será funcional e intuitiva, mas sem foco em design elaborado ou customizações visuais complexas. A prioridade será a usabilidade básica para realizar as operações definidas.  
  * **Exclusões do MVP:** Não serão incluídos no MVP: relatórios complexos, histórico detalhado de transações (além da atualização do saldo corrente), gerenciamento de múltiplos usuários com diferentes perfis, integrações com outros sistemas, alertas de estoque, ou funcionalidades avançadas de WMS.

  2.1.3 Público-Alvo e Casos de Uso SimplificadosO público-alvo primário do Mini WMS Simplificado inclui:

  * Pequenos comerciantes (lojas de bairro, pequenos e-commerces).  
  * Profissionais autônomos que gerenciam estoque de materiais ou produtos para seu trabalho.  
  * Uso pessoal para organização de coleções, ferramentas, ou outros itens.

  Os casos de uso simplificados que o MVP deve atender são:

  * "Como usuário, quero cadastrar um novo produto (com código, nome, unidade) para que ele possa ser gerenciado pelo sistema."  
  * "Como usuário, quero visualizar uma lista de todos os produtos cadastrados com suas quantidades atuais em estoque."  
  * "Como usuário, quero poder editar as informações de um produto existente."  
  * "Como usuário, quero poder excluir um produto que não será mais gerenciado."  
  * "Como usuário, quero registrar a entrada de X unidades do produto Y para que o saldo em estoque seja atualizado."  
  * "Como usuário, quero registrar a saída de Z unidades do produto W para que o saldo em estoque seja atualizado."  
  * "Como usuário, quero consultar o saldo atual de um produto específico para saber quantas unidades tenho disponíveis."

  2.2 Arquitetura do SoftwareA arquitetura define a estrutura fundamental do sistema, e sua escolha deve ser pragmática, especialmente considerando o prazo de desenvolvimento.2.2.1 Arquitetura GeralO Mini WMS Simplificado adotará uma arquitetura Cliente-Servidor:

  * **Frontend (Cliente):** Uma aplicação móvel desenvolvida com Flutter. Será responsável por toda a interface com o usuário (UI), captura de dados de entrada e apresentação das informações obtidas do backend.  
  * **Backend (Servidor):** Uma API (Interface de Programação de Aplicações) RESTful desenvolvida em Go. Será responsável por toda a lógica de negócios, validação de dados, processamento das requisições do frontend e comunicação com o banco de dados para persistência.  
  * **Comunicação:** A interação entre o frontend Flutter e o backend Go ocorrerá exclusivamente através de requisições HTTP, com dados transacionados no formato JSON.

  Diagrama de Blocos da Arquitetura Geral (Sugestão para Apêndice A):Um diagrama simples mostrando um bloco "Aplicação Flutter (Cliente)" comunicando-se via HTTP/JSON com um bloco "API Go (Servidor)", que por sua vez interage com um bloco "Banco de Dados SQLite".2.2.2 Escolhas de Design e Justificativas (Foco na Viabilidade em 10 Dias)A "melhor" arquitetura é aquela que atende aos requisitos do projeto, incluindo o prazo, de forma eficaz e eficiente. Para este MVP, a simplicidade e a velocidade de desenvolvimento são prioritárias em relação a arquiteturas excessivamente "puras" ou projetadas para escalabilidade massiva, que seriam inviáveis no tempo disponível.

  * **Monólito Simplificado para o Backend (Go):** Dada a natureza "mini" do sistema e o prazo exíguo de 10 dias, uma arquitetura monolítica para a API Go é a abordagem mais pragmática e eficiente.2 A lógica da API será contida em uma única aplicação Go. Deve-se evitar a complexidade inerente a arquiteturas de microserviços, que introduzem desafios adicionais de deploy, comunicação inter-serviços e gerenciamento, incompatíveis com o cronograma do MVP. Um monólito bem organizado, com clara separação de responsabilidades internas (handlers, serviços, repositórios), será suficiente e mais rápido de desenvolver.  
  * **Backend API RESTful Stateless:** A API Go será projetada para ser *stateless*, ou seja, não manterá o estado da sessão do cliente entre requisições. Cada requisição do cliente para o servidor deverá conter todas as informações necessárias para ser processada, e o servidor não dependerá de interações anteriores. Se a autenticação for implementada (mesmo que de forma básica), será baseada em tokens (por exemplo, um JWT simples). Esta abordagem facilita a escalabilidade horizontal futura (embora não seja um foco primário do MVP) e, mais importante para o prazo, simplifica o desenvolvimento, o teste e o raciocínio sobre o comportamento da API.  
  * **Frontend Flutter com Gerenciamento de Estado Localizado/Simples:** A aplicação Flutter focará em gerenciar o estado necessário para cada tela ou funcionalidade de forma isolada ou utilizando um gerenciador de estado simples e de rápida implementação, como o Provider. O objetivo é evitar soluções de gerenciamento de estado global excessivamente complexas (como BLoC com muita cerimônia para este escopo) que demandariam mais tempo de configuração, aprendizado e implementação, o que seria contraproducente para um MVP rápido.  
  * **Persistência de Dados Leve (SQLite):** O banco de dados escolhido para o sistema será o SQLite.3 Ele será embutido diretamente na aplicação Go, eliminando a necessidade de um servidor de banco de dados separado. Esta escolha simplifica drasticamente a configuração do ambiente de desenvolvimento, o processo de desenvolvimento em si e o futuro deploy da aplicação, especialmente para um "Mini WMS". SQLite é ideal para aplicações de pequeno a médio porte, cenários com um único escritor ou concorrência de escrita baixa, e onde a portabilidade e a facilidade de setup são cruciais.

  A combinação de um monólito Go para o backend, uma aplicação Flutter com gerenciamento de estado simples para o frontend, e SQLite para a persistência de dados, representa a arquitetura mais direta e rápida para entregar o valor funcional essencial do Mini WMS dentro do prazo de 10 dias.2.3 Backend (Go)O backend será o cérebro do Mini WMS, responsável por processar todas as operações e gerenciar os dados.2.3.1 Estrutura do Projeto GoA estrutura de pastas do projeto Go será organizada para promover clareza e separação de responsabilidades, seguindo práticas comuns na comunidade Go, mas adaptada para a simplicidade exigida pelo MVP.2 A filosofia é começar simples e permitir que a estrutura evolua se necessário, evitando a criação de diretórios apenas por uma questão de organização visual sem um propósito de pacotes distinto.2A estrutura de diretórios principal será:

  * **cmd/mini-wms-api/main.go**: Este será o ponto de entrada da aplicação backend. Ele será responsável por inicializar o servidor HTTP, configurar as rotas da API (utilizando o router Chi), injetar dependências (como instâncias de serviços e repositórios) e iniciar o listener do servidor.  
  * **internal/**: Este diretório conterá toda a lógica principal da aplicação que não se destina a ser importada por outros projetos externos. É uma convenção do Go para código específico do projeto.  
    * **internal/core/** (ou alternativamente internal/domain/): Aqui serão definidas as structs que representam as entidades de domínio do sistema (ex: Produto, EntradaDeEstoque, SaidaDeEstoque), bem como as interfaces para os serviços (lógica de negócio) e repositórios (acesso a dados).  
    * **internal/services/**: Conterá as implementações concretas da lógica de negócio. Por exemplo, product\_service.go implementaria as operações de CRUD para produtos, validando regras de negócio e coordenando com o repositório.  
    * **internal/repositories/**: Conterá as implementações para o acesso ao banco de dados SQLite. Por exemplo, product\_repository\_sqlite.go implementaria as queries SQL para interagir com a tabela de produtos.  
    * **internal/handlers/** (ou alternativamente internal/api/http/): Agrupará os manipuladores de requisições HTTP. Estes handlers serão responsáveis por receber as requisições da API, decodificar os dados de entrada (JSON), chamar os métodos apropriados nos serviços e formatar as respostas (JSON) a serem enviadas de volta ao cliente. Exemplo: product\_handler.go.  
  * **pkg/**: Este diretório é opcional para o MVP. Se houver utilitários genéricos que poderiam, teoricamente, ser compartilhados com outros projetos (embora improvável neste contexto), eles poderiam residir aqui.4 Para o MVP, se não houver tal necessidade, pode ser omitido para manter a simplicidade. Exemplos poderiam ser pkg/jsonutil para helpers de resposta JSON ou pkg/config para leitura de configurações.  
  * **migrations/** (Opcional, mas recomendado): Conterá scripts SQL para criar e atualizar o schema do banco de dados SQLite. Isso facilita o versionamento do schema e a configuração inicial do banco.  
  * **go.mod, go.sum**: Arquivos padrão do Go para gerenciamento de módulos e suas dependências.  
  * **README.md**: Um arquivo com instruções básicas sobre como compilar, executar e testar o backend.  
  * **Makefile** (Opcional): Pode ser usado para simplificar comandos comuns de build, execução e teste.

Esta estrutura, especialmente com o uso de cmd/ e internal/, é um padrão bem estabelecido que oferece uma boa separação de preocupações sem introduzir complexidade desnecessária para um projeto do tamanho e prazo do Mini WMS

**Tabela 1 – Estrutura de Diretórios inicial para o Backend Go**

| Caminho | Descrição |
| :---- | :---- |
| cmd/mini-wms-api/main.go | Ponto de entrada da aplicação, inicialização do servidor. |
| internal/core/ | Modelos de domínio (structs), interfaces de serviço/repositório. |
| internal/services/ | Implementação da lógica de negócios. |
| internal/repositories/ | Implementação do acesso a dados (interação com SQLite). |
| internal/handlers/ | Manipuladores de requisições HTTP, interação com serviços. |
| pkg/utils/ (Opcional) | Funções utilitárias compartilhadas (ex: formatação de resposta JSON, erros HTTP). |
| migrations/ (Recomendado) | Scripts SQL para criação/atualização do schema do banco de dados SQLite. |
| go.mod, go.sum | Arquivos de módulos Go para gerenciamento de dependências. |
| README.md | Documentação básica do backend (setup, execução, API). |

2.3.2 Design da API RESTful

A API será projetada seguindo os princípios REST (\*Representational State Transfer\*) para garantir uma interface padronizada, previsível e fácil de consumir pelo frontend Flutter.\[5\]

\*   \*\*Uso Correto de Métodos HTTP:\*\*  
    \*   \`GET\`: Para recuperar recursos (listar produtos, obter um produto específico, consultar estoque).  
    \*   \`POST\`: Para criar novos recursos (cadastrar um novo produto, registrar uma entrada ou saída de estoque).  
    \*   \`PUT\`: Para atualizar integralmente um recurso existente (modificar os dados de um produto).  
    \*   \`DELETE\`: Para remover um recurso (excluir um produto).  
\*   \*\*Endpoints Baseados em Substantivos no Plural:\*\* Os caminhos da API representarão os recursos e devem usar substantivos no plural. Por exemplo, \`/products\` para o recurso "produto", \`/inventory/entries\` para o recurso "entrada de inventário".  
\*   \*\*Respostas em JSON:\*\* Todas as respostas da API que contêm dados serão formatadas em JSON. O cabeçalho \`Content-Type\` da resposta será \`application/json\`.  
\*   \*\*Uso de Códigos de Status HTTP Padrão:\*\* A API utilizará códigos de status HTTP padrão para indicar o resultado das requisições:  
    \*   \`200 OK\`: Requisição bem-sucedida (geralmente para GET, PUT).  
    \*   \`201 Created\`: Recurso criado com sucesso (geralmente para POST).  
    \*   \`204 No Content\`: Requisição bem-sucedida, mas sem corpo de resposta (geralmente para DELETE).  
    \*   \`400 Bad Request\`: A requisição do cliente é inválida (ex: dados faltando, formato incorreto).  
    \*   \`401 Unauthorized\`: O cliente não está autenticado (se autenticação for implementada).  
    \*   \`403 Forbidden\`: O cliente está autenticado, mas não tem permissão para acessar o recurso.  
    \*   \`404 Not Found\`: O recurso solicitado não foi encontrado.  
    \*   \`500 Internal Server Error\`: Um erro genérico ocorreu no servidor.  
\*   \*\*Versionamento Simplificado:\*\* Para o MVP, não será implementado um versionamento explícito na URL da API (ex: \`/api/v1/products\`). A API desenvolvida será considerada a primeira versão. Futuras alterações incompatíveis poderiam introduzir versionamento.  
\*   \*\*Tratamento de Erros:\*\* As respostas de erro (códigos 4xx e 5xx) incluirão um corpo JSON com uma mensagem clara descrevendo o erro, para auxiliar na depuração pelo cliente ou desenvolvedor.

\*\*Endpoints Principais (Exemplos):\*\*  
A seguir, uma lista dos principais endpoints planejados para o MVP:

\*   \*\*Produtos:\*\*  
    \*   \`POST /products\`: Cria um novo produto.  
    \*   \`GET /products\`: Lista todos os produtos cadastrados.  
    \*   \`GET /products/{id}\`: Obtém os detalhes de um produto específico pelo seu ID.  
    \*   \`PUT /products/{id}\`: Atualiza os dados de um produto específico.  
    \*   \`DELETE /products/{id}\`: Remove um produto específico.  
\*   \*\*Inventário:\*\*  
    \*   \`POST /inventory/entries\`: Registra uma nova entrada de produtos no estoque.  
    \*   \`POST /inventory/exits\`: Registra uma nova saída de produtos do estoque.  
    \*   \`GET /inventory/products/{productId}/stock\`: Consulta a quantidade atual em estoque de um produto específico.

\*\*Tabela 2 – Principais Endpoints da API RESTful\*\*

| Método | Rota | Descrição | Exemplo de Payload (Requisição/Resposta Simplificada) |
| :---- | :---- | :---- | :---- |
| POST | /products | Cria um novo produto | Req: {"code": "SKU001", "name": "Caneta Azul", "unit": "UN", "quantity": 100}\<br\>Res: {Product Object} |
| GET | /products | Lista todos os produtos | Res: \[{Product Object 1}, {Product Object 2},...\] |
| GET | /products/{id} | Obtém detalhes de um produto por ID | Res: {Product Object} |
| PUT | /products/{id} | Atualiza um produto existente por ID | Req: {"name": "Caneta Azul BIC", "description": "Esferográfica"}\<br\>Res: {Product Object} |
| DELETE | /products/{id} | Deleta um produto por ID | Res: (Status 204 No Content) |
| POST | /inventory/entries | Registra uma entrada de produto no estoque | Req: {"product\_id": 1, "quantity": 50, "entry\_date": "2024-07-29T10:00:00Z"}\<br\>Res: {Updated Stock Status} |
| POST | /inventory/exits | Registra uma saída de produto do estoque | Req: {"product\_id": 1, "quantity": 10, "exit\_date": "2024-07-29T11:00:00Z"}\<br\>Res: {Updated Stock Status} |
| GET | /inventory/products/{productId}/stock | Consulta o estoque de um produto específico | Res: {"product\_id": 1, "current\_quantity": 140} |

2.3.3 Escolha do Router Go

Para o roteamento das requisições HTTP na API Go, a escolha recai sobre a biblioteca \*\*Chi (\`go-chi/chi\`)\*\*.

\*Justificativa:\* Chi é um router leve, idiomático, rápido e altamente compatível com a biblioteca padrão \`net/http\` do Go.\[6\] Ele oferece um bom suporte a middlewares (como logging, recuperação de panics, CORS), permite a definição de rotas de forma clara e concisa, e possui uma curva de aprendizado suave. Para um MVP a ser desenvolvido em 10 dias, sua simplicidade e eficiência são ideais, evitando a sobrecarga e a complexidade que frameworks mais "opinativos" ou com mais funcionalidades embutidas (como Gin \[7\]) poderiam trazer. Embora Gin seja uma opção popular e performática, a leveza de Chi e sua aderência aos padrões \`net/http\` minimizam o tempo gasto em aprendizado de idiossincrasias do framework, permitindo maior foco na lógica de negócios do WMS.\[6\] A flexibilidade de Chi para compor middlewares e sua performance são adequadas para as necessidades do Mini WMS.

Um exemplo básico de configuração de rotas com Chi seria:  
\`\`\`go  
// import (  
//     "net/http"  
//     "github.com/go-chi/chi/v5"  
//     "github.com/go-chi/chi/v5/middleware"  
//     "yourproject/internal/handlers"  
// )

// func main() {  
//     r := chi.NewRouter()

//     // Middlewares básicos  
//     r.Use(middleware.Logger)        // Loga cada requisição  
//     r.Use(middleware.Recoverer)     // Recupera de panics e retorna 500  
//     r.Use(middleware.RealIP)        // Obtém o IP real do cliente  
//     r.Use(middleware.RequestID)     // Adiciona um ID único a cada requisição

//     // Instanciar handlers (exemplo)  
//     // productHandler := handlers.NewProductHandler(productService)

//     // Rotas para produtos  
//     r.Route("/products", func(r chi.Router) {  
//         r.Post("/", productHandler.CreateProduct)  
//         r.Get("/", productHandler.ListProducts)  
//         r.Get("/{id}", productHandler.GetProductByID)  
//         r.Put("/{id}", productHandler.UpdateProduct)  
//         r.Delete("/{id}", productHandler.DeleteProduct)  
//     })

//     // Rotas para inventário  
//     r.Route("/inventory", func(r chi.Router) {  
//         r.Post("/entries", inventoryHandler.RegisterEntry)  
//         r.Post("/exits", inventoryHandler.RegisterExit)  
//         r.Get("/products/{productId}/stock", inventoryHandler.GetProductStock)  
//     })

//     http.ListenAndServe(":3000", r)  
// }  
\`\`\`

2.3.4 Camada de Lógica de Negócios (Serviços)

A camada de serviços (\`internal/services/\`) será responsável por encapsular a lógica de negócios principal do Mini WMS. Esta camada atuará como um intermediário entre os handlers HTTP e os repositórios de dados.

\*   \*\*Responsabilidades:\*\*  
    \*   Orquestrar as operações de negócio (ex: ao criar um produto, pode haver validações de regras de negócio específicas antes de persistir).  
    \*   Implementar validações de regras de negócio que vão além da simples validação de formato de dados (esta última é feita nos handlers ou com bibliotecas de validação).  
    \*   Coordenar interações com um ou mais repositórios, se necessário (embora para o MVP, provavelmente cada serviço interaja com um repositório correspondente).  
    \*   Manter os handlers (camada HTTP) "enxutos", focados apenas em lidar com a requisição e resposta HTTP, e delegando a lógica de negócio para os serviços.  
    \*   Manter os repositórios focados exclusivamente nas interações com o banco de dados (queries SQL).

\*   \*\*Estrutura:\*\*  
    \*   Serão definidas interfaces para cada serviço (ex: \`ProductService\`, \`InventoryService\`) no pacote \`internal/core/\`.  
    \*   As implementações concretas desses serviços residirão em \`internal/services/\`.  
    \*   Exemplo de interface e método:  
        \`\`\`go  
        // Em internal/core/product\_service.go  
        // package core  
        // type ProductService interface {  
        //     CreateProduct(ctx context.Context, product \*Product) (\*Product, error)  
        //     GetProductByID(ctx context.Context, id int) (\*Product, error)  
        //     ListProducts(ctx context.Context) (\*Product, error)  
        //     UpdateProduct(ctx context.Context, id int, productUpdateData \*ProductUpdate) (\*Product, error)  
        //     DeleteProduct(ctx context.Context, id int) error  
        // }

        // Em internal/services/product\_service\_impl.go  
        // package services  
        // type productServiceImpl struct {  
        //     repo core.ProductRepository  
        // }  
        // func NewProductService(repo core.ProductRepository) core.ProductService {  
        //     return \&productServiceImpl{repo: repo}  
        // }  
        // func (s \*productServiceImpl) CreateProduct(ctx context.Context, product \*core.Product) (\*core.Product, error) {  
        //     // Lógica de negócio aqui (ex: validações, transformações)  
        //     return s.repo.Create(ctx, product)  
        // }  
        \`\`\`  
Esta separação de preocupações é uma prática recomendada que, mesmo em um projeto pequeno, contribui para a testabilidade e manutenibilidade do código.

2.3.5 Persistência de Dados

A persistência dos dados do Mini WMS será realizada utilizando \*\*SQLite\*\*.

\*Justificativa:\* Conforme mencionado na seção de arquitetura, SQLite é uma escolha pragmática e eficiente para o MVP.\[3\] Suas principais vantagens neste contexto são:  
\*   \*\*Leveza e Simplicidade:\*\* Não requer um processo de servidor de banco de dados separado, sendo um arquivo no sistema de arquivos. Isso simplifica enormemente a configuração do ambiente de desenvolvimento e o deploy.  
\*   \*\*Integração com Go:\*\* A biblioteca padrão \`database/sql\` do Go, em conjunto com um driver SQLite como \`github.com/mattn/go-sqlite3\`, oferece uma interface robusta e familiar para interações com o banco de dados.\[3, 8\]  
\*   \*\*Adequação ao Escopo:\*\* Para um "Mini WMS Simplificado" com um único usuário (ou baixa concorrência no futuro), as capacidades do SQLite são mais do que suficientes.

\*   \*\*Implementação:\*\*  
    \*   \*\*Criação do Banco e Tabelas:\*\* O banco de dados (ex: \`miniwms.db\`) será criado na primeira execução da aplicação, se não existir. As tabelas serão criadas utilizando instruções SQL \`CREATE TABLE\`. É recomendável usar um sistema de migrações simples (mesmo que scripts SQL versionados manualmente na pasta \`migrations/\`) para gerenciar a evolução do schema.  
    \*   \*\*Operações CRUD:\*\*  
        \*   \`db.Exec()\`: Será utilizado para instruções SQL que não retornam linhas, como \`INSERT\`, \`UPDATE\`, \`DELETE\`, e \`CREATE TABLE\`.\[3\]  
        \*   \`db.Query()\` e \`db.QueryRow()\`: Serão utilizados para instruções \`SELECT\`. Os resultados serão mapeados para as structs de domínio usando \`rows.Scan()\`.\[3\]  
        \*   Transações (\`db.Begin()\`, \`tx.Commit()\`, \`tx.Rollback()\`): Serão usadas para operações que exigem atomicidade (ex: atualizar o estoque de um produto e registrar a movimentação, se um log fosse implementado).

\*   \*\*Modelo de Dados Simplificado (Foco no MVP):\*\*  
    Para o MVP, o modelo de dados será minimalista, focando em atualizar diretamente a quantidade em estoque na tabela de produtos, em vez de manter um log detalhado de transações de entrada e saída separadamente.

    \*\*Tabela 3 – Modelo de Dados Simplificado (SQLite)\*\*

| Tabela | Coluna | Tipo | Restrições | Descrição |
| :---- | :---- | :---- | :---- | :---- |
| products | id | INTEGER | PRIMARY KEY AUTOINCREMENT | Identificador único do produto |
|  | code | TEXT | UNIQUE NOT NULL | Código único do produto (SKU) |
|  | name | TEXT | NOT NULL | Nome do produto |
|  | description | TEXT |  | Descrição opcional do produto |
|  | unit | TEXT | NOT NULL | Unidade de medida (ex: UN, KG, L, M) |
|  | quantity | INTEGER | NOT NULL DEFAULT 0 | Quantidade atual em estoque |
|  | created\_at | DATETIME | DEFAULT CURRENT\_TIMESTAMP | Data e hora de criação do registro do produto |
|  | updated\_at | DATETIME | DEFAULT CURRENT\_TIMESTAMP | Data e hora da última atualização do produto |

    \*Notas sobre o modelo:\*  
    \*   A coluna \`quantity\` na tabela \`products\` será diretamente incrementada ou decrementada pelas operações de entrada e saída.  
    \*   As colunas \`created\_at\` e \`updated\_at\` são úteis para rastreamento e podem ser gerenciadas por triggers no SQLite ou pela aplicação Go. Para \`updated\_at\`, um trigger como \`CREATE TRIGGER update\_products\_updated\_at AFTER UPDATE ON products FOR EACH ROW BEGIN UPDATE products SET updated\_at \= CURRENT\_TIMESTAMP WHERE id \= OLD.id; END;\` pode ser útil.

Esta abordagem de persistência é robusta o suficiente para o MVP e alinhada com o objetivo de desenvolvimento rápido.

2.3.6 Validação de Requisições

A validação dos dados de entrada recebidos pela API é crucial para garantir a integridade dos dados e a segurança do sistema. Para isso, será utilizada a biblioteca \*\*\`go-playground/validator/v10\`\*\*.

\*Justificativa:\* \`go-playground/validator/v10\` é uma biblioteca popular, eficiente e flexível para validação de structs em Go.\[9, 10\] Ela permite definir regras de validação diretamente nas tags das structs que representam os payloads da API (ex: dados de um novo produto). Isso torna o código de validação declarativo, conciso e fácil de manter, reduzindo significativamente o código boilerplate que seria necessário para validações manuais em cada handler.\[9\] Comparada a outras opções, ela oferece um bom equilíbrio entre funcionalidades e performance para as necessidades do projeto.\[11, 12\]

\*   \*\*Implementação:\*\*  
    \*   \*\*Definição de Structs de Requisição com Tags:\*\* Serão criadas structs específicas para os dados esperados no corpo das requisições \`POST\` e \`PUT\`. Essas structs utilizarão tags \`validate\` para especificar as regras.  
        \`\`\`go  
        // Exemplo para a criação de um produto  
        // type CreateProductRequest struct {  
        //     Code        string \`json:"code"        validate:"required,alphanum,min=1,max=50"\`  
        //     Name        string \`json:"name"        validate:"required,min=3,max=100"\`  
        //     Description string \`json:"description" validate:"omitempty,max=255"\` // omitempty significa opcional  
        //     Unit        string \`json:"unit"        validate:"required,min=1,max=10"\`  
        //     Quantity    int    \`json:"quantity"    validate:"gte=0"\` // gte=0 significa maior ou igual a zero  
        // }  
        \`\`\`  
    \*   \*\*Uso nos Handlers:\*\*  
        1\.  Decodificar o corpo da requisição JSON para a struct de requisição.  
        2\.  Criar uma instância do validador: \`validate := validator.New()\`.  
        3\.  Chamar \`validate.Struct(requestStruct)\` para executar a validação.\[10\]  
        4\.  Se \`validate.Struct()\` retornar um erro, este erro (do tipo \`validator.ValidationErrors\`) pode ser iterado para fornecer mensagens de erro detalhadas ao cliente (ex: quais campos falharam e por quê).  
        \`\`\`go  
        // Exemplo simplificado no handler  
        // var req CreateProductRequest  
        // if err := json.NewDecoder(r.Body).Decode(\&req); err\!= nil {  
        //     http.Error(w, "Invalid request body", http.StatusBadRequest)  
        //     return  
        // }

        // validate := validator.New()  
        // if err := validate.Struct(req); err\!= nil {  
        //     // Tratar erros de validação (ex: retornar 400 com detalhes)  
        //     validationErrors := err.(validator.ValidationErrors)  
        //     // Construir resposta de erro baseada em validationErrors  
        //     http.Error(w, buildValidationErrorMessage(validationErrors), http.StatusBadRequest)  
        //     return  
        // }  
        // // Se a validação passar, prosseguir com a lógica de serviço  
        \`\`\`  
    \*   \*\*Validadores Customizados (Opcional para MVP):\*\* A biblioteca também permite registrar validadores customizados para regras mais complexas que não são cobertas pelas tags padrão \[9\], mas isso provavelmente estará fora do escopo do MVP de 10 dias.

A utilização desta biblioteca acelera o desenvolvimento ao automatizar uma parte significativa do processo de validação de entrada.

2.3.7 Testes

A estratégia de testes para o backend Go focará em garantir a corretude da lógica de negócios e a funcionalidade dos endpoints da API, priorizando o que é factível no prazo de 10 dias.

\*   \*\*Testes Unitários:\*\*  
    \*   \*\*Foco:\*\* Testar as menores unidades de código de forma isolada, principalmente funções e métodos dentro dos pacotes de serviço (\`internal/services/\`) e repositório (\`internal/repositories/\`).  
    \*   \*\*Ferramentas:\*\*  
        \*   Pacote \`testing\` da biblioteca padrão do Go: Para a estrutura básica dos testes (arquivos \`\_test.go\`, funções \`TestXxx\`).  
        \*   \`testify/assert\`: Para asserções mais ricas e legíveis (ex: \`assert.NoError(t, err)\`, \`assert.Equal(t, expected, actual)\`).\[13\]  
        \*   \`testify/mock\`: Para criar mocks (dublês de teste) de dependências, especialmente interfaces. Por exemplo, ao testar um serviço, a interface do repositório correspondente será mockada para simular o comportamento do banco de dados sem realmente acessá-lo.\[13, 14\]  
    \*   \*\*Abordagem:\*\*  
        \*   \*\*Table-Driven Tests:\*\* Esta abordagem será utilizada sempre que possível para testar uma função ou método com múltiplos conjuntos de entradas e saídas esperadas de forma organizada e concisa.\[15, 16\] Cada caso de teste é uma struct dentro de um slice, contendo nome, entradas e resultado esperado.  
            \`\`\`go  
            // Exemplo de Table-Driven Test para um serviço (simplificado)  
            // func TestProductService\_CreateProduct(t \*testing.T) {  
            //     mockRepo := new(mocks.ProductRepository) // Mock gerado com testify/mock ou mockery  
            //     service := services.NewProductService(mockRepo)

            //     tests :=struct {  
            //         name          string  
            //         inputProduct  \*core.Product  
            //         mockRepoSetup func() // Configura o comportamento do mockRepo  
            //         expectedError bool  
            //     }{  
            //         // Casos de teste aqui  
            //     }  
            //     for \_, tt := range tests {  
            //         t.Run(tt.name, func(t \*testing.T) {  
            //             tt.mockRepoSetup()  
            //             \_, err := service.CreateProduct(context.Background(), tt.inputProduct)  
            //             if tt.expectedError {  
            //                 assert.Error(t, err)  
            //             } else {  
            //                 assert.NoError(t, err)  
            //             }  
            //             mockRepo.AssertExpectations(t) // Verifica se os métodos mockados foram chamados como esperado  
            //         })  
            //     }  
            // }  
            \`\`\`

\*   \*\*Testes de Integração (Nível API):\*\*  
    \*   \*\*Foco:\*\* Testar a integração entre os handlers HTTP, serviços e (potencialmente) uma instância real do banco de dados de teste (SQLite em memória ou arquivo). O objetivo é verificar se os endpoints da API se comportam conforme o esperado quando recebem requisições HTTP.  
    \*   \*\*Ferramentas:\*\*  
        \*   Pacote \`net/http/httptest\` da biblioteca padrão: Para criar um servidor de teste que executa os handlers da aplicação e para construir requisições HTTP simuladas (\`httptest.NewRequest\`, \`httptest.NewRecorder\`).\[17, 18\]  
        \*   \`testify/assert\`: Para verificar os códigos de status da resposta, cabeçalhos e o corpo da resposta JSON.  
    \*   \*\*Abordagem:\*\*  
        \*   Para cada endpoint, serão criados testes que enviam requisições HTTP com diferentes payloads (válidos e inválidos) e verificam se a resposta (código de status, corpo) está correta.  
        \*   O banco de dados pode ser configurado para cada conjunto de testes (ex: populado com dados iniciais ou limpo).  
            \`\`\`go  
            // Exemplo de Teste de Integração de API (simplificado)  
            // func TestProductHandlers\_CreateProductAPI(t \*testing.T) {  
            //     // Setup: router, dependências (serviço com repo real ou mockado para DB)  
            //     router := chi.NewRouter()  
            //     //... configurar productHandler e rotas...  
            //     // router.Post("/products", productHandler.CreateProduct)

            //     productJSON := \`{"code": "TEST01", "name": "Test Product", "unit": "UN", "quantity": 10}\`  
            //     req := httptest.NewRequest("POST", "/products", strings.NewReader(productJSON))  
            //     req.Header.Set("Content-Type", "application/json")  
            //     rr := httptest.NewRecorder()

            //     router.ServeHTTP(rr, req)

            //     assert.Equal(t, http.StatusCreated, rr.Code)  
            //     // Verificar o corpo da resposta JSON para o produto criado  
            //     var createdProduct core.Product  
            //     err := json.Unmarshal(rr.Body.Bytes(), \&createdProduct)  
            //     assert.NoError(t, err)  
            //     assert.Equal(t, "Test Product", createdProduct.Name)  
            // }  
            \`\`\`

Para o MVP de 10 dias, o foco será em testes unitários robustos para a lógica de negócios e testes de integração para os principais fluxos da API. Testes \*end-to-end\* (E2E) completos, envolvendo a interação com o frontend Flutter, são desejáveis, mas provavelmente muito complexos e demorados para serem implementados extensivamente neste prazo. A prioridade é garantir que o backend funcione corretamente de forma isolada.

2.4 Frontend (Flutter)

O frontend será a interface com a qual o usuário interagirá para gerenciar o Mini WMS, desenvolvido utilizando Flutter para garantir uma experiência móvel nativa e fluida.

2.4.1 Estrutura do Projeto Flutter

Para a organização do código Flutter, será adotada a abordagem \*\*Feature-First (Camadas Dentro das Funcionalidades)\*\*.

\*Justificativa:\* Mesmo para um projeto de escopo reduzido como o Mini WMS, a estrutura Feature-First promove uma melhor organização, manutenibilidade e escalabilidade futura em comparação com a abordagem Layer-First.\[19, 20\] Nesta abordagem, todos os arquivos relacionados a uma funcionalidade específica (como UI, lógica de estado, modelos de dados específicos da feature) são agrupados em um único diretório de alto nível. Isso facilita a localização de código, reduz o acoplamento entre features e simplifica a adição, modificação ou remoção de funcionalidades. A alternativa Layer-First, onde o código é agrupado por tipo de camada (ex: todas as telas em uma pasta, todos os providers em outra), pode levar à dispersão dos arquivos de uma mesma feature, tornando a navegação e a compreensão do código mais difíceis à medida que o projeto cresce.\[20\] Adotar a estrutura Feature-First desde o início, mesmo para um MVP, estabelece uma base sólida.

A estrutura de pastas principal dentro do diretório \`lib/\` será:

\*   \*\*\`main.dart\`\*\*: Ponto de entrada da aplicação Flutter. Responsável pela inicialização de configurações globais, como o \`ProviderScope\` (se Riverpod for usado) ou \`MultiProvider\` (se Provider for usado), e por chamar o widget raiz da aplicação.  
\*   \*\*\`app/\`\*\* (ou \`app.dart\` contendo o widget \`MyApp\`):  
    \*   \`view/app.dart\`: Contém o widget raiz da aplicação (ex: \`MaterialApp\` ou \`CupertinoApp\`), onde são definidas configurações globais como tema e rotas iniciais.  
\*   \*\*\`core/\`\*\* (ou \`common/\`): Este diretório abrigará código que é compartilhado e reutilizado por múltiplas features.  
    \*   \`utils/\`: Funções utilitárias genéricas (ex: formatadores de data, validadores de string, constantes globais).  
    \*   \`widgets/\`: Widgets customizados que são genéricos e podem ser usados em várias partes da aplicação (ex: \`CustomAppBar\`, \`LoadingIndicatorWidget\`, \`EmptyStateWidget\`).  
    \*   \`services/\`: Serviços compartilhados, como o \`api\_service.dart\`, que encapsulará a lógica de comunicação com a API backend Go (usando o pacote \`http\`).  
    \*   \`theme/\`: Definições do tema da aplicação (cores, tipografia, estilos de widgets).  
    \*   \`models/\` (Opcional aqui, ou dentro de cada feature): Modelos de dados que são usados por múltiplas features (ex: \`ProductModel\`). Se um modelo for muito específico de uma feature, pode residir dentro da pasta da feature.  
\*   \*\*\`features/\`\*\*: Diretório raiz para todas as funcionalidades principais do Mini WMS. Cada subdiretório aqui representa uma feature.  
    \*   \*\*\`products/\`\*\*: Feature de Gerenciamento de Produtos.  
        \*   \`data/\` (Opcional para o MVP, se a lógica de dados for simples e puder ser contida nos providers/notifiers):  
            \*   \`models/\`: \`product\_model.dart\` (se não estiver em \`core/models/\`).  
            \*   \`repositories/\`: (Opcional para o MVP) Interface e implementação do repositório de produtos, responsável por abstrair as chamadas à API relacionadas a produtos.  
        \*   \`presentation/\`: Camada de apresentação da feature.  
            \*   \`providers/\` (ou \`notifiers/\`, \`cubits/\`, \`blocs/\` dependendo da escolha de gerenciamento de estado): Contém a lógica de estado específica da feature de produtos (ex: \`product\_list\_provider.dart\`, \`product\_form\_provider.dart\`).  
            \*   \`screens/\` (ou \`pages/\`, \`views/\`): Widgets que representam telas completas da feature (ex: \`product\_list\_screen.dart\`, \`product\_form\_screen.dart\`, \`product\_detail\_screen.dart\`).  
            \*   \`widgets/\`: Widgets menores, específicos para a feature de produtos, reutilizados dentro de suas telas (ex: \`product\_list\_item\_widget.dart\`, \`product\_search\_bar.dart\`).  
    \*   \*\*\`inventory/\`\*\*: Feature de Gerenciamento de Inventário (Entradas, Saídas, Consulta de Estoque).  
        \*   (Estrutura similar à da feature \`products/\`, com seus próprios providers, screens e widgets para registrar entradas, saídas e visualizar o estoque).  
\*   \*\*\`navigation/\`\*\* (ou \`routes/\`):  
    \*   \`app\_router.dart\`: Configuração centralizada da navegação e das rotas da aplicação (ex: usando rotas nomeadas do Flutter ou um pacote como GoRouter, se a complexidade justificar para o MVP).

\*\*Tabela 4 – Estrutura de Diretórios Sugerida para o Frontend Flutter (Feature-First)\*\*

| Caminho | Descrição |
| :---- | :---- |
| lib/main.dart | Ponto de entrada da aplicação, inicialização de serviços globais/providers. |
| lib/app/view/app.dart | Widget raiz da aplicação (MaterialApp/CupertinoApp), configuração de tema e rota inicial. |
| lib/core/ | Código compartilhado entre features. |
| lib/core/utils/ | Funções utilitárias, constantes globais. |
| lib/core/widgets/ | Widgets genéricos reutilizáveis em toda a aplicação. |
| lib/core/services/api\_service.dart | Lógica centralizada para comunicação com a API Go (usando o pacote http). |
| lib/core/theme/app\_theme.dart | Definições de tema da aplicação (cores, fontes, etc.). |
| lib/features/ | Diretório raiz para todas as features do aplicativo. |
| lib/features/products/ | Feature de Gerenciamento de Produtos. |
| lib/features/products/presentation/screens/ | Telas da feature de produtos (ex: product\_list\_screen.dart, product\_form\_screen.dart). |
| lib/features/products/presentation/widgets/ | Widgets específicos da feature de produtos (ex: product\_card.dart). |
| lib/features/products/presentation/providers/ | Lógica de estado para a feature de produtos (ex: product\_provider.dart). |
| lib/features/inventory/ | Feature de Gerenciamento de Inventário (Entrada, Saída, Consulta de Estoque). |
| lib/features/inventory/presentation/screens/ | Telas da feature de inventário (ex: inventory\_transaction\_screen.dart, stock\_query\_screen.dart). |
| lib/features/inventory/presentation/widgets/ | Widgets específicos da feature de inventário. |
| lib/features/inventory/presentation/providers/ | Lógica de estado para a feature de inventário (ex: inventory\_provider.dart). |
| lib/navigation/app\_router.dart | Configuração de navegação e definição de rotas nomeadas. |

Esta estrutura visa um equilíbrio entre organização e simplicidade, adequada para o desenvolvimento ágil do MVP.

2.4.2 Gerenciamento de Estado

A escolha do gerenciamento de estado é crucial para a manutenibilidade e escalabilidade de uma aplicação Flutter. Para o MVP do Mini WMS, a decisão deve ponderar a curva de aprendizado, a quantidade de boilerplate e a adequação às necessidades do projeto.\[21\]

\*\*Opções Consideradas:\*\*  
\*   \*\*\`setState()\`:\*\* Adequado apenas para estado local de widgets muito simples. Insuficiente para compartilhar estado entre telas ou gerenciar dados de API de forma eficaz.  
\*   \*\*Provider (\`provider\` package):\*\* Uma solução madura, recomendada pela equipe do Flutter para muitos cenários, e relativamente simples de aprender e usar.\[22\] Utiliza \`ChangeNotifier\`, \`ChangeNotifierProvider\` e \`Consumer\` (ou \`context.watch/read\`) para propagar o estado e reconstruir a UI. É bem adequado para aplicações de pequeno a médio porte.  
\*   \*\*Riverpod (\`flutter\_riverpod\` package):\*\* Considerado uma evolução do Provider, oferece maior segurança de compilação, flexibilidade com diferentes tipos de providers (ex: \`StateProvider\` para estado simples, \`FutureProvider\` para dados assíncronos de APIs, \`NotifierProvider\` para lógica de estado mais complexa) e resolve algumas das limitações do Provider.\[23\]  
\*   \*\*BLoC/Cubit (\`flutter\_bloc\` package):\*\* BLoC é um padrão robusto para separar lógica de negócios da UI, ideal para aplicações complexas, mas geralmente envolve mais código boilerplate (eventos, estados, blocos).\[24, 25\] Cubit é uma versão simplificada do BLoC, com menos boilerplate, mas Provider ou Riverpod ainda podem ser mais diretos para o escopo do MVP.

\*\*Escolha para o MVP: Provider.\*\*

\*Justificativa:\* Embora Riverpod seja uma excelente e moderna solução, para um MVP com prazo de 10 dias e com o objetivo de maximizar a velocidade de desenvolvimento (assumindo que o desenvolvedor pode não ter profunda experiência prévia com Riverpod), \*\*Provider\*\* oferece uma curva de aprendizado mais suave e é amplamente documentado.\[22\] Sua simplicidade é uma vantagem significativa sob restrição de tempo. Ele é capaz de lidar com as necessidades de estado do Mini WMS, como gerenciar listas de produtos, o estado de formulários e os resultados de chamadas de API (carregando, dados, erro).  
Se o desenvolvedor já possuir proficiência em Riverpod, esta seria uma escolha igualmente válida e potencialmente superior devido às suas melhorias sobre o Provider. No entanto, para este plano, a recomendação pende para a opção com menor barreira de entrada para garantir a viabilidade do prazo.

\*   \*\*Implementação com Provider:\*\*  
    \*   \*\*\`ChangeNotifier\`:\*\* Classes que contêm o estado e a lógica para modificá-lo, notificando os ouvintes sobre as mudanças usando \`notifyListeners()\`. Por exemplo, \`ProductProvider extends ChangeNotifier\` gerenciaria a lista de produtos e o estado de carregamento/erro ao buscar da API.  
    \*   \*\*\`ChangeNotifierProvider\`:\*\* Usado para fornecer uma instância de um \`ChangeNotifier\` para seus descendentes na árvore de widgets. Geralmente colocado acima dos widgets que precisam acessar o estado.  
    \*   \*\*\`Consumer\` / \`context.watch\<T\>()\` / \`context.read\<T\>()\`:\*\* Usados dentro dos widgets para ouvir as mudanças no \`ChangeNotifier\` e reconstruir a UI, ou para chamar métodos do notifier.

2.4.3 Comunicação com API (Backend Go)

A comunicação entre o frontend Flutter e o backend Go será realizada através de requisições HTTP para a API RESTful. O pacote \*\*\`http\`\*\* será utilizado para esta finalidade.

\*Justificativa:\* O pacote \`http\` (\`package:http/http.dart\`) é a solução padrão e recomendada pela comunidade Flutter para realizar chamadas de rede.\[26, 27, 28\] Ele é robusto, bem testado e fornece todas as funcionalidades necessárias para interagir com uma API RESTful.

\*   \*\*Operações HTTP:\*\*  
    \*   \*\*\`http.get(Uri.parse(url))\`\*\*: Para buscar dados, como a lista de produtos, detalhes de um produto específico ou o saldo de estoque.  
    \*   \*\*\`http.post(Uri.parse(url), headers: headers, body: body)\`\*\*: Para criar novos recursos (ex: cadastrar produto, registrar entrada/saída).  
        \*   \`headers\`: Deverá incluir \`{'Content-Type': 'application/json; charset=UTF-8'}\` para indicar que o corpo da requisição é JSON.\[26, 29\]  
        \*   \`body\`: Os dados a serem enviados (geralmente um \`Map\<String, dynamic\>\`) devem ser codificados para uma string JSON usando \`jsonEncode(data)\` da biblioteca \`dart:convert\`.\[26, 29\]  
    \*   \*\*\`http.put(Uri.parse(url), headers: headers, body: body)\`\*\*: Para atualizar um recurso existente (ex: modificar dados de um produto). Similar ao \`POST\` em termos de cabeçalhos e corpo.  
    \*   \*\*\`http.delete(Uri.parse(url), headers: headers)\`\*\*: Para remover um recurso (ex: excluir um produto).  
\*   \*\*Tratamento de Resposta:\*\*  
    \*   O objeto \`http.Response\` retornado pelas chamadas contém \`statusCode\` e \`body\`.  
    \*   \`response.statusCode\`: Será verificado para determinar o sucesso ou falha da operação (ex: 200 OK, 201 Created, 400 Bad Request, 404 Not Found).\[26\]  
    \*   \`response.body\`: Se a requisição for bem-sucedida e retornar dados, o corpo (uma string JSON) será decodificado para objetos Dart (geralmente \`Map\<String, dynamic\>\` ou \`List\<dynamic\>\`) usando \`jsonDecode(response.body)\` da biblioteca \`dart:convert\`.\[26\]  
\*   \*\*Gerenciamento de Estado da Chamada:\*\* A lógica de estado associada às chamadas de API (como indicar carregamento, exibir dados recebidos ou mostrar mensagens de erro) será gerenciada pelo Provider (ou Riverpod, se escolhido). Por exemplo, um \`ProductProvider\` teria métodos como \`fetchProducts()\`, que internamente usaria o pacote \`http\`, atualizaria seu estado interno (lista de produtos, isLoading, errorMessage) e notificaria os widgets consumidores.

Um serviço de API centralizado (\`api\_service.dart\` em \`lib/core/services/\`) pode ser criado para encapsular a lógica de todas as chamadas HTTP, tornando o código dos providers mais limpo e focado na gestão do estado.

2.4.4 Design de UI/UX (Interface do Usuário e Experiência do Usuário)

Para o MVP de 10 dias, o design da UI/UX priorizará a funcionalidade, clareza e usabilidade básica sobre a estética elaborada. Serão seguidos princípios de design para aplicações CRUD (Create, Read, Update, Delete) em dispositivos móveis.\[30, 31\]

\*   \*\*Princípios Fundamentais para o MVP:\*\*  
    \*   \*\*Simplicidade e Clareza:\*\* A interface será mantida limpa, com foco nas tarefas essenciais de gerenciamento de produtos e estoque. Evitar-se-á poluição visual e excesso de informações em tela.\[30\]  
    \*   \*\*Consistência (Material Design):\*\* Será utilizado o conjunto de widgets e diretrizes do Material Design, fornecido pelo Flutter, para garantir uma aparência e comportamento consistentes e familiares aos usuários de Android (e aceitável para usuários de iOS no contexto do MVP).\[30, 32\] Isso inclui o uso de componentes como \`AppBar\`, \`FloatingActionButton\`, \`Card\`, \`ListTile\`, \`TextFormField\`, etc.  
    \*   \*\*Responsividade Básica:\*\* O layout das telas deverá se adaptar minimamente a diferentes tamanhos de tela de smartphones. Widgets como \`Expanded\`, \`Flexible\`, \`MediaQuery\`, \`LayoutBuilder\` e o uso de \`FractionallySizedBox\` ou \`AspectRatio\` podem ajudar a alcançar uma responsividade básica sem grande complexidade.\[31\]  
    \*   \*\*Feedback ao Usuário:\*\* Serão fornecidos feedbacks visuais para as ações do usuário, como indicadores de carregamento (\`CircularProgressIndicator\`) durante chamadas de API, e mensagens de sucesso ou erro (usando \`SnackBar\` ou diálogos simples).  
    \*   \*\*Navegação Intuitiva:\*\* A transição entre telas será clara e previsível.  
    \*   \*\*Foco no Toque:\*\* Elementos interativos como botões e campos de formulário terão tamanhos adequados para interação por toque, com espaçamento suficiente para evitar toques acidentais.\[33\]

\*   \*\*Principais Telas (Exemplos de Componentes e Layouts):\*\*  
    \*   \*\*Tela de Listagem de Produtos:\*\*  
        \*   Layout: \`Scaffold\` com \`AppBar\` (título "Produtos", talvez um ícone de busca simples) e \`FloatingActionButton\` para adicionar um novo produto.  
        \*   Corpo: \`ListView.builder\` para exibir a lista de produtos de forma eficiente.\[34, 35\] Cada item da lista pode ser um \`Card\` ou \`ListTile\` exibindo o nome do produto, código e quantidade em estoque. Um toque no item pode navegar para detalhes/edição.  
    \*   \*\*Tela de Formulário de Produto (Cadastro/Edição):\*\*  
        \*   Layout: \`Scaffold\` com \`AppBar\` (título "Novo Produto" ou "Editar Produto").  
        \*   Corpo: Um \`Form\` widget contendo \`TextFormField\`s para os campos do produto (código, nome, descrição, unidade, quantidade inicial para cadastro). Rótulos claros (\`labelText\`) e dicas (\`hintText\`) serão usados.\[33\] Botões para "Salvar" e "Cancelar". Validação de formulário básica (campos obrigatórios, formato numérico para quantidade).  
    \*   \*\*Tela de Registro de Entrada/Saída de Estoque:\*\*  
        \*   Layout: \`Scaffold\` com \`AppBar\`.  
        \*   Corpo: Formulário simples para:  
            \*   Selecionar o produto: Pode ser um \`DropdownButtonFormField\` listando os produtos existentes ou um campo de texto com busca simples.  
            \*   Informar a quantidade da movimentação (\`TextFormField\` numérico).  
            \*   Botão para "Registrar Entrada" ou "Registrar Saída".  
    \*   \*\*Tela de Consulta de Estoque (ou integrada):\*\*  
        \*   Pode ser uma funcionalidade integrada na tela de listagem de produtos (exibindo a quantidade ao lado de cada produto) ou uma tela separada onde o usuário busca por um produto e vê seu saldo. Para o MVP, a integração na listagem é mais simples.

\*   \*\*Considerações de Design Adicionais:\*\*  
    \*   \*\*Não Bloquear Orientação:\*\* O aplicativo não deve ter sua orientação bloqueada para retrato, permitindo flexibilidade.\[31\]  
    \*   \*\*Material Design como Padrão:\*\* Para o MVP, não haverá uma distinção visual significativa entre o design para iOS (Cupertino) e Android. O Material Design será o padrão para ambas as plataformas para acelerar o desenvolvimento, sendo uma abordagem aceitável conforme discutido em \[32\] como uma das alternativas para desenvolvimento multiplataforma rápido.

O foco será em criar uma experiência de usuário que seja funcional e não frustrante, permitindo que as tarefas principais sejam realizadas de forma eficiente.

2.4.5 Navegação

A navegação entre as diferentes telas da aplicação Flutter será implementada de forma simples e direta, adequada ao escopo do MVP.

\*Justificativa:\* Para um aplicativo com um número limitado de telas como o Mini WMS, a API de navegação padrão do Flutter, utilizando \`Navigator.push()\`, \`Navigator.pop()\` e \`MaterialPageRoute\`, é suficiente e a mais rápida de implementar.\[36, 37\] Soluções de roteamento mais avançadas como \`GoRouter\` ou \`auto\_route\` oferecem mais funcionalidades (deep linking, navegação tipada, guards), mas podem introduzir uma complexidade e curva de aprendizado desnecessárias para o prazo de 10 dias, a menos que o desenvolvedor já seja altamente proficiente nelas.

\*   \*\*Abordagem de Navegação:\*\*  
    \*   \*\*Rotas Nomeadas (\`Navigator.pushNamed()\`):\*\* Mesmo para um aplicativo pequeno, o uso de rotas nomeadas é uma boa prática, pois melhora a legibilidade do código, facilita a manutenção e o desacoplamento das telas.\[36\] As rotas serão definidas em um local centralizado (ex: no widget \`MaterialApp\` ou em um arquivo \`app\_router.dart\`).  
    \*   \*\*Passagem de Argumentos:\*\* Argumentos para as rotas (ex: ID de um produto para a tela de edição) serão passados através do parâmetro \`arguments\` do \`Navigator.pushNamed()\` ou utilizando soluções de passagem de argumentos tipadas, se um gerador de rotas simples for usado.  
    \*   \*\*Estrutura de Rotas (Exemplos):\*\*  
        \*   \`/\` (ou \`/products\`): Tela de listagem de produtos (tela inicial).  
        \*   \`/product\_new\`: Tela de formulário para cadastrar um novo produto.  
        \*   \`/product\_edit\`: Tela de formulário para editar um produto existente (espera um ID de produto como argumento).  
        \*   \`/inventory\_entry\`: Tela para registrar entrada de estoque.  
        \*   \`/inventory\_exit\`: Tela para registrar saída de estoque.

\*   \*\*Exemplo de Definição de Rota Nomeada (em \`MaterialApp\`):\*\*  
    \`\`\`dart  
    // MaterialApp(  
    //   initialRoute: '/products', // Ou a rota inicial desejada  
    //   routes: {  
    //     '/products': (context) \=\> ProductListScreen(),  
    //     '/product\_new': (context) \=\> ProductFormScreen(),  
    //     '/product\_edit': (context) {  
    //       final productId \= ModalRoute.of(context)\!.settings.arguments as String; // Exemplo  
    //       return ProductFormScreen(productId: productId);  
    //     },  
    //     // Outras rotas...  
    //   },  
    // )  
    \`\`\`

\*   \*\*Navegação:\*\*  
    \*   Para ir para uma nova tela: \`Navigator.pushNamed(context, '/product\_new');\`  
    \*   Para voltar à tela anterior: \`Navigator.pop(context);\`

Esta abordagem de navegação é simples, eficaz para o escopo do MVP e alinha-se com as práticas padrão do Flutter para gerenciamento de rotas em aplicações de pequeno a médio porte.

2.4.6 Testes

A estratégia de testes para o frontend Flutter no MVP se concentrará em garantir a funcionalidade básica dos widgets e a corretude da lógica de estado.

\*   \*\*Testes de Widget (\`widget\_test.dart\`):\*\*  
    \*   \*\*Foco:\*\* Testar individualmente os principais widgets e telas da aplicação para verificar se eles são renderizados corretamente com base em diferentes estados e se respondem a interações básicas do usuário (como toques em botões, preenchimento de formulários).  
    \*   \*\*Ferramentas:\*\* Pacote \`flutter\_test\` da SDK do Flutter. Utiliza-se \`WidgetTester\` para interagir com os widgets e \`expect\` com \`finders\` e \`matchers\` para fazer asserções sobre o estado da UI.  
    \*   \*\*Exemplos de Cenários de Teste de Widget:\*\*  
        \*   Verificar se a tela de listagem de produtos exibe uma \`CircularProgressIndicator\` quando está carregando.  
        \*   Verificar se a lista de produtos é exibida corretamente após os dados serem carregados.  
        \*   Verificar se o formulário de produto exibe mensagens de erro de validação quando campos obrigatórios não são preenchidos.  
        \*   Simular o toque em um botão "Salvar" e verificar se a ação esperada (ex: chamada a um método do provider) ocorre.  
    \*   \*Justificativa para o MVP:\* Testes de widget são relativamente rápidos de escrever e executar, fornecendo um bom nível de confiança sobre a corretude dos componentes da UI e sua lógica básica.

\*   \*\*Testes Unitários (\`test/\`):\*\*  
    \*   \*\*Foco:\*\* Testar a lógica dentro das classes de \`ChangeNotifier\` (ou equivalentes, se outra solução de estado for usada). Isso inclui testar métodos que fazem chamadas de API (mockando o serviço de API), manipulam o estado interno e notificam os ouvintes.  
    \*   \*\*Ferramentas:\*\* Pacote \`test\` da SDK do Dart. Se os notifiers dependerem de serviços (como \`ApiService\`), estes serviços serão mockados usando bibliotecas como \`mockito\` ou implementações de mock manuais.  
    \*   \*\*Exemplos de Cenários de Teste Unitário:\*\*  
        \*   Testar se o \`ProductProvider\`, ao chamar \`fetchProducts()\`, atualiza corretamente seu estado para \`isLoading \= true\`, depois para \`isLoading \= false\` e popula a lista de produtos em caso de sucesso, ou define uma mensagem de erro em caso de falha da API.  
        \*   Testar a lógica de incremento/decremento de um contador, se houver.  
    \*   \*Justificativa para o MVP:\* Garantir que a lógica de estado, que é o coração da interatividade da aplicação, funcione corretamente de forma isolada.

Para o MVP de 10 dias, testes de UI mais complexos e demorados, como testes de integração completos (usando \`flutter\_driver\` para simular interações do usuário em um dispositivo/emulador real) ou testes \*end-to-end\* com o backend, estarão fora do escopo devido à restrição de tempo.\[38\] O foco será em testes de widget e unitários para cobrir os aspectos mais críticos da UI e da lógica de estado.

2.5 Cronograma de Desenvolvimento (MVP \- 10 dias)

O desenvolvimento do MVP do Mini WMS Simplificado será conduzido em um cronograma intensivo de 10 dias úteis. Uma abordagem ágil simplificada será adotada, com foco em entregas funcionais incrementais e comunicação constante (mesmo que seja um auto-alinhamento para um desenvolvedor solo).

\*   \*\*Abordagem Ágil Simplificada:\*\*  
    \*   \*\*Ciclos Curtos:\*\* O trabalho será dividido em metas diárias ou a cada dois dias.  
    \*   \*\*Priorização Contínua:\*\* As funcionalidades CRUD de Produto e Consulta de Estoque são a maior prioridade, seguidas pelas funcionalidades de Entrada e Saída.  
    \*   \*\*Flexibilidade:\*\* Embora haja um plano, pequenas adaptações podem ser necessárias com base no progresso e nos desafios encontrados.

\*   \*\*Divisão de Tarefas (Backend/Frontend):\*\*  
    Para um desenvolvedor solo, as tarefas de backend e frontend serão realizadas de forma sequencial, com possível alternância para manter o progresso em ambas as frentes ou para desbloquear dependências. Se mais de um desenvolvedor estiver envolvido, algumas tarefas poderão ser paralelizadas. O cronograma abaixo assume um esforço concentrado.

\*   \*\*Metas e Atividades:\*\*

\*\*Tabela 5 – Cronograma Detalhado do MVP (10 Dias)\*\*

| Dia | Foco Principal (B \= Backend, F \= Frontend) | Tarefas Chave | Entregáveis Esperados |
| :---- | :---- | :---- | :---- |
| 1 | B: Setup e Modelo de Dados | Configurar projeto Go (estrutura de pastas, go.mod). Definir modelo de dados Product (struct Go). Criar schema da tabela products no SQLite. Configurar driver SQLite. | Projeto Go inicializado com main.go. Schema da tabela products definido e script de migração inicial (se usado). |
| 2 | B: API CRUD Produtos | Implementar handlers, serviços e repositórios para as operações CRUD (Create, Read, Update, Delete) de Produtos. Escrever testes unitários e de integração básicos. | Endpoints RESTful /products (GET all, GET by ID, POST, PUT, DELETE) funcionais e minimamente testados. |
| 3 | F: Setup e Listagem de Produtos | Configurar projeto Flutter (estrutura de pastas Feature-First). Implementar tela de listagem de produtos, consumindo o endpoint GET /products da API Go. Configurar Provider. | Projeto Flutter inicializado. Tela de listagem exibindo produtos (mockados ou da API, se o backend estiver pronto). |
| 4 | F: Formulário de Produtos | Implementar formulário de cadastro/edição de produtos, com chamadas aos endpoints POST /products e PUT /products/{id}. Incluir validação básica no formulário. | Funcionalidade de criar e editar produtos no frontend, com comunicação com o backend. |
| 5 | B: Lógica de Inventário (Entrada/Saída) | Modificar/Confirmar a tabela Product para incluir o campo quantity. Implementar lógica de atualização de estoque na API para registrar entradas e saídas. | Endpoints RESTful para registrar entrada (/inventory/entries), saída (/inventory/exits) e consultar estoque (/inventory/products/{productId}/stock) funcionais. |
| 6 | B: Testes de Inventário / F: Preparação Telas de Inventário | Escrever testes unitários e de integração para a lógica e endpoints de inventário. No frontend, começar a estruturar as telas de entrada/saída de estoque. | Backend de inventário testado. Esqueletos das telas de inventário no Flutter. |
| 7 | F: Telas de Entrada/Saída de Estoque | Implementar as telas no Flutter para registrar entrada e saída de produtos, consumindo os respectivos endpoints da API de inventário. | Funcionalidade de registrar entrada e saída de produtos no frontend, atualizando o estoque via backend. |
| 8 | F: Consulta de Estoque e Navegação | Implementar a funcionalidade de consulta de estoque no frontend (integrada à listagem ou tela separada). Configurar a navegação entre todas as telas da aplicação. | Consulta de estoque funcional no frontend. Navegação completa e fluida entre as telas do MVP. |
| 9 | B+F: Testes de Fluxo e Correções | Realizar testes de fluxo completo (ex: cadastrar produto \-\> registrar entrada \-\> registrar saída \-\> consultar estoque). Identificar e corrigir bugs. Realizar pequenos ajustes de UI/UX. | MVP funcional com os principais fluxos de usuário testados e a maioria dos bugs críticos corrigidos. |
| 10 | Documentação e "Entrega" | Finalizar os arquivos README.md para o backend e frontend (instruções de setup, build, run). Limpar código, remover comentários de debug. Preparar para demonstração ou um deploy simples (ex: build de APK Android). | Código limpo e comentado (onde necessário). READMEs completos. MVP pronto para avaliação/demonstração. |

Este cronograma é agressivo e exige foco total e um escopo rigidamente controlado. Qualquer desvio ou adição de complexidade não planejada pode comprometer o prazo. A simplicidade das escolhas tecnológicas (Go/Chi, Flutter/Provider, SQLite) é fundamental para a viabilidade deste plano.

2.6 Versionamento e Boas Práticas de Código

A adoção de boas práticas de versionamento e documentação de código é essencial, mesmo em projetos de curto prazo e com equipe reduzida, para garantir a rastreabilidade, colaboração (mesmo que futura) e manutenibilidade.

\*   \*\*Sistema de Controle de Versão (Git):\*\* \[39\]  
    \*   \*\*Repositórios Separados:\*\* Serão criados dois repositórios Git distintos: um para o backend Go (\`mini-wms-api\`) e outro para o frontend Flutter (\`mini-wms-app\`).  
    \*   \*\*Commits Frequentes e Atômicos:\*\* As alterações no código serão commitadas frequentemente. Cada commit deve ser "atômico", ou seja, representar uma unidade lógica de trabalho pequena e coesa (ex: implementação de uma função, correção de um bug específico, adição de um campo em um formulário). Isso facilita a revisão, o rastreamento de alterações e a reversão, se necessário.\[39\]  
    \*   \*\*Mensagens de Commit Descritivas:\*\* As mensagens de commit seguirão um padrão claro e informativo, preferencialmente começando com um verbo no imperativo que descreva a ação realizada (ex: "Adiciona endpoint para criar produto", "Corrige validação do campo de email", "Implementa listagem de produtos na tela inicial"). O corpo da mensagem pode fornecer mais detalhes sobre o "porquê" da mudança, se necessário.\[39\]  
    \*   \*\*Uso de Branches:\*\* Mesmo para um desenvolvedor solo, o trabalho em novas funcionalidades ou correções significativas deve ser feito em branches separados (ex: \`feature/product-crud\`, \`fix/stock-update-bug\`). O branch principal (\`main\` ou \`master\`) deve ser mantido estável e conter apenas código funcional e testado. Os branches de feature são mesclados ao principal após a conclusão e teste.\[39\] Esta prática organiza o desenvolvimento e isola o trabalho em andamento.

\*   \*\*Documentação de Código:\*\*  
    \*   \*\*Comentários:\*\* Código complexo, algoritmos não triviais ou decisões de design que não são imediatamente óbvias devem ser acompanhados de comentários claros e concisos.  
    \*   \*\*Go:\*\* Seguir as convenções para comentários de documentação que podem ser processados pela ferramenta \`godoc\`. Funções e tipos exportados devem ter comentários explicando seu propósito e uso.  
    \*   \*\*Flutter (Dart):\*\* Usar comentários de documentação de três barras (\`///\`) para classes, métodos e propriedades públicas. Estes comentários podem ser usados para gerar documentação com \`dartdoc\`.

\*   \*\*Arquivo \`README.md\`:\*\* \[39\]  
    Cada repositório (backend e frontend) terá seu próprio arquivo \`README.md\` na raiz. Este arquivo é a porta de entrada do projeto e deve conter informações essenciais.  
    \*   \*\*Estrutura Sugerida para \`README.md\`:\*\*  
        1\.  \*\*Título do Projeto:\*\* (Ex: Mini WMS Simplificado \- API Backend Go)  
        2\.  \*\*Breve Descrição:\*\* Um ou dois parágrafos sobre o propósito do componente (API ou App).  
        3\.  \*\*Tecnologias Utilizadas:\*\* Lista das principais tecnologias, frameworks e bibliotecas com suas versões (Ex: Go 1.xx, Chi v5.x, SQLite3; Flutter 3.xx, Provider 6.x, http 0.13.x).  
        4\.  \*\*Pré-requisitos:\*\* Softwares e ferramentas que precisam estar instalados para rodar o projeto (Ex: Go, Flutter SDK, Git).  
        5\.  \*\*Configuração do Ambiente:\*\* Passos para configurar o ambiente de desenvolvimento.  
        6\.  \*\*Como Compilar e Rodar o Projeto:\*\* Comandos detalhados para build e execução (Ex: \`go run cmd/mini-wms-api/main.go\` para o backend; \`flutter run\` para o frontend).  
        7\.  \*\*(Para Backend) Endpoints da API:\*\* Uma breve lista dos principais endpoints e como usá-los (pode referenciar uma documentação de API mais detalhada, como Postman Collection ou Swagger, se criada).  
        8\.  \*\*Estrutura do Projeto:\*\* Uma visão geral da organização das pastas e arquivos importantes.  
        9\.  \*\*Decisões de Design Chave:\*\* Um resumo conciso das principais escolhas arquiteturais e tecnológicas e suas justificativas.  
        10\. \*\*Testes:\*\* Como executar os testes unitários e de integração.  
        11\. \*\*Como Contribuir (Opcional):\*\* Diretrizes para futuras contribuições.  
        12\. \*\*Licença (Opcional):\*\* Informação sobre a licença do software (ex: MIT, Apache 2.0).

    Um \`README.md\` bem elaborado é crucial para a usabilidade e longevidade do projeto, facilitando o entendimento e a execução por qualquer pessoa (incluindo o próprio desenvolvedor no futuro).

A adesão a estas práticas de versionamento e documentação, mesmo de forma simplificada, contribuirá significativamente para a qualidade e organização do projeto Mini WMS Simplificado.

\*\*3 CONSIDERAÇÕES FINAIS\*\*

Este plano de projeto delineou as etapas, escolhas tecnológicas e arquiteturais, e o cronograma para o desenvolvimento do Produto Mínimo Viável (MVP) do "Mini WMS Simplificado". A abordagem adotada priorizou a simplicidade, a rapidez de desenvolvimento e a utilização de tecnologias robustas e bem estabelecidas, como Go para o backend e Flutter para o frontend, com persistência de dados em SQLite.

3.1 Reafirmação da Viabilidade do Plano para o MVP de 10 Dias

O desenvolvimento do MVP do Mini WMS Simplificado no prazo estipulado de 10 dias é um objetivo desafiador, mas considerado exequível sob as condições e escolhas detalhadas neste plano. A viabilidade se sustenta em:  
\*   \*\*Escopo Rigidamente Controlado:\*\* A definição precisa das funcionalidades mínimas essenciais, excluindo quaisquer complexidades desnecessárias para a primeira entrega.  
\*   \*\*Escolhas Tecnológicas Pragmáticas:\*\* A seleção de Go com Chi, Flutter com Provider e SQLite visa minimizar a curva de aprendizado e o tempo de configuração, permitindo foco na implementação da lógica de negócio.  
\*   \*\*Arquitetura Simplificada:\*\* A adoção de um monólito para o backend e uma gestão de estado direta no frontend evita a sobrecarga de padrões arquiteturais mais complexos, que seriam inadequados para o prazo.  
\*   \*\*Cronograma Detalhado:\*\* A divisão do trabalho em metas diárias fornece um roteiro claro e permite o monitoramento do progresso.

É imperativo que, durante a execução, haja um compromisso contínuo em manter o foco nas funcionalidades essenciais do MVP, resistindo à tentação de adicionar funcionalidades "interessantes, mas não cruciais" que poderiam comprometer o prazo. A disciplina na execução do plano será um fator determinante para o sucesso.

3.2 Próximos Passos Sugeridos Após a Aprovação do Plano

Com a aprovação deste plano de projeto, os próximos passos imediatos incluem:  
1\.  \*\*Revisão Final do Escopo do MVP:\*\* Uma última validação com os stakeholders (se houver) ou pelo próprio desenvolvedor para confirmar o entendimento e o aceite do escopo mínimo.  
2\.  \*\*Configuração dos Ambientes de Desenvolvimento:\*\* Preparação das máquinas de desenvolvimento com as versões necessárias de Go, Flutter, Git e quaisquer outras ferramentas auxiliares.  
3\.  \*\*Criação dos Repositórios Git:\*\* Inicialização dos repositórios para o backend e frontend, conforme as boas práticas de versionamento.  
4\.  \*\*Início do Desenvolvimento:\*\* Começar a execução das tarefas conforme o cronograma detalhado na Tabela 5, iniciando pelo setup do backend e a definição do modelo de dados.  
5\.  \*\*Comunicação e Monitoramento:\*\* Estabelecer um ritmo de acompanhamento do progresso (mesmo que auto-avaliações diárias para um desenvolvedor solo) para identificar rapidamente quaisquer impedimentos ou desvios do plano.

3.3 Potenciais Evoluções Futuras do Mini WMS (Pós-MVP)

O MVP de 10 dias é concebido como um ponto de partida funcional que valida a ideia central do Mini WMS Simplificado. Após a entrega e avaliação do MVP, o sistema possui um grande potencial para evoluções futuras, agregando mais valor e atendendo a necessidades mais complexas. Algumas das potenciais evoluções incluem:

\*   \*\*Autenticação e Autorização Avançadas:\*\* Implementação de um sistema de múltiplos usuários com diferentes níveis de acesso e papéis (ex: administrador, operador).  
\*   \*\*Relatórios Gerenciais:\*\* Desenvolvimento de funcionalidades para gerar relatórios sobre movimentações de estoque, produtos mais movimentados, níveis de estoque ao longo do tempo, etc.  
\*   \*\*Histórico Detalhado de Movimentações:\*\* Criação de um log completo de todas as transações de entrada e saída, permitindo auditoria e rastreabilidade.  
\*   \*\*Alertas de Estoque:\*\* Configuração de alertas para notificar quando o estoque de determinados produtos atingir níveis mínimos pré-definidos.  
\*   \*\*Interface Web para Administração:\*\* Desenvolvimento de uma interface web para gerenciamento do sistema, possivelmente com funcionalidades mais ricas que a aplicação móvel.  
\*   \*\*Integração com Leitor de Código de Barras:\*\* Adição de suporte para leitura de códigos de barras através da câmera do dispositivo móvel ou leitores dedicados para agilizar o registro de produtos.  
\*   \*\*Múltiplos Armazéns/Localizações:\*\* Capacidade de gerenciar o estoque em diferentes locais físicos.  
\*   \*\*Deploy em Ambiente de Produção Robusto:\*\* Migração do backend para uma infraestrutura mais escalável e resiliente (ex: utilizando Docker, serviços de nuvem como AWS, Google Cloud, Azure) e um banco de dados mais parrudo se o volume de dados e transações aumentar significativamente.  
\*   \*\*Melhorias na UI/UX:\*\* Refinamento da interface do usuário com base no feedback dos usuários do MVP, focando em otimizar a experiência e a estética.

Estas evoluções podem ser priorizadas e implementadas em iterações subsequentes, transformando gradualmente o "Mini WMS Simplificado" em uma solução de gerenciamento de estoque cada vez mais completa e adaptada às necessidades de seus usuários.

\*\*REFERÊNCIAS\*\*

ASSOCIAÇÃO BRASILEIRA DE NORMAS TÉCNICAS. \*\*NBR 6022\*\*: Informação e documentação – Artigo em publicação periódica técnica e/ou científica – Apresentação. Rio de Janeiro, 2018\. \[40\]

ASSOCIAÇÃO BRASILEIRA DE NORMAS TÉCNICAS. \*\*NBR 6023\*\*: Informação e documentação – Referências – Elaboração. Rio de Janeiro, 2018\. \[41\]

ASSOCIAÇÃO BRASILEIRA DE NORMAS TÉCNICAS. \*\*NBR 10719\*\*: Informação e documentação – Relatórios técnico-científicos – Apresentação. Rio de Janeiro, 2015\. \[1, 42\]

ASSOCIAÇÃO BRASILEIRA DE NORMAS TÉCNICAS. \*\*NBR ISO/IEC 12207\*\*: Sistemas e engenharia de software – Processos de ciclo de vida de software. Rio de Janeiro, 2017\. \[43\]

EDWARDS, Alex. \*11 Tips for Structuring Your Go Projects\*., 2023\. Disponível em: \[https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects\](https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects). Acesso em:. \[2\]

GEEKSFORGEEKS. \*Flutter \- State Management Provider\*. Disponível em: \[https://www.geeksforgeeks.org/flutter-state-management-provider/\](https://www.geeksforgeeks.org/flutter-state-management-provider/). Acesso em:. \[22\]

GEEKSFORGEEKS. \*Implementing Rest API in Flutter\*. Disponível em: \[https://www.geeksforgeeks.org/implementing-rest-api-in-flutter/\](https://www.geeksforgeeks.org/implementing-rest-api-in-flutter/). Acesso em:. \[26\]

KODY TECHNOLAB. \*Layer-first or Feature-first Flutter project structure?\* Disponível em: \[https://kodytechnolab.com/blog/layer-first-or-feature-first-flutter-project-structure/\](https://kodytechnolab.com/blog/layer-first-or-feature-first-flutter-project-structure/). Acesso em:. \[20\]

LEAPCELL. \*Exploring Golang’s Validation Libraries: Validator vs Ozzo-Validation\*. Disponível em: \[https://leapcell.io/blog/exploring-golang-s-validation-libraries\](https://leapcell.io/blog/exploring-golang-s-validation-libraries). Acesso em:. \[11\]

PUC MINAS. Biblioteca. \*ABNT NBR 10719:2015: Como elaborar e formatar Relatório Técnico e ou Científico\*. Disponível em: \[https://www.pucminas.br/biblioteca/DocumentoBiblioteca/ABNT-Elaborar-formatar-relatorio-tecnico-e-ou-cientifico.pdf\](https://www.pucminas.br/biblioteca/DocumentoBiblioteca/ABNT-Elaborar-formatar-relatorio-tecnico-e-ou-cientifico.pdf). Acesso em:. \[1\]

RELIASOFTWARE. \*Guide for Golang Unit Test: From Basics to Advanced\*. Disponível em: \[https://reliasoftware.com/blog/guide-for-golang-unit-test\](https://reliasoftware.com/blog/guide-for-golang-unit-test). Acesso em:. \[15\]

SPEEDSCALE. \*Testing Golang HTTP Handlers and Clients with httptest\*. Disponível em: \[https://speedscale.com/blog/testing-golang-with-httptest/\](https://speedscale.com/blog/testing-golang-with-httptest/). Acesso em:. \[17\]

STACK OVERFLOW BLOG. \*Best practices for REST API design\*. Disponível em: \[https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/\](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/). Acesso em:. \[5\]

SUHAASYA LINGAM. \*A Beginner's Guide to Simplifying Go API Validation\*. Hashnode, 2024\. Disponível em: \[https://suhaasya.hashnode.dev/a-beginners-guide-to-simplifying-go-api-validation\](https://suhaasya.hashnode.dev/a-beginners-guide-to-simplifying-go-api-validation). Acesso em:. \[10\]

TILLITSDONE. \*Chi vs Other Go Routers: Features & Use Cases\*. Disponível em: \[https://tillitsdone.com/blogs/chi-vs-other-go-routers-guide/\](https://tillitsdone.com/blogs/chi-vs-other-go-routers-guide/). Acesso em:. \[6\]

TWILIO BLOG. \*How To Use SQLite with GoLang\*. Disponível em: \[https://www.twilio.com/en-us/blog/use-sqlite-go\](https://www.twilio.com/en-us/blog/use-sqlite-go). Acesso em:. \[3\]

BETTERSTACK COMMUNITY. \*Golang Testify: A Deep Dive Into Go's Most Popular Testing Toolkit\*. Disponível em: \[https://betterstack.com/community/guides/scaling-go/golang-testify/\](https://betterstack.com/community/guides/scaling-go/golang-testify/). Acesso em:. \[13\]

BLUP. \*Flutter State Management Explained: How to Choose the Right Approach\*. Disponível em: \[https://www.blup.in/blog/flutter-state-management-explained-how-to-choose\](https://www.blup.in/blog/flutter-state-management-explained-how-to-choose). Acesso em:. \[21\]

CODINGEXPLORATIONS. \*Managing Files in a Go API: Folder Structure Best Practices for Organizing Your Project\*. Disponível em: \[https://www.codingexplorations.com/blog/managing-files-in-a-go-api-folder-structure-best-practices-for-organizing-your-project\](https://www.codingexplorations.com/blog/managing-files-in-a-go-api-folder-structure-best-practices-for-organizing-your-project). Acesso em:. \[4\]

CODINGEXPLORATIONS. \*Using Go Validator for Efficient Data Validation in Go Applications\*. Disponível em: \[https://www.codingexplorations.com/blog/using-go-validator-for-efficient-data-validation-in-go-applications\](https://www.codingexplorations.com/blog/using-go-validator-for-efficient-data-validation-in-go-applications). Acesso em:. \[9\]

CODEPARROT. \*Riverpod Flutter: A Beginner's Guide\*. Disponível em: \[https://codeparrot.ai/blogs/riverpod-flutter-a-beginners-guide\](https://codeparrot.ai/blogs/riverpod-flutter-a-beginners-guide). Acesso em:. \[23\]

DHIWISE. \*Flutter Project Structure for Building Future-Proof Flutter Apps\*. Disponível em: \[https://www.dhiwise.com/post/flutter-project-structure-for-building-future-proof-flutter-apps\](https://www.dhiwise.com/post/flutter-project-structure-for-building-future-proof-flutter-apps). Acesso em:. \[44\]

FLUTTER. \*Layout in Flutter\*. Disponível em: \[https://docs.flutter.dev/ui/layout\](https://docs.flutter.dev/ui/layout). Acesso em:. \[35\]

FLUTTER. \*ListView class \- widgets library \- Dart API\*. Disponível em: \[https://api.flutter.dev/flutter/widgets/ListView-class.html\](https://api.flutter.dev/flutter/widgets/ListView-class.html). Acesso em:. \[34\]

FORMS ON FIRE. \*Mobile Form Design: 15 Best Practices For Creating User-Friendly Forms\*. Disponível em: \[https://www.formsonfire.com/blog/mobile-form-design\](https://www.formsonfire.com/blog/mobile-form-design). Acesso em:. \[33\]

GITLAB. \*Práticas recomendadas de controle de versão\*. Disponível em: \[https://about.gitlab.com/pt-br/topics/version-control/version-control-best-practices/\](https://about.gitlab.com/pt-br/topics/version-control/version-control-best-practices/). Acesso em:. \[39\]

GOLANG COMPANY. \*A Comparison Between Gin, Gorilla Mux, and Net/HTTP in Golang\*. Disponível em: \[https://www.golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http\](https://www.golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http). Acesso em:. \[7\]

\*Observação: As datas de acesso e URLs completas devem ser preenchidas conforme o momento da consulta final. Os exemplos de código Go e Flutter foram simplificados e comentados para fins ilustrativos no plano; a implementação real exigirá mais detalhes e tratamento de erros completo.\*

#### **Referências citadas**

1. www.pucminas.br, acessado em junho 4, 2025, [https://www.pucminas.br/biblioteca/DocumentoBiblioteca/ABNT-Elaborar-formatar-relatorio-tecnico-e-ou-cientifico.pdf](https://www.pucminas.br/biblioteca/DocumentoBiblioteca/ABNT-Elaborar-formatar-relatorio-tecnico-e-ou-cientifico.pdf)  
2. 11 tips for structuring your Go projects – Alex Edwards, acessado em junho 4, 2025, [https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects](https://www.alexedwards.net/blog/11-tips-for-structuring-your-go-projects)  
3. How to Use SQLite With Go | Twilio, acessado em junho 4, 2025, [https://www.twilio.com/en-us/blog/use-sqlite-go](https://www.twilio.com/en-us/blog/use-sqlite-go)  
4. Managing Files in a Go API Folder Structure: Best Practices for ..., acessado em junho 4, 2025, [https://www.codingexplorations.com/blog/managing-files-in-a-go-api-folder-structure-best-practices-for-organizing-your-project](https://www.codingexplorations.com/blog/managing-files-in-a-go-api-folder-structure-best-practices-for-organizing-your-project)  
5. Best practices for REST API design \- Stack Overflow, acessado em junho 4, 2025, [https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/)  
6. Chi vs Other Go Routers: Features & Use Cases \- Till it's done, acessado em junho 4, 2025, [https://tillitsdone.com/blogs/chi-vs-other-go-routers-guide/](https://tillitsdone.com/blogs/chi-vs-other-go-routers-guide/)  
7. Comparison between Gin, Gorilla Mux and Net/Http \- Golang Company, acessado em junho 4, 2025, [https://www.golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http](https://www.golang.company/blog/comparison-between-gin-gorilla-mux-and-net-http)  
8. go sqlite tutorial example \- GitHub Gist, acessado em junho 4, 2025, [https://gist.github.com/NaniteFactory/9554eb82a88f8d5ecc0a7a76ea5fe1c3](https://gist.github.com/NaniteFactory/9554eb82a88f8d5ecc0a7a76ea5fe1c3)  
9. Using Go Validator for Efficient Data Validation in Go Applications ..., acessado em junho 4, 2025, [https://www.codingexplorations.com/blog/using-go-validator-for-efficient-data-validation-in-go-applications](https://www.codingexplorations.com/blog/using-go-validator-for-efficient-data-validation-in-go-applications)  
10. Go API Validation Made Easy \- suhaasya, acessado em junho 4, 2025, [https://suhaasya.hashnode.dev/a-beginners-guide-to-simplifying-go-api-validation](https://suhaasya.hashnode.dev/a-beginners-guide-to-simplifying-go-api-validation)  
11. Exploring Golang's Validation Libraries | Leapcell, acessado em junho 4, 2025, [https://leapcell.io/blog/exploring-golang-s-validation-libraries](https://leapcell.io/blog/exploring-golang-s-validation-libraries)  
12. Benchmark Result of go-playground/validator | ozzo-validation | GoValidator : r/golang, acessado em junho 4, 2025, [https://www.reddit.com/r/golang/comments/1i19ru9/benchmark\_result\_of\_goplaygroundvalidator/](https://www.reddit.com/r/golang/comments/1i19ru9/benchmark_result_of_goplaygroundvalidator/)  
13. Testing in Go with Testify | Better Stack Community, acessado em junho 4, 2025, [https://betterstack.com/community/guides/scaling-go/golang-testify/](https://betterstack.com/community/guides/scaling-go/golang-testify/)  
14. Test with Testify and Mockery in Go \- Outcome School, acessado em junho 4, 2025, [https://outcomeschool.com/blog/test-with-testify-and-mockery-in-go](https://outcomeschool.com/blog/test-with-testify-and-mockery-in-go)  
15. Comprehensive Guide for Golang Unit Test with Code Examples ..., acessado em junho 4, 2025, [https://reliasoftware.com/blog/guide-for-golang-unit-test](https://reliasoftware.com/blog/guide-for-golang-unit-test)  
16. Go Unit Testing: A Complete Guide to Writing Reliable Tests \- Ceos3c, acessado em junho 4, 2025, [https://www.ceos3c.com/golang/go-unit-testing-a-complete-guide-to-writing-reliable-tests/](https://www.ceos3c.com/golang/go-unit-testing-a-complete-guide-to-writing-reliable-tests/)  
17. Testing Golang with httptest | Speedscale, acessado em junho 4, 2025, [https://speedscale.com/blog/testing-golang-with-httptest/](https://speedscale.com/blog/testing-golang-with-httptest/)  
18. Use httptest when testing HTTP handlers in Go/Golang \- willem.dev, acessado em junho 4, 2025, [https://www.willem.dev/articles/testing-http-handlers-using-httptest/](https://www.willem.dev/articles/testing-http-handlers-using-httptest/)  
19. github.com, acessado em junho 4, 2025, [https://github.com/bizz84/flutter-tips-and-tricks/blob/main/tips/0039-flutter-project-structure-feature-first-or-layer-first/index.md\#:\~:text=For%20any%20given%20feature%2C%20files,the%20same%20top%2Dlevel%20folder.](https://github.com/bizz84/flutter-tips-and-tricks/blob/main/tips/0039-flutter-project-structure-feature-first-or-layer-first/index.md#:~:text=For%20any%20given%20feature%2C%20files,the%20same%20top%2Dlevel%20folder.)  
20. Layer-First Or Feature-First To Structure Your Flutter App?, acessado em junho 4, 2025, [https://kodytechnolab.com/blog/layer-first-or-feature-first-flutter-project-structure/](https://kodytechnolab.com/blog/layer-first-or-feature-first-flutter-project-structure/)  
21. Flutter State Management: How to Choose the Right Approach \- Blup, acessado em junho 4, 2025, [https://www.blup.in/blog/flutter-state-management-explained-how-to-choose](https://www.blup.in/blog/flutter-state-management-explained-how-to-choose)  
22. Flutter – State Management Provider | GeeksforGeeks, acessado em junho 4, 2025, [https://www.geeksforgeeks.org/flutter-state-management-provider/](https://www.geeksforgeeks.org/flutter-state-management-provider/)  
23. Riverpod Flutter: A Beginner's Guide \- CodeParrot, acessado em junho 4, 2025, [https://codeparrot.ai/blogs/riverpod-flutter-a-beginners-guide](https://codeparrot.ai/blogs/riverpod-flutter-a-beginners-guide)  
24. Bloc vs. Cubit: Which One Should You Use? \- Insight, acessado em junho 4, 2025, [https://insight.vayuz.com/insight-detail/bloc-vs.-cubit:-which-one-should-you-use-/bmV3c18xNzQyODEwNDQxNTM5](https://insight.vayuz.com/insight-detail/bloc-vs.-cubit:-which-one-should-you-use-/bmV3c18xNzQyODEwNDQxNTM5)  
25. Cubit vs BLoC: A Comprehensive Comparison for Flutter Developers \- DhiWise, acessado em junho 4, 2025, [https://www.dhiwise.com/post/cubit-vs-bloc-a-comprehensive-comparison](https://www.dhiwise.com/post/cubit-vs-bloc-a-comprehensive-comparison)  
26. Implementing Rest API in Flutter | GeeksforGeeks, acessado em junho 4, 2025, [https://www.geeksforgeeks.org/implementing-rest-api-in-flutter/](https://www.geeksforgeeks.org/implementing-rest-api-in-flutter/)  
27. Flutter Accessing REST API \- Tutorialspoint, acessado em junho 4, 2025, [https://www.tutorialspoint.com/flutter/flutter\_accessing\_rest\_api.htm](https://www.tutorialspoint.com/flutter/flutter_accessing_rest_api.htm)  
28. How To Make HTTP Requests With Flutter And Parse JSON Result Data \- QuickCoder, acessado em junho 4, 2025, [https://quickcoder.org/flutter-http-requests/](https://quickcoder.org/flutter-http-requests/)  
29. How to Nail Flutter Post Request: A Beginner-to-Advanced Developer's Guide \- DhiWise, acessado em junho 4, 2025, [https://www.dhiwise.com/post/how-to-nail-flutter-post-request-a-developers-guid](https://www.dhiwise.com/post/how-to-nail-flutter-post-request-a-developers-guid)  
30. 10 Best Practices For Designing User Interfaces in Flutter \- GeeksforGeeks, acessado em junho 4, 2025, [https://www.geeksforgeeks.org/best-practices-for-designing-user-interfaces-in-flutter/](https://www.geeksforgeeks.org/best-practices-for-designing-user-interfaces-in-flutter/)  
31. Best practices for adaptive design \- Flutter Documentation, acessado em junho 4, 2025, [https://docs.flutter.dev/ui/adaptive-responsive/best-practices](https://docs.flutter.dev/ui/adaptive-responsive/best-practices)  
32. What's your approach to building for iOS and Android? : r/FlutterDev \- Reddit, acessado em junho 4, 2025, [https://www.reddit.com/r/FlutterDev/comments/1hws52k/whats\_your\_approach\_to\_building\_for\_ios\_and/](https://www.reddit.com/r/FlutterDev/comments/1hws52k/whats_your_approach_to_building_for_ios_and/)  
33. 13 Mobile Form Design Best Practices and Examples for Beginners, acessado em junho 4, 2025, [https://www.formsonfire.com/blog/mobile-form-design](https://www.formsonfire.com/blog/mobile-form-design)  
34. ListView class \- widgets library \- Dart API, acessado em junho 4, 2025, [https://api.flutter.dev/flutter/widgets/ListView-class.html](https://api.flutter.dev/flutter/widgets/ListView-class.html)  
35. Layout | Flutter, acessado em junho 4, 2025, [https://docs.flutter.dev/ui/layout](https://docs.flutter.dev/ui/layout)  
36. Navigating the Basics: A Comprehensive Guide to Flutter Routing \- DhiWise, acessado em junho 4, 2025, [https://www.dhiwise.com/post/deep-dive-into-flutter-routing-everything-you-need-to-know](https://www.dhiwise.com/post/deep-dive-into-flutter-routing-everything-you-need-to-know)  
37. Routing and Navigation Basics \- Flutter Training, acessado em junho 4, 2025, [https://fluttertraining.dev/article/Flutter\_Navigation\_Routing\_and\_Navigation\_Basics.html](https://fluttertraining.dev/article/Flutter_Navigation_Routing_and_Navigation_Basics.html)  
38. Flutter App Best Practices To Follow in 2025 | Blog Miquido, acessado em junho 4, 2025, [https://www.miquido.com/blog/flutter-app-best-practices/](https://www.miquido.com/blog/flutter-app-best-practices/)  
39. Quais são as melhores práticas de controle de versão do Git? \- GitLab, acessado em junho 4, 2025, [https://about.gitlab.com/pt-br/topics/version-control/version-control-best-practices/](https://about.gitlab.com/pt-br/topics/version-control/version-control-best-practices/)  
40. www.cetam.am.gov.br, acessado em junho 4, 2025, [https://www.cetam.am.gov.br/wp-content/uploads/2022/02/01\_ABNT\_NBR\_6022-2018\_Artigo-em-publicacao-periodica-tecnica-eou-cientifica-1.pdf](https://www.cetam.am.gov.br/wp-content/uploads/2022/02/01_ABNT_NBR_6022-2018_Artigo-em-publicacao-periodica-tecnica-eou-cientifica-1.pdf)  
41. NBR 6023 Informação e documentação \- Referências \- Elaboração \- UFPE, acessado em junho 4, 2025, [https://www.ufpe.br/documents/40070/848544/abntnbr6023.pdf/092b145a-7dce-4b97-8514-364793d8877e](https://www.ufpe.br/documents/40070/848544/abntnbr6023.pdf/092b145a-7dce-4b97-8514-364793d8877e)  
42. INSTITUTO FEDERAL DE SANTA CATARINA \- IFSC CURSO DE XXXX \- XXX NOME NBR 10719: apresentação de relatórios técnico-científic, acessado em junho 4, 2025, [https://www.ifsc.edu.br/documents/30721/897853/Anexo+III+-+Relatorio+tecnico.pdf/2a4e7b77-c293-482e-1c08-4ce8762709c8](https://www.ifsc.edu.br/documents/30721/897853/Anexo+III+-+Relatorio+tecnico.pdf/2a4e7b77-c293-482e-1c08-4ce8762709c8)  
43. ABNT NBR ISO/IEC 12207 NBRISO/IEC12207 ... \- Target Normas, acessado em junho 4, 2025, [https://www.normas.com.br/visualizar/abnt-nbr-nm/10985/abnt-nbriso-iec12207-sistemas-e-engenharia-de-software-processos-de-ciclo-de-vida-de-software](https://www.normas.com.br/visualizar/abnt-nbr-nm/10985/abnt-nbriso-iec12207-sistemas-e-engenharia-de-software-processos-de-ciclo-de-vida-de-software)  
44. Flutter Project Structure Revamped for 2024 Flutter Projects \- DhiWise, acessado em junho 4, 2025, [https://www.dhiwise.com/post/flutter-project-structure-for-building-future-proof-flutter-apps](https://www.dhiwise.com/post/flutter-project-structure-for-building-future-proof-flutter-apps)