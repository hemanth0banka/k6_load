# k6 Load Tester

A custom load testing platform inspired by k6, built using Go for the backend and React for the frontend.
This project demonstrates how a real load testing tool works internally, including script management, concurrent execution, and result analysis.

---

1. ## 📌 Project Objective

The goal of this project is to:

- Understand the internal working of load testing tools like k6
- Build a simplified k6-like system without using the actual k6
- Apply real-world backend architecture principles
- Implement concurrency and performance measurement in Go

---

2. ## 🧩 Key Features

- Create and manage load test scripts
- Configure load tests (virtual users, iterations)
- Execute load tests concurrently using goroutines
- Collect and store execution metrics
- View test execution history
- REST API–driven backend
- Clean, layered backend architecture

---

3. ## 🛠️ Tech Stack

### Backend

- Language: Go
- Architecture: Layered / MVC
- API Style: REST (HTTP)
- Concurrency: Goroutines
- Storage: File-based & in-memory repositories

### Frontend

- Framework: React
- HTTP Client: Axios
- Build Tool: Vite
---
4. ### High-Level Architecture
```
Frontend (React)
        ↓ HTTP (REST APIs)
Backend (Go API Server)
        ↓
Load Engine (Concurrent Execution)
        ↓
Repositories (File / Memory Storage)
```
---

5. ### Project Structure
```
k6_Load_Tester/
│
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── handlers/      # HTTP controllers
│   │   ├── middleware/    # CORS & middleware
│   │   ├── router/        # API routes
│   │   ├── service/       # Business logic
│   │   ├── engine/        # Load execution engine
│   │   ├── generator/     # Script & execution generators
│   │   ├── repository/    # Data access layer
│   │   └── model/         # Domain models
│   │
│   ├── scripts/           # Stored load test scripts
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.js
│
└── design.md
```
---

6. ## Application Flow

- User creates a script from the frontend
- Script is stored via backend repository
- User creates a test configuration
- Test execution is triggered
- Load engine runs concurrent requests
- Metrics are collected during execution
- Results are stored
- Test history is displayed in UI

---

7. ## API Endpoints

### Script APIs
```
Method	Endpoint	      Description
POST	  /scripts	      Create a new script
GET	    /scripts	      Get all scripts
GET	    /scripts/{id}	  Get script by ID
```
### Test APIs
```
Method	Endpoint	      Description
POST	  /tests	        Create a new test
POST	  /tests/{id}/run	Run a specific test
```
### History APIs
```
Method	Endpoint	Description
GET	    /history	Get test execution history
```
8. ## Design Principles Used

- Separation of Concerns
- Single Responsibility Principle
- Open / Closed Principle
- Liskov Substitution Principle
- Interface Segregation Principle
- Dependency Inversion Principle

---

9. ## How to Run the Project

### Backend
```
cd backend
go run cmd/server/main.go
```

### Frontend
```
cd frontend
npm install
npm run dev
```
10. ### Conclusion
This project demonstrates a real-world backend architecture and a custom load testing engine inspired by k6.
While simplified, it shows strong understanding of:
- Backend system design
- Concurrency in Go
- Clean API architecture
- Full-stack integration
