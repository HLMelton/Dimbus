# 🔗 Dimbus - Link Shortener

A modern, full-stack URL shortener built with Go, HTMX, and PostgreSQL. Features server-side rendering with templ templates and a clean, responsive UI.

## 🛠️ Prerequisites

Before you begin, ensure you have the following installed:

- [Go 1.23+](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [templ](https://templ.guide/quick-start/installation) - Install with:
  ```bash
  go install github.com/a-h/templ/cmd/templ@latest
  ```

## 🏗️ Project Structure

```
.
├── cmd/server/          # Application entry point
├── internal/            # Internal packages
│   └── database/       # Database operations
├── templates/          # UI templates
│   ├── components/    # Reusable UI components
│   └── pages/        # Page templates
└── database/          # Database migrations
```

## 🚀 Running the Project

### Using Docker Compose (Recommended)

1. Clone the repository:
   ```bash
   git clone https://github.com/hlmelton/dimbus.git
   cd dimbus
   ```

2. Start the application:
   ```bash
   docker-compose up --build
   ```

3. Access the application at [http://localhost](http://localhost)

### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Generate templ files:
   ```bash
   templ generate
   ```

3. Start PostgreSQL:
   ```bash
   docker-compose up db
   ```

4. Run the application:
   ```bash
   go run ./cmd/server
   ```

5. Access the application at [http://localhost:8080](http://localhost:8080)

## 🔧 Configuration

The application can be configured using environment variables:

- `DATABASE_URL`: PostgreSQL connection string (default: "postgres://dimbus_user:yourpassword@localhost:5432/dimbus?sslmode=disable")
- `PORT`: Server port (default: 8080)

## 🛡️ Features

- ✨ Modern, responsive UI with TailwindCSS
- 🔄 Real-time interactions using HTMX
- 🔒 Secure by default with proper HTTP headers
- 📊 PostgreSQL for reliable data storage
- 🚦 Reverse proxy with Caddy
- 🐳 Docker support for easy deployment

## 🧪 Development

### Running Tests
```bash
go test ./...
```

### Code Generation
```bash
# Generate templ files
templ generate
```

### Database Migrations
The initial schema is automatically applied when running with Docker Compose. For manual application:

```sql
psql -U dimbus_user -d dimbus -f database/migrations/001_initial_schema.sql
```

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

![Go](https://img.shields.io/badge/Go-1.23-blue)
![Docker](https://img.shields.io/badge/Docker-yes-blue)
![Postgres](https://img.shields.io/badge/PostgreSQL-15-blue)
![HTMX](https://img.shields.io/badge/HTMX-1.9.10-blue)
![templ](https://img.shields.io/badge/templ-0.3.865-blue)

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

```


[![DigitalOcean Referral Badge](https://web-platforms.sfo2.cdn.digitaloceanspaces.com/WWW/Badge%202.svg)](https://www.digitalocean.com/?refcode=2fe7735e3150&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge)