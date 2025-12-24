# RSS Feed Service

A robust backend service designed to aggregate blog posts from various RSS feeds and expose them through a clean, well-documented HTTP API. Built with Go and PostgreSQL, it focuses on reliability, type-safety, and ease of deployment.

---

## Overview

Modern information consumption often requires monitoring dozens of disparate sources. This project provides a centralized engine to automate that process. It doesn't just fetch data; it cleans, persists, and serves it through a structured API, making it an ideal foundational component for blog readers or data analysis tools.

### Core Capabilities
- **Automated Ingestion**: Seamlessly handles XML-to-JSON conversion for various RSS formats.
- **Relational Persistence**: Stores blog and post metadata in a highly available PostgreSQL database.
- **Type-Safe Access**: Leverages SQLC to ensure that every database interaction is validated at compile-time.
- **Background Synchronization**: A dedicated worker periodically refreshes feeds without manual intervention.
- **Container-First**: Optimized for Docker, ensuring consistent behavior from local development to production.

---

## Architecture

The service is built on a clean, layered architecture that separates network concerns from business logic and data persistence.

```mermaid
graph TD
    User["User/Client"] --> API["HTTP API (chi)"]
    API --> Handlers["Handlers"]
    Handlers --> DB["Database Layer (SQLC)"]
    DB --> Postgres[("PostgreSQL")]
    Scraper["Background Scraper"] --> RSS["External RSS Feeds"]
    Scraper --> DB
```

### Design Principles
- **Separation of Concerns**: Each layer—HTTP, logic, and database—has a single, well-defined responsibility.
- **Scalability**: The modular design allows you to easily plug in new feed parsers or background tasks.
- **Reliability**: Comprehensive error handling at the handler level ensures the API remains responsive even if external feeds fail.

---

## Technical Stack

- **Language**: Go
- **Database**: PostgreSQL
- **SQL Generation**: SQLC (Type-safe SQL)
- **Migrations**: Goose
- **Containerization**: Docker & Docker Compose
- **Routing**: Chi (Lightweight and fast)

---

## Project Structure

```text
rss-feed/
├── crons/           # Background scheduling and ticker logic
├── handlers/        # HTTP request fulfillment and endpoint logic
├── internal/        # Internal database access and SQLC generated code
├── models/          # Shared domain entities
├── sql/             # SQL migrations and query definitions
├── utils/           # Shared helper functions (JSON responses, errors)
├── main.go          # Application entry point and router configuration
├── Dockerfile       # Container image definition
└── docker-compose.yml # Service orchestration
```

---

## Getting Started

### Prerequisites
- Go 1.20+
- Docker & Docker Compose
- PostgreSQL (or use the provided containerized setup)

### Installation & Setup

1. **Clone the Project**
   ```bash
   git clone https://github.com/MohamedAklamaash/rss-feed.git
   cd rss-feed
   ```

2. **Environment Configuration**
   Create a `.env` file in the root based on your environment:
   ```env
   PORT=8080
   DATABASE_URL=postgres://aklamaash:akla123@localhost:5432/blogs?sslmode=disable
   ```

3. **Infrastructure Launch**
   Launch PostgreSQL and pgAdmin using Docker Compose:
   ```bash
   docker-compose up -d
   ```

4. **Database Initialization**
   Apply migrations using Goose:
   ```bash
   go install github.com/pressly/goose/v3/cmd/goose@latest
   goose -dir sql/schema postgres "$DATABASE_URL" up
   ```

5. **Code Generation**
   Generate the database interface with SQLC:
   ```bash
   go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
   sqlc generate
   ```

6. **Start the Service**
   ```bash
   go run .
   ```

---

## API & Usage

All endpoints are versioned under `/v1`. Below are the primary ways to interact with the service.

### Example: Creating a User
```bash
curl -X POST http://localhost:8080/v1/user/create \
     -H "Content-Type: application/json" \
     -d '{"name": "Jane Doe"}'
```

### Endpoint Reference

#### System
- `GET /v1/healthz`: Health check endpoint.
- `GET /v1/errorz`: Error simulation for testing middleware.

#### Users & Authentication
- `POST /v1/user/create`: Register a new user and receive an API key.
- `GET /v1/user/getuser`: Retrieve the current user's profile (Requires `Authorization: ApiKey <key>`).

#### Feeds & Posts
- `POST /v1/feed/create`: Register a new source URL.
- `GET /v1/feed/all`: View all globally tracked feeds.
- `GET /v1/feed/posts`: Fetch aggregated posts from all followed feeds.
- `GET /v1/user/feed/follow`: Subscribe to a specific feed.

---

## Contributing

We value community contributions. To contribute:
1. Fork the repository.
2. Create a feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes and push to the branch.
4. Open a Pull Request for review.

---

## License

This project is licensed under the MIT License.

## Author

**Mohamed Aklamaash M.R**
Backend Engineer focused on Go, Systems, and Data Engineering.
