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
.
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── internal
│   └── chat
│       ├── adapter
│       │   ├── middleware
│       │   │   └── logging.go
│       │   ├── persistence
│       │   │   ├── migrations
│       │   │   │   ├── 20250625163230_create_rooms_table.down.sql
│       │   │   │   └── 20250625163230_create_rooms_table.up.sql
│       │   │   ├── postgresql.go
│       │   │   └── room_repository_impl.go
│       │   ├── rest
│       │   │   ├── dto
│       │   │   │   └── room
│       │   │   │       ├── create_room_request.go
│       │   │   │       └── update_room_request.go
│       │   │   └── room_router.go
│       │   └── utils
│       │       └── error_response.go
│       ├── controller
│       │   └── room
│       │       ├── command
│       │       │   └── room_command.go
│       │       └── query
│       │           └── room_query.go
│       └── domain
│           ├── model
│           │   ├── error_response.go
│           │   ├── room.go
│           │   └── user.go
│           ├── port
│           │   └── room_repository.go
│           └── service
│               └── room
│                   ├── create_room.go
│                   ├── delete_room.go
│                   ├── get_all_rooms.go
│                   ├── get_room_by_id.go
│                   └── update_room_name.go
├── README.md
└── sample.env
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
   createdb your_database_name
