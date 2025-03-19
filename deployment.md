# DerbyOps Deployment & Infrastructure

## 📌 Overview
This document outlines the **deployment strategy** for DerbyOps, covering **self-hosting**, **SaaS deployment**, and **scalability considerations**.

## **1️⃣ Deployment Models**
DerbyOps supports two primary deployment models:

### **🔹 Self-Hosting (Docker & Kubernetes)**
- Suitable for **leagues that prefer to manage their own infrastructure**.
- Uses **Docker Compose** for local deployments.
- Provides **Helm charts** for Kubernetes-based hosting.
- Supports PostgreSQL as the primary database.

### **🔹 SaaS Deployment (Multi-Tenant Model)**
- Each league operates on its **own subdomain or custom domain**.
- **DigitalOcean Kubernetes** (initial deployment) with scalability in mind.
- Separate PostgreSQL instances per league for data isolation.
- Custom Helm charts for rapid provisioning.

---

## **2️⃣ Deployment Stack**
| **Component**      | **Technology**        | **Purpose**                     |
|-------------------|----------------------|---------------------------------|
| **Backend API**   | Golang (Gin)         | Core application logic         |
| **Database**      | PostgreSQL           | Stores league & user data      |
| **Frontend**      | React (Next.js)      | Web-based admin panel          |
| **Mobile Apps**   | Swift (iOS), Kotlin (Android) | (Future feature) Mobile companion apps |
| **Containerization** | Docker, Kubernetes | Deployment automation         |
| **Orchestration** | Helm                 | Manages Kubernetes deployments |
| **Reverse Proxy** | Nginx, Traefik       | Manages HTTPS & routing       |
| **Monitoring**    | Prometheus, Grafana  | Observability & performance    |

---

## **3️⃣ Self-Hosting Guide (Docker Compose)**
### **1️⃣ Install Dependencies**
Ensure you have **Docker** and **Docker Compose** installed:
```sh
sudo apt update && sudo apt install docker docker-compose -y
```

### **2️⃣ Clone the Repository**
```sh
git clone https://github.com/your-org/derbyops.git
cd derbyops
```

### **3️⃣ Configure Environment Variables**
Copy the example environment file:
```sh
cp .env.example .env
```
Modify `.env` with database credentials and API secrets.

### **4️⃣ Start Services**
```sh
docker-compose up --build -d
```
This will start the **API, PostgreSQL database, and other services**.

---

## **4️⃣ Kubernetes Deployment Guide**
### **1️⃣ Prerequisites**
- **Kubernetes cluster** (Managed K8s on DigitalOcean, AWS, GCP, etc.)
- **Helm 3 installed**
- **Configured domain name for subdomain-based leagues**

### **2️⃣ Deploy with Helm**
```sh
helm repo add derbyops https://your-org.github.io/helm-charts/
helm install my-league derbyops/derbyops --set domain=myleague.derbyops.com
```
This will provision **a fully isolated instance** for `myleague.derbyops.com`.

---

## **5️⃣ Scaling Considerations**
### **Database Scaling**
- Each league gets its **own PostgreSQL instance**.
- Optional: Use **managed PostgreSQL** for performance and backups.

### **Application Scaling**
- **Auto-scaling API pods** using Kubernetes **Horizontal Pod Autoscaler (HPA)**.
- **Load balancing with Traefik/Nginx** to distribute traffic.
- **Redis caching (future)** for performance optimizations.

### **Security Considerations**
- **HTTPS enforcement** via Let’s Encrypt.
- **Role-based access control (RBAC)** for API security.
- **Audit logging** for admin actions.

---

## **6️⃣ Next Steps**
1️⃣ Finalize Helm charts & Kubernetes manifests.  
2️⃣ Automate database provisioning for multi-league SaaS deployment.  
3️⃣ Set up CI/CD for automated deployments.  

This document will evolve as deployment strategies are refined. 🚀
