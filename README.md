
# 🏥 Hospital Management System Backend (GoFiber + MySQL)

A scalable, modular monolithic backend for managing hospital operations, built with **GoFiber** and **MySQL**. Designed for future microservice architecture while keeping the initial system simple and maintainable.

---

## 📦 Features
- ✅ Modular folder structure (monolith with future microservices in mind)
- ✅ OTP-based patient login & appointment booking
- ✅ Doctor portal for managing schedules & prescriptions
- ✅ Billing & Pharmacy management portals
- ✅ HRMS, Reports, Inventory & Utility tracking (Back office)
- ✅ Notification system (SMS, Email, Push ready)
- ✅ Clean RESTful APIs (JSON-based)
- ✅ MySQL as primary datastore
- ✅ Production-ready project structure

---

## 🛠 Tech Stack
- **Language**: Go (GoFiber framework)
- **Database**: MySQL
- **ORM**: GORM
- **Configuration**: .env files
- **API Standard**: REST (JSON)
- **Deployment**: Docker (coming soon)

---

## 🚀 Getting Started

### 1️⃣ Clone the Repo
\`\`\`bash
git clone https://github.com/yourusername/hospital-management-backend.git
cd hospital-management-backend
\`\`\`

### 2️⃣ Install Dependencies
\`\`\`bash
go mod tidy
\`\`\`

### 3️⃣ Configure Environment
Create a \`.env\` file in the root directory:
\`\`\`env
DB_USER=root
DB_PASS=yourpassword
DB_NAME=hospital
DB_HOST=127.0.0.1
DB_PORT=3306
\`\`\`

### 4️⃣ Run the Application
\`\`\`bash
go run main.go
\`\`\`
Server will start at \`http://localhost:3000\`.

---




## ✅ TODO
- Authentication module (JWT/OTP)
- Full CRUD for each module
- Notification service integration
- Docker & docker-compose setup
- API documentation (Swagger)
- Testing & CI/CD pipeline

---

## 📄 License
[MIT License](LICENSE)

---

## 👨‍💻 Author
Made with ❤️ by [Md Jannatul Nayem](https://github.com/iamjnayem)

