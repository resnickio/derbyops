# DerbyOps

## 📌 Overview
DerbyOps is an **open-source, API-first roller derby league management platform** designed for **self-hosting** or as a **SaaS solution**. It helps leagues **organize memberships, schedule events, track attendance, manage teams, and more!**

## 🚀 Features
- **Role-Based Access Control (RBAC)**: Different access levels for admins, staff, officials, and members.
- **API-First Design**: RESTful API with JWT authentication & API key support.
- **Scheduling & Attendance**: Calendar-based event tracking with attendance check-ins.
- **Team & Roster Management**: Customizable teams, positions, and bout tracking.
- **Gameday Staffing**: Assigning officials, referees, and volunteers.
- **Internationalization Support**: Built with i18n in mind.

## 📂 Project Structure
```
derbyops/
│── .github/workflows/      # CI/CD workflows
│── api/
│   ├── handlers/           # API route handlers
│   ├── auth/               # JWT & API key authentication
│   ├── middleware/         # Middleware (RBAC, rate limiting)
│   ├── models/             # Database models
│   ├── database/           # Database connection
│   ├── main.go             # Main Go application entrypoint
│── config/                 # Config files (env, Docker, Helm)
│── migrations/             # Database migrations
│── docs/                   # Documentation
│── tests/                  # API test cases
│── web/                    # (Future) Web frontend code
│── mobile/                 # (Future) iOS & Android code
│── README.md               # Project overview
│── .env.example            # Example environment variables
│── .gitignore              # Git ignore file
│── docker-compose.yml      # Local Docker setup
│── Dockerfile              # Docker build file
│── helm/                   # Helm charts for Kubernetes
```

## 🛠️ Tech Stack
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL
- **Frontend**: React (Future)
- **Mobile**: Swift (iOS) & Kotlin (Android) (Future)
- **Deployment**: Docker, Helm (Kubernetes support)

## 🔧 Installation & Setup
### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/your-org/derbyops.git
cd derbyops
```

### **2️⃣ Set Up Environment Variables**
Copy the `.env.example` file to `.env` and configure it:
```sh
cp .env.example .env
```

### **3️⃣ Run with Docker (Recommended)**
```sh
docker-compose up --build
```

### **4️⃣ Run Locally (Go)**
```sh
go run api/main.go
```

## 🛠️ Development
- **Linting**: `golangci-lint run`
- **Tests**: `go test ./tests`
- **Database Migrations**: `migrate -path migrations -database "postgres://user:pass@localhost/dbname?sslmode=disable" up`

## 🤝 Contributing
We welcome contributors! Feel free to submit PRs, file issues, and help improve DerbyOps.

## 📜 License
**MIT License** - Open source and free to use.

## 📬 Contact
For questions or support, open an issue or reach out via GitHub discussions.