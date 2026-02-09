# Golang Simple Authentication System

This project is a demonstration a complete Authentication System (Login/Register/Profile) using modern Golang practices.

## 🛠 Tech Stack

- **Backend:** Go (Golang) 1.25.6, Gin Framework, GORM
- **Database:** PostgreSQL 18 (via Docker)
- **Frontend:** HTML5, Vanilla JavaScript
- **DevOps:** Docker & Docker Compose
- **Testing:** Testify (Unit Testing with Mocking)

---

## 📂 Project Structure (Clean Architecture)

The project is organized to separate concerns and ensure maintainability:

```text
.
├── backend
│   ├── cmd/api          # Entry point (main.go)
│   ├── internal
│   │   ├── domain       # Data models & interfaces
│   │   ├── handler      # HTTP Transport layer (Gin)
│   │   ├── service      # Business logic (Hashing, JWT)
│   │   ├── repository   # Database access layer (GORM)
│   │   └── middleware   # Auth middleware (JWT Validation)
│   ├── pkg
│   │   ├── database     # DB Connection logic
│   │   └── utils        # Helper functions
│   └── go.mod
├── frontend             # User Interface (HTML/JS)
├── docker-compose.yml   # Database orchestration
├── .env                 # Environment variables
└── README.md

## 📚 Documentation
- [Setup & Installation Guide](./SETUP.md)
- [API Documentation](./API_DOC.md)