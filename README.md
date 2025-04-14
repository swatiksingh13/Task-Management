# Task Manager 📝

A full-stack Task Manager application built with:

- 🚀 **Backend:** Golang (Gin framework, GORM, SQLite)
- 🌐 **Frontend:** React.js, Axios
- 📁 **Database:** SQLite (for easy setup)
- ⚙️ **API Communication:** RESTful APIs
- 🧪 **Extras:** Logging and Unit Test Cases for Backend

---
## 🚀 Getting Started

### 🔧 Backend (Golang)

1. Navigate to the backend folder:

cd backend
Run the Go server: go run main.go
Make sure you have Go installed and modules initialized (go mod init, go mod tidy)

💻 Frontend (React)
Navigate to the frontend folder: cd frontend
Install dependencies:npm start

Axios base URL should point to your backend (e.g., http://localhost:8081)

---
✨ Features
Create, Read, Update, Delete (CRUD) tasks

Set due date and status (Pending, In-progress, Completed)

Responsive UI

Full-stack integration

---

🛠️ API Endpoints (CRUD Examples)
🔹 Create Task
POST : http://localhost:8081/tasks
```json
{
  "title": "Write README",
  "description": "Document everything clearly",
  "status": "Pending",
  "due_date": "2025-04-20T12:00:00Z"
}
```

🔹 Get All Tasks
GET: http://localhost:8081/tasks
```json
Response:
  {
    "id": 1,
    "title": "Write README",
    "description": "Document everything clearly",
    "status": "Pending",
    "due_date": "2025-04-20T12:00:00Z"
  }
]
````

🔹 Update Task
PUT : http://localhost:8081/tasks/{id}

```json
{
  "title": "Write README",
  "description": "Update with API details",
  "status": "In-Progress",
  "due_date": "2025-04-21T14:00:00Z"
}
```

🔹 Delete Task
DELETE http://localhost:8081/tasks/{id}

---
🎥 Demo Screens / Videos 
[📽️ Click to watch the demo video] 
-- (https://drive.google.com/drive/folders/1jLpzxIFcMDmWjfk-X_98uAm9-fq3IB15?usp=drive_link)

---
## 💽 SQLite DB Commands

To view and interact with the SQLite database used in the backend, follow these steps:

1. Open the SQLite database: sqlite3 taskmanager.db
2. List all tables: .tables
3. View all tasks:SELECT * FROM tasks;
   
---
🧠 Tech Stack
Technology	Description
Golang + Gin	Web server and API routing
GORM + SQLite	ORM and lightweight DB
React.js	Frontend framework
Axios	API calls from React
VS Code	Development IDE

Thank you for checking out the Task Manager project! Happy Coding 😃
