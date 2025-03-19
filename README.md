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
│── .github/                # GitHub workflows (CI/CD)
│── api/                    # Backend API (Golang)
│   ├── handlers/           # API route handlers
│   ├── auth/               # Authentication (JWT, API Keys)
│   ├── middleware/         # Middleware (RBAC, rate limiting, logging)
│   ├── models/             # Database models (GORM)
│   ├── database/           # Database connection & migrations
│   ├── services/           # Business logic services
│   ├── routes/             # Route registration
│   ├── config/             # Config files (YAML, ENV parsing)
│   ├── tests/              # Unit & integration tests
│   ├── main.go             # Main application entrypoint
│── web/                    # Web frontend (React, Next.js)
│   ├── src/                # Source code
│   ├── components/         # Reusable UI components
│   ├── pages/              # Page-based routing (Next.js)
│   ├── hooks/              # React Hooks
│   ├── styles/             # CSS, Tailwind, etc.
│   ├── tests/              # Unit & integration tests
│   ├── package.json        # Frontend dependencies
│   ├── tsconfig.json       # TypeScript config
│   ├── next.config.js      # Next.js config
│── mobile/                 # Mobile app (React Native / Swift & Kotlin)
│   ├── ios/                # iOS app (Swift)
│   ├── android/            # Android app (Kotlin)
│   ├── src/                # React Native components (if hybrid approach)
│   ├── package.json        # Mobile dependencies
│   ├── tests/              # Unit & integration tests
│── infra/                  # Infrastructure & deployment
│   ├── docker/             # Docker configuration
│   ├── helm/               # Helm charts for Kubernetes
│   ├── terraform/          # IaC (if needed for cloud setup)
│── migrations/             # Database migration scripts
│── docs/                   # Documentation
│   ├── api/                # API documentation (OpenAPI / Swagger)
│   ├── dev/                # Development guidelines
│   ├── user/               # User guides
│   ├── contributing.md     # Contribution guidelines
│── tests/                  # Global tests (end-to-end, API, integration)
│── scripts/                # Utility scripts (DB resets, setup scripts)
│── .env.example            # Example environment variables
│── .gitignore              # Git ignore file
│── docker-compose.yml      # Local Docker setup
│── Dockerfile              # Backend Docker build file
│── README.md               # Project overview

```

## 🛠️ Tech Stack
- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL
- **Frontend**: React (Future)
- **Mobile**: Swift (iOS) & Kotlin (Android) (Future)
- **Deployment**: Docker, Helm (Kubernetes support)

## 📜 Requirements & Business Rules
### **1. User Roles & Permissions**
- **Admin**: Full control over league settings, teams, events, and user management.
- **Staff**: Can create and manage events, approve attendance.
- **Officials & NSOs**: Can access bout assignments and officiating schedules.
- **Members (Skaters)**: Limited access, can view schedules and check attendance.

### **2. Membership Management**
- Users must provide **legal name, derby name, email, phone, emergency contact, and WFTDA insurance number**.
- Leagues can define **custom fields** for additional member information.
- Memberships can have different **statuses** (active, suspended, retired, guest, alumni).

### **3. Scheduling & Attendance**
- Events can be **practices, games, meetings, or fundraisers**.
- **Event visibility is role-based** (e.g., only referees see officiating schedules).
- Attendance can be tracked via **printed sign-in sheets, QR code scans, or passphrase-based check-ins**.

### **4. Team & Roster Management**
- Users can belong to **multiple teams** (home team, travel team, special events).
- Positions are tracked separately for each team (e.g., a skater may be a jammer on one team and a blocker on another).
- Custom team categories (home teams, travel teams, official crews, NSO crews).

### **5. Gameday Staffing & Bout Management**
- Events may have **multiple bouts**, each with different team combinations.
- Referees and NSOs may be assigned across **multiple bouts in a tournament**.
- Leagues can define **custom roles** for staffing (head referee, penalty tracker, scoreboard operator, etc.).

### **6. Security & Compliance**
- **GDPR-compliant** data handling for European leagues.
- **Password hashing** with bcrypt.
- **API security** enforced with JWT authentication and API keys.
- Audit logs track important actions (user role changes, event updates).

## Installation & Setup

```
# Clone the repository
git clone https://github.com/your-org/derbyops.git
cd derbyops

# Set up environment variables
cp .env.example .env

# Start services
docker-compose up --build
```

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

