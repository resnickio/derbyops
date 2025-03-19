# DerbyOps Use Case Scenarios

## 📌 Overview
This document outlines the **primary use cases** for DerbyOps, ensuring that the platform meets real-world league management needs.

---

## **1️⃣ User Management & Authentication**
### **Use Case 1: User Registration & Login**
- **Actors**: League members, officials, league administrators
- **Preconditions**: User must provide a valid email and password
- **Flow**:
  1. User registers with email, legal name, and derby name
  2. System hashes the password and stores user information
  3. User logs in and receives a JWT token
  4. JWT is used to authenticate subsequent API requests
- **Postconditions**: User is authenticated and can access the platform based on their role

### **Use Case 2: Role-Based Access Management**
- **Actors**: League administrators
- **Preconditions**: User must have an `admin` role
- **Flow**:
  1. Admin assigns roles to users (admin, staff, official, member)
  2. API enforces access controls for different features
- **Postconditions**: Users have the correct permissions to manage leagues

---

## **2️⃣ Team & Roster Management**
### **Use Case 3: Creating & Managing Teams**
- **Actors**: League administrators, team captains
- **Preconditions**: League must exist
- **Flow**:
  1. Admin creates a team and assigns a category (home, travel, official crew, NSO crew)
  2. Admin adds or removes skaters to/from the team
- **Postconditions**: Teams are available for event scheduling and roster assignments

### **Use Case 4: Assigning Skaters to Rosters**
- **Actors**: Team captains, officials
- **Preconditions**: Team and event must exist
- **Flow**:
  1. Captain assigns skaters to a bout roster
  2. Officials review the roster for compliance with league rules
- **Postconditions**: Roster is saved and visible for event management

---

## **3️⃣ Event Scheduling & Attendance Tracking**
### **Use Case 5: Scheduling a Practice or Game**
- **Actors**: League administrators, team captains
- **Preconditions**: Team must exist
- **Flow**:
  1. Admin schedules an event (practice, game, meeting)
  2. Users with access see the event in their dashboard
- **Postconditions**: Event is scheduled and visible

### **Use Case 6: Checking in to an Event**
- **Actors**: League members, officials
- **Preconditions**: User must be rostered for the event
- **Flow**:
  1. User checks in using a QR code, passphrase, or manual check-in
  2. Attendance is logged in the system
- **Postconditions**: Attendance is recorded for tracking eligibility

---

## **4️⃣ Gameday Staffing & Officiating**
### **Use Case 7: Assigning Officials to a Bout**
- **Actors**: League administrators, officials
- **Preconditions**: Event must be scheduled
- **Flow**:
  1. Admin assigns referees and NSOs to officiating roles
  2. Assignments are saved and visible to officials
- **Postconditions**: Officiating crew is confirmed for the event

---

## **5️⃣ Future Use Cases (Post-MVP)**
### **Use Case 8: Training Progression Tracking**
- **Actors**: League administrators, training coordinators
- **Flow**:
  1. Admin defines training cohorts and assessment criteria
  2. Skaters log training progress and assessment completion
- **Postconditions**: Training data is tracked for league development

### **Use Case 9: Calendar Integration**
- **Actors**: League members, captains
- **Flow**:
  1. User links DerbyOps calendar with Google Calendar or Outlook
  2. Events sync automatically
- **Postconditions**: Skaters see league events in external calendars

---

## 📌 Next Steps
1️⃣ Validate these use cases with real-world leagues  
2️⃣ Ensure API endpoints fully cover these workflows  
3️⃣ Expand scenarios as new features are developed  

This document will evolve as DerbyOps grows! 🚀
