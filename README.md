# Go RSS Aggregator

A high-performance RSS feed aggregator backend built with Go and PostgreSQL. This application allows users to create accounts, follow RSS feeds, and automatically fetch/store posts from those feeds for offline reading or centralizing news sources.

## Features

- **User Management**: Create users and authenticate using API keys.
- **Feed Management**: Add new RSS feeds to the system.
- **Feed Following**: Follow/unfollow feeds to personalize your post feed.
- **Automated Scraper**: A background worker that periodically fetches the latest posts from all added feeds.
- **RESTful API**: Clean API endpoints for interacting with users, feeds, and posts.
- **Type-Safe Database Access**: Built using `sqlc` for compile-time safe SQL queries.

## Tech Stack

- **Language**: [Go](https://go.dev/)
- **Router**: [go-chi/chi](https://github.com/go-chi/chi)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **SQL Toolkit**: [sqlc](https://sqlc.dev/)
- **Migrations**: [goose](https://github.com/pressly/goose)
- **Containerization**: Docker (for database)

## Getting Started

### Prerequisites

- Go 1.22+
- Docker and Docker Compose
- `goose` (for database migrations)

### Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/abhyuday-fr/rss-aggregator.git
   cd go-rss-aggregator
   ```

2. **Environment Configuration**:
   Create a `.env` file in the root directory:
   ```env
   PORT=8080
   DB_URL=postgres://postgres:1709@localhost:5432/rssagg?sslmode=disable
   ```
   *(Note: The default password in `docker-compose.yml` is `1709`)*

3. **Start the Database**:
   ```bash
   cd sql/schema
   docker-compose up -d
   ```

4. **Run Migrations**:
   Make sure you have `goose` installed, then run:
   ```bash
   cd sql/schema
   make migrate-up
   ```
   *Or manually:*
   ```bash
   goose postgres <YOUR_DB_URL> up
   ```

5. **Build and Run**:
   ```bash
   go build -o rss-aggregator && ./rss-aggregator
   ```

## API Endpoints

### Public Endpoints
- `GET /v1/healthz`: Check if the server is running.
- `GET /v1/err`: Test error handling.
- `POST /v1/users`: Create a new user (`{"name": "Your Name"}`).
- `GET /v1/feeds`: List all available feeds.

### Authenticated Endpoints
*Include `Authorization: ApiKey <your_api_key>` in the headers.*

- `GET /v1/users`: Get current user details.
- `POST /v1/feeds`: Create a new feed (`{"name": "Feed Name", "url": "https://example.com/rss"}`).
- `GET /v1/posts`: Get the latest posts from feeds you follow.
- `POST /v1/feed_follows`: Follow a feed (`{"feed_id": "uuid"}`).
- `GET /v1/feed_follows`: List feeds you are following.
- `DELETE /v1/feed_follows/{feedFollowID}`: Unfollow a feed.

## Database Schema

The database consists of the following tables:
- `users`: Stores user information and API keys.
- `feeds`: Stores RSS feed metadata.
- `feed_follows`: Maps users to the feeds they follow.
- `posts`: Stores individual posts fetched from RSS feeds.

