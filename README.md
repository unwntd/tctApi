# tctAPI

A robust REST API built with Go, designed for managing users, organizers, and roles. This project implements clean architecture principles, JWT-based authentication, and uses PostgreSQL as the database backend.

## Features

- **User Management**: Registration and login with secure password hashing
- **Authentication**: JWT access tokens with refresh token support
- **Organizer Management**: Full CRUD operations for organizers
- **Role Management**: Full CRUD operations for roles
- **Clean Architecture**: Organized into layers (handler, usecase, repository)
- **Database**: PostgreSQL with GORM ORM
- **Configuration**: Environment-based configuration using Viper
- **CORS Support**: Configured for cross-origin requests
- **Migration**: Automated database migration

## Tech Stack

- **Language**: Go 1.25.1
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT (golang-jwt/jwt/v5)
- **Configuration**: Viper
- **Password Hashing**: bcrypt
- **CORS**: gin-contrib/cors

## Prerequisites

- Go 1.25.1 or later
- PostgreSQL database
- Git

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd tctAPI
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your PostgreSQL database and update the configuration (see Configuration section).

## Configuration

Create a `.env` file in the root directory with the following variables:

```env
APP_NAME=tctAPI
APP_PORT=8080
DB_HOST=localhost
DB_USER=your_db_user
DB_PASS=your_db_password
DB_NAME=tctapi
DB_PORT=5432
JWT_SECRET=your-secret-key
APP_MODE=debug
```

### Default Values
If `.env` is not found, the application will use default values:
- APP_NAME: tctAPI
- APP_PORT: 8080
- DB_HOST: localhost
- DB_USER: root
- DB_PASS: (empty)
- DB_NAME: tctapi
- DB_PORT: 3306 (Note: Config shows 3306 but code uses PostgreSQL)
- JWT_SECRET: your-secret-key
- APP_MODE: debug

## Running the Application

1. Run database migration (optional, runs automatically on first start):
   ```bash
   go run cmd/main.go migrate
   ```

2. Start the server:
   ```bash
   go run cmd/main.go
   ```

The API will be available at `http://localhost:8080` (or the port specified in config).

## API Endpoints

### Authentication
- `POST /api/v1/users/register` - Register a new user
- `POST /api/v1/users/login` - Login user
- `POST /api/v1/auth/login` - Alternative login endpoint
- `GET /api/v1/auth/refresh` - Refresh access token

### Organizers (Protected)
- `POST /api/v1/organizers/create` - Create organizer
- `GET /api/v1/organizers/all` - Get all organizers
- `PUT /api/v1/organizers/update/:id` - Update organizer
- `DELETE /api/v1/organizers/delete/:id` - Delete organizer

### Roles (Protected)
- `POST /api/v1/roles/create` - Create role
- `GET /api/v1/roles/all` - Get all roles
- `PUT /api/v1/roles/update/:id` - Update role
- `DELETE /api/v1/roles/delete/:id` - Delete role

**Note**: Protected endpoints require JWT authentication via Authorization header.

## Project Structure

```
tctAPI/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── auth/                   # Authentication module
│   │   ├── handler.go
│   │   ├── jwt_service.go
│   │   ├── middlerware.go
│   │   ├── model.go
│   │   ├── repository.go
│   │   └── service.go
│   ├── organizer/              # Organizer module
│   │   ├── dto.go
│   │   ├── model.go
│   │   ├── handler/
│   │   │   └── handler.go
│   │   ├── repository/
│   │   │   └── repo.go
│   │   └── usecase/
│   │       └── usecase.go
│   ├── role/                   # Role module
│   │   ├── dto.go
│   │   ├── model.go
│   │   ├── handler/
│   │   │   └── handler.go
│   │   ├── repository/
│   │   │   └── repo.go
│   │   └── usecase/
│   │       └── usecase.go
│   ├── router/
│   │   └── router.go           # Route setup
│   └── user/                   # User module
│       ├── dto.go
│       ├── model.go
│       ├── handler/
│       │   └── handler.go
│       ├── repository/
│       │   └── repo.go
│       └── usecase/
│           └── usecase.go
├── pkg/
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── database/
│   │   ├── db.go               # Database connection
│   │   └── migration.go        # Database migration
│   ├── helper/
│   │   └── helper.go           # Utility functions
│   └── response/
│       └── response.go         # Response helpers
├── .env                        # Environment variables (create this)
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Usage Example

### Register a User
```bash
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "password": "password123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{"email": "john@example.com", "password": "password123"}'
```

### Create Organizer (with JWT token)
```bash
curl -X POST http://localhost:8080/api/v1/organizers/create \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"name": "Event Corp", "address": "123 Main St", "pic_name": "Jane Smith", "pic_phone_number": "1234567890", "pic_email": "jane@eventcorp.com", "logo_url": "https://example.com/logo.png"}'
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
