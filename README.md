# go_crud

Go (Gin) + MySQL (Docker) Todo CRUD API.

## Requirements
- Go
- Docker / Docker Compose

## Setup & Run

### 1) Start MySQL
```bash
docker compose up -d
```

### 2) Run API
```bash
DB_HOST=127.0.0.1 \
DB_PORT=3307 \
DB_USER=app \
DB_PASS=app \
DB_NAME=go_crud \
go run .
```

API listens on `http://localhost:3003`.

## Endpoints (curl)

### Health
```bash
curl -i http://localhost:3003/health
```

### List todos
```bash
curl -i http://localhost:3003/todos
```

### Create todo
```bash
curl -i -X POST http://localhost:3003/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"first todo"}'
```

### Update done
```bash
curl -i -X PUT http://localhost:3003/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"done":true}'
```

### Delete todo

- success: 204
```bash
curl -i -X DELETE http://localhost:3003/todos/1
```

- not found: 404
```bash
curl -i -X DELETE http://localhost:3003/todos/99999
```

- invalid id: 400
```bash
curl -i -X DELETE http://localhost:3003/todos/abc
```



