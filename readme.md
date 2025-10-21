# Dockerized Go Web Application

A simple Go web server demonstrating Docker containerization with multi-stage builds.

## 🚀 Features

- Simple HTTP web server with multiple endpoints
- Health check endpoint for monitoring
- JSON API endpoint
- Multi-stage Docker build for optimized image size
- Runs on Alpine Linux for minimal footprint


**Direct link to Docker Hub image:**
```
https://hub.docker.com/repository/docker/mohamedmosaad1/go-web-app/general
```

**Pull command:**
```bash
docker pull mohamedmosaad1/go-web-app:latest
```

## 📁 Project Structure

```
.
├── main.go           # Go application source code
├── Dockerfile        # Docker configuration
├── .dockerignore     # Files to exclude from Docker build
└── README.md         # This file
```
