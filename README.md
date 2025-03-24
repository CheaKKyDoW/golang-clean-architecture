# Overview



websocket-chat/
│── internal/
│   ├── chat/               # Chat Domain
│   │   ├── handler/        # HTTP / WebSocket Handler
│   │   │   ├── http.go     # HTTP Handler
│   │   │   ├── websocket.go # WebSocket Handler
│   │   ├── repository/     # Data Layer
│   │   │   ├── chat_repo.go  # Chat Repository
│   │   ├── usecase/        # Business Logic
│   │   │   ├── chat_usecase.go  # Chat Usecase
│   │   ├── model/          # Chat Models
│   │   │   ├── chat.go     # Chat Model
│   │   ├── websocket/      # WebSocket Logic
│   │   │   ├── manager.go  # WebSocket Manager
│   │   ├── routes.go       # Route Register
│── pkg/                    # Shared Utilities
│── cmd/                    # Entrypoint (main.go)
│── configs/                # Configurations
│── build/                  # Docker / CI Configs
│── README.md               # Project Documentation
│── go.mod                  # Dependencies
