# Golang Clean Architecture Template

This project is a boilerplate/template for building robust, scalable, and testable web applications using Go (Golang) enabling **Clean Architecture** principles.

## 🚀 Features

- **Clean Architecture Implementation**: Separation of concerns using handlers, usecases, repositories, and entities.
- **Dependency Injection**: Manual dependency injection for clear component wiring.
- **RESTful API**: Built with [Fiber](https://gofiber.io/) web framework.
- **Database ORM**: Uses [GORM](https://gorm.io/) for database interactions with PostgreSQL.
- **Database Transaction Support**: Example implementation of atomic transactions spanning multiple repositories (see `OrderUseCase`).
- **Configuration Management**: Environment variable handling with `godotenv`.
- **Middleware**: Error handling, logging, panic recovery, and CORS.

## 🛠️ Tech Stack

- **Language**: Go (1.24+)
- **Web Framework**: Fiber v2
- **Database**: PostgreSQL
- **ORM**: GORM
- **Authentication**: JWT (JSON Web Tokens) - *Structure ready*
- **Encryption**: Bcrypt for password hashing

## 📂 Project Structure

```
.
├── cmd
│   └── api
│       └── main.go           # Application entry point
├── config                    # Configuration load logic
├── internal
│   ├── delivery              
│   │   └── http              # HTTP handlers and routers (Delivery Layer)
│   ├── domain                # Entities and interfaces (Domain Layer)
│   ├── infrastructure        
│   │   ├── database          # DB connection & migrations
│   │   └── persistence       # Repository implementations (Infrastructure Layer)
│   └── usecase               # Business logic (Usecase Layer)
├── pkg                       # Shared packages / utils
└── .env                      # Environment variables
```

## ⚡ Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL database

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. **Setup Environment Variables**
   Copy the example environment file:
   ```bash
   cp .env.example .env
   ```
   Edit `.env` and configure your database credentials (DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, etc.).

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

### Running the Application

1. **Run locally**
   ```bash
   go run cmd/api/main.go
   ```
   The server will start on port `8080` (or as defined in .env). Auto-migration will create necessary database tables.

## 🔗 API Endpoints

### Auth / Users
- `POST /api/v1/users/register` - Register new user
- `POST /api/v1/users/login` - Login user
- `GET /api/v1/users/:id` - Get user profile

### Products
- `GET /api/v1/products` - List all products
- `POST /api/v1/products` - Create a product
- `GET /api/v1/products/:id` - Get product details
- `PUT /api/v1/products/:id` - Update product
- `DELETE /api/v1/products/:id` - Delete product

### Orders
- `POST /api/v1/orders` - Create a new order (Transactional)
- `GET /api/v1/orders/:id` - Get order details
- `GET /api/v1/orders/user/:user_id` - List orders for a user

## 🧪 Testing
Coming soon...

## 🤝 Contributing\
Contributions are welcome! Please feel free to submit a Pull Request.
