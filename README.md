````markdown
# ChainForge

**ChainForge** is a production-grade Web3 backend platform built with **Golang microservices**, focusing on **blockchain infrastructure, on-chain data indexing, and secure distributed systems**.

This project is developed incrementally as a public repository to demonstrate real-world backend, blockchain, and Web3 engineering practices.

---

## Vision

ChainForge aims to showcase:

- Custom blockchain implementation (PoW → PoS)
- Wallets, transactions, and cryptographic signing
- Ethereum & Solana on-chain data indexing
- Token price tracking & alerting
- Secure, scalable Golang microservices
- Event-driven architecture with message brokers
- Dockerized & CI-ready services

---

## Architecture Overview

- **Language:** Golang  
- **Architecture:** Microservices + Clean Architecture  
- **APIs:** REST (external), gRPC (internal)  
- **Blockchain:** Custom chain + Ethereum + Solana  
- **Messaging:** Kafka (event-driven)  
- **Storage:** PostgreSQL, Redis, BadgerDB  
- **Security:** JWT, RBAC, cryptographic signatures  
- **DevOps:** Docker, Docker Compose, Kubernetes (planned)

---

## Repository Structure

```text
chainforge/
├── api/           # gRPC & OpenAPI specs
├── cmd/           # Service entry points
├── internal/      # Core business logic
├── pkg/           # Shared libraries
├── scripts/       # Utility scripts
├── deployments/   # Docker & Kubernetes configs
├── docs/          # Architecture & design docs
└── README.md
````

---

## Roadmap

### Phase 1 – Blockchain Core

* Block & Blockchain structures
* Proof-of-Work consensus
* Persistent storage

### Phase 2 – Wallets & Transactions

* Key generation
* Transaction signing
* UTXO model

### Phase 3 – APIs & Services

* Blockchain service API
* Wallet service
* Auth service (JWT, RBAC)

### Phase 4 – Indexing & Pricing

* Ethereum & Solana indexers
* Token price tracking
* Price alerts & notifications

---

## Status

**Active development**
Each day introduces a new production-quality feature.

```
```
