# 🌟 Tarzan - Revolutionary Email-Powered Blogging Platform

<div align="center">

<!-- Project Status & Quality Badges -->

[![CI/CD Pipeline](https://github.com/meysam81/tarzan/actions/workflows/ci.yml/badge.svg)](https://github.com/meysam81/tarzan/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/meysam81/tarzan)](https://goreportcard.com/report/github.com/meysam81/tarzan)
[![Vulnerability Scan](https://img.shields.io/badge/🛡️_Zero_Vulnerabilities-Kubescape_Verified-brightgreen?style=flat-square)](https://github.com/meysam81/tarzan/actions)

<!-- Release & Distribution -->

[![Latest Release](https://img.shields.io/github/v/release/meysam81/tarzan?style=flat-square&logo=github&color=blue)](https://github.com/meysam81/tarzan/releases/latest)
[![Docker Image](https://img.shields.io/badge/docker-meysam81%2Ftarzan-blue?style=flat-square&logo=docker)](https://ghcr.io/meysam81/tarzan)
[![Docker Pulls](https://img.shields.io/docker/pulls/meysam81/tarzan?style=flat-square&logo=docker)](https://ghcr.io/meysam81/tarzan)
[![Docker Image](https://img.shields.io/docker/image-size/meysam81/tarzan?style=flat-square&logo=docker)](https://ghcr.io/meysam81/tarzan)
[![Container Security](https://img.shields.io/badge/🔐_Container-Cosign_Signed-green?style=flat-square)](https://github.com/meysam81/tarzan/actions)

<!-- License & Community -->

[![License](https://img.shields.io/badge/License-Apache_2.0-green.svg?style=flat-square&logo=apache)](LICENSE)
[![GitHub Sponsors](https://img.shields.io/github/sponsors/meysam81?style=flat-square&logo=github&color=pink)](https://github.com/sponsors/meysam81)
[![Contributors](https://img.shields.io/github/contributors/meysam81/tarzan?style=flat-square&logo=github)](https://github.com/meysam81/tarzan/graphs/contributors)
[![Community Chat](https://img.shields.io/badge/💬_Chat-Zulip-brightgreen?style=flat-square)](https://developer-friendly.zulipchat.com/)

<!-- Technical Stack -->

[![Go Version](https://img.shields.io/github/go-mod/go-version/meysam81/tarzan?style=flat-square&logo=go)](go.mod)
[![Vue.js](https://img.shields.io/badge/Frontend-Vue.js_3-4FC08D?style=flat-square&logo=vue.js)](https://vuejs.org/)
[![TailwindCSS](https://img.shields.io/badge/Styling-TailwindCSS-38B2AC?style=flat-square&logo=tailwind-css)](https://tailwindcss.com/)
[![SQLite | Redis](https://img.shields.io/badge/Database-SQLite_|_Redis-003B57?style=flat-square&logo=sqlite)](https://sqlite.org/)

<!-- Code Quality & Tools -->

[![oxlint](https://img.shields.io/badge/Linting-oxlint-FF6B6B?style=flat-square&logo=eslint)](https://oxc-project.github.io/docs/guide/usage/linter.html)
[![Pre-commit](https://img.shields.io/badge/Quality-pre--commit_hooks-FAB040?style=flat-square&logo=pre-commit)](https://pre-commit.com/)
[![12-Factor App](https://img.shields.io/badge/Architecture-12--Factor_App-blueviolet?style=flat-square)](https://12factor.net/)
[![PWA Ready](https://img.shields.io/badge/📱_PWA-Ready-purple?style=flat-square)](https://web.dev/progressive-web-apps/)

<!-- Features & Capabilities -->

[![Email Powered](https://img.shields.io/badge/✉️_Email--Powered-Revolutionary-ff6b35?style=flat-square)](https://postmarkapp.com/inbound-email)
[![Self Hosted](https://img.shields.io/badge/🏠_Self--Hosted-Zero_Dependencies-success?style=flat-square)](docs/email-setup.md)
[![Air Gap Ready](https://img.shields.io/badge/🔒_Air--Gap-Compatible-darkgreen?style=flat-square)](#deployment-excellence)
[![RSS & Sitemap](https://img.shields.io/badge/📡_RSS_&_Sitemap-Native_Support-orange?style=flat-square)](#technical-superiority)

<!-- Demo & Documentation -->

[![Demo](https://img.shields.io/badge/🚀_Live_Demo-tarzan.meysam.io-blue?style=for-the-badge)](https://tarzan.meysam.io)
[![Sponsor](https://img.shields.io/badge/💖_Sponsor_on_GitHub-@meysam81-pink?style=for-the-badge)](https://github.com/sponsors/meysam81)

</div>

> **Transform your inbox into a publishing powerhouse with revolutionary email-powered blogging.**

Tarzan revolutionizes content creation by turning every email you send into a beautifully rendered blog post. Built with cutting-edge technology and an unwavering commitment to simplicity, Tarzan empowers creators to publish content effortlessly while maintaining complete control over their data and infrastructure.

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [✨ Why Choose Tarzan?](#-why-choose-tarzan)
  - [🎯 **Effortless Publishing**](#-effortless-publishing)
  - [🔒 **Complete Ownership & Privacy**](#-complete-ownership--privacy)
  - [⚡ **Lightning-Fast Performance**](#-lightning-fast-performance)
  - [🎨 **Stunning Visual Design**](#-stunning-visual-design)
- [🚀 Key Features](#-key-features)
  - [**📦 Deployment Excellence**](#-deployment-excellence)
  - [**🛠 Technical Superiority**](#-technical-superiority)
  - [**🔐 Security & Authentication**](#-security--authentication)
  - [**🌟 Open Source Excellence**](#-open-source-excellence)
- [🚀 Quick Start](#-quick-start)
  - [Docker Compose (Recommended)](#docker-compose-recommended)
  - [Docker](#docker)
  - [Binary Download](#binary-download)
  - [Kubernetes Deployment (Enterprise Ready)](#kubernetes-deployment-enterprise-ready)
- [⚙️ Configuration](#-configuration)
- [📧 Email Integration](#-email-integration)
- [🎨 Customization](#-customization)
- [🤝 Community & Support](#-community--support)
  - [Get Help](#get-help)
  - [Support Development](#support-development)
- [📈 Performance & Reliability](#-performance--reliability)
- [🔄 Migration & Vendor Freedom](#-migration--vendor-freedom)
- [📱 Modern Web Standards](#-modern-web-standards)
- [🏗️ Architecture](#-architecture)
- [📄 License](#-license)
- [🎯 What's Next?](#-whats-next)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## ✨ Why Choose Tarzan?

### 🎯 **Effortless Publishing**

Simply send an email to your configured domain, and watch as Tarzan transforms it into a polished blog post. No complex interfaces, no learning curves—just pure, intuitive content creation.

### 🔒 **Complete Ownership & Privacy**

Your content, your rules. Tarzan is completely self-hostable with no vendor lock-in, ensuring your valuable content remains under your complete control.

### ⚡ **Lightning-Fast Performance**

Built with Go and powered by native SQLite, Tarzan delivers exceptional speed and efficiency, making your blog blazingly fast for both you and your readers.

### 🎨 **Stunning Visual Design**

Experience exceptional styling and UI with beautiful colors and responsive design that adapts seamlessly across all devices and screen sizes.

## 🚀 Key Features

### **📦 Deployment Excellence**

- **Single Binary Deployment**: Entire application packaged in one compiled binary with all batteries included
- **No Dynamic Dependencies**: Perfect for air-gapped environments and secure deployments
- **12-Factor App Compliance**: Highly configurable and cloud-native ready
- **Docker Support**: Containerized with native Docker Compose support - deploy with one command
- **PWA Ready**: Install as a Progressive Web App on any desktop or mobile device

### **🛠 Technical Superiority**

- **Go Backend**: Ultra-fast, minimal dependencies, exceptional performance
- **Vue.js Frontend**: Modern, component-based UI compiled with Vite into minified static assets
- **SQLite Native**: Local data storage with configurable database location
- **TailwindCSS**: Future-proof styling with extensive customization capabilities
- **RSS & Sitemap**: Built-in support for content syndication and SEO
- **Oxlint Integration**: Blazingly fast ESLint alternative for superior code quality and performance
- **Pre-commit Hooks**: Comprehensive automated checks ensuring consistent code formatting, security, and maintainability

### **🔐 Security & Authentication**

- **Basic Authentication**: Secure webhook endpoints with configurable credentials
- **Inbound Email Processing**: Seamless integration with Postmark's inbound email system
- **SEO Optimized**: Configurable SEO settings (create your own robots.txt)

### **🌟 Open Source Excellence**

- **Apache 2.0 License**: OSI-approved permissive licensing
- **Professional CI/CD**: State-of-the-art pipeline with cosign-signed releases
- **Zero Vulnerabilities**: Kubescape security scans on every commit ensure container images & compiled binaries remain vulnerability-free
- **Exceptional Code Organization**: DRY, KISS principles with dependency injection
- **Long-term Maintainability**: Static typing, separation of concerns, modular architecture

## 🚀 Quick Start

### Docker Compose (Recommended)

Get Tarzan running in seconds with our one-command deployment:

```bash
curl -o .env https://raw.githubusercontent.com/meysam81/tarzan/main/.env.example
curl -o docker-compose.yml https://raw.githubusercontent.com/meysam81/tarzan/main/compose.yml
# modify the .env as desired
docker compose up -d
```

Your blog will be accessible at `http://localhost:8000`

### Docker

```bash
docker run -d \
  --name tarzan \
  -p 8000:8000 \
  -v tarzan-data:/data \
  ghcr.io/meysam81/tarzan:latest
```

### Binary Download

Download the latest release for your platform:

```bash
# Linux/macOS
curl -L https://github.com/meysam81/tarzan/releases/latest/download/tarzan_linux_amd64.tar.gz | tar xz
./tarzan

# Windows
# Download from GitHub releases page
```

### Kubernetes Deployment (Enterprise Ready)

Tarzan includes official Kubernetes deployment manifests that are security-hardened for production environments:

```yaml
# tarzan/kustomization.yml
resources:
  - https://github.com/meysam81/tarzan//deploy/k8s?ref=v1.0.0&timeout=2m

namespace: default
```

And then apply this resource:

```shell
kubectl apply -k ./tarzan/
```

The Kubernetes deployment includes:

- **Security-hardened containers**: Non-root user, read-only filesystem, dropped capabilities
- **Health checks**: Configurable liveness and readiness probes
- **Resource management**: Optimized for efficient cluster resource usage
- **Persistent storage**: Configurable PVC for data persistence
- **Production-ready**: Rolling updates with zero-downtime deployment strategy

## ⚙️ Configuration

Tarzan follows 12-factor app principles for maximum flexibility:

```bash
# Basic configuration
export TARZAN_PORT=8000
export TARZAN_BASE__URL=https://yourdomain.com
export TARZAN_AUTH_USERNAME=your-username
export TARZAN_AUTH_PASSWORD=your-secure-password

# Advanced options
export TARZAN_DIR_DB=/path/to/database.db
export TARZAN_DIR_STORAGE=/path/to/attachments
```

## 📧 Email Integration

Configure your domain's MX record to point to Postmark's inbound servers:

1. **Set MX Record**: Point your domain to `inbound.postmarkapp.com`
2. **Configure Webhook**: Set your Tarzan instance as the webhook endpoint
3. **Start Publishing**: Send emails to your domain and watch them become blog posts

_For detailed setup instructions, visit our [Email Integration Guide](docs/email-setup.md)_

## 🎨 Customization

Tarzan's beautiful interface is built with TailwindCSS, making customization straightforward:

- Modify color schemes and typography
- Adjust responsive breakpoints
- Customize component layouts
- Add your brand identity

## 🤝 Community & Support

### Get Help

- **Community Chat**: Join our [Zulip Server](https://developer-friendly.zulipchat.com/) for real-time support
- **Issues**: Report bugs or request features on GitHub

### Support Development

Love Tarzan? Consider supporting continued development:

[![Sponsor](https://img.shields.io/badge/💖_Sponsor_on_GitHub-@meysam81-pink?style=for-the-badge)](https://github.com/sponsors/meysam81)

Your support helps maintain this project and develop new features that benefit the entire community.

## 📈 Performance & Reliability

- **Sub-millisecond Response Times**: Go's efficiency meets SQLite's speed
- **Minimal Resource Usage**: Runs efficiently on modest hardware
- **Air-Gap Compatible**: No external dependencies
- **Automated Testing**: Comprehensive CI/CD with release automation
- **Production Ready**: Battle-tested in real-world deployments

## 🔄 Migration & Vendor Freedom

Tarzan believes in data portability and vendor freedom:

- Export your content in multiple formats
- SQLite database ensures easy data access
- No proprietary formats or lock-in mechanisms
- Migrate to/from other platforms effortlessly

## 📱 Modern Web Standards

- **Progressive Web App**: Install and use like a native application
- **Responsive Design**: Perfect experience across all device types
- **Accessibility**: WCAG compliant with screen reader support
- **Modern JavaScript**: ES6+ with optimal browser compatibility

## 🏗️ Architecture

Tarzan's architecture prioritizes maintainability and extensibility:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Vue.js SPA    │    │   Go HTTP Server │    │  SQLite Database│
│  (TailwindCSS)  │◄──►│                  │◄──►│   (Local File)  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         │                       ▼                       │
         │              ┌──────────────────┐             │
         └─────────────►│  Static Assets   │◄────────────┘
                        │  (Embedded)      │
                        └──────────────────┘
```

## 📄 License

Tarzan is proudly open source under the [Apache License 2.0](LICENSE). This OSI-approved license ensures maximum freedom for both personal and commercial use while maintaining attribution to the original creators.

## 🎯 What's Next?

- **Plugin System**: Extensible architecture for community contributions
- **Multi-language Support**: Internationalization for global audiences
- **Advanced Theming**: Visual theme marketplace and customization tools
- **Enhanced Analytics**: Detailed insights into content performance

---

<div align="center">

**Ready to revolutionize your blogging experience?**

[🚀 **Start Your Journey**](https://tarzan.meysam.io) | [📖 **Read the Docs**](docs/) | [💬 **Join Community**](https://developer-friendly.zulipchat.com/)

_Built with ❤️ by [@meysam81](https://github.com/meysam81) and the open source community_

</div>
