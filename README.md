# Pokédex CLI - Interactive Pokémon Explorer #

This sophisticated command-line tool integrates with the PokéAPI, featuring intelligent caching, interactive navigation, and comprehensive Pokémon data management.

## 🚀 Key Features

- **Interactive CLI**: REPL-style interface with command completion and help system
- **HTTP API Integration**: Seamless interaction with the PokéAPI for real-time data
- **Intelligent Caching**: 5-minute TTL cache with automatic cleanup to optimize performance
- **Location Exploration**: Navigate through Pokémon world locations and areas
- **Concurrent Operations**: Thread-safe operations with goroutine-based cache management
- **Pokémon Management**: Bask in the glory once more as you catch, inspect, and manage your Pokémon collection

## 🎯 Skills Demonstrated

- **HTTP Client Development**: Custom clients with timeout management and error handling
- **Concurrent Programming**: Goroutines, mutexes, and thread-safe data structures
- **Caching Strategies**: TTL-based cache with automatic cleanup and memory management
- **API Integration**: Complex JSON parsing and HTTP response handling
- **CLI Design**: Interactive command-line interfaces with state management
- **Testing**: Unit test coverage for critical components
- **Memory Management**: Preventing memory leaks with proper resource cleanup

## 🏗️ Technical Architecture

### HTTP Client Design
- **Custom HTTP Client**: Configurable timeouts and connection management
- **Error Handling**: Comprehensive HTTP error detection and user feedback
- **JSON Processing**: Robust unmarshaling of complex API response structures

### Caching System
- **TTL-Based Cache**: 5-minute time-to-live with automatic expiration
- **Concurrent Safety**: `sync.Mutex` protection for thread-safe cache operations
- **Memory Management**: Background goroutine cleanup prevents memory leaks
- **Cache Optimization**: Reduces API calls and improves response times

### Command Architecture
- **Command Pattern**: Extensible command registration and execution system
- **REPL Interface**: Interactive read-eval-print loop with persistent state
- **State Management**: Location tracking and Pokémon collection persistence

## 💻 Technologies Used

- **Go 1.21+**: Modern Go with generics and advanced features
- **HTTP Client**: Standard library with custom configuration
- **JSON Processing**: Native encoding/json for API response handling
- **Concurrency**: Goroutines and mutexes for safe concurrent operations
- **Testing**: Comprehensive unit tests for cache functionality

## 🔧 Technical Highlights

### Intelligent Caching Implementation
```go
type Cache struct {
    entries map[string]cacheEntry
    mu      *sync.Mutex
}

// TTL-based cache with goroutine cleanup
func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    defer ticker.Stop()
    
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}
```

### Thread-Safe Operations
```go
// Concurrent-safe cache operations
func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    c.entries[key] = cacheEntry{
        createdAt: time.Now().UTC(),
        val:       val,
    }
}
```

### Complex JSON Unmarshaling
```go
// Nested API response structures
type LocationAreasResponse struct {
    Count    int    `json:"count"`
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}
```


## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Internet connection (for PokéAPI access)

### Installation & Running
```bash
# Clone and build
git clone https://github.com/DNewmanDev/pokedex.git
cd pokedex
go build -o pokedex

# Run interactive CLI
./pokedex
```

## 📋 Available Commands

### Navigation
- `help` - Display all available commands and usage
- `exit` - Quit the application
- `map` - Show next 20 location areas
- `mapb` - Show previous 20 location areas

### Activities
- `explore <area_name>` - Explore a specific location area
- `catch <pokemon_name>` - Attempt to catch a Pokémon
- `inspect <pokemon_name>` - View detailed Pokémon information
- `pokedex` - Display your caught Pokémon collection

## 🔍 Example Usage

```bash
$ ./pokedex
Welcome to the Pokedex!
Pokedex > help
Available commands:
- help: Displays a help message
- exit: Exit the Pokedex
- map: Display next locations
- mapb: Display previous locations
- explore <location>: Explore area
- catch <pokemon>: Catch a pokemon
- inspect <pokemon>: Inspect a pokemon
- pokedex: Show caught pokemon

Pokedex > map
Location areas:
- canalave-city-area
- eterna-city-area
- pastoria-city-area
...

Pokedex > explore canalave-city-area
Found Pokemon:
- tentacool
- tentacruel
- pelipper

Pokedex > catch tentacool
Throwing a Pokeball at tentacool...
tentacool was caught!

Pokedex > pokedex
Your Pokedex:
- tentacool
```

## 🏗️ Project Structure

```
pokedex/
├── internal/pokecache/    # Caching implementation with tests
├── *.go                   # Command handlers and HTTP client
├── go.mod                 # Go module configuration
└── pokecache_test.go     # Comprehensive cache testing
```

