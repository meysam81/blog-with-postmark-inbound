# Tarzan - Email-Powered Blogging Platform

<!-- START doctoc -->
<!-- END doctoc -->

[![GitHub Release](https://img.shields.io/github/v/release/meysam81/tarzan)](https://github.com/meysam81/tarzan/releases/latest)
[![License](https://img.shields.io/github/license/meysam81/tarzan)](https://github.com/meysam81/tarzan/blob/main/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/meysam81/tarzan)](https://github.com/meysam81/tarzan/issues)
[![GitHub pull requests](https://img.shields.io/github/issues-pr/meysam81/tarzan)](https://github.com/meysam81/tarzan/pulls)

> **Transforming Your Emails into Blogs!**
> _For Those Who Still Remember SMTP Commands_

Tarzan is a minimalist blogging platform designed for developers who prefer the command line over fancy interfaces. Write your blog posts in your favorite email client and publish them instantly by sending an email.

## 🚀 Features

- **Email-to-Blog Publishing**: Send an email, get a blog post
- **Markdown Support**: Write your content in markdown format
- **Image Attachments**: Include images as email attachments
- **Authorization System**: Secure contributor management via GitHub
- **Responsive Design**: Mobile-friendly reading experience
- **Dark Mode**: Eye-friendly theme switching
- **SQLite Backend**: Lightweight and efficient data storage

## 🎯 Quick Start

### Prerequisites

- Go 1.21 or higher
- SQLite
- Email service with webhook support (e.g., Postmark, SendGrid)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/meysam81/tarzan.git
cd tarzan
```

2. Build the application:

```bash
go build -o tarzan
```

3. Create configuration files:

```bash
cp config.example.json config.json
echo '["your-email@example.com"]' > deploy/authorized_emails.json
```

4. Configure your environment:

```bash
export TARZAN_AUTH_USERNAME=your_webhook_username
export TARZAN_AUTH_PASSWORD=your_webhook_password
```

5. Run the application:

```bash
./tarzan
```

## 📝 How to Publish

### Step 1: Get Authorized

Open a pull request to add your email to the authorized contributors list:

1. Fork this repository
2. Edit `deploy/authorized_emails.json` and add your email
3. Submit a pull request

### Step 2: Write and Send

Once authorized, compose your blog post and send it to your configured webhook endpoint:

**Email Format:**

- **Subject**: Your post title
- **Body**: Your post content (HTML or plain text, markdown supported)
- **Attachments**: Images will be automatically processed and embedded

**Example Email:**

```
To: blog@tarzan.meysam.io
Subject: My First Tarzan Post

# Welcome to Tarzan!

This is my first blog post using **Tarzan**.

## Features I Love

- Simple email interface
- No complex CMS
- Perfect for developers

Check out this awesome screenshot!

[Attach your image to the email - it will be automatically embedded]
```

## 🔧 Configuration

Tarzan uses environment variables for configuration:

| Variable                        | Description            | Default                           |
| ------------------------------- | ---------------------- | --------------------------------- |
| `TARZAN_PORT`                   | Server port            | `8080`                            |
| `TARZAN_AUTH_USERNAME`          | Webhook auth username  | `postmarkapp`                     |
| `TARZAN_AUTH_PASSWORD`          | Webhook auth password  | `Secr3t!`                         |
| `TARZAN_STORAGE_PATH`           | Storage directory path | `./storage`                       |
| `TARZAN_DB_PATH`                | SQLite database path   | `./tarzan.db`                     |
| `TARZAN_AUTHORIZED_EMAILS_PATH` | Authorized emails file | `./deploy/authorized_emails.json` |

## 🛡️ Security

- **Authentication**: Webhook endpoints require HTTP Basic Auth
- **Authorization**: Only pre-approved email addresses can publish
- **Input Validation**: All email content is sanitized
- **File Security**: Attachments are stored securely with UUID filenames

## 🏗️ Architecture

```plaintext
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Email Client  │───▶│  Email Service   │───▶│     Tarzan      │
│                 │    │   (Webhook)      │    │                 │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                                                        │
                                                        ▼
                                               ┌─────────────────┐
                                               │  SQLite Database│
                                               └─────────────────┘
```

## 📁 Project Structure

```plaintext
tarzan/
├── cmd/
│   └── config/          # Configuration management
├── dist/                # Frontend assets
├── src/                 # Frontend source
├── storage/             # Data storage
│   ├── attachments/     # Uploaded images
│   └── tarzan.db        # SQLite database
├── main.go              # Main application
├── index.html           # Frontend template
└── authorized_emails.json # Authorized contributors
```

## 🔗 API Endpoints

- `GET /` - Blog homepage
- `GET /api/posts` - List all posts (JSON)
- `POST /webhook` - Email webhook (authenticated)
- `GET /api/assets/*` - Static file serving

## 🚀 Deployment

### Docker

```bash
docker build -t tarzan .
docker run -p 8080:8080 \
  -e TARZAN_AUTH_USERNAME=user \
  -e TARZAN_AUTH_PASSWORD=pass \
  -v tarzan-data:/data \
  tarzan
```

### Systemd Service

```ini
[Unit]
Description=Tarzan Blog Service
After=network.target

[Service]
Type=simple
User=tarzan
WorkingDirectory=/opt/tarzan
ExecStart=/opt/tarzan/tarzan
Environment=TARZAN_PORT=8080
Environment=TARZAN_AUTH_USERNAME=your_username
Environment=TARZAN_AUTH_PASSWORD=your_password
Restart=always

[Install]
WantedBy=multi-user.target
```

## 🤝 Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes and add tests
4. Commit your changes: `git commit -am 'Add new feature'`
5. Push to the branch: `git push origin feature-name`
6. Submit a pull request

## 📜 License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.

## 🐛 Issues and Support

- **Bug Reports**: [GitHub Issues](https://github.com/meysam81/tarzan/issues)
- **Feature Requests**: [GitHub Discussions](https://github.com/meysam81/tarzan/discussions)

## 🌟 Acknowledgments

- Built with Go and vanilla JavaScript
- Inspired by the simplicity of email and the power of webhooks
- Perfect for developers who think GUIs are overrated

---

**Made with ❤️ by [Meysam Azad](https://github.com/meysam81)**

_Because Real Geeks Send Blog Posts via Postfix_
