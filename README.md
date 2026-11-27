# smurl

smurl is a small URL shortener written in Go with Gin. It generates seven-character, case-sensitive codes using letters and digits, and reserves `healthz` and `metrics`. Links are stored in memory and expire when the server stops.

## Run locally

Install Go 1.26.6 or newer. From the repository root, run:

```sh
go run .
```

The server listens on port 8080 by default and builds short links with the fixed prefix `http://localhost:8080/`.

## API

| Method and path | Result |
| --- | --- |
| `POST /api/v1/urls` | `201 Created` with the new link |
| `GET /{code}` | `302 Found` with the original URL in the `Location` header |
| `GET /healthz` | `200 OK` with `{"status":"ok"}` |

### Create a short link

```sh
curl -i http://localhost:8080/api/v1/urls \
  -H 'Content-Type: application/json' \
  -d '{"original_url":"https://example.com/articles?tag=go"}'
```

Example response body:

```json
{
  "code": "aB3xY7z",
  "short_url": "http://localhost:8080/aB3xY7z",
  "original_url": "https://example.com/articles?tag=go",
  "created_at": "2026-11-27T09:00:00Z"
}
```

The server trims surrounding whitespace from `original_url`. The trimmed URL must use HTTP or HTTPS, include a host, and be at most 2,048 bytes long. `created_at` is a UTC timestamp in RFC 3339 format.

### Use a short link

Replace the example below with the `short_url` from your response. This sends a GET request and prints the redirect status and headers.

```sh
curl -sS -D - -o /dev/null http://localhost:8080/aB3xY7z
```

### Errors

| HTTP status | Error code | When it occurs |
| --- | --- | --- |
| `400 Bad Request` | `invalid_json` | The body cannot be decoded or `original_url` is missing or blank |
| `400 Bad Request` | `invalid_url` | The URL is too long, cannot be parsed, uses an unsupported scheme, or has no host |
| `404 Not Found` | `not_found` | The requested code is unknown |
| `500 Internal Server Error` | `internal_error` | Short-code generation fails |

API handlers return errors in this shape:

```json
{
  "error": {
    "code": "not_found",
    "message": "short_url not found"
  }
}
```

## Tests

```sh
go test -race ./...
```

The suite covers shortcode generation and reserved codes, duplicate store entries and lookups, and HTTP creation, validation, redirects, and unknown codes.
