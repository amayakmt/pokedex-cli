# Pokedex CLI

A terminal-based Pokedex built in Go. Explore the Pokemon world through an interactive REPL — browse location areas, find wild Pokemon, attempt to catch them, and build your personal Pokedex. Powered by the [PokeAPI](https://pokeapi.co/).

---

## Features

- Interactive REPL with command parsing
- Browse paginated location areas (`map` / `mapb`)
- Explore a specific area to see wild Pokemon encounters
- Catch Pokemon with randomized catch mechanics (harder Pokemon escape more often)
- Inspect your caught Pokemon's stats and types
- In-memory response cache with a configurable TTL to reduce redundant API calls

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+

### Installation

```bash
git clone https://github.com/amayakmt/pokedex-cli.git
cd pokedex-cli
go build -o pokedex
./pokedex
```

Or run directly without building:

```bash
go run .
```

---

## Usage

Once launched, you will see the prompt:

```
Pokedex >
```

Type any command and press Enter.

### Commands

| Command | Description |
|---|---|
| `help` | Display all available commands |
| `exit` | Quit the Pokedex |
| `map` | Show the next 20 location areas |
| `mapb` | Show the previous 20 location areas |
| `explore <area-name>` | List all Pokemon found in a location area |
| `catch <pokemon-name>` | Attempt to catch a Pokemon |
| `inspect <pokemon-name>` | Show stats of a Pokemon you've caught |
| `pokedex` | List all Pokemon in your Pokedex |

### Example Session

```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
...

Pokedex > explore pastoria-city-area
Exploring pastoria-city-area...
Found Pokemon:
 - tentacool
 - tentacruel
 - shellos

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > inspect tentacool
Name: tentacool
Height: 9
Weight: 455
Stats:
  - hp: 40
  - attack: 40
  - defense: 35
  ...
Types:
  water
  poison

Pokedex > pokedex
Your Pokedex:
 - tentacool
```

---

## Catch Mechanics

When you throw a Pokeball, the catch attempt uses the Pokemon's `base_experience` from the API. A random number between `0` and `base_experience` is generated — if it exceeds `40`, the Pokemon escapes. This means high-experience Pokemon like legendaries are much harder to catch.

---

## Project Structure

```
.
├── main.go                        # Entry point and REPL loop
├── repl.go                        # Input parsing (cleanInput)
├── repl_test.go                   # Tests for cleanInput
├── commands.go                    # All command definitions and handlers
└── internal/
    ├── pokeapi/
    │   ├── client.go              # HTTP client wired with cache
    │   ├── location_areas.go      # GET /location-area (paginated list)
    │   ├── location_area.go       # GET /location-area/:name (single area)
    │   └── pokemon.go             # GET /pokemon/:name
    └── pokecache/
        ├── cache.go               # Thread-safe TTL cache
        └── cache_test.go          # Cache unit tests
```

---

## Caching

API responses are cached in memory by URL. The cache uses a background goroutine to evict entries older than the configured interval (default: 5 minutes). This avoids redundant network requests when navigating back to previously visited areas or re-inspecting the same Pokemon.

---

## Running Tests

```bash
go test ./...
```

Tests cover:
- `cleanInput` — whitespace trimming, lowercasing, and splitting
- `Cache.Add` / `Cache.Get` — basic storage and retrieval
- Cache reap loop — automatic eviction after TTL expires

---

## Dependencies

No third-party dependencies. Uses the Go standard library only.

---

## License

[MIT](LICENSE)
