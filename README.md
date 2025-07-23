# Real-Time Chat Application in Go

This project is a **real-time WebSocket-based chat application** built using Go. It follows **Clean Architecture**, **Hexagonal Architecture**, and **Domain-Driven Design (DDD)** principles to ensure maintainability, modularity, and scalability.

---

## 🚀 Features

- Real-time, bi-directional communication using WebSockets
- Unique user/session identification using UUIDs
- Join/Leave notifications
- Multiple predefined chat rooms (e.g., general, tech, random)
- Broadcast messages to all users in a room
- REST API (optional) for room and participant management
- Graceful handling of WebSocket disconnections

---

## 🛠️ Tech Stack

| Component         | Technology                     |
|------------------|--------------------------------|
| Language          | Go (Golang)                    |
| WebSocket Library | `github.com/gorilla/websocket` |
| Router (optional) | `github.com/gorilla/mux`       |
| UUID Generator    | `github.com/satori/go.uuid`    |
| Architecture      | Clean Architecture + DDD + Hexagonal |

---

## 📁 Folder Structure

```bash
├── cmd
├── docs
├── internal
│   └── chat
│       ├── adapter
│       │   ├── middleware
│       │   ├── persistence
│       │   │   ├── migrations
│       │   ├── rest
│       │   │   ├── dto
│       │   │   │   ├── room
│       │   │   │   └── user
│       │   ├── utils
│       │   └── websocket
│       ├── controller
│       │   ├── room
│       │   │   ├── command
│       │   │   └── query
│       │   └── user
│       │       ├── command
│       │       └── query
│       └── domain
│           ├── model
│           ├── port
│           └── service
│               ├── room
│               └── user
└── test
```
---

## ⚙️ Project Setup

1. Copy the example environment file:

   ```bash
   cp sample.env .env

2. Open the newly created .env file and fill in the required values based on your environment.

#####  Setup Postgresql
   ```bash
   brew install postgresql
   brew services start postgresql
   createdb chat_app

## 📋 Documentation:
localhost:8080/docs/index.html
