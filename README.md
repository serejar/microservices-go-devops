# Go Microservices Architecture

This project demonstrates a microservice architecture built with Go, featuring containerized services, API communication, Prometheus metrics, and GitHub CI/CD integration.

## Architecture Overview

This system consists of the following services:

- **API Gateway**: Entry point for external requests, routes to appropriate microservices
- **User Service**: Handles user management
- **Product Service**: Manages product data
- **Notification Service**: Sends notifications

Each service is containerized and communicates via RESTful APIs. The system includes monitoring with Prometheus and Grafana.

## Features

- **Modular Architecture**: Each service has a single responsibility
- **API Communication**: Services communicate via RESTful APIs
- **Containerization**: Each service runs in its own Docker container
- **Prometheus Metrics**: Monitoring system performance
- **CI/CD Pipeline**: Automated testing and deployment
- **Structured Logging**: Consistent logging across services
- **Graceful Shutdown**: All services handle shutdown properly

## Getting Started

### Prerequisites

- Go 1.20+
- Docker
- Docker Compose

### Running the Application

1. Clone the repository
2. Run the application:

```bash
npm run start
```

This will start all services using Docker Compose.

### Development

To build the containers:

```bash
npm run build
```

To run tests:

```bash
npm run test
```

To run the linter:

```bash
npm run lint
```

## Service Endpoints

### API Gateway (Port 8080)

- `GET /health`: Health check
- `GET /metrics`: Prometheus metrics
- `GET /api/users`: Get all users
- `GET /api/users/:id`: Get a single user
- `POST /api/users`: Create a user
- `GET /api/products`: Get all products
- `GET /api/products/:id`: Get a single product
- `POST /api/products`: Create a product

### User Service (Port 8081)

- `GET /health`: Health check
- `GET /metrics`: Prometheus metrics
- `GET /users`: Get all users
- `GET /users/:id`: Get a single user
- `POST /users`: Create a user
- `PUT /users/:id`: Update a user
- `DELETE /users/:id`: Delete a user

### Product Service (Port 8082)

- Similar structure to User Service

### Notification Service (Port 8083)

- Similar structure to User Service

## Monitoring

Access Prometheus at: http://localhost:9090
Access Grafana at: http://localhost:3000 (default credentials: admin/admin)

## CI/CD Pipeline

The GitHub Actions workflow includes:

1. Linting the code
2. Running tests
3. Building Docker images
4. Pushing images to Docker Hub
5. Deployment steps (would connect to your environment)

## Project Structure

```
.
├── .github
│   └── workflows           # GitHub Actions workflows
├── prometheus              # Prometheus configuration
├── scripts                 # Utility scripts
└── services                # Microservices
    ├── api-gateway         # API Gateway service
    ├── user-service        # User management service
    ├── product-service     # Product management service
    └── notification-service # Notification service
```

Each service follows a similar structure:

```
service/
├── cmd/                    # Main entry points
├── internal/               # Private application code
│   ├── config/             # Configuration
│   ├── database/           # Database access
│   ├── handlers/           # HTTP handlers
│   ├── middleware/         # HTTP middleware
│   ├── models/             # Data models
│   ├── repository/         # Data access layer
│   └── service/            # Business logic
├── migrations/             # Database migrations
├── Dockerfile              # Container definition
└── go.mod                  # Dependencies
```