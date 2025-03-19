# DerbyOps Database Design

## 📌 Overview
DerbyOps uses **PostgreSQL** as its primary database for handling user authentication, team management, event scheduling, attendance tracking, and gameday staffing. The database follows **a relational schema with strict foreign key constraints** for **data integrity and security**.

## 📐 Database Schema
### **1️⃣ Users Table**
Stores user profiles and authentication data.
```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    legal_name VARCHAR(255) NOT NULL,
    derby_name VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address TEXT,
    emergency_contact JSONB, -- { "name": "John Doe", "phone": "555-1234" }
    wftda_insurance_number VARCHAR(50) UNIQUE,
    password_hash TEXT NOT NULL,
    role VARCHAR(50) CHECK (role IN ('admin', 'staff', 'member', 'official', 'nso')) NOT NULL,
    membership_status VARCHAR(50) CHECK (membership_status IN ('active', 'suspended', 'retired', 'guest', 'alumni')) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL -- Soft delete tracking
);
```

### **2️⃣ Teams Table**
Stores team information.
```sql
CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    category VARCHAR(50) CHECK (category IN ('home', 'travel', 'officials', 'nsos')) NOT NULL,
    custom_fields JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
```

### **3️⃣ Team Membership Table**
Tracks which users belong to which teams and their roles.
```sql
CREATE TABLE team_membership (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams(id) ON DELETE CASCADE,
    preferred_position VARCHAR(50) CHECK (preferred_position IN ('jammer', 'blocker', 'pivot', 'referee', 'nso', 'none')),
    assigned_position VARCHAR(50) CHECK (assigned_position IN ('jammer', 'blocker', 'pivot', 'referee', 'nso', 'none')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    UNIQUE(user_id, team_id)
);
```

### **4️⃣ Events Table**
Stores practices, games, and other events.
```sql
CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR(255) NOT NULL,
    description TEXT,
    event_type VARCHAR(50) CHECK (event_type IN ('practice', 'game', 'meeting', 'tournament', 'fundraiser')),
    visibility VARCHAR(50) CHECK (visibility IN ('public', 'private', 'restricted')) DEFAULT 'public',
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    custom_fields JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);
```

### **5️⃣ Event Access Table**
Manages role-based access control for events.
```sql
CREATE TABLE event_access (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID REFERENCES events(id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams(id) ON DELETE CASCADE,
    role VARCHAR(50) CHECK (role IN ('skaters', 'officials', 'leadership', 'nso', 'public')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **6️⃣ Attendance Table**
Tracks attendance for events.
```sql
CREATE TABLE attendance (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID REFERENCES events(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    check_in_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    check_in_method VARCHAR(50) CHECK (check_in_method IN ('manual', 'qr_code', 'passphrase')),
    UNIQUE(event_id, user_id), -- Ensures only one check-in per event per user
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### **7️⃣ Bouts Table**
Tracks multiple bouts per event.
```sql
CREATE TABLE bouts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID REFERENCES events(id) ON DELETE CASCADE,
    bout_number INT NOT NULL CHECK (bout_number > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **8️⃣ Bout Teams Table**
Associates teams with bouts.
```sql
CREATE TABLE bout_teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bout_id UUID REFERENCES bouts(id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams(id) ON DELETE CASCADE,
    role VARCHAR(50) CHECK (role IN ('home', 'away', 'scrimmage', 'referee', 'nso')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **9️⃣ Bout Rosters Table**
Tracks individual skaters, referees, and NSOs assigned to bouts.
```sql
CREATE TABLE bout_rosters (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bout_id UUID REFERENCES bouts(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    team_id UUID REFERENCES teams(id) ON DELETE CASCADE, -- Null if the person is unaffiliated with a team
    role VARCHAR(50) CHECK (role IN ('jammer', 'blocker', 'pivot', 'referee', 'nso', 'coach')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### **🔹 Future Enhancements**
- **Webhook Support** for real-time event updates.
- **GraphQL API** to allow for flexible querying.
- **Microservices Data Isolation** for modular scaling.

## 📌 Next Steps
- Implement schema migrations using **Go Migrate**.
- Set up **unit tests** for database queries.
- Optimize **indexes** for performance.
