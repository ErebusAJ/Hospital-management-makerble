# Hospital Management System Backend

A simple, secure, and scalable backend for a Hospital Management System, built with Go, Gin, PostgreSQL, and SQLC. This service provides RESTful APIs for authentication, patient management, receptionist operations, doctor portals, and patient medical histories.

## 🔍 Features

- **Unified Authentication** via JWT for Receptionists and Doctors
- **Receptionist APIs**: Register, list, update, delete patients
- **Doctor APIs**: View assigned patients, update medical history
- **Patient History**: CRUD operations on patient visit records
- **Secure**: Role-based access control, password hashing (bcrypt)
- **Type-safe** database queries**: SQLC generates Go code from SQL
- **Migration**: Goose for managing schema changes
- **Dockerized**: Easily build and deploy with Docker

## 📦 Technology Stack

- **Language**: Go
- **Framework**: Gin HTTP web framework
- **Database**: PostgreSQL
- **ORM/Query**: SQLC (type-safe SQL to Go)
- **Migrations**: Goose
- **Containerization**: Docker

## 🚀 Getting Started

### Prerequisites

- Go (>= 1.18)
- PostgreSQL (>= 12)
- Docker (optional)

### Clone the Repository

```bash
git clone https://github.com/yourusername/hospital-backend.git
cd hospital-backend
```

### Environment Variables

Create a `.env` file in project root:

```dotenv
# Server
PORT_NO=8080

DB_URL="host=dbHost port=5432 user=username password=pass dbname=db sslmode=require"
#JWT
SECRET_KEY=secret

API=hostedEndpoint
```

### Database Setup & Migrations

Install Goose CLI:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run migrations:
```bash
goose -dir migrations postgres "user=$DB_USER password=$DB_PASSWORD host=$DB_HOST dbname=$DB_NAME sslmode=disable" up
```

### Generate SQLC Code

```bash
sqlc generate
```

### Run the Server

```bash
go build cmd/main.go
./main
```

The API will be available at `http://localhost:8080` or your hosted address.

## 🛠️ Docker

Build the image:
```bash
docker build -t hospital-backend .
```

Run the container:
```bash
docker run --name hospital-backend \
  -p 8080:8080 \
  --env-file .env \
  hospital-backend
```

## 📑 API Endpoints

Open Postman import the ***api-endpoint-details.json***

* It contains the documentation of endpoints

> Note: All protected routes require `Authorization: Bearer <token>` header.


*Last updated: 1st May 2025*

