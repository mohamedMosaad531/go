# Dockerized Go Web Application

A simple Go web server demonstrating Docker containerization with multi-stage builds.

## 🚀 Features

- Simple HTTP web server with multiple endpoints
- Health check endpoint for monitoring
- JSON API endpoint
- Multi-stage Docker build for optimized image size
- Runs on Alpine Linux for minimal footprint

## 📋 Prerequisites

- Go 1.21 or higher (for local development)
- Docker installed and running
- Docker Hub account (for pushing images)

## 🛠️ Local Development

### Run without Docker

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Available Endpoints

- `GET /` - Home page with application info
- `GET /health` - Health check endpoint (returns JSON)
- `GET /api/info` - API endpoint with app information (returns JSON)

## 🐳 Docker Instructions

### Build the Docker Image

```bash
docker build -t your-dockerhub-username/go-web-app:latest .
```

Replace `your-dockerhub-username` with your actual Docker Hub username.

### Run the Docker Container Locally

```bash
docker run -p 8080:8080 your-dockerhub-username/go-web-app:latest
```

Access the application at `http://localhost:8080`

### Run in Detached Mode

```bash
docker run -d -p 8080:8080 --name go-app your-dockerhub-username/go-web-app:latest
```

### View Logs

```bash
docker logs go-app
```

### Stop the Container

```bash
docker stop go-app
docker rm go-app
```

## 📤 Push to Docker Hub

### Login to Docker Hub

```bash
docker login
```

### Tag the Image (if not already tagged)

```bash
docker tag go-web-app:latest your-dockerhub-username/go-web-app:latest
```

### Push to Docker Hub

```bash
docker push your-dockerhub-username/go-web-app:latest
```

## 🔗 Docker Hub Image

**Direct link to Docker Hub image:**
```
https://hub.docker.com/r/your-dockerhub-username/go-web-app
```

**Pull command:**
```bash
docker pull your-dockerhub-username/go-web-app:latest
```

## 📁 Project Structure

```
.
├── main.go           # Go application source code
├── Dockerfile        # Docker configuration
├── .dockerignore     # Files to exclude from Docker build
└── README.md         # This file
```

## 🏗️ Docker Build Details

This project uses a **multi-stage build** approach:

1. **Builder Stage**: Uses `golang:1.21-alpine` to compile the Go application
2. **Final Stage**: Uses `alpine:latest` for a minimal runtime environment

This approach results in a significantly smaller final image (~15MB vs ~300MB+).

### Image Optimization

- Multi-stage build reduces final image size
- Alpine Linux base for minimal footprint
- Only includes compiled binary and necessary certificates
- No Go toolchain in final image

## 🧪 Testing the Application

### Test Endpoints

```bash
# Home page
curl http://localhost:8080/

# Health check
curl http://localhost:8080/health

# API info
curl http://localhost:8080/api/info
```

### Expected Responses

**Health Check:**
```json
{
  "status": "healthy",
  "time": "2025-10-21T10:30:00Z"
}
```

**API Info:**
```json
{
  "message": "Hello from Dockerized Go Application!",
  "timestamp": "2025-10-21T10:30:00Z",
  "version": "1.0.0"
}
```

## 🔧 Environment Variables

- `PORT` - Server port (default: 8080)

Example with custom port:
```bash
docker run -p 3000:3000 -e PORT=3000 your-dockerhub-username/go-web-app:latest
```

## 📝 Notes

- The application is stateless and suitable for container orchestration
- Health check endpoint can be used with Docker health checks or Kubernetes probes
- The image is publicly accessible on Docker Hub

## 🤝 Contributing

Feel free to fork this repository and submit pull requests for any improvements.

## 📄 License

This project is open source and available under the MIT License.