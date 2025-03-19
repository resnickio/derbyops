# DerbyOps Architecture

## 📌 Overview
DerbyOps is a **multi-instance, API-first roller derby league management platform** designed for **self-hosting** and a **SaaS model**. Each league operates as an independent instance with its own **subdomain or custom domain**, ensuring full data separation. The system supports **team and roster management, event scheduling, attendance tracking, and role-based access control (RBAC)**.

## **System Architecture**
### **1️⃣ Core Components**
- **Backend API**: Golang (Gin framework), PostgreSQL, JWT authentication
- **Frontend**: React (Next.js), consuming the API
- **Mobile Apps**: Swift (iOS) & Kotlin (Android) (Future feature)
- **Infrastructure**: Docker, Helm, Kubernetes (DigitalOcean initially, with possible future migration to a larger cloud provider)

### **2️⃣ Deployment Model**
- **Self-Hosting**: Leagues can deploy and manage their own DerbyOps instance
- **SaaS Model**: Each league gets its own isolated API and database instance via a subdomain (e.g., `myleague.derbyops.com`)
- **Database Backend**: PostgreSQL (only supported database for now)

### **3️⃣ Key Functional Components**
| **Component**             | **Description** |
|---------------------------|-----------------|
| **Authentication**        | JWT-based user authentication, API keys for third-party access |
| **User Management**       | Profiles, roles, hierarchical permissions (admins, captains, officials, league leadership) |
| **Team & Roster Management** | Multi-team support per league, ad-hoc teams, historical roster tracking |
| **Event Scheduling**      | Bout, practice, and meeting scheduling with **calendar sync (future)** |
| **Attendance Tracking**   | QR code/manual check-ins, passphrase sign-in, eligibility tracking |
| **Gameday Staffing**      | Assign volunteers, officials, and NSOs to events |
| **Activity Logging**      | Tracks user actions for security & audit purposes |
| **Historical Data**       | Player history, past team assignments, officiating records |
| **Future Expansion**      | Training tracking, game statistics import, external integrations |

## **📌 Data Model Considerations**
1. **Each league is independent** (separate database instances in SaaS model)
2. **Historical tracking** is enabled for:
   - Past rosters and bout assignments
   - Team membership changes
   - Officiating records
3. **Hierarchical roles** include:
   - **League Leadership** (Admins, Team Captains)
   - **Officials** (Head Officials, Regular Officials, NSOs)
   - **Skaters** (Regular members, training cohorts [Future])
4. **Training Progression Tracking (Future Feature)**
   - Leagues can define their own **cohort names & assessment rules**
   - Example: **Bronze → Silver → Gold** system with skill assessment & scrimmage hours

## **📌 Scalability Considerations**
- Each instance should support **hundreds of users** per league
- Supports **5-12 internal teams per league** with **ad-hoc teams for one-off matches**
- Official crews & NSO crews managed separately for officiating history
- Long-term scalability allows migration to larger cloud providers if needed

## **📌 Security & Compliance**
- **GDPR Compliance**: User data is handled securely per EU guidelines
- **User Activity Logging**: All key user actions (role changes, event modifications) are logged
- **RBAC Permissions**: Fine-grained access control at league, team, and event levels

## 📌 Next Steps
1. **Refine API documentation to ensure it meets multi-instance SaaS requirements**
2. **Define database migrations for historical tracking**
3. **Prepare deployment strategy (Helm charts, Kubernetes setup)**
