# 📚 Go Books CRUD API

A simple **CRUD REST API built in Go** to understand core backend concepts such as routing, HTTP handlers, request/response handling, and JSON encoding/decoding.

This is a personal learning project, but suggestions and improvements are welcome. 🤝

Please Note - This project intentionally keeps things **basic and in-memory** to focus on learning fundamentals rather than production concerns.

---

## 🚀 Features

- RESTful CRUD operations for books
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
├── books/
│   └── books.go
└── operations/
    └── operations.go
```

The project is organized with separate packages for data models (`books`) and API operations (`operations`) for better code organization and maintainability.

## 📌 Data Model

### Book

```json
{
  "id": 12345,
  "title": "Book Title",
  "publication": "Publisher Name",
  "ratings": 5,
  "author": {
    "firstname": "Author",
    "lastname": "Name"
  }
}
```

## 📌 API Endpoints

| Method | Endpoint        | Description               |
|--------|-----------------|---------------------------|
| GET    | `/books`        | Get all books             |
| GET    | `/book/{id}`    | Get a single book         |
| POST   | `/books`        | Create a new book         |
| PUT    | `/book/{id}`    | Update an existing book   |
| DELETE | `/book/{id}`    | Delete a book             |


## ▶️ Running the Project

### 1. Clone this repository
```bash
git clone https://github.com/ritushinde36/GolangProjects.git
cd golang-books-crud-api
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
http://localhost:8000
```

## ⚠️ Notes & Intentional Limitations

This project is built for learning purposes:

❌ No database

❌ No authentication

❌ No concurrency handling

❌ Minimal error handling
