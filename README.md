# Go RSS Aggregator

A robust, high-performance RSS feed aggregator backend built with Go and PostgreSQL. This application provides a RESTful API to manage users, feeds, and followings, complemented by a concurrent background worker that automatically scrapes and stores the latest posts from subscribed feeds.

## 🚀 Features

- **User Authentication**: Secure API key-based authentication.
- **Feed Management**: Support for adding and listing RSS feeds.
- **Feed Following**: Personalized feed subscriptions for each user.
- **Concurrent Scraper**: Background worker utilizing goroutines to efficiently fetch and parse RSS feeds.
- **Deduplication**: Intelligent post storage that avoids duplicate entries.
- **Type-Safe DB**: Compile-time safe SQL queries using `sqlc`.
- **Database Migrations**: Managed schema changes with `goose`.

## 🛠 Tech Stack

- **Backend**: [Go](https://go.dev/) (v1.22+)
- **Router**: [go-chi](https://github.com/go-chi/chi)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **SQL Toolkit**: [sqlc](https://sqlc.dev/)
- **Migrations**: [goose](https://github.com/pressly/goose)
- **Environment**: [godotenv](https://github.com/joho/godotenv)

## 📁 Project Structure

```text
.
├── internal/
│   ├── auth/         # Authentication logic
│   └── database/     # SQLC generated type-safe database code
├── sql/
│   ├── queries/      # SQL query definitions for SQLC
│   └── schema/       # Database migrations (goose)
├── handler_*.go      # HTTP handlers for different resources
├── main.go           # Application entry point & router setup
├── scraper.go        # Background RSS scraping logic
├── rss.go            # RSS XML parsing utilities
└── models.go         # API response models
```

## 🏁 Getting Started

### Prerequisites

- Go 1.22+
- Docker (for PostgreSQL)
- [goose](https://github.com/pressly/goose) installed locally

### Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/abhyuday-fr/go-rss-aggregator.git
   cd go-rss-aggregator
   ```

2. **Environment Setup**
   Create a `.env` file in the root:
   ```env
   PORT=8080
   DB_URL=postgres://postgres:password@localhost:5432/rssagg?sslmode=disable
   ```

3. **Database Setup**
   Run the database container:
   ```bash
   cd sql/schema
   docker-compose up -d
   ```
   Apply migrations:
   ```bash
   make migrate-up
   ```

4. **Run the Application**
   ```bash
   go build -o rss-aggregator && ./rss-aggregator
   ```

## 📡 API Endpoints

### Public Endpoints
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/healthz` | Health check |
| GET | `/v1/err` | Error test |
| POST | `/v1/users` | Register a new user |
| GET | `/v1/feeds` | List all system feeds |

### Authenticated Endpoints
*Requires header: `Authorization: ApiKey <YOUR_API_KEY>`*

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/v1/users` | Get current user info |
| POST | `/v1/feeds` | Create a new feed |
| POST | `/v1/feed_follows` | Follow a feed |
| GET | `/v1/feed_follows` | List followed feeds |
| DELETE | `/v1/feed_follows/{id}`| Unfollow a feed |
| GET | `/v1/posts` | Get posts from followed feeds |

## 🔄 Background Scraper

The application starts a background worker that:
1. Identifies feeds that haven't been fetched recently.
2. Fetches them concurrently using a configurable number of goroutines.
3. Parses the RSS XML and saves new posts to the database.
4. Automatically skips already existing posts to avoid duplicates.

## 🤝 Contributing

Feel free to open issues or submit pull requests to improve the project!
