
# 🏥 Happy Health Backend System  (GoFiber + MySQL)

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
```
git clone https://github.com/iamjnayem/happy-health-backend.git
cd happy-health-backend
```

### 2️⃣ Install Dependencies
```
go mod tidy
```

### 3️⃣ Configure Environment
Create a ```.env``` file in the root directory:
```
DB_USER=root
DB_PASS=yourpassword
DB_NAME=hospital
DB_HOST=127.0.0.1
DB_PORT=3306
```

### 4️⃣ Run the Application
```
go run main.go
```
Server will start at ```http://localhost:3000```.

---

## ✅ TODO
1. Informational Website (Public Website)
- Display hospital name, address, phone, email.
- Service list and departments.
- Doctor list with specialty and filters.
- Hospital branches with Google Maps.
- Insight section showing patient count, doctor count.
- Campaigns, promotions, and offers.
- Appointment booking option with OTP verification.
- Auto redirect to patient portal after OTP verification.

2. Patient Portal 
- OTP based sign-in and login.
- Search doctors by specialty and branch.
- Today’s available doctors listing.
- Book appointment with SMS verification and 10% optional payment.
- View reports, prescriptions, and medical history.
- Step-by-step process guide (counter → payment → doctor → report → pharmacy).
- Billing and test processing status update.
- Test preparation instructions.
- Notification if test not available.

3. Doctor Portal(web & mobile)
- Doctor profile creation by admin (invitation-based login).
- Profile and basic info update.
- Weekly schedule management.
- Today’s patient waiting list.
- Create and update prescriptions.
- View income and earnings.

4. Bill Counter Portal 
- Separate web portal for billing.
- Cash and digital payment handling.
- Unique login for each accountant.
- Report and bill delivery status update.
- Hardcopy report delivery tracking.

5. Pharmacy Portal 
- Separate pharmacy login.
- Sell medicine, create invoice, accept payments.
- Stock check and management.
- Pre-defined stock level alarm setup.
- Inventory and stock history reporting.

6. Operations Portal (Back Office/Head Office)
- HRMS module (staff onboarding, doctor, pharmacist, other roles).
- Business profit-loss reporting.
- Utility bill tracking (electricity, gas, generator).
- Machine/equipment damage log and accounting.
- Test sample agent purchase and stock management.
- Approval system (cash, purchase, payments).
- Document management (tax, license, agreements).
- Centralized supervision and reporting layer.

7. Notification System (Cross Application Communication)
- SMS/Email/Push Notification integration.
- Patient booking confirmation.
- Doctor notifications.
- Billing status notifications.
- Report delivery notifications.
- Stock alarms.
- Test reminders.

---

## 📄 License
[MIT License](LICENSE)

---

## 👨‍💻 Author
Made with ❤️ by [Md Jannatul Nayem](https://github.com/iamjnayem)

