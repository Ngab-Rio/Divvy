# Divvy

Divvy adalah **RESTful API** untuk aplikasi **management** yang dibangun menggunakan **Golang** dengan PostgreSQL sebagai database utama.  
API ini menyediakan fitur untuk mengelola user, group, serta manajemen data yang dapat digunakan oleh aplikasi client (misalnya mobile atau web).

---

## 🚀 Fitur Utama
- **Autentikasi & Otorisasi**
  - Login & Register user
  - JWT Authentication
- **Manajemen User**
  - CRUD user
- **Manajemen Group**
  - Membuat dan mengelola group
  - Menambahkan anggota ke group
- **Relasi Data**
  - Relasi user dengan group
- **Middleware**
  - Logging
  - CORS
  - JWT Validator

---

## 🛠️ Tech Stack
- **Backend:** Go (Golang) dengan [Fiber](https://gofiber.io/)
- **Database:** PostgreSQL
- **ORM/Query Builder:** [Goqu](https://github.com/doug-martin/goqu)
- **Authentication:** JWT
- **Deployment:** Railway / VPS (opsional)

---

## 📂 Struktur Project (ringkas)
