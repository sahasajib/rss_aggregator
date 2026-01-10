# rss_aggregator

Lightweight RSS aggregator / API written in Go.

Features
- User creation + API key authentication (header: `Authorization: ApiKey <key>`)
- Create and list feeds
- Follow/unfollow feeds
- Background scraper that fetches feeds and exposes posts via API
- PostgreSQL + SQL schema + generated queries (sqlc)

Prerequisites
- Go 1.18+ installed
- PostgreSQL database
- (Optional) `sqlc` if you want to regenerate typed DB queries

Environment
The server reads configuration from environment variables (you can use a `.env` file).
- `PORT` — port to listen on (example: `8080`)
- `DB_URL` — PostgreSQL connection string (example: `postgres://user:pass@localhost:5432/dbname?sslmode=disable`)

Quick start
1. Clone the repo

   git clone <repo-url>
   cd rss_aggregator

2. Create a `.env` file or export env vars. Example `.env`:

   PORT=8080
   DB_URL=postgres://user:pass@localhost:5432/rssdb?sslmode=disable

3. Install dependencies

   go mod download

4. Create the database schema

You can apply the SQL files in `sql/schema` to your database. Example:

```bash
for f in sql/schema/*.sql; do psql "$DB_URL" -f "$f"; done
```

5. (Optional) Generate DB code if you modify queries

   sqlc generate

6. Run the server

   go run .

The server will start and the background scraper will run periodically.

API
Base path: `/v1`

- `GET /v1/ready` — readiness check
- `GET /v1/err` — error test endpoint

User
- `POST /v1/users` — create a user (returns user with API key)
- `GET /v1/users` — get user info (requires `Authorization: ApiKey <key>`)

Feeds
- `POST /v1/feeds` — create a feed (requires API key)
- `GET /v1/feeds` — list available feeds

Feed follows
- `POST /v1/feed_follows` — follow a feed (requires API key)
- `GET /v1/feed_follows` — list followed feeds (requires API key)
- `DELETE /v1/feed_follows/{id}` — unfollow a feed (requires API key)

Posts
- `GET /v1/posts` — list posts for the authenticated user (requires API key)

Auth
The service expects the API key in the `Authorization` header using the `ApiKey` scheme, e.g.:

  Authorization: ApiKey 0123456789abcdef

Development notes
- Handlers are implemented in files like `handler_user.go`, `handler_feed.go`, and `handler_feed_follow.go`.
- DB queries are under `sql/` and generated code lives under `internal/database`.
- Background scraping logic is in `scraper.go` and starts automatically in `main.go`.

Contributing
Feel free to open issues and pull requests. For DB changes, update the `sql/schema` and `sql/queries` and run `sqlc generate`.


