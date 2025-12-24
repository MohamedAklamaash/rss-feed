# RSS Feed Service

> A backend service that fetches top blog posts from RSS feeds and exposes them through a clean HTTP API.

---

## 📌 Overview

This project is a **Go-based RSS Feed Aggregator** that:
- Fetches blog posts from multiple RSS feed sources
- Parses and stores them in **PostgreSQL**
- Serves the data via HTTP endpoints
- Uses **SQLC** for type-safe database access
- Is fully containerized using **Docker**

The goal of this project is to demonstrate clean backend architecture, data ingestion, and API design using Go.

---

## ✨ Features

- Fetch and parse RSS feeds (XML)
- Store blog metadata in PostgreSQL
- Expose feed data via REST API
- SQLC-generated database queries
- Docker & Docker Compose support
- Cron-based feed refresh support

---

## 🏗️ Architecture

The application follows a **layered architecture** to keep responsibilities clear and the codebase maintainable.

┌─────────────────────┐
│ HTTP Server │
│ (main.go) │
└─────────┬───────────┘
│
▼
┌─────────────────────┐
│ Handlers │
│ (HTTP Routes) │
└─────────┬───────────┘
│
▼
┌─────────────────────┐
│ Business Logic │
│ (RSS Fetch & Parse)│
└─────────┬───────────┘
│
▼
┌─────────────────────┐
│ Database Layer │
│ (Postgres + SQLC) │
└─────────────────────┘


### 🔹 Why this Architecture?

- **Separation of Concerns** – HTTP, business logic, and DB access are independent
- **Scalability** – Easy to add more feeds or background workers
- **Maintainability** – Each layer has a single responsibility
- **Safety** – SQLC ensures compile-time query validation

---

## ⚙️ How It Works

1. **RSS Feed Source**
    - One or more RSS feed URLs are configured as sources.

2. **Fetcher**
    - The service fetches RSS XML data from these URLs.
    - XML is parsed to extract blog title, link, publish date, etc.

3. **Persistence**
    - Parsed feed items are stored in PostgreSQL using SQLC-generated queries.

4. **API Layer**
    - Clients call HTTP endpoints to retrieve the latest blog posts.
    - Data is returned in structured JSON format.

5. **Cron Jobs**
    - Scheduled jobs can periodically refresh feeds to keep data up-to-date.

---

## 📁 Project Structure

rss-feed/
├── .github/workflows/ # CI/CD pipelines
├── crons/ # Scheduled jobs for feed refresh
├── handlers/ # HTTP request handlers
├── internal/database/ # SQLC generated DB code
├── models/ # Data models
├── sql/ # SQL migrations and queries
├── utils/ # Helper utilities
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go # Application entry point


---

## 🛠️ Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **ORM:** SQLC (type-safe SQL)
- **Containerization:** Docker, Docker Compose
- **Migrations:** Goose
- **API:** REST

---

## 🚀 Getting Started

### Prerequisites

- Go 1.20+
- Docker & Docker Compose
- PostgreSQL

---

### 🔧 Local Setup

#### 1️⃣ Clone the repository
```bash
git clone https://github.com/MohamedAklamaash/rss-feed.git
cd rss-feed

```
## **Spin up a database container**

```
docker run -d \
  --name rss-blogs \
  -e POSTGRES_USER=aklamaash \
  -e POSTGRES_PASSWORD=akla123 \
  -e POSTGRES_DB=blogs \
  -p 5432:5432 \
  postgres
```

## PGAdmin

```
docker run -d \
  --name pgadmin \
  -e PGADMIN_DEFAULT_EMAIL=admin@local.dev \
  -e PGADMIN_DEFAULT_PASSWORD=admin123 \
  -p 5050:80 \
  dpage/pgadmin4:latest

```

## Install db tools

```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

```
go get github.com/sqlc-dev/sqlc/cmd/sqlc@lates
```

## *Goose is for tracking migrations*

## *SQLC is interesting*

we can set the sqlc.yaml file in the root of the project, set folder for 
queries and schema and when we run sqlc generate we get all base boiler-plate code

🤝 Contributing

Contributions are welcome!

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Open a Pull Request

📄 License

This project is licensed under the MIT License

✍️ Author

Mohamed Aklamaash M.R

Backend Engineer | Go | Systems | Data
