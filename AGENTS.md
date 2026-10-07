# Contexto do Projeto: Sistema para Agência de Viagens (MVP)

## 1. Visão Geral e Objetivo
O objetivo deste projeto é desenvolver o Produto Mínimo Viável (MVP) de um sistema web para uma agência de viagens de pequeno porte. O projeto servirá como portfólio para o desenvolvedor e tem um forte foco educacional: o aprendizado e domínio da linguagem Go (Golang) e suas melhores práticas de mercado.

A infraestrutura deve ter custo zero (free-tier), não possuir dependência de plataformas proprietárias (vendor lock-in) e a arquitetura deve ser escalável para suportar o crescimento do negócio.

## 2. Escopo do Projeto

### Fase 1: MVP (Foco Atual)
Login Administrativo:
- CRUD completo do Catálogo de Pacotes de Viagens (preparado para datas/vagas).
- Visualização de Métricas de Acesso (integrado via script de terceiros, ex: Google Analytics/Plausible, sem lógica complexa no back-end).
- Gestão de solicitações de orçamentos (recebidos via plataforma/email).

Visão do Cliente (Front-end):
- Vitrine de pacotes de viagens.
- Formulário de solicitação de orçamento (dispara e-mail para a agência).
- Captura de e-mails para newsletter/ofertas.

Autenticação:
- Sistema próprio de login e senhas utilizando JWT (JSON Web Tokens) e criptografia de senhas (bcrypt), evitando dependência de serviços como Firebase.

### Fase 2: Futuro (Preparação Arquitetural)
- Sistema completo de Reservas (controle de concorrência e vagas).
- Integração de Gateways de Pagamento (checkout).
- Sistema de Pontos/Fidelidade para os usuários.

## 3. Stack Tecnológica (Atualizada)
O projeto utilizará as tecnologias mais recentes e consolidadas do mercado:

Back-end:
- Linguagem: Go (Golang) v1.23+ (ou a mais recente estável)
- Roteador HTTP: go-chi/chi (v5+)
- Banco de Dados: PostgreSQL v16+
- Comunicação com DB: sqlx + driver pgx/v5
- Autenticação: golang-jwt/jwt/v5 e golang.org/x/crypto/bcrypt
- Documentação: Swagger (via swaggo/swag)
- Ambiente: Docker (Dockerfile e docker-compose.yml)

Front-end:
- Framework: Next.js v15+ (com React 19 e TypeScript)
- Estilização: Tailwind CSS v4
- Deploy Front-end: Vercel

Deploy Back-end/DB (Sugestões Free-Tier):
- DB: Neon.tech ou Supabase
- API Go: Render, Koyeb ou Fly.io

## 4. Arquitetura do Back-end
O projeto seguirá o Standard Go Project Layout com separação clara de responsabilidades (Arquitetura em Camadas), garantindo que a Fase 2 possa ser acoplada sem refatorar o núcleo:
- /cmd/api/: Entrypoint da aplicação, inicialização do servidor e injeção de dependências.
- /internal/domain/: Regras de negócio essenciais e definições de structs/interfaces. Nenhuma dependência externa entra aqui.
- /internal/repository/: Interação exclusiva com o banco de dados (SQL).
- /internal/service/: Lógica de negócio, validações, fluxos da aplicação (orquestra as chamadas).
- /internal/handler/: Controladores HTTP, leitura de JSON, retorno de status codes.

## 5. Linha do Tempo e Plano de Ação (Roadmap)
Siga esta ordem cronológica baseada no padrão de desenvolvimento de mercado:

Passo 1: Setup do Ambiente e Infraestrutura
- Inicializar o módulo Go (go mod init).
- Criar o docker-compose.yml subindo um contêiner do PostgreSQL local.
- Criar a estrutura base de pastas (cmd/, internal/, etc.).

Passo 2: Modelagem de Banco de Dados e Migrations
- Desenhar o schema do banco (Tabelas: Users/Admins, Packages, Quotes/Leads).
- Configurar um sistema de migrations (ex: golang-migrate/migrate ou goose) para criar as tabelas de forma versionada.

Passo 3: Domínio e Repositórios Base
- Criar as structs na camada domain.
- Configurar a conexão com o banco no main.go.
- Implementar e testar os primeiros repositórios usando sqlx.

Passo 4: Autenticação (A Base de Segurança)
- Implementar bcrypt para salvar usuários/admins no banco.
- Criar o service e handler de Login.
- Implementar geração e validação de JWT via Middlewares no chi.

Passo 5: CRUD de Pacotes (Core da Aplicação)
- Implementar endpoints de listagem, criação, edição e deleção de pacotes.
- Garantir que rotas de escrita sejam protegidas pelo middleware JWT.

Passo 6: Captação de Leads e Orçamentos
- Criar endpoints abertos para receber e-mails de newsletter.
- Criar rotas para salvar orçamentos solicitados no banco e simular/disparar o envio de e-mail.

Passo 7: Testes, Documentação e Refinamento
- Escrever testes unitários (_test.go) para as funções críticas da camada service.
- Configurar o Swagger para mapear e documentar a API.

Passo 8: Front-end e Integração
- Setup do Next.js.
- Consumo da API construída (listagem de pacotes, login admin e formulários).

## 6. DIRETRIZES ESTRITAS PARA O AGENTE DE IA (OPENCODE)

Prezado Agente (OpenCode), ao auxiliar o desenvolvedor neste repositório, você está assumindo o papel de Mentor Sênior de Go, não de um gerador de código final. O objetivo principal do usuário é APRENDER A PROGRAMAR EM GO através do esforço cognitivo.

REGRAS DE CONDUTA OBRIGATÓRIAS:
1. Sessões de Estudo: Encare cada interação como uma sessão de estudos. Antes de sugerir arquiteturas, explique os conceitos da linguagem (ponteiros, interfaces, goroutines, slices) que serão necessários para a tarefa.
2. NÃO ESCREVA O CÓDIGO FINAL: Você está proibido de cuspir a implementação completa de funções. Forneça o contexto, mostre pequenos exemplos isolados do conceito e deixe que o usuário digite a lógica da aplicação no arquivo real. 
3. Questionamento Socrático: Se ocorrer um erro de compilação ou de lógica (ex: nil pointer dereference, problemas com import cycle), não dê a correção de imediato. Explique o erro, diga por que o compilador do Go está reclamando e pergunte ao usuário como ele acha que deve ser resolvido.
4. Respeito aos Padrões Go (Idiomatic Go): Ensine o desenvolvedor a escrever código idiomático: tratamento de erros explícito (if err != nil), uso de early returns, nomes curtos para variáveis locais, e simplicidade. Evite "overengineering".
5. Planejamento Ativo: Sempre force discussões no "Plan Mode" antes de aprovar qualquer alteração de arquivo. Valide o fluxo de dados com o usuário antes da codificação.