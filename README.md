# test2

A simple task management REST API built with Go.

## Getting Started

### Prerequisites

- [Go 1.24+](https://go.dev/dl/)

### Run the server

```bash
go run .
```

The server listens on port `8080` by default. Override with the `PORT` environment variable:

```bash
PORT=9090 go run .
```

## API

| Method | Path         | Description    |
|--------|--------------|----------------|
| GET    | /tasks       | List all tasks |
| POST   | /tasks       | Create a task  |
| GET    | /tasks/{id}  | Get a task     |
| PUT    | /tasks/{id}  | Update a task  |
| DELETE | /tasks/{id}  | Delete a task  |

### Examples

**Create a task**
```bash
curl -s -X POST http://localhost:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk","description":"2% milk"}'
```

**List tasks**
```bash
curl -s http://localhost:8080/tasks
```

**Update a task**
```bash
curl -s -X PUT http://localhost:8080/tasks/1 \
  -H 'Content-Type: application/json' \
  -d '{"title":"Buy milk","description":"2% milk","done":true}'
```

**Delete a task**
```bash
curl -s -X DELETE http://localhost:8080/tasks/1
```

## Testing

```bash
go test ./...
```
