# 🛠 Setup & Installation Guide

Follow these steps to set up and run the **Authentication System** (Golang + PostgreSQL + Docker).

## 1. Prerequisites
Ensure you have the following installed on your machine:
- **Go** (v1.18 or higher) - [Download](https://go.dev/dl/)
- **Docker & Docker Compose** - [Download](https://www.docker.com/products/docker-desktop)
- **Git** (Optional, for cloning)

---

## 2. Environment Configuration
The project uses a single `.env` file at the root directory for both Docker and Backend configuration.

Create a file named `.env` in the root folder (`thai-bev-test/`) and paste the following:

```env
# --- Database Config ---
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=password123
DB_NAME=userdb

# --- Server Config ---
PORT=8080
JWT_SECRET=secretjaa1234
```

---

## 3. Start Database (Docker)
We use Docker Compose to spin up a PostgreSQL 16 container.

1. Open your terminal at the root directory.
2. Run the following command:

```bash
docker-compose up -d
```

3. Verify that the database is running:

```bash
docker ps
```

You should see a container named `thai_bev_db` running on port `5432`.

---

## 4. Run Backend Server (Golang)

1. Navigate to the backend directory:

```bash
cd backend
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run cmd/api/main.go
```

Output should indicate:

```
🔌 Connecting to: host=localhost port=5432 ...
🚀 Database Connected Successfully!
🔥 Server running on port 8080
```

---

## 5. Access Frontend (UI)

The frontend is built with pure HTML/JS and requires no installation.

1. Go to the `frontend` folder.
2. Open `index.html` in your web browser (Chrome/Edge recommended).
- Login Page: index.html (IT 02-1)
- Register Page: register.html (IT 02-2)
- Welcome Page: welcome.html (IT 02-3)

---

## 6. Run Unit Tests (Bonus) 🧪

To verify the business logic (Service Layer) with mocked dependencies:

```bash
cd backend
go test ./internal/service/... -v
```