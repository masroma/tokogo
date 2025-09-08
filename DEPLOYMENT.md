# TokoGo Deployment Guide

Panduan lengkap untuk men-deploy aplikasi TokoGo ke production server menggunakan VPS.

## 📋 Prerequisites

- VPS dengan Ubuntu 22.04 LTS
- Domain name (opsional, bisa menggunakan IP)
- SSH access ke server
- Git repository dengan kode aplikasi

## 🚀 Quick Start

### 1. Setup Server

```bash
# Update system
apt update && apt upgrade -y

# Install dependencies
apt install -y curl wget git unzip software-properties-common

# Install Go 1.21
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### 2. Setup Database

```bash
# Install MySQL
apt install -y mysql-server
systemctl start mysql
systemctl enable mysql
mysql_secure_installation

# Create database and user
mysql -u root -p
```

```sql
CREATE DATABASE tokogo_production CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'tokogo_user'@'localhost' IDENTIFIED BY 'your_secure_password_here';
GRANT ALL PRIVILEGES ON tokogo_production.* TO 'tokogo_user'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```

### 3. Setup Nginx

```bash
# Install Nginx
apt install -y nginx
systemctl start nginx
systemctl enable nginx

# Copy nginx configuration
cp nginx.conf /etc/nginx/sites-available/default
nginx -t
systemctl reload nginx
```

### 4. Setup SSL (Optional)

```bash
# Install Certbot
apt install -y certbot python3-certbot-nginx

# Generate SSL certificate
certbot --nginx -d yourdomain.com -d www.yourdomain.com

# Setup auto-renewal
echo "0 12 * * * /usr/bin/certbot renew --quiet" | crontab -
```

### 5. Deploy Application

```bash
# Clone repository
cd /root
git clone https://github.com/yourusername/tokogo.git
cd tokogo

# Setup environment
cp .env.production .env
nano .env  # Edit environment variables

# Build application
go mod tidy
go build -o main .

# Setup systemd service
cp tokogo.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable tokogo
systemctl start tokogo
```

## 📁 File Structure

```
tokogo/
├── Dockerfile              # Docker configuration
├── nginx.conf              # Nginx configuration
├── tokogo.service          # Systemd service file
├── backup.sh               # Database backup script
├── .env.production         # Production environment variables
├── .env.example            # Environment variables template
└── .github/workflows/      # CI/CD configuration
    └── deploy.yml
```

## 🔧 Configuration

### Environment Variables

Copy `.env.example` to `.env` and configure:

```bash
# Database
DB_HOST=localhost
DB_PORT=3306
DB_USER=tokogo_user
DB_PASSWORD=your_secure_password
DB_NAME=tokogo_production

# JWT
JWT_SECRET=your-super-secret-jwt-key

# Server
SERVER_PORT=8080
GIN_MODE=release

# CORS
ALLOWED_ORIGINS=https://yourdomain.com
```

### Nginx Configuration

Update `nginx.conf` with your domain:

```nginx
server_name yourdomain.com www.yourdomain.com;
```

## 📊 Monitoring

### Health Check

```bash
curl https://yourdomain.com/health
```

### Service Status

```bash
systemctl status tokogo
systemctl status nginx
systemctl status mysql
```

### Logs

```bash
# Application logs
journalctl -u tokogo -f

# Nginx logs
tail -f /var/log/nginx/access.log
tail -f /var/log/nginx/error.log

# MySQL logs
tail -f /var/log/mysql/error.log
```

## 🔄 CI/CD

### GitHub Actions

1. Add secrets to GitHub repository:
   - `HOST`: Server IP address
   - `USERNAME`: SSH username (usually `root`)
   - `SSH_KEY`: Private SSH key

2. Push to main branch to trigger deployment

### Manual Deployment

```bash
cd /root/tokogo
git pull origin main
go mod tidy
go build -o main .
systemctl restart tokogo
```

## 💾 Backup

### Database Backup

```bash
# Make backup script executable
chmod +x backup.sh

# Run backup
./backup.sh

# Setup cron job for daily backup
echo "0 2 * * * /root/tokogo/backup.sh" | crontab -
```

## 🚨 Troubleshooting

### Common Issues

1. **Application won't start**
   ```bash
   journalctl -u tokogo -f
   netstat -tulpn | grep :8080
   ```

2. **Database connection failed**
   ```bash
   systemctl status mysql
   mysql -u tokogo_user -p tokogo_production
   ```

3. **Nginx 502 Bad Gateway**
   ```bash
   tail -f /var/log/nginx/error.log
   systemctl status tokogo
   nginx -t
   ```

## 📚 Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Gin Framework](https://gin-gonic.com/docs/)
- [GORM Documentation](https://gorm.io/docs/)
- [DigitalOcean Tutorials](https://www.digitalocean.com/community/tutorials)
- [Nginx Documentation](https://nginx.org/en/docs/)
- [Let's Encrypt](https://letsencrypt.org/docs/)

## 🎯 Performance Optimization

### Production Tips

1. **Enable Gzip compression** in Nginx
2. **Setup caching** for static files
3. **Use CDN** for static assets
4. **Monitor resource usage** regularly
5. **Setup log rotation** to prevent disk full
6. **Use connection pooling** for database
7. **Enable HTTP/2** in Nginx

### Scaling

For high traffic applications:

1. **Load Balancer**: Use multiple server instances
2. **Database**: Setup read replicas
3. **Caching**: Implement Redis for session storage
4. **CDN**: Use CloudFlare or similar
5. **Monitoring**: Setup Prometheus + Grafana
