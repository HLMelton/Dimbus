# 🔗 Dimbus - Link Shortener

A minimal, production-ready URL shortener built in Go with PostgreSQL and Docker Compose.  
Features HTTPS via [Caddy](https://caddyserver.com), security headers, and clean modular code.

![Go](https://img.shields.io/badge/Go-1.22-blue)
![Docker](https://img.shields.io/badge/Docker-yes-blue)
![Postgres](https://img.shields.io/badge/PostgreSQL-15-blue)

---

## 🚀 Features

- ✅ Written in idiomatic Go
- ✅ In-memory or Postgres-based URL mapping
- ✅ Secure by default (TLS via Caddy + middleware headers)
- ✅ Production-friendly with Docker Compose
- ✅ Easy to extend and deploy

---

## 🧱 Tools

- **Go**: HTTP server + PostgreSQL connection
- **PostgreSQL**: Stores shortened URLs
- **Caddy**: Handles HTTPS and reverse proxying
- **Docker Compose**: Simplifies development and deployment

---

## 📦 Running the Project (Local Dev)

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/)
- A domain name (optional for Caddy TLS)

---

### 1. Clone and build the stack

```bash
git clone https://github.com/HLMelton/Dimbus.git
cd dimbus
docker-compose up --build
