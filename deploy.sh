#!/bin/bash

# Update system
apt-get update
apt-get upgrade -y

# Install required packages
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg \
    lsb-release \
    git

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Install Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

# Create app directory
mkdir -p /opt/dimbus
cd /opt/dimbus

# Clone repository (replace with your repository URL)
git clone https://github.com/hlmelton/dimbus.git .

# Create .env file
cat > .env << EOL
POSTGRES_DB=dimbus
POSTGRES_USER=dimbus_user
POSTGRES_PASSWORD=$(openssl rand -base64 32)
DOMAIN=${DOMAIN:-localhost}
CADDY_EMAIL=${CADDY_EMAIL:-your@email.com}
EOL

# Copy production files
cp docker-compose.prod.yml docker-compose.yml
cp Caddyfile.prod Caddyfile

# Start services
docker-compose pull
docker-compose up -d

# Print success message
echo "Deployment completed! Your application should be available at https://$DOMAIN"
echo "Please make sure your domain's DNS is pointing to this server's IP address." 