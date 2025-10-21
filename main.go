package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type HealthCheck struct {
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/info", infoHandler)

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Dockerized Go App</title>
		<style>
			body {
				font-family: Arial, sans-serif;
				max-width: 800px;
				margin: 50px auto;
				padding: 20px;
				background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
				color: white;
			}
			.container {
				background: rgba(255,255,255,0.1);
				padding: 30px;
				border-radius: 10px;
				backdrop-filter: blur(10px);
			}
			h1 { margin-top: 0; }
			.endpoint {
				background: rgba(255,255,255,0.2);
				padding: 15px;
				margin: 10px 0;
				border-radius: 5px;
			}
			code {
				background: rgba(0,0,0,0.3);
				padding: 2px 8px;
				border-radius: 3px;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<h1>🐳 Dockerized Go Application</h1>
			<p>Welcome! This is a simple Go web server running in a Docker container.</p>
			
			<h2>Available Endpoints:</h2>
			<div class="endpoint">
				<strong>GET /</strong> - This page
			</div>
			<div class="endpoint">
				<strong>GET /health</strong> - Health check endpoint
			</div>
			<div class="endpoint">
				<strong>GET /api/info</strong> - JSON API endpoint with app information
			</div>
			
			<p style="margin-top: 30px; font-size: 0.9em; opacity: 0.8;">
				Running on Go | Containerized with Docker 🚀
			</p>
		</div>
	</body>
	</html>
	`
	fmt.Fprint(w, html)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	health := HealthCheck{
		Status: "healthy",
		Time:   time.Now(),
	}
	json.NewEncoder(w).Encode(health)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message:   "Hello from Dockerized Go Application!",
		Timestamp: time.Now(),
		Version:   "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}