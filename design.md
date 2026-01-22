# k6 Load Tester

## 1. Project Overview

This project is a **custom load testing backend** that I designed and developed in Go, inspired by tools like **k6**. The main goal of the project was to understand how load testing systems work internally—from handling concurrent requests to collecting metrics and storing execution history.

I structured the backend using a **layered (MVC-style) architecture** so that each part of the system has a clear responsibility. This makes the codebase easier to maintain, extend, and test as the project grows.

The backend exposes REST APIs that allow:

* Creating and managing load test scripts
* Running load tests with configurable concurrency
* Storing and viewing test execution history

---

## 2. Tech Stack Used

### Backend

* **Language:** Go
* **Architecture:** MVC / Layered Architecture
* **API Style:** REST (HTTP)

### Frontend (Planned / Assumed)

* **React** for UI
* **Axios** for API communication

### Storage

* **File-based storage** for persisting scripts and test results
* **In-memory storage** used during development and testing

---

## 3. High-Level System Design

```
 Frontend (React)
        ↓ HTTP Requests
 Go Backend API Server
        ↓
 Load Execution Engine (Concurrency)
        ↓
 Repository Layer (File / Memory)
```

The frontend communicates with the Go backend via REST APIs. The backend triggers the load engine, which executes concurrent requests and collects metrics. All data is stored through a repository layer, allowing storage implementation to be swapped easily.

---

## 4. Project Folder Structure

```
 k6_Load_Tester/
 │
 ├── cmd/
 │   └── server/
 │       └── main.go        # Application entry point
 │
 ├── internal/
 │   ├── handlers/          # HTTP request handlers (controllers)
 │   ├── middleware/        # Middleware (CORS, logging, etc.)
 │   ├── router/            # Route definitions
 │   ├── service/           # Business logic layer
 │   ├── engine/            # Load execution engine
 │   ├── generator/         # Request/script generators
 │   ├── repository/        # Data access layer (file / memory)
 │   └── model/             # Domain models
 │
 ├── scripts/               # Stored load test scripts
 ├── go.mod
 └── go.sum
```

This structure helped me keep the code clean and enforce separation of concerns across the application.

---

## 5. Application Flow

1. The user creates a load test script from the frontend
2. The script is stored using the repository layer
3. The user creates a test configuration
4. A test execution is triggered via API
5. The load engine runs concurrent HTTP requests
6. Execution metrics (requests, errors, timing) are collected
7. Test results are stored
8. Execution history is returned to the frontend

---

## 6. API Endpoints

### Script Management APIs

* **POST** `/scripts` – Create a new load test script
* **GET** `/scripts` – Fetch all scripts
* **GET** `/scripts/{id}` – Fetch a specific script

### Test Execution APIs

* **POST** `/tests` – Create a new test configuration
* **POST** `/tests/{id}/run` – Run a specific test

### History APIs

* **GET** `/history` – Fetch load test execution history

---

## 7. Design Principles Applied

While building this project, I consciously applied the following design principles:

* **MVC / Layered Architecture** for clarity and maintainability
* **Single Responsibility Principle** for cleaner components
* **Open–Closed Principle** to allow extension without modification
* **Liskov Substitution Principle** through interface-based design
* **Interface Segregation** to avoid large, tightly coupled interfaces
* **Dependency Inversion** to decouple business logic from storage and execution details

---

## 8. Conclusion

This backend is a custom load testing platform built using Go. It follows a layered architecture with clear separation of concerns, making the system modular, testable, and scalable.



