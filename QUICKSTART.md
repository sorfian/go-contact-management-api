# ⚡ Quick Start Guide

Panduan cepat untuk menjalankan aplikasi Go Todo List.

## 🏃 Development Mode

### Option 1: Menggunakan Make (Recommended)

```bash
# 1. Install dependencies
make install

# 2. Generate Wire code
make wire

# 3. Setup .env file
cp .env.example .env
# Edit .env sesuai konfigurasi database Anda

# 4. Run aplikasi
make dev
```

### Option 2: Tanpa Make

```bash
# 1. Install dependencies
go mod download
go mod tidy

# 2. Generate Wire code
go run github.com/google/wire/cmd/wire
cd test && go run github.com/google/wire/cmd/wire && cd ..

# 3. Setup .env file
cp .env.example .env
# Edit .env sesuai konfigurasi database Anda

# 4. Run aplikasi
go run .
```

Aplikasi akan berjalan di `http://localhost:3000`

---

## 🚀 Production Mode

### Build Binary

```bash
make build
```

### Run Binary

```bash
# Linux/Mac
./bin/app

# Windows
.\bin\app.exe
```

### Dengan Environment Variables

```bash
# Set environment variables
export APP_ENV=production
export APP_PORT=8080
export DB_HOST=your_host
export DB_USER=your_user
export DB_PASSWORD=your_password
export DB_NAME=go_todo_list

# Run
./bin/app
```

---

## 🧪 Testing

```bash
# Run all tests
make test

# Run integration tests
make test-integration
```

---

## 📝 Available Commands

```bash
make help              # Show all available commands
make install           # Install dependencies
make wire              # Generate Wire code
make dev               # Run in development mode
make build             # Build binary
make run               # Build and run
make test              # Run all tests
make test-integration  # Run integration tests
make clean             # Clean build artifacts
```

---

## 🔧 Configuration

Edit file `.env`:

```env
# Application
APP_ENV=development
APP_PORT=3000

# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=go_todo_list
```

---

## 📚 API Documentation

Lihat file `apispec.json` untuk dokumentasi lengkap API.

---

## 🆘 Troubleshooting

**Port sudah digunakan:**
```bash
# Ganti port di .env
APP_PORT=8080
```

**Database connection error:**
```bash
# Pastikan MySQL running
# Periksa kredensial di .env
# Test koneksi:
mysql -h localhost -u root -p go_todo_list
```

**Wire error:**
```bash
# Regenerate wire
make wire
```

---

## 📖 Dokumentasi Lengkap

Lihat [DEPLOYMENT.md](DEPLOYMENT.md) untuk panduan deployment lengkap.
