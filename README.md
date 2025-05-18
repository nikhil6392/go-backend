# Go Backend Starter – WIP 🚧

This is a simple, lightweight backend starter built with **Go**, **Gin**, **pgx**, and **PostgreSQL (Neon)**.  
I'm building this project as a learning exercise and plan to expand it gradually with full authentication, Docker support, Kubernetes, Redis, and more.

---

## ✅ Features So Far

- 🌐 REST API with [Gin](https://github.com/gin-gonic/gin)
- 🛢️ PostgreSQL database hosted on [Neon](https://neon.tech/)
- ⚡ DB connection via [pgx](https://github.com/jackc/pgx)
- 🔒 Environment variables via `.env` with [godotenv](https://github.com/joho/godotenv)
- 🧪 Tested with Postman
- 🧱 Basic user model and signup route
- ✅ DB status and health check routes


## 📦 Setup Instructions

### 1. Clone the repo
```bash
git clone <repo-url>
cd <project-name>
```
### 2. Install Dependencies
```
go mod tidy
```

### 3. Add in .env File
```
DATABASE_URL=postgres://<user>:<password>@<host>/<db>?sslmode=require
```

### 4. Create User Table in Neon Manually

```
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 5. Run the server
```
go run main.go
```



