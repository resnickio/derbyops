# Contributing to DerbyOps

## 📌 Overview
Thank you for considering contributing to **DerbyOps**! This guide will help you get started with development, submitting contributions, and following best practices.

---

## **1️⃣ How to Contribute**
### **🔹 Reporting Issues**
- Use [GitHub Issues](https://github.com/your-org/derbyops/issues) to report bugs, feature requests, or documentation improvements.
- Include detailed information, reproduction steps, and screenshots if applicable.

### **🔹 Submitting Code Changes**
1. **Fork the Repository**: Click the `Fork` button on GitHub.
2. **Clone Your Fork**:
   ```sh
   git clone https://github.com/your-username/derbyops.git
   cd derbyops
   ```
3. **Create a Feature Branch**:
   ```sh
   git checkout -b feature/your-feature-name
   ```
4. **Write Code & Tests**:
   - Follow the **coding guidelines**.
   - Ensure tests pass before submitting PR.
5. **Commit & Push**:
   ```sh
   git add .
   git commit -m "Add feature X"
   git push origin feature/your-feature-name
   ```
6. **Open a Pull Request (PR)**:
   - Go to `Pull Requests` on GitHub and create a PR.
   - Fill in the description and link any related issues.
   - Wait for code review and make necessary changes.

---

## **2️⃣ Coding Standards**
### **🔹 Backend (Golang)**
- Follow idiomatic Go practices (`golangci-lint` enforced).
- Use `Gin` framework for API routes.
- Database queries use `GORM` ORM.
- Tests are required for new features (`go test`).

### **🔹 Frontend (React, Next.js)**
- Use functional components and hooks (`useState`, `useEffect`).
- Follow accessibility best practices (`a11y`).
- Use TypeScript for type safety.

### **🔹 Mobile (Swift & Kotlin) [Future]**
- Maintain consistent UI with web version.
- Use platform-native best practices.

---

## **3️⃣ Testing Requirements**
- Run **unit tests** before submitting a PR.
  ```sh
  go test -v ./...
  ```
- If applicable, include **integration tests** for new API features.
- Use **Postman** or **cURL** to validate API changes manually.
- Frontend tests should pass before PR approval.

---

## **4️⃣ Branching & Versioning**
- `main` → Stable, production-ready code.
- `develop` → Active development branch.
- Feature branches follow `feature/your-feature-name`.
- Bug fixes follow `fix/issue-description`.
- Releases are tagged using `vX.Y.Z` format.

---

## **5️⃣ Commit Message Format**
| Type     | Description                                      |
|----------|--------------------------------------------------|
| `feat:`  | New feature implementation                      |
| `fix:`   | Bug fix                                        |
| `docs:`  | Documentation changes                          |
| `test:`  | Adding or modifying tests                      |
| `chore:` | General maintenance (formatting, refactoring)  |

Example:
```sh
git commit -m "feat: add attendance tracking API"
```

---

## **6️⃣ Reviewing & Merging PRs**
- PRs require **at least one approval** before merging.
- Ensure **all tests pass** before merging.
- Merge strategies:
  - **Squash and Merge** (default for small changes).
  - **Rebase and Merge** (for cleaner history).

---

## **7️⃣ Community Guidelines**
- Be respectful and collaborative.
- Provide constructive feedback in reviews.
- Follow the [Code of Conduct](https://github.com/your-org/derbyops/blob/main/CODE_OF_CONDUCT.md).

---

## 📌 Next Steps
1️⃣ Read the [README](README.md) for project setup.  
2️⃣ Check out open issues labeled [`good first issue`](https://github.com/your-org/derbyops/issues?q=is%3Aissue+is%3Aopen+label%3A%22good+first+issue%22).  
3️⃣ Start contributing! 🚀
