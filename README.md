# Question Answer Web Application

A web application built with Go that provides a platform for questions and answers. This project uses the Gin web framework and implements session-based authentication.

## Project Structure

```
.
├── controllers/    # Request handlers and business logic
├── globals/       # Global variables and configurations
├── middleware/    # Custom middleware functions
├── models/        # Data models and database schemas
├── routes/        # Route definitions and grouping
├── service/       # Business service layer
├── templates/     # HTML templates (.gohtml files)
├── main.go        # Application entry point
├── go.mod         # Go module dependencies
└── go.sum         # Go module checksum file
```

## Technologies Used

- [Go](https://golang.org/) - Programming language
- [Gin](https://gin-gonic.com/) - Web framework
- [Gin Sessions](https://github.com/gin-contrib/sessions) - Session management
- HTML Templates - Server-side rendering

## Prerequisites

- Go 1.16 or higher
- Environment variables:
  - `PORT`: The port number for the server to listen on

## Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/emre-guler/question-answer.git
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up environment variables:
   ```bash
   export PORT=8080
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The application will start on `localhost:PORT`.

## Features

- Public and private routes
- Session-based authentication
- Static file serving
- HTML template rendering
- Middleware support

## Project Structure Details

- `controllers/`: Contains the request handlers that process incoming HTTP requests
- `globals/`: Stores global variables and configuration settings
- `middleware/`: Custom middleware including authentication checks
- `models/`: Data structures and database models
- `routes/`: Route definitions split into public and private groups
- `service/`: Business logic layer
- `templates/`: GoHTML templates for rendering views

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request