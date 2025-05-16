# UrlShorter

A simple URL shortener service written in Go. This project allows you to create short aliases for long URLs and redirect users to the original URLs using those aliases. It uses SQLite for storage and provides a RESTful API.

## Features
- Shorten long URLs with custom or auto-generated aliases
- Redirect to original URLs using short aliases
- Basic authentication for API endpoints
- Configurable via YAML file
- Simple SQLite storage backend

## Getting Started

### Prerequisites
- Go 1.20 or higher

### Installation
1. Clone the repository:
   ```sh
   git clone https://github.com/Lacky1234union/UrlShorter.git
   cd UrlShorter
   ```
2. Build the project:
   ```sh
   go build -o url-shortener ./cmd/url-shortener
   ```

### Configuration
Edit the `config/local.yaml` file to set up your environment:
```yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
  address: "localhost:8082"
  timeout: 4s
  idle_timeout: 30s
  user: "my_user"
  password: "my_pass"
```

### Running the Service
Start the server:
```sh
./url-shortener
```
The service will run on the address specified in your config (default: `localhost:8082`).

## API Usage

### 1. Shorten a URL
**Endpoint:** `POST /url`

**Auth:** Basic Auth (`user` and `password` from config)

**Request Body:**
```json
{
  "url": "https://example.com/very/long/url",
  "alias": "customalias" // optional
}
```
If `alias` is omitted, a random alias will be generated.

**Response:**
```json
{
  "status": "ok",
  "alias": "customalias"
}
```

### 2. Redirect to Original URL
**Endpoint:** `GET /{alias}`

Redirects to the original URL associated with the alias.

**Example:**
```
GET http://localhost:8082/youralias
```

If the alias is not found, a JSON error is returned.

## Testing
Run tests with:
```sh
go test ./tests
```

## License
MIT
