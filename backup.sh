#!/bin/bash
# backup.sh

DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR="/root/backups"
DB_NAME="tokogo_production"
DB_USER="tokogo_user"

# Create backup directory
mkdir -p $BACKUP_DIR

# Backup database
mysqldump -u $DB_USER -p$DB_PASSWORD $DB_NAME > $BACKUP_DIR/tokogo_$DATE.sql

# Compress backup
gzip $BACKUP_DIR/tokogo_$DATE.sql

# Keep only last 7 days of backups
find $BACKUP_DIR -name "tokogo_*.sql.gz" -mtime +7 -delete

echo "Backup completed: tokogo_$DATE.sql.gz"
