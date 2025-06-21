# Smart Plane

![Smart Plane Banner](docs/assets/top_banner.png)

[![Go](https://img.shields.io/badge/Go-1.19+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://github.com/rafa-mori/smart_plane/blob/main/LICENSE)
[![Automation](https://img.shields.io/badge/automation-smart%20contracts-blue)](#features)

---

**Smart Plane é uma plataforma modular para contratos inteligentes, autenticação, validação e gestão de documentos sobre Hyperledger Fabric, com API extensível e arquitetura Go idiomática.**

---

## **Table of Contents**

1. [About the Project](#about-the-project)
2. [Features](#features)
3. [Project Structure](#project-structure)
4. [Core Components](#core-components)
5. [Installation](#installation)
6. [Usage](#usage)
7. [Roadmap](#roadmap)
8. [Contributing](#contributing)
9. [Contact](#contact)

---

## **About the Project**

Smart Plane combina contratos inteligentes, autenticação robusta, validação extensível e integração com Hyperledger Fabric para criar uma base segura e escalável para aplicações de documentos, identidades e transações. O núcleo é altamente modular, com tipos e interfaces exportados para fácil extensão e integração.

---

## **Features**

- 🔗 **Contratos inteligentes plugáveis** (Approval, Signature, Traffic, etc.)
- 🔒 **Autenticação JWT e gerenciamento de chaves RSA**
- 🧩 **Validação extensível e listeners de eventos**
- 🛠️ **API Go idiomática e interfaces para integração**
- 📦 **Arquitetura modular e fácil de manter**
- 📝 **Documentação e tipos exportados para uso externo**

---

## **Project Structure**

```plaintext
./
├── api                  # Exposed API for external integrations
├── cmd                  # Command line interface
├── flight.go            # Main exported interface for the module
├── internal
│   ├── authentication   # Authentication related functionalities
│   ├── interfaces       # Types abstraction for modularity and exported API
│   └── smart_contracts  # Smart contract related functionalities
├── logger               # Logging utilities
├── smart_plane.go       # Main entry point for the smart plane module
├── types                # Type definitions for the module (exported)
└── version              # Versioning services and utilities
```

---

## **Core Components**

### `flight.go`

- Interface principal exportada do módulo. Centraliza a inicialização e integração dos principais serviços do Smart Plane.

### `internal/authentication/auth_manager.go`

- Gerenciamento de autenticação JWT.
- Geração e validação de tokens de acesso e refresh.
- Integração com serviços de certificados RSA.

### `internal/smart_contracts/`

- **injection.go**: Estruturas e métodos para injeção de contratos, gerenciamento de chaves, requests e erros.
- **metadata.go**: Estruturas base para metadados de contratos inteligentes (ID, nome, versão, owner, etc).
- **smart_plane.go**: BlockchainManager para registro, consulta, aprovação, assinatura e exclusão de documentos em contratos inteligentes (Approval, Signature, Traffic).
- **state_content.go**: Estrutura genérica para resposta de contratos, com tipagem dinâmica.

### `types/`

- **reference.go**: Tipos e utilitários para identificação única e nomeação de entidades.
- **validation.go**: Infraestrutura de validação extensível, com funções, resultados, prioridades e integração com interfaces.
- **validation_listener.go**: Sistema de listeners para eventos de validação, com filtros, handlers e registro dinâmico.

---

## **Installation**

Requisitos:

- Go 1.19+
- Hyperledger Fabric (para uso blockchain)

Clone o repositório e compile:

```sh
git clone https://github.com/rafa-mori/smart_plane.git
cd smart_plane
go build -o smart_plane .
```

---

## **Usage**

- Importe o módulo em seu projeto Go ou utilize como serviço standalone.
- Exemplo de inicialização do BlockchainManager:

```go
bm := smart_plane.NewBlockchainManager()
err := bm.RegisterDocument("ApprovalContract", "doc123", "conteúdo do documento")
```

- Para autenticação, utilize o AuthManager para geração e validação de tokens JWT.

---

## **Roadmap**

- [x] Núcleo modular para contratos inteligentes
- [x] Autenticação JWT e gerenciamento de chaves
- [x] Infraestrutura de validação extensível
- [x] Listeners de eventos de validação
- [ ] Suporte a novos tipos de contratos
- [ ] Dashboard web para monitoramento

---

## **Contributing**

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests. Veja o [Guia de Contribuição](docs/CONTRIBUTING.md) para mais detalhes.

---

## **Contact**

💌 **Developer**:  
[Rafael Mori](mailto:faelmori@gmail.com)  
💼 [Follow me on GitHub](https://github.com/rafa-mori)  
Estou aberto a colaborações e novas ideias. Se achou o projeto interessante, entre em contato!

---

**Made with care by the Mori family!** ❤️
