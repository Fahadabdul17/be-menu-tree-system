# Menu Tree Management System

A production-ready hierarchical menu management system built with Go, Gin, and GORM. Designed for high performance and clean architecture.

## 🚀 Features
- **O(N) Tree Construction**: High-performance algorithm for building nested structures.
- **Recursive Operations**: Efficiently delete subtrees using PostgreSQL CTE.
- **Circular Dependency Check**: Prevents infinite loops in the menu hierarchy.
- **Docker Ready**: Easy deployment with Docker and Docker Compose.
- **Clean Architecture**: Separation of concerns between Handler, Service, and Repository.

## 🛠 Tech Stack
- **Backend**: Go 1.22, Gin, GORM
- **Database**: PostgreSQL 15
- **Docs**: Swagger UI
- **DevOps**: Docker, Docker Compose

## 📦 Getting Started

### 1. Using Docker (Recommended)
```bash
# Start the application and database
docker-compose up --build
```
The API will be available at `http://localhost:8080`.
Swagger UI: `http://localhost:8080/swagger/index.html`

### 2. Manual Setup
1.  Copy `.env.example` to `.env` and configure your DB.
2.  Install dependencies: `go mod download`
3.  Run the app: `go run cmd/api/main.go`

## 🧪 Testing
Run unit tests for business logic:
```bash
go test ./internal/service/... -v
```

## 🏗 Architecture
The project follows **Clean Architecture** principles:
- **cmd/api**: Application entry point.
- **internal/api**: HTTP Handlers.
- **internal/service**: Business logic and tree building.
- **internal/repository**: Optimized database operations.
- **pkg/**: Shared utilities (logger, db connection).

## 📊 Performance
- **Tree Building**: Optimized to $O(N)$ complexity using hash maps, making it suitable for large menu structures.
- **Recursive Deletion**: Uses PostgreSQL **Recursive CTE** to perform subtree deletions in a single database transaction.
- **Indexing**: `parent_id` and `order` fields are indexed for fast retrieval and sorting.

## 📸 Screenshots
![Architecture Diagram](architecture_diagram.png)
*(Run the application to see the Swagger UI in action)*
