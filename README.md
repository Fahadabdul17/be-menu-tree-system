# Menu Tree Management System

A production-ready hierarchical menu management system built with **Go**, **Gin**, and **GORM**. This system is designed for high performance, utilizing optimized algorithms for tree construction and recursive database operations.

## 🚀 Features

- **O(N) Tree Construction**: High-performance algorithm for building nested structures from flat database results.
- **Recursive Subtree Operations**: Efficiently delete entire subtrees using PostgreSQL **Recursive CTE**.
- **Circular Dependency Check**: Robust business logic to prevent infinite loops in the menu hierarchy (e.g., a node cannot be its own ancestor).
- **Clean Architecture**: Strict separation of concerns (API Handler -> Service -> Repository).
- **Structured Logging**: Integrated production-ready logging with Uber-Zap.
- **API Documentation**: Automated Swagger documentation for easy integration.
- **Docker Ready**: Fully containerized with Docker and Docker Compose.

## 🛠 Tech Stack

- **Backend**: Go 1.25, Gin Web Framework
- **Database**: PostgreSQL 15 (using GORM ORM)
- **Documentation**: Swagger (swaggo)
- **Logging**: Zap Logger
- **DevOps**: Docker, Docker Compose

## 📦 Getting Started

### Prerequisites
- [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.25+](https://go.dev/dl/) (if running manually)
- [PostgreSQL](https://www.postgresql.org/) (if running manually)

### 1. Using Docker (Recommended)
This will set up both the Go API and the PostgreSQL database automatically.

```bash
# Clone the repository
git clone https://github.com/Fahadabdul17/be-menu-tree-system.git
cd be-menu-tree-system

# Start the application and database
docker-compose up --build
```
- API Endpoint: `http://localhost:8080/api`
- Swagger UI: `http://localhost:8080/swagger/index.html`

### 2. Manual Setup
1.  **Environment Variables**: Copy `.env.example` to `.env` and configure your database credentials.
    ```bash
    cp .env.example .env
    ```
2.  **Install Dependencies**:
    ```bash
    go mod download
    ```
3.  **Run Migrations**: The application automatically migrates the schema on startup.
4.  **Start Server**:
    ```bash
    go run cmd/api/main.go
    ```

## 🛤 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/api/menus` | Get the full hierarchical menu tree |
| `POST` | `/api/menus` | Create a new menu item |
| `GET` | `/api/menus/:id` | Get a single menu item by ID |
| `PUT` | `/api/menus/:id` | Update a menu item |
| `DELETE` | `/api/menus/:id` | Delete a menu item and all its descendants |
| `PATCH` | `/api/menus/:id/move` | Move a menu to a new parent |
| `PATCH` | `/api/menus/:id/reorder` | Update the sorting order of a menu item |

## 🏗 Architecture

The project follows the **Clean Architecture** pattern to ensure testability and maintainability:

- **`cmd/api`**: Application entry point and dependency injection.
- **`internal/api`**: HTTP Handlers and Middleware.
- **`internal/service`**: Business logic, tree construction, and validation.
- **`internal/repository`**: Database persistence layer with optimized SQL/GORM queries.
- **`internal/model`**: Domain entities and database schema.
- **`pkg/`**: Reusable utility packages (logger, db connection, custom response).

## 📸 Screenshots & Demos

### Architecture Overview
![Architecture Diagram](architecture_diagram.png)

### API Documentation (Swagger)
![Swagger UI Placeholder](https://raw.githubusercontent.com/swaggo/swag/master/assets/swagger-ui.png)
*(Run the application and visit `/swagger/index.html` to see the interactive documentation)*

## 🧪 Testing

The system includes unit tests for core business logic, including tree building and circular dependency checks.

```bash
go test ./internal/service/... -v
```

## 📊 Performance Optimization

- **Tree Building**: The algorithm uses a single-pass hash map approach, ensuring **O(N)** time complexity even for thousands of nodes.
- **Database**: PostgreSQL **Recursive Common Table Expressions (CTE)** are used for subtree deletions, avoiding multiple round-trips to the database.
- **Indexing**: Database indexes on `parent_id` and `order` ensure fast hierarchical queries.
