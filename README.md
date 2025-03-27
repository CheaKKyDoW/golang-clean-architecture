# Clean Architecture Monorepo WebSocket Chat

## Overview
This is a simple WebSocket-based chat application built with Golang. It's a personal project designed for learning and practice, focusing on real-time communication and backend development. The project follows clean architecture principles and uses WebSocket, Redis, and PostgreSQL for handling real-time messaging and data storage.

## Why I Built This
I started this project to:
- Get hands-on experience with WebSockets in Golang.
- Explore how to structure a backend using clean architecture.
- Learn how Redis can help with caching and pub/sub messaging.
- Improve my skills in building scalable backend systems.

## Project Structure

```
websocket-chat/
│── cmd/server              # Entrypoint (main.go)
│── internal/
│   ├── chat/               # Chat Domain
│   │   ├── handler/        # HTTP / WebSocket Handler
│   │   ├── repository/     # Data Layer
│   │   ├── usecase/        # Business Logic
│   │   ├── model/          # Chat Models
│   │   ├── server.go       # Route Register
│   │   ├── infrastructure/ 
│   │       ├── client/      # Client to make HTTP calls
│   │       ├── database/    # PostgreSQL, Redis configurations
│   │       ├── server/      # Main server settings & domain configs
│   │       ├── validate/    # Validators
│   │       ├── websocket/   # WebSocket management
│── pkg/                    # Shared Utilities
│── configs/                # Configurations
│── mock/                   # Mock Data
│── test/                   # Tests
│── README.md               # Project Documentation
│── go.mod                  # Dependencies
```

## Features
- Real-time messaging using WebSockets
- Persistent message storage with PostgreSQL
- Redis integration for caching and pub/sub communication
- Clean architecture for better code organization

## Prerequisites
- Golang 1.21+
- PostgreSQL (for storing messages)
- Redis (for handling real-time events)
- Docker (optional, but helpful for running dependencies)

## How to Run
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/websocket-chat.git
   cd websocket-chat
   ```
2. Install dependencies (PostgreSQL, Redis) manually if not using Docker.
3. Run the application:
   ```sh
   go run cmd/main.go
   ```

## Configuration
Before running the application, update the environment variables in the `configs/` directory to match your setup. Here’s an example `.env` file:

```
APP_ENV="local"
APP_DEBUG=true
DATABASE_HOST="localhost"
DATABASE_USER="youruser"
DATABASE_PASSWORD="yourpassword"
DATABASE_NAME="chatdb"
DATABASE_PORT=5432
DATABASE_SSLMODE="disable"
REDIS_HOST="localhost:6379"
JWT_SECRET="your-secret-key"
JWT_EXPIRE="1h"
```
Additionally, you may need to configure your settings.json for gopls:
```
"gopls": {
    "build.buildFlags": [
        "-tags=domain name tag"
    ]
}
```
## Testing API Endpoints
### WebSocket Connection
- `ws://localhost:8080/ws` – Connect to the WebSocket server

### HTTP Endpoints
- `POST /messages` – Send a new chat message
- `GET /messages` – Retrieve chat history
