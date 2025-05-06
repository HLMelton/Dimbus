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

1. Create a `.env` file with the following content:
   ```env
   POSTGRES_USER=dimbus_user
   POSTGRES_PASSWORD=yourpassword
   POSTGRES_DB=dimbus
   DATABASE_URL=postgres://dimbus_user:yourpassword@localhost:5432/dimbus?sslmode=disable
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate templ files:
   ```bash
   templ generate
   ```

4. Start PostgreSQL:
   ```bash
   docker-compose up db -d
   ```

5. Set environment variables (PowerShell):
   ```powershell
   $env:DATABASE_URL="postgres://dimbus_user:yourpassword@localhost:5432/dimbus?sslmode=disable"
   ```
   Or for Bash:
   ```bash
   export DATABASE_URL="postgres://dimbus_user:yourpassword@localhost:5432/dimbus?sslmode=disable"
   ```

6. Run the application:
   ```bash
   go run ./cmd/server
   ```

7. Access the application at [http://localhost:8080](http://localhost:8080)

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

## 🚀 Deployment to Digital Ocean

### Prerequisites
1. A Digital Ocean account
2. A domain name pointing to your Digital Ocean droplet
3. Basic knowledge of SSH and command line

### Steps

1. Create a new Ubuntu droplet on Digital Ocean:
   - Choose Ubuntu 22.04 LTS
   - Select Basic plan (minimum 1GB RAM)
   - Choose a datacenter region close to your users
   - Add your SSH key
   - Choose a hostname

2. Update your domain's DNS:
   - Add an A record pointing to your droplet's IP address
   - Wait for DNS propagation (can take up to 24 hours)

3. SSH into your droplet:
   ```bash
   ssh root@your-droplet-ip
   ```

4. Set environment variables:
   ```bash
   export DOMAIN=your-domain.com
   export CADDY_EMAIL=your@email.com
   ```

5. Download and run the deployment script:
   ```bash
   curl -O https://raw.githubusercontent.com/hlmelton/dimbus/main/deploy.sh
   chmod +x deploy.sh
   ./deploy.sh
   ```

6. Monitor the deployment:
   ```bash
   docker-compose logs -f
   ```

### Post-Deployment

1. Your application will be available at `https://your-domain.com`
2. SSL certificates will be automatically provisioned by Caddy
3. The database is persisted in a Docker volume
4. Logs can be viewed using `docker-compose logs`

### Backup

To backup the database:
```bash
docker-compose exec db pg_dump -U dimbus_user dimbus > backup.sql
```

### Updating

To update the application:
```bash
cd /opt/dimbus
git pull
docker-compose build
docker-compose up -d
```