# ⚙️ Configuration Reference

This comprehensive guide covers all configuration options available in Tarzan, following the 12-factor app methodology for maximum flexibility and deployment convenience.

## Overview

Tarzan uses environment variables for configuration, making it cloud-native and container-friendly. All configuration options can be set via environment variables with the `TARZAN_` prefix.

## Configuration Methods

### Environment Variables

The primary and recommended method:

```bash
export TARZAN_PORT=8000
export TARZAN_BASE__URL=https://yourdomain.com
```

### .env File

For development and Docker Compose deployments:

```bash
# Copy the example file
cp .env.example .env
# Edit with your values
nano .env
```

### Docker Environment

```bash
docker run -e TARZAN_PORT=8000 -e TARZAN_BASE__URL=https://yourdomain.com ghcr.io/meysam81/tarzan:latest
```

## Configuration Options

### Server Configuration

#### `TARZAN_PORT`

- **Description**: Port number for the HTTP server
- **Type**: Integer
- **Default**: `8000`
- **Example**: `TARZAN_PORT=3000`

#### `TARZAN_BASE__URL`

- **Description**: Base URL where Tarzan is accessible (used for generating links, RSS feeds, and sitemaps)
- **Type**: String
- **Default**: `"https://tarzan.meysam.io"`
- **Example**: `TARZAN_BASE__URL=https://blog.yourdomain.com`
- **Note**: Include protocol (http/https) and exclude trailing slash

### Authentication Configuration

#### `TARZAN_AUTH_USERNAME`

- **Description**: Username for webhook authentication
- **Type**: String
- **Default**: `"postmarkapp"`
- **Example**: `TARZAN_AUTH_USERNAME=webhook_user`
- **Security**: Use a unique username for better security

#### `TARZAN_AUTH_PASSWORD`

- **Description**: Password for webhook authentication
- **Type**: String
- **Default**: `"Secr3t!"`
- **Example**: `TARZAN_AUTH_PASSWORD=MyVerySecurePassword123!`
- **Security**: Use a strong, unique password with special characters

### Directory Configuration

Control where Tarzan stores its data and finds required files.

#### `TARZAN_DIR_DB`

- **Description**: Path to SQLite database file
- **Type**: String
- **Default**: `"tarzan.db"`
- **Example**: `TARZAN_DIR_DB=/data/tarzan.db`
- **Note**: Ensure directory exists and is writable

#### `TARZAN_DIR_STORAGE`

- **Description**: Directory path for storing uploaded files and attachments
- **Type**: String
- **Default**: `"storage"`
- **Example**: `TARZAN_DIR_STORAGE=/var/www/uploads`
- **Note**: Must be writable by the application

## Configuration Examples

### Development Setup

```bash
# .env file for local development
TARZAN_PORT=8000
TARZAN_BASE__URL=http://localhost:8000
TARZAN_AUTH_USERNAME=dev_user
TARZAN_AUTH_PASSWORD=dev_password
TARZAN_DIR_DB=./dev.db
TARZAN_DIR_STORAGE=./dev_storage
```

### Production Docker Setup

```bash
# .env file for production
TARZAN_PORT=8000
TARZAN_BASE__URL=https://blog.company.com
TARZAN_AUTH_USERNAME=webhook_prod
TARZAN_AUTH_PASSWORD=ProductionSecurePassword123!
TARZAN_DIR_DB=/data/tarzan.db
TARZAN_DIR_STORAGE=/data/storage
```

## Special Configuration Files

### Docker Compose Configuration

```yaml
services:
  tarzan:
    image: ghcr.io/meysam81/tarzan:latest
    ports:
      - "8000:8000"
    environment:
      - TARZAN_BASE__URL=https://yourdomain.com
      - TARZAN_AUTH_USERNAME=postmarkapp
      - TARZAN_AUTH_PASSWORD=SecurePassword123!
    volumes:
      - tarzan_data:/data
      - ./authorized-emails.json:/data/authorized-emails.json:ro

volumes:
  tarzan_data:
```

## Security Best Practices

### Authentication

- Use strong, unique passwords with special characters
- Change default usernames from `postmarkapp`
- Rotate credentials regularly
- Use environment variables, never hardcode secrets

### Database Security

- Store database files in secure, backed-up locations
- Ensure proper file permissions (600 for database files)
- Regular database backups

## Performance Tuning

### Database Optimization

```bash
# Use SSD storage for database
TARZAN_DIR_DB=/fast-ssd/tarzan.db

# Separate storage location
TARZAN_DIR_STORAGE=/large-storage/uploads
```

## Troubleshooting

### Common Issues

#### Port Conflicts

```bash
# Check if port is in use
netstat -tlnp | grep :8000

# Use different port
TARZAN_PORT=8080
```

#### Permission Errors

```bash
# Fix database permissions
chmod 644 /path/to/tarzan.db
chown app:app /path/to/tarzan.db

# Fix storage permissions
chmod -R 755 /path/to/storage
chown -R app:app /path/to/storage
```

### Configuration Validation

Tarzan validates configuration on startup. Check logs for:

```plaintext
Invalid configuration: missing required field
Database connection failed
Authorized emails file not found
```

## Environment-Specific Configurations

### Development

- Enable debug logging
- Use HTTP instead of HTTPS for base URL

### Staging

- Mirror production configuration
- Test webhook endpoints thoroughly

### Production

- Enable all security features
- Configure proper monitoring
- Regular backup procedures

## Configuration Management

### Version Control

- Never commit `.env` files with real credentials
- Use `.env.example` as template
- Document required variables in README

### Configuration Updates

- Changes require application restart
- Test configuration changes in staging first
- Monitor application logs after configuration updates

## Monitoring Configuration

Track configuration health with these patterns:

### Health Checks

```bash
# Basic health check
curl http://localhost:$TARZAN_PORT/health

# Database connectivity
curl http://localhost:$TARZAN_PORT/api/status
```

### Logging Configuration

- Monitor configuration errors in application logs
- Set up alerts for authentication failures
- Track database connection issues

---

For additional support with configuration, visit our [community chat](https://developer-friendly.zulipchat.com/) or [email setup guide](email-setup.md).
