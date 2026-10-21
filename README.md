# Memory Match

Memory Match is a Go + Gin + PostgreSQL card-matching game with Easy, Medium and Hard layouts, a server-enforced five-minute timer, and a leaderboard.

## Requirements

- Go version compatible with the `go` directive in `go.mod`
- PostgreSQL
- A `DATABASE_URL` environment variable
- All SQL migrations applied to the database

## Configure the database

Set the connection string for your local database. Example only; replace the username, password and database with your own:

```bash
export DATABASE_URL="postgres://YOUR_USER:YOUR_PASSWORD@localhost:5432/YOUR_DATABASE?sslmode=disable"
```

Apply the migrations in `db/migrations` in numerical order. If you already applied migrations 1 and 2, apply migration 3:

```bash
psql "$DATABASE_URL" -f db/migrations/000003_add_timer_leaderboard_support.up.sql
```

Migration 3 allows the `expired` game state and updates the dimensions constraint so new Easy games use the screenshot's 4 x 4 board. Legacy Easy games with the old 3 x 4 dimensions remain valid.

## Run

Run from the project root so the server can find `web/index.html`:

```bash
go run .
```

Then open [http://localhost:8082](http://localhost:8082).

## API

- `GET /health` — health check
- `POST /games` — create a game
- `GET /games/:id` — fetch game state and remaining time
- `POST /games/:id/flip` — flip a card
- `POST /games/:id/resolve` — hide the two cards after a mismatch
- `GET /leaderboard?limit=10` — leaderboard, with a limit from 1 to 100

Create a game example:

```json
{
  "player_name": "Player",
  "difficulty": "easy"
}
```

The API returns `time_limit_seconds`, `time_left_seconds`, and `elapsed_seconds`. The backend rejects card flips after the five-minute limit and records the game as `expired`. Only completed games are eligible for the leaderboard. Scores are ordered by fewest moves, then shortest completion time.

## Board sizes

| Difficulty | Board | Pairs |
|---|---:|---:|
| Easy | 4 x 4 | 8 |
| Medium | 4 x 5 | 10 |
| Hard | 5 x 6 | 15 |
