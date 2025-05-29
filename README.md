# Smart Plane

Smart Plane é um projeto inovador que combina a tecnologia de contratos inteligentes com a infraestrutura de Hyperledger Fabric para criar uma plataforma segura e escalável para gestão de documentos e transações. Este repositório contém o código-fonte, documentação e recursos necessários para desenvolver, implantar e manter o sistema.

## Estrutura do Repositório

```plaintext
smart-plane
│
├── chaincode                    # 📜 Código dos contratos inteligentes (Hyperledger Fabric)
│   ├── document_chaincode       # 📝 Chaincode para registros de documentos
│   ├── identity_chaincode       # 🔐 Chaincode para gestão de identidades e permissões
│   ├── transaction_chaincode    # 🔄 Chaincode para validação de transações
│   ├── utils.go                 # 🛠️ Funções auxiliares e segurança
│   └── main.go                  # 🚀 Arquivo de inicialização do chaincode
│
├── network                      # ⚡ Configuração da rede Hyperledger Fabric
│   ├── config.yaml              # 🔧 Configuração dos peers, canais e ordens
│   ├── crypto-config            # 🔑 Certificados TLS e Autoridades Certificadoras
│   ├── docker-compose.yml       # 🐳 Arquivo para orquestração via Docker
│   ├── scripts                  # 🔄 Scripts para inicialização e deploy
│   └── start-network.sh         # 🚀 Script para levantar a rede Fabric
│
├── api                          # 🌐 Interface RESTful para comunicação com `smart_plane`
│   ├── handlers                 # 📩 Manipuladores de requisições HTTP
│   ├── middleware               # 🔐 Controle de autenticação e segurança
│   ├── routes.go                # 🛣️ Definição das rotas da API
│   └── main.go                  # 🚀 Inicialização do servidor REST
│
├── clients                      # 🏛️ Aplicações cliente para interação com `smart-plane`
│   ├── web                      # 🌍 Interface web para visualizar registros
│   ├── cli                      # 🖥️ Ferramenta CLI para gerenciar contratos via terminal
│   ├── mobile                   # 📱 Aplicação mobile para usuários finais
│   └── sdk-go                   # 🏗️ SDK em Go para integração de terceiros
│
├── tests                          # ✅ Testes automatizados de contrato e API
│   ├── chaincode_test.go        # 🔍 Testes unitários dos contratos inteligentes
│   ├── api_test.go              # 🔥 Testes de integração da API REST
│   ├── security_test.go         # 🛡️ Testes de segurança e permissões
│   └── performance_test.go      # 🚀 Testes de desempenho da rede Fabric
│
├── docs                           # 📖 Documentação do projeto
│   ├── architecture.md          # 🏗️ Especificação arquitetural do `smart-plane`
│   ├── api-reference.md         # 🌐 Documentação da API REST
│   ├── smart-contracts.md       # 📜 Explicação dos contratos inteligentes
│   ├── roadmap.md               # 🚀 Planejamento de evolução do projeto
│   └── diagrams                 # 📊 Diagramas técnicos de integração
│
├── README.md                    # 📌 Guia inicial do projeto
├── LICENSE                      # ⚖️ Licença do repositório
├── .gitignore                   # 🚫 Ignorar arquivos irrelevantes no Git
└── go.mod                       # 🏗️ Dependências do projeto em Go
```

## Tecnologias Utilizadas

### Infraestrutura e Ferramentas

- **Go**: Linguagem de programação para o backend, escolhida por sua eficiência e escalabilidade.
- **Hyperledger Fabric**: Framework de blockchain para a criação de redes permissionadas.
- **Docker**: Plataforma para containerização e orquestração de serviços.
- **Kubernetes**: Sistema de orquestração de containers para automação de implantação, escalonamento e gerenciamento.

### Banco de Dados e Armazenamento

- **PostgreSQL**: Sistema de gerenciamento de banco de dados relacional utilizado para armazenamento de dados.
- **MongoDB**: Banco de dados NoSQL para armazenamento de documentos e dados não estruturados.
- **RabbitMQ**: Sistema de mensageria para comunicação assíncrona entre serviços.
- **gRPC**: Framework de comunicação eficiente entre serviços, utilizado para a API.
- **Redis**: Armazenamento em cache para melhorar a performance e reduzir latência.

### Segurança e Autenticação

- **Sigstore**: Sistema de assinatura digital para garantir a autenticidade dos registros.
- **Cosign**: Ferramenta para assinatura de imagens de container, garantindo integridade e autenticidade.

### Interface do Usuário

- **React**: Biblioteca JavaScript para construção de interfaces de usuário dinâmicas e responsivas.
- **Flutter**: Framework para desenvolvimento de aplicativos móveis multiplataforma, permitindo uma experiência consistente em iOS e Android.

## Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues, enviar pull requests ou discutir melhorias no projeto. Consulte o arquivo `CONTRIBUTING.md` para mais detalhes sobre como contribuir.

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo `LICENSE` para mais informações.

## Contato

Para dúvidas, sugestões ou colaborações, entre em contato!  

[Gmail](mailto:faelmori@gmail.com)

[GitHub](https://github.com/faelmori)

[Linkedin](https://www.linkedin.com/in/rafa-mori)
