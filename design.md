# k6 Load Tester
---

## 1. Overview

This backend is a custom load testing platform inspired by k6, built using Go.
It follows a layered architecture with clear separation of concerns, making the system modular, testable, and scalable.
The backend exposes REST APIs for:

- Managing load test scripts
- Running load tests
- Fetching test execution history

---

## 2. Tech Stack

### Backend
- Language: Go
- Architecture: MVC Architecture
- APIs: REST (HTTP)

### Frontend (Assumed)
- React
- Axios for API calls

### Storage
- File-based storage
- In-memory storage (for testing)

---

## 3. System Components
```
 Frontend (React)
        ↓ HTTP
 Backend (Go API Server)
        ↓
 Load Engine (Concurrency + Execution)
        ↓
 Repositories (File / Memory)
```
---

## 4. Project Folder Structure
```
 k6_Load_Tester/
 │
 ├── cmd/
 │   └── server/
 │       └── main.go
 │
 ├── internal/
 │   │── handlers/
 │   │── middleware/
 │   ├── router/
 │   ├── service/
 │   ├── engine/
 │   ├── generator/
 │   ├── repository/
 │   └── model/
 │
 ├── scripts/
 ├── go.mod
 └── go.sum
```
---

## 5. Application Flow

- User creates a script from frontend
- Script is stored using repository
- User creates a test configuration
- Test execution is triggered
- Load engine runs concurrent requests
- Metrics are collected
- Results are stored
- History is returned to frontend

---

## 6. API Documentation

### Script APIs
- POST   /scripts         - creates a new script
- GET    /scripts         - returns all scripts
- GET    /scripts/{id}    - returns a specific script

### Test APIs
- POST   /tests           - Create a new test
- POST   /tests/{id}/run  - Run a specific test

### History APIs
- GET    /history         - Returns all test history

---

## 7. Design Principles Followed

- MVC Architecture
- Single Responsibility
- Open Closed
- Liskov Substitution
- Interface Segregation
- Dependency Inversion

---

## 8. Conclusion 

This backend is a custom load testing platform built using Go. It follows a layered architecture with clear separation of concerns, making the system modular, testable, and scalable.
