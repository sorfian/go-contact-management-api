# 🚀 Deployment Guide - Go Contact Management App

Panduan lengkap untuk menjalankan aplikasi di mode development dan production.

## 📋 Daftar Isi

- [Requirements](#requirements)
- [Setup Awal](#setup-awal)
- [Development](#development)
- [Production](#production)
- [Environment Variables](#environment-variables)
- [Troubleshooting](#troubleshooting)

---

## Requirements

- Go 1.25 atau lebih baru
- MySQL 5.7+ atau MariaDB 10.2+
- Make (optional, untuk menggunakan Makefile)
- Git

---

## Setup Awal

### 1. Clone Repository

```bash
git clone <repository-url>
cd go-todo-list-app
```

### 2. Install Dependencies

```bash
# Menggunakan make
make install

# Atau manual
go mod download
go mod tidy
```

### 3. Setup Database

Buat database MySQL:

```sql
CREATE DATABASE go_todo_list CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. Konfigurasi Environment

Copy file `.env.example` ke `.env`:

```bash
cp .env.example .env
```

Edit `.env` dan sesuaikan dengan konfigurasi lokal Anda:

```env
APP_ENV=development
APP_PORT=3000

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=go_todo_list
```

### 5. Generate Wire Dependencies

```bash
# Menggunakan make
make wire

# Atau manual
go run github.com/google/wire/cmd/wire
cd test && go run github.com/google/wire/cmd/wire && cd ..
```

---

## Development

### Menjalankan Aplikasi

#### Metode 1: Menggunakan Make (Recommended)

```bash
make dev
```

#### Metode 2: Go Run

```bash
go run .
```

#### Metode 3: Build & Run

```bash
# Build
make build

# Run
./bin/app
```

### Hot Reload (Optional)

Untuk development dengan hot reload, install Air:

```bash
go install github.com/air-verse/air@latest
```

Buat file `.air.toml`:

```toml
root = "."
tmp_dir = "tmp"

[build]
  cmd = "go build -o ./tmp/main ."
  bin = "tmp/main"
  include_ext = ["go", "tpl", "tmpl", "html"]
  exclude_dir = ["assets", "tmp", "vendor", "test"]
  delay = 1000
  stop_on_error = true
```

Jalankan dengan:

```bash
air
```

### Testing

```bash
# Run semua test
make test

# Run integration test saja
make test-integration

# Run test dengan verbose
go test -v ./test

# Run specific test
cd test && go test -v -run TestRegisterSuccess
```

---

## Production

### 1. Prepare Environment

Buat file `.env` untuk production atau set environment variables:

```bash
export APP_ENV=production
export APP_PORT=3000
export DB_HOST=your_production_host
export DB_PORT=3306
export DB_USER=your_production_user
export DB_PASSWORD=your_production_password
export DB_NAME=go_todo_list
```

### 2. Build Aplikasi

```bash
# Build binary
make build

# Atau dengan optimasi
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/app .
```

### 3. Deployment Options

#### A. Direct Server Deployment

```bash
# Copy binary ke server
scp bin/app user@server:/path/to/app/

# Copy .env.production
scp .env.production user@server:/path/to/app/.env

# SSH ke server dan jalankan
ssh user@server
cd /path/to/app
./app
```

#### B. Systemd Service (Linux)

Buat file `/etc/systemd/system/go-contact-management.service`:

```ini
[Unit]
Description=Go Contact Management API
After=network.target mysql.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/var/www/go-todo-list
ExecStart=/var/www/go-todo-list/bin/app
Restart=on-failure
RestartSec=5s

# Environment variables
Environment="APP_ENV=production"
Environment="APP_PORT=3000"
Environment="DB_HOST=localhost"
Environment="DB_PORT=3306"
Environment="DB_USER=your_user"
Environment="DB_PASSWORD=your_password"
Environment="DB_NAME=go_todo_list"

[Install]
WantedBy=multi-user.target
```

Enable dan start service:

```bash
sudo systemctl daemon-reload
sudo systemctl enable go-contact-management
sudo systemctl start go-contact-management
sudo systemctl status go-contact-management
```

#### C. Docker Deployment

Buat `Dockerfile`:

```dockerfile
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Generate wire
RUN go install github.com/google/wire/cmd/wire@latest
RUN wire

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/main .

EXPOSE 3000

CMD ["./main"]
```

Buat `docker-compose.yml`:

```yaml
version: '3.8'

services:
  app:
    build: .
    ports:
      - "3000:3000"
    environment:
      - APP_ENV=production
      - APP_PORT=3000
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=root
      - DB_PASSWORD=rootpassword
      - DB_NAME=go_todo_list
    depends_on:
      - mysql
    restart: unless-stopped

  mysql:
    image: mysql:8.0
    environment:
      - MYSQL_ROOT_PASSWORD=rootpassword
      - MYSQL_DATABASE=go_todo_list
    volumes:
      - mysql_data:/var/lib/mysql
    ports:
      - "3306:3306"
    restart: unless-stopped

volumes:
  mysql_data:
```

Jalankan dengan:

```bash
docker-compose up -d
```

### 4. Nginx Reverse Proxy (Optional)

Konfigurasi Nginx `/etc/nginx/sites-available/go-contact-management`:

```nginx
server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://localhost:3000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

Enable dan restart Nginx:

```bash
sudo ln -s /etc/nginx/sites-available/go-contact-management /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
```

---

## Environment Variables

### Aplikasi

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `APP_ENV` | Environment mode (development/production) | development | No |
| `APP_PORT` | Port aplikasi | 3000 | No |
| `LOG_LEVEL` | Log level (info/error) | info | No |

### Database

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `DB_HOST` | Database host | localhost | Yes |
| `DB_PORT` | Database port | 3306 | Yes |
| `DB_USER` | Database user | root | Yes |
| `DB_PASSWORD` | Database password | - | Yes |
| `DB_NAME` | Database name | go_todo_list | Yes |
| `DB_MAX_IDLE_CONNS` | Max idle connections | 10 | No |
| `DB_MAX_OPEN_CONNS` | Max open connections | 100 | No |
| `DB_CONN_MAX_LIFETIME` | Connection max lifetime | 30m | No |
| `DB_CONN_MAX_IDLE_TIME` | Connection max idle time | 10m | No |

---

## Troubleshooting

### Port Already in Use

```bash
# Check what's using the port
lsof -i :3000

# Or on Windows
netstat -ano | findstr :3000

# Kill the process or change APP_PORT in .env
```

### Database Connection Failed

1. Pastikan MySQL service running
2. Periksa kredensial di `.env`
3. Periksa firewall/security groups
4. Test koneksi manual:

```bash
mysql -h localhost -u root -p go_todo_list
```

### Wire Generation Error

```bash
# Regenerate wire files
make wire

# Or manually
go run github.com/google/wire/cmd/wire
cd test && go run github.com/google/wire/cmd/wire
```

### Module/Dependency Issues

```bash
# Clear cache dan reinstall
go clean -modcache
go mod download
go mod tidy
```

---

## Monitoring & Logs

### Application Logs

```bash
# Systemd service
sudo journalctl -u go-todo-list -f

# Docker
docker-compose logs -f app
```

### Database Logs

```bash
# MySQL logs
sudo tail -f /var/log/mysql/error.log
```

---

## Backup & Recovery

### Database Backup

```bash
# Backup
mysqldump -u root -p go_todo_list > backup_$(date +%Y%m%d).sql

# Restore
mysql -u root -p go_todo_list < backup_20250125.sql
```

---

## Security Checklist

- [ ] Change default database password
- [ ] Use HTTPS in production
- [ ] Set proper file permissions (644 for files, 755 for directories)
- [ ] Don't commit `.env` file to git
- [ ] Use strong database passwords
- [ ] Regularly update dependencies
- [ ] Enable firewall rules
- [ ] Use reverse proxy (Nginx/Apache)

---

## Performance Optimization

### Database

- Enable query cache
- Add proper indexes
- Optimize connection pool settings
- Use read replicas for high traffic

### Application

- Build with optimizations: `-ldflags="-s -w"`
- Use production mode: `APP_ENV=production`
- Enable Gzip compression in reverse proxy
- Implement caching strategy

---

## Support

Untuk pertanyaan atau issues, silakan buat issue di repository atau hubungi tim development.
