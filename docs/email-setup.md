# 📧 Email Integration Setup Guide

This guide walks you through setting up email integration with Tarzan, allowing you to transform emails into blog posts seamlessly.

## Overview

Tarzan integrates with [Postmark's inbound email service](https://postmarkapp.com/inbound-email) to process incoming emails and convert them into blog posts. This setup requires:

1. **Postmark Account**: Configure inbound processing
1. **Domain Configuration**: Point your domain's MX records to Postmark
1. **Webhook Setup**: Connect Postmark to your Tarzan instance
1. **Authentication**: Secure your webhook endpoint

## Prerequisites

- A domain you control (e.g., `yourdomain.com`)
- Access to your domain's DNS settings
- A [Postmark account](https://account.postmarkapp.com/sign_up) (free tier available)
- A running Tarzan instance (see [Quick Start](../README.md#-quick-start))

## Step 1: Choose Your Email Domain

### Option A: Subdomain (Recommended)

Use a dedicated subdomain for blog emails:

```plaintext
blog.yourdomain.com
posts.yourdomain.com
```

### Option B: Root Domain

Use your main domain directly:

```plaintext
yourdomain.com
```

### Option C: Wildcard Domain

Accept emails to any subdomain:

```plaintext
*.yourdomain.com
```

## Step 2: Configure DNS MX Record

In your DNS management console, create an MX record:

### For Subdomain Setup

| Name/Host | Type | Value                   | Priority | TTL               |
| --------- | ---- | ----------------------- | -------- | ----------------- |
| blog      | MX   | inbound.postmarkapp.com | 10       | 3600 (or default) |

### For Root Domain Setup

| Name/Host    | Type | Value                   | Priority | TTL               |
| ------------ | ---- | ----------------------- | -------- | ----------------- |
| @ (or blank) | MX   | inbound.postmarkapp.com | 10       | 3600 (or default) |

### For Wildcard Setup

| Name/Host | Type | Value                   | Priority | TTL               |
| --------- | ---- | ----------------------- | -------- | ----------------- |
| \*        | MX   | inbound.postmarkapp.com | 10       | 3600 (or default) |

### Common DNS Providers

#### Cloudflare

1. Go to DNS settings
2. Click "Add record"
3. Select "MX" type
4. Enter the values above

#### Namecheap

1. Navigate to Advanced DNS
2. Add new record type "MX Record"
3. Enter the configuration

#### GoDaddy

1. Go to DNS Management
2. Add new MX record
3. Configure as shown above

## Step 3: Set Up Postmark Account

### Create Postmark Server

1. Sign up at [Postmark](https://account.postmarkapp.com/sign_up)
2. Create a new server for inbound processing
3. Note your Server ID and API Token

### Configure Inbound Domain

1. Go to your server's settings
2. Navigate to "Inbound" tab
3. Set your inbound domain (e.g., `blog.yourdomain.com`)
4. Enable "Inbound processing"

### Set Webhook URL

Configure the webhook to point to your Tarzan instance:

```plaintext
https://yourdomain.com/webhook
```

If testing locally, use ngrok or similar:

```plaintext
https://abc123.ngrok.io/webhook
```

## Step 4: Configure Tarzan

### Environment Variables

Set the following environment variables for Tarzan:

```bash
# Authentication for webhook endpoint
export TARZAN_AUTH_USERNAME=postmarkapp
export TARZAN_AUTH_PASSWORD=your-secure-password

# Base URL for your Tarzan instance
export TARZAN_BASE__URL=https://yourdomain.com

# Path to authorized emails list
export TARZAN_DIR_AUTHORIZED_EMAILS=./authorized-emails.json
```

### Create Authorized Emails List

Create a JSON file listing allowed sender email addresses:

```json
["your-email@gmail.com", "team@yourcompany.com", "admin@yourdomain.com"]
```

Save this as `authorized-emails.json` in your Tarzan directory.

### Docker Compose Configuration

If using Docker Compose, update your `.env` file:

```env
TARZAN_AUTH_USERNAME=postmarkapp
TARZAN_AUTH_PASSWORD=SecurePassword123!
TARZAN_BASE__URL=https://yourdomain.com
TARZAN_DIR_AUTHORIZED_EMAILS=/data/authorized-emails.json
```

## Step 5: Configure Postmark Webhook Authentication

In your Postmark server settings:

1. Go to "Webhooks" section
2. Set "Inbound webhook URL": `https://yourdomain.com/webhook`
3. Configure Basic Authentication:
   - Username: `postmarkapp` (or your custom username)
   - Password: Your secure password from Step 4

## Step 6: Test Your Setup

### DNS Propagation Check

Verify your MX record is active:

```bash
dig MX blog.yourdomain.com
```

You should see `inbound.postmarkapp.com` in the response.

### Send Test Email

1. Send an email to any address at your configured domain:

   ```plaintext
   To: test@blog.yourdomain.com
   Subject: My First Blog Post
   Body: This is my first email-powered blog post!
   ```

2. Check your Tarzan instance for the new post
3. Monitor logs for any processing errors

### Troubleshooting Common Issues

#### Email Not Received

- Verify MX record is properly configured
- Check DNS propagation (can take up to 24 hours)
- Ensure sender email is in authorized list

#### Webhook Errors

- Verify webhook URL is accessible
- Check authentication credentials match
- Review Tarzan logs for error details

#### 401 Authentication Errors

- Confirm username/password in both Postmark and Tarzan
- Ensure no special characters causing issues

## Advanced Configuration

### Multiple Domains

Configure multiple domains by:

1. Setting up additional MX records
2. Using wildcard domains
3. Adding multiple inbound domains in Postmark

### Custom Email Processing

Customize how emails become posts by:

- Modifying authorized senders list
- Using email filters in Postmark
- Implementing custom parsing logic

### Security Considerations

#### Webhook Security

- Use HTTPS for webhook URLs
- Implement strong authentication
- Consider IP whitelisting for Postmark IPs

#### Email Security

- Maintain a strict authorized senders list
- Monitor for suspicious activity
- Consider SPF/DKIM validation

## Production Deployment

### SSL Certificate

Ensure your Tarzan instance has a valid SSL certificate:

```bash
# Using Let's Encrypt with Certbot
certbot --nginx -d yourdomain.com
```

### Firewall Configuration

Open necessary ports:

```bash
# HTTP/HTTPS for webhook
ufw allow 80
ufw allow 443

# Tarzan port (if different)
ufw allow 8000
```

### Monitoring

Set up monitoring for:

- Webhook endpoint availability
- Email processing success rate
- Storage space for attachments

## Email Formatting Tips

### Subject Line

The email subject becomes your blog post title:

```plaintext
Subject: How to Build Amazing Web Applications
```

### Content

- Use HTML for rich formatting
- Include images as attachments
- Structure content with headings

### Example Email

```plaintext
To: posts@blog.yourdomain.com
Subject: Welcome to Email-Powered Blogging

<h2>Getting Started</h2>
<p>This is a sample blog post sent via email.</p>

<h3>Features</h3>
<ul>
  <li>Easy publishing</li>
  <li>Rich formatting</li>
  <li>Image attachments</li>
</ul>
```

## Support and Resources

### Documentation

- [Postmark Inbound Email Guide](https://postmarkapp.com/developer/user-guide/inbound)
- [Webhook Configuration](https://postmarkapp.com/developer/webhooks/inbound-webhook)
- [Tarzan Configuration Reference](configuration.md)

### Community Support

- [Zulip Community](https://developer-friendly.zulipchat.com/)
- [GitHub Issues](https://github.com/meysam81/tarzan/issues)
- [Email Setup Discussions](https://github.com/meysam81/tarzan/discussions)

### Professional Support

For enterprise deployments, consider:

- [GitHub Sponsors](https://github.com/sponsors/meysam81) for priority support
- Custom integration assistance
- Production deployment consulting

---

**Need help?** Join our [community chat](https://developer-friendly.zulipchat.com/) or [open an issue](https://github.com/meysam81/tarzan/issues) on GitHub.
