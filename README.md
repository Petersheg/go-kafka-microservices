# Go Kafka Microservices

This project is an **event-driven microservices architecture** implemented in Go, using Docker and Kafka. It demonstrates how multiple services communicate via Kafka topics using a shared Go package for Kafka producers and consumers.

---

## Table of Contents

- [Overview](#overview)  
- [Architecture](#architecture)  
- [Services](#services)  
- [Setup](#setup)  
- [Running the Project](#running-the-project)  

---

## Overview

The project consists of three main services:

1.  **Order Service** – Receives HTTP requests to create orders and publishes `order_created` events to Kafka.  
2.  **Payment Service** – Listens to `order_created` events, processes payments, and publishes `payment_completed` events.  
3.  **Notification Service** – Listens to `payment_completed` events and sends notifications.

All services use a **shared package** located at `pkg/kafka` for Kafka producer and consumer logic.

---

## Architecture

```text
Client
  │
  ▼
Order Service (HTTP API) 
  │ publishes "order_created"
  ▼
Kafka
  │ consumed by Payment Service
Payment Service
  │ publishes "payment_completed"
  ▼
Kafka
  │ consumed by Notification Service
Notification Service
```
---

Dockerized: All services run in isolated containers.

Orchestration: Kafka and Zookeeper are managed via Docker Compose.

Shared Logic: Services use a shared Go module (pkg/kafka) to avoid code duplication.

---

## Services

| Service | Role | Port |
| :--- | :--- | :--- |
| **order** | Creates orders and publishes events | 8080 |
| **payment** | Processes payments for orders | N/A |
| **notification** | Sends notifications after payments | N/A |

## Shared Kafka Topics
* `order_created`
* `payment_completed`

---

## Setup

## Requirements
* **Docker & Docker Compose**
* **Go >= 1.25**
* **Internet connection** (to pull Docker images)

## File Structure
```text
go-kafka-microservices/
├── docker-compose.yml
├── go.mod
├── go.sum
├── pkg/
│   └── kafka/             # Shared Kafka code
└── services/
    ├── order/
    ├── payment/
    └── notification/
```
---

## Running the Project

Follow these steps to spin up the environment and verify the services.

## 1. Build and start all services

Run the following command to build the Go binaries and start the Docker containers:
```bash
docker compose up --build
```

Verify all containers are running:
```bash
docker ps
```

Create an order:
```bash
curl localhost:8080/orders
```