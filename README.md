# 📚 BDU OAuth Server

Secure OAuth 2.0 Server for Bahir Dar University (BDU) with institutional login, CSRF protection, CAPTCHA verification, and password reset via email.

---

## ✨ Features

- OAuth 2.0 Authorization Server (Authorization Code Flow, Client Credentials Flow)
- JWT-based access and refresh tokens
- CSRF protection for forms
- CAPTCHA integration to block bots
- Secure password reset system via institutional or custom emails
- Rate limiting and brute-force attack protection
- Modular, scalable codebase (built with Fiber + Golang)

---

## 🚀 Tech Stack

| Component | Technology |
|:----------|:------------|
| Framework | [Fiber](https://gofiber.io/) |
| Language | Golang |
| Database | PostgreSQL / MySQL |
| Token Signing | JWT |
| Password Hashing | Bcrypt |
| Email Service | SMTP / SendGrid |
| Security Layers | CSRF, CAPTCHA, Rate Limiting, HTTPS |

---

## 📦 Project Structure

```bash
bdu-oauth-server/
├── cmd/
│   └── server/       # Fiber server setup
├── internal/
│   ├── auth/         # Authentication and OAuth flows
│   ├── user/         # User management (login, reset password)
│   ├── email/        # Email service (password reset emails)
│   ├── middleware/   # Security middlewares (CSRF, Rate Limiter, CAPTCHA validator)
│   └── token/        # JWT and token management
├── pkg/
│   └── utils/        # Utility functions (password hashing, random token generation)
├── configs/
│   └── config.yaml   # Server, database, SMTP configurations
├── docs/
│   └── architecture.md  # System design and flow diagrams
├── go.mod
├── go.sum
└── README.md
