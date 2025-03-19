# DerbyOps API Architecture

## 📌 Overview
DerbyOps is an **API-first roller derby league management platform** designed to be modular, secure, and scalable. The API follows **RESTful design principles** and is built in **Golang (Gin Framework)** with **PostgreSQL** as the primary database.

## 📐 High-Level Architecture
- **Backend**: Golang (Gin) + PostgreSQL
- **Authentication**: JWT-based user authentication + API keys for third-party services
- **Authorization**: Role-Based Access Control (RBAC)
- **Event-Driven**: Uses **Webhooks** for notifications
- **Scalable Deployment**: Supports Docker, Kubernetes, and Helm

## 🔑 Authentication & Security
- **Users authenticate via JWT tokens**
- **API Keys for external integrations**
- **RBAC (Role-Based Access Control)** ensures users only access authorized data
- **Audit Logs** track key actions (e.g., user role changes, event modifications)

## 📂 API Modules & Endpoints

### **1️⃣ Authentication API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/auth/register` | User registration |
| `POST` | `/auth/login` | User login, returns JWT |
| `POST` | `/auth/logout` | Logout, invalidate refresh token |
| `POST` | `/auth/refresh` | Get a new access token using refresh token |

---

### **2️⃣ User Management API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET`  | `/users` | List all users (Admin only) |
| `GET`  | `/users/:id` | Get a user's profile |
| `PATCH` | `/users/:id` | Update user profile |
| `DELETE` | `/users/:id` | Soft delete user (Admin only) |

---

### **3️⃣ Team Management API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/teams` | Create a new team |
| `GET`  | `/teams` | List all teams |
| `GET`  | `/teams/:id` | Get team details |
| `PATCH` | `/teams/:id` | Update team details |
| `POST` | `/teams/:id/members` | Add a user to a team |
| `DELETE` | `/teams/:id/members/:user_id` | Remove a user from a team |

---

### **4️⃣ Event Scheduling API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/events` | Create an event |
| `GET`  | `/events` | List all visible events |
| `GET`  | `/events/:id` | Get event details |
| `PATCH` | `/events/:id` | Update event details |
| `DELETE` | `/events/:id` | Soft delete an event |

---

### **5️⃣ Attendance API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/events/:event_id/attendance` | User check-in |
| `GET`  | `/events/:event_id/attendance` | Get attendance list |
| `DELETE` | `/events/:event_id/attendance/:user_id` | Remove check-in |

---

### **6️⃣ Roster & Bout Management API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/bouts` | Create a bout |
| `GET`  | `/bouts` | List all bouts |
| `POST` | `/bouts/:id/teams` | Add teams to bout |
| `POST` | `/bouts/:id/rosters` | Add users to bout roster |

---

### **7️⃣ Gameday Staffing API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/events/:event_id/staffing` | Assign a role |
| `GET`  | `/events/:event_id/staffing` | List assigned staff |
| `DELETE` | `/events/:event_id/staffing/:user_id` | Remove a staff role |

---

### **8️⃣ Notifications API**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/notifications/email` | Send an email |
| `POST` | `/notifications/push` | Send a push notification |

## 📜 Future Enhancements
- **Microservices Architecture**: Breaking API into independent services
- **GraphQL Support**: Enabling more flexible queries
- **Webhook-based Notifications**: Real-time event updates

## 📌 Next Steps
- Implement authentication first (`/auth` module)
- Develop user management & role-based access controls
- Build event scheduling & attendance tracking

---
This document should be updated as the API evolves to reflect new endpoints, security updates, and structural changes.
