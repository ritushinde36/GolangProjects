# 🎬 Go Movies CRUD API

A simple **CRUD REST API built in Go** to understand core backend concepts such as routing, HTTP handlers, request/response handling, and JSON encoding/decoding.

This is a personal learning project, but suggestions and improvements are welcome. 🤝

Please Note - This project intentionally keeps things **basic and in-memory** to focus on learning fundamentals rather than production concerns.

---

## 🚀 Features

- RESTful CRUD operations for movies
- In-memory data storage (no database)
- Gorilla Mux for routing
- JSON request & response handling
- Random seed data for quick testing

---

## 🛠️ Tech Stack

- **Language:** Go
- **Router:** Gorilla Mux
- **Data:** In-memory slice
- **Utilities:** `go-randomdata` (for mock data)

---

## 📁 Project Structure

```text
.
├── main.go
├── go.mod
└── go.sum
```
Everything is intentionally kept in a single file to make learning and understanding easier.

## 📌 Data Model

### Movie

```json
{
  "id": "123",
  "isbn": "45678",
  "title": "Movie Title",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

## 📌 API Endpoints

| Method | Endpoint        | Description               |
|--------|-----------------|---------------------------|
| GET    | `/movies`       | Get all movies            |
| GET    | `/movies/{id}`  | Get a single movie        |
| POST   | `/movies`       | Create a new movie        |
| PUT    | `/movies/{id}`  | Update an existing movie  |
| DELETE | `/movies/{id}`  | Delete a movie            |


## ▶️ Running the Project

### 1. Clone this repository
```bash
git clone https://github.com/ritushinde36/GolangProjects.git
cd golang-movies-crud-api
```
### 2. Install dependencies
```bash
go mod download
```

### 3. Run the server
```bash
go run main.go
```

### 4. Server will start at:
```bash
Server will start at:
```

## ⚠️ Notes & Intentional Limitations

This project is built for learning purposes:

❌ No database

❌ No authentication

❌ No concurrency handling

❌ Minimal error handling

❌ No HTTP status codes

These are intentional omissions to keep the focus on understanding Go fundamentals and REST concepts.
