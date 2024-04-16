# URL Shortener

This is a simple URL shortener project implemented in Go, featuring SQLite as the database and Chi as the router.

## Features

- Shorten long URLs into manageable links.
- Retrieve original URLs from shortened links.
- Delete unnecessary shortened links.
- Middleware for request identification and logging.
- Functional and unit testing for robustness.

## Configuration

The project utilizes a `config.yaml` file for configuration:

```yaml
env: "local"
storage_path: "./storage/storage.db"
http_server:
  address: "localhost:8082"
  timeout: 4s
  idle_timeout: 30s
```

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Alhanaqtah/url-shortener.git
   ```
2. Navigate to the project directory:
   ```bash
   cd url-shortener
   ```
3. Build and run the project:
   ```bash
   make run
   ```
4. The API will be accessible at [http://localhost:8080](http://localhost:8080).

## Getting Started

To start using the URL shortener API, you can use tools like Postman or cURL to make HTTP requests to the provided endpoints.

For example, to shorten a new URL:

**POST** http://localhost:8080/save

Request Body:
```json
{
  "url": "https://example.com/very/long/url/that/needs/shortening"
}
```
