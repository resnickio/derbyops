# DerbyOps API

A RESTful API for managing roller derby events and leagues.

## Features

- User authentication and authorization
- CRUD operations for derby events
- User management
- JWT-based authentication
- PostgreSQL database
- RESTful API design

## Prerequisites

- Go 1.21 or later
- PostgreSQL 14 or later
- Make (optional, for using Makefile commands)

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/resnickio/derbyops.git
   cd derbyops
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a PostgreSQL database:
   ```sql
   CREATE DATABASE derbyops;
   ```

4. Copy the environment file and update the values:
   ```bash
   cp .env.example .env
   ```

5. Update the `.env` file with your database credentials and JWT secret.

## Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

2. The server will start on `http://localhost:8080` by default.

## API Endpoints

### Authentication

- `POST /api/v1/auth/register` - Register a new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/logout` - Logout user

### Users

- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Derbies

- `GET /api/v1/derbies` - Get all derbies
- `POST /api/v1/derbies` - Create new derby
- `GET /api/v1/derbies/:id` - Get derby by ID
- `PUT /api/v1/derbies/:id` - Update derby
- `DELETE /api/v1/derbies/:id` - Delete derby

## Development

### Project Structure

```
.
├── api/
│   ├── auth/        # Authentication related code
│   ├── database/    # Database connection and setup
│   ├── handlers/    # HTTP request handlers
│   ├── middleware/  # Custom middleware
│   └── models/      # Data models
├── .env             # Environment variables
├── .gitignore      # Git ignore file
├── go.mod          # Go module file
├── main.go         # Application entry point
└── README.md       # Project documentation
```

### Adding New Features

1. Create new models in `api/models/`
2. Add handlers in `api/handlers/`
3. Update routes in `api/handlers/api.go`
4. Add any necessary middleware in `api/middleware/`

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.